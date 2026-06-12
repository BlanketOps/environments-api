//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/magefile/mage/sh"
)

const (
	colorReset = "\033[0m"
	colorGreen = "\033[32m"
	colorCyan  = "\033[36m"
	colorRed   = "\033[31m"
)

func step(msg string)    { fmt.Printf("%s▶ %s%s\n", colorCyan, msg, colorReset) }
func success(msg string) { fmt.Printf("%s✔ %s%s\n", colorGreen, msg, colorReset) }
func fail(msg string)    { fmt.Printf("%s✘ %s%s\n", colorRed, msg, colorReset) }

// Generate runs deepcopy code generation.
func Generate() error {
	step("Generating deepcopy...")
	if err := installControllerGen(); err != nil {
		return err
	}
	if err := sh.RunV("bin/controller-gen", `object:headerFile=hack/boilerplate.go.txt`, `paths=./...`); err != nil {
		fail("Generate failed")
		return err
	}
	success("Generate complete")
	return nil
}

// Manifests generates CRDs from API types.
func Manifests() error {
	step("Generating CRDs...")
	if err := installControllerGen(); err != nil {
		return err
	}
	if err := sh.RunV("bin/controller-gen", "crd", `paths=./api/...`, `output:crd:artifacts:config=config/crd/bases`); err != nil {
		fail("Manifests failed")
		return err
	}
	success("Manifests complete")
	return nil
}

// Fmt runs go fmt.
func Fmt() error {
	step("Running go fmt...")
	if err := sh.RunV("go", "fmt", "./..."); err != nil {
		fail("Fmt failed")
		return err
	}
	success("Fmt complete")
	return nil
}

// Vet runs go vet.
func Vet() error {
	step("Running go vet...")
	if err := sh.RunV("go", "vet", "./..."); err != nil {
		fail("Vet failed")
		return err
	}
	success("Vet complete")
	return nil
}

// Verify runs fmt, vet, generate, and manifests.
func Verify() error {
	step("Verifying API contracts...")
	if err := Fmt(); err != nil {
		return err
	}
	if err := Vet(); err != nil {
		return err
	}
	if err := Generate(); err != nil {
		return err
	}
	return Manifests()
}

// Bundle assembles the OCI contract bundle: CRDs, optional OpenAPI, metadata.
func Bundle() error {
	if err := Manifests(); err != nil {
		return err
	}
	step("Building contract bundle...")
	if err := os.MkdirAll("bundle", 0755); err != nil {
		return err
	}
	if err := bundleCRDs("config/crd/bases", "bundle/crds.yaml"); err != nil {
		fail("Bundle failed")
		return err
	}
	for _, f := range []string{"openapi.yaml", "openapi.json"} {
		if err := copyIfExists(f, filepath.Join("bundle", f)); err != nil {
			fail("Bundle failed")
			return err
		}
	}
	if err := writeMetadata("bundle/metadata.yaml"); err != nil {
		fail("Bundle failed")
		return err
	}
	success("Bundle complete")
	return nil
}

// Publish builds the bundle and pushes it to the OCI registry via ORAS.
func Publish() error {
	if err := Bundle(); err != nil {
		return err
	}
	ref, err := contractRef()
	if err != nil {
		fail("Publish failed")
		return err
	}
	step(fmt.Sprintf("Publishing %s...", ref))
	err = sh.RunV("oras", "push", ref,
		"bundle/crds.yaml:application/yaml",
		"bundle/metadata.yaml:application/yaml",
	)
	if err != nil {
		fail("Publish failed")
		return err
	}
	success(fmt.Sprintf("Published %s", ref))
	return nil
}

func bundleCRDs(srcDir, dst string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("reading %s: %w", srcDir, err)
	}
	var names []string
	for _, e := range entries {
		if !e.IsDir() && (strings.HasSuffix(e.Name(), ".yaml") || strings.HasSuffix(e.Name(), ".yml")) {
			names = append(names, e.Name())
		}
	}
	if len(names) == 0 {
		return fmt.Errorf("no CRD manifests found in %s", srcDir)
	}
	sort.Strings(names)
	var out strings.Builder
	for _, name := range names {
		data, err := os.ReadFile(filepath.Join(srcDir, name))
		if err != nil {
			return err
		}
		doc := strings.TrimRight(string(data), "\n")
		if !strings.HasPrefix(doc, "---") {
			out.WriteString("---\n")
		}
		out.WriteString(doc)
		out.WriteString("\n")
	}
	return os.WriteFile(dst, []byte(out.String()), 0644)
}

func copyIfExists(src, dst string) error {
	data, err := os.ReadFile(src)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

func writeMetadata(dst string) error {
	meta := fmt.Sprintf(
		"version: %s\nrepository: %s\npublishedAt: %s\n",
		version(),
		envOr("GITHUB_REPOSITORY", "ntlaletsi70/blanketops-environments-contract"),
		time.Now().UTC().Format(time.RFC3339),
	)
	return os.WriteFile(dst, []byte(meta), 0644)
}

func contractRef() (string, error) {
	v := version()
	if v == "" {
		return "", fmt.Errorf("no version: set VERSION or GITHUB_REF_NAME, or tag the commit")
	}
	registry := envOr("CONTRACT_REGISTRY", "ghcr.io")
	owner := envOr("GITHUB_REPOSITORY_OWNER", "ntlaletsi70")
	name := envOr("CONTRACT_PACKAGE", "api-contract")
	return fmt.Sprintf("%s/%s/%s:%s", registry, owner, name, v), nil
}

func version() string {
	if v := os.Getenv("VERSION"); v != "" {
		return v
	}
	if v := os.Getenv("GITHUB_REF_NAME"); v != "" {
		return v
	}
	if v, err := sh.Output("git", "describe", "--tags", "--abbrev=0"); err == nil {
		return strings.TrimSpace(v)
	}
	return ""
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func installControllerGen() error {
	const version = "v0.20.0"
	if _, err := os.Stat("bin/controller-gen"); err == nil {
		return nil
	}
	if err := os.MkdirAll("bin", 0755); err != nil {
		return err
	}
	step(fmt.Sprintf("Installing controller-gen %s...", version))
	wd, _ := os.Getwd()
	cmd := exec.Command("go", "install", "sigs.k8s.io/controller-tools/cmd/controller-gen@"+version)
	cmd.Env = append(os.Environ(), "GOBIN="+wd+"/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
