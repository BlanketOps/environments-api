# BlanketOps Environments API

## Overview

`BlanketOps Environments API` defines the **canonical Kubernetes API contracts**
for BlanketOps Environments.

This repository contains **versioned Custom Resource Definitions (CRDs)** and
their corresponding Go types. It is the **source of truth** for all environment,
build, event, and source APIs used by BlanketOps components.

---

## What This Repository Is

This repository:

- defines Kubernetes API schemas
- owns CRD generation
- provides Go types for consumers
- enforces semantic versioning of contracts

It is a **contract repository**, not a runtime.

---

## What This Repository Is Not

This repository does **not**:

- run controllers
- build container images
- deploy to clusters
- define RBAC, managers, or webhooks
- contain reconciliation logic

Those concerns live in downstream projects.

---

## API Groups and Versions

Currently supported API groups:

- `environments.blanketops.dev/v1alpha1`
- `events.blanketops.dev/v1alpha1`
- `networks.blanketops.dev/v1alpha1`
- `sources.blanketops.dev/v1alpha1`

All APIs are currently **alpha** and may change in backward-incompatible ways
until stabilized.

---

## Versioning Policy

This repository follows **semantic versioning at the contract level**.

- `v1alpha1` — experimental, breaking changes allowed
- `v1beta1` — feature complete, limited breaking changes
- `v1` — stable, backward-compatible

Once an API version is stabilized, it is **append-only**.

Breaking changes require a new API version.

---

## Code Generation

CRDs and deepcopy code are generated using `controller-gen`, invoked via
[mage](https://magefile.org) targets (this repo has no Makefile).

To generate all artifacts:

```sh
mage generate
mage manifests
```
