//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

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
