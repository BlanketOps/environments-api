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

### New Kinds Follow the Existing Envelope Pattern, Not Raw Kubebuilder Scaffolding
Every Kind is a fixed `Contract runtime.RawExtension` envelope (see API
Design below), not a conventionally-typed kubebuilder resource. Copy an
existing file (e.g. `domain.go`) rather than trust `kubebuilder create api`'s
default scaffold output, which produces typed fields that don't match this
repo's pattern.

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

## API Design

Every Kind in this repo is a thin, uniform **envelope around an opaque
contract**, not a conventionally-typed kubebuilder resource. Kubernetes
stores and round-trips the envelope; it does not understand or validate the
`Contract` payload itself — that's owned by the contract layer downstream.
All nine existing Kinds (`Environment`, `Build`, `Deployment`, `ServiceUnit`,
`Package`, `Route`, `Domain`, `GitHubEvent`, `GitRepository`) follow this
exact shape, e.g. `api/networks/v1alpha1/domain.go`:

```go
// +k8s:openapi-gen=true
package v1alpha1

type <Kind>Spec struct {
	// Contract is the canonical BlanketOps <Kind> specification.
	// Opaque to Kubernetes; preserved verbatim by the API server.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:pruning:PreserveUnknownFields
	Contract runtime.RawExtension `json:"contract"`
}

type <Kind>Status struct {
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Contract runtime.RawExtension `json:"contract,omitempty"`

	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type <Kind> struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   <Kind>Spec   `json:"spec,omitempty"`
	Status <Kind>Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type <Kind>List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []<Kind> `json:"items"`
}

func init() {
	SchemeBuilder.Register(&<Kind>{}, &<Kind>List{})
}
```

**Rules that follow from this:**
- Do NOT add typed fields to `Spec`/`Status` with per-field validation
  markers (`Minimum`, `MaxLength`, `Pattern`, `default`, etc.) — there is
  exactly one field, `Contract runtime.RawExtension`, marked `Required` +
  `PreserveUnknownFields` on Spec, and `optional` + `PreserveUnknownFields`
  on Status. Field-level schema validation belongs to the contract layer,
  not this envelope.
- `Status.Conditions []metav1.Condition` is the only other status field;
  it's what `kubectl`/UIs read, not `Contract`.
- Every file starts with a `// +k8s:openapi-gen=true` marker (added by
  `hack/generate-api-types.sh`, not scaffolded by kubebuilder — see below).

### Adding a Kind to an existing group
Reuse the group's existing version directory (`api/<group>/v1alpha1`); do
not create a new version unless intentionally starting a new API version.
`kubebuilder create api` scaffolds a *conventional* Spec/Status (typed
fields, no `Contract` envelope) — do not keep its default output. Instead:
1. Copy an existing file in the target group (e.g. `domain.go`) to
   `<kind>.go` and rename `Domain` → `<Kind>` throughout.
2. Add the new file's path to the list in `hack/generate-api-types.sh` so
   it gets the `+k8s:openapi-gen=true` marker, then run that script.
3. Add the new Kind's entry to `PROJECT` under the correct `group`/`path`.
4. Add the corresponding block to
   `pkg/generated/openapi/generated.openapi.go` (`<Kind>`, `<Kind>List`,
   `<Kind>Spec`, `<Kind>Status`) — this file is hand-maintained, not
   generated by a tool, despite the header comment.
5. Run `mage generate manifests` and `gofmt -w` the OpenAPI file.

### Adding a new group
Same as above, plus a new `api/<group>/v1alpha1/groupversion_info.go`
declaring `+groupName=<group>.blanketops.dev` (see
`api/networks/v1alpha1/groupversion_info.go` for the pattern).

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
