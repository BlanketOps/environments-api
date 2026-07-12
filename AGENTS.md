# blanketops-environments-api - AI Agent Guide

## What This Repository Is

This is a **contract repository**, not a runtime. It defines the canonical
Kubernetes API schemas (CRDs + Go types) for BlanketOps Environments. It does
**not** contain controllers, webhooks, a manager, RBAC, or deployment
manifests — those live in downstream projects that consume this API.

## Project Structure

```
api/<group>/<version>/*_types... .go   CRD schemas (+kubebuilder markers), one file per Kind
api/<group>/<version>/zz_generated.*   Auto-generated deepcopy (DO NOT EDIT)
api/<group>/<version>/groupversion_info.go  Group/version registration
config/crd/bases/*                     Generated CRDs (DO NOT EDIT)
pkg/generated/openapi/                 Generated OpenAPI definitions (DO NOT EDIT)
hack/boilerplate.go.txt                License header used by codegen
magefile.go                            Build/generate/bundle/publish commands (mage, not make)
PROJECT                                Kubebuilder metadata (DO NOT EDIT)
```

This project uses the **multigroup layout**: `api/<group>/<version>/`.

### Current API groups and kinds

| Group | Version | Kinds |
|---|---|---|
| `environments.blanketops.dev` | v1alpha1 | Environment, Build, Deployment, ServiceUnit, Package |
| `networks.blanketops.dev` | v1alpha1 | Route, Domain |
| `events.blanketops.dev` | v1alpha1 | GitHubEvent |
| `sources.blanketops.dev` | v1alpha1 | GitRepository |

All APIs are currently `v1alpha1` and may change in backward-incompatible ways.

## Critical Rules

### Never Edit These (Auto-Generated)
- `config/crd/bases/*.yaml` - from `mage manifests`
- `**/zz_generated.*.go` - from `mage generate`
- `pkg/generated/openapi/generated.openapi.go` - from OpenAPI generation
- `PROJECT` - from `kubebuilder [OPTIONS]` (edit by hand only for a Kind rename/move, matching the actual `api/` layout, as done here)

### Keep Project Structure
Do not move files around outside of a deliberate API group change. The CLI
expects files in specific locations (`api/<group>/<version>/`).

### Always Use CLI Commands to Scaffold
Use `kubebuilder create api` (with `--controller=false` since this repo has
no controllers) to scaffold new Kinds. Do NOT create type files manually.

### No Makefile
This repo does not have a `Makefile`. All build/codegen/release tasks are
[mage](https://magefile.org) targets defined in `magefile.go`. Run `go install
github.com/magefile/mage@latest` once, then `mage <target>`.

### No Controllers, Webhooks, or Manager
Do not add `internal/controller`, `internal/webhook`, `cmd/main.go`,
`config/rbac`, or `config/samples` here — this repo only owns API types and
generated CRDs/OpenAPI. Reconciliation and deployment concerns belong in
downstream operator repos.

## After Making Changes

**After editing `*_types.go` or markers:**
```
mage generate    # Regenerate DeepCopy methods
mage manifests   # Regenerate CRDs from markers
```

**Before committing:**
```
mage verify      # fmt + vet + generate + manifests
```

## Mage Targets Cheat Sheet

```
mage generate    # controller-gen deepcopy generation
mage manifests   # controller-gen CRD generation -> config/crd/bases
mage fmt         # go fmt ./...
mage vet         # go vet ./...
mage verify      # fmt, vet, generate, manifests (what CI runs)
mage bundle      # assemble OCI contract bundle (CRDs + OpenAPI + metadata) -> bundle/
mage publish     # bundle, then push to OCI registry via oras (needs VERSION/tag + registry env)
```

`controller-gen` is installed automatically into `bin/controller-gen` by the
mage targets that need it.

## CLI Commands Cheat Sheet

### Create a new API type (no controller, this is a types-only repo)
```bash
kubebuilder create api --group <group> --version v1alpha1 --kind <Kind> --controller=false
```

### Adding a Kind to an existing group
Reuse the group's existing version directory (`api/<group>/v1alpha1`); do not
create a new version unless intentionally starting a new API version.

## API Design

**Key markers for** `api/<group>/<version>/*.go`:

```go
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].status"

// On fields:
// +kubebuilder:validation:Required
// +kubebuilder:validation:Minimum=1
// +kubebuilder:validation:MaxLength=100
// +kubebuilder:validation:Pattern="^[a-z]+$"
// +kubebuilder:default="value"
```

- **Use** `metav1.Condition` for status (not custom string fields)
- **Use predefined types**: `metav1.Time` instead of `string` for dates
- **Follow K8s API conventions**: Standard field names (`spec`, `status`, `metadata`)

## Versioning Policy

Semantic versioning at the contract level:
- `v1alpha1` — experimental, breaking changes allowed
- `v1beta1` — feature complete, limited breaking changes
- `v1` — stable, backward-compatible, append-only

Breaking changes require a new API version, not an in-place change to a
stabilized version.

## References

### Essential Reading
- **Kubebuilder Book**: https://book.kubebuilder.io (comprehensive guide)
- **Markers Reference**: https://book.kubebuilder.io/reference/markers.html
- **Good Practices**: https://book.kubebuilder.io/reference/good-practices.html
- **API Conventions**: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md
- **Mage**: https://magefile.org

### Tools & Libraries
- **controller-tools** (controller-gen): https://github.com/kubernetes-sigs/controller-tools
- **Kubebuilder Repo**: https://github.com/kubernetes-sigs/kubebuilder
