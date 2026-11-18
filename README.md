# bare-metal-cloud/client-go

Go SDK client for the Bare Metal Cloud API. Generated from the platform's
OpenAPI specification, not written by hand — every endpoint, request body and
response model in this repository mirrors the API that the control plane
actually serves.

Two audiences:

- **External developers** building Go integrations against the BMC platform
  API.
- **BMC's own backend**, which consumes the same client so internal callers
  and third-party integrations share one typed surface.

Status: generated from `bare-metal-cloud/backend` at the `dev-release`
integration branch. The client regenerates automatically on every green merge
to `dev-release` (within the hour at most), and each regeneration merge tags the
next patch release, which publishes straight to the Go module proxy
(proxy.golang.org / pkg.go.dev).

## Install

```bash
go get github.com/bare-metal-cloud/client-go@v0.1.0
```

## Quick start

```go
package main

import (
	"context"
	"fmt"

	bmc "github.com/bare-metal-cloud/client-go"
)

func main() {
	cfg := bmc.NewConfiguration()
	cfg.Servers = bmc.ServerConfigurations{
		{URL: "https://api.example.com", Description: "BMC control plane"},
	}
	client := bmc.NewAPIClient(cfg)

	// Authenticate per request with a platform API token.
	auth := context.WithValue(context.Background(),
		bmc.ContextAccessToken, "YOUR_API_TOKEN")

	// Example call — see docs/ for the full endpoint list.
	// resp, r, err := client.DefaultAPI.SomeEndpoint(auth).Execute()
	_ = auth
	_ = client
	fmt.Println("see docs/ for available endpoints")
}
```

## How this repository works

- `spec/swagger.json` is the normalized snapshot of the API (Swagger 2.0)
  that the current generated code was built from. It is **not edited by
  hand** here; regeneration overwrites it.
- `spec/source-state.json` records the backend `dev-release` commit the client
  currently mirrors (`source_sha`), and the hash of the spec at that commit.
- `scripts/generate.sh` normalizes the snapshot and regenerates the client
  using `openapi-generator` in Docker (no local Java needed).
- `.github/workflows/generate.yml` keeps the client in lockstep with the
  backend API. Every hour it checks the `dev-release` head. A
  `repository_dispatch` of type `api-spec-updated` carrying a commit `sha`
  also triggers it; the backend sends that once its notify change for
  `dev-release` is merged (BAR-2638), and until then the hourly check alone
  picks up merges. It fetches the spec **at that commit**, compares its hash
  with `spec/source-state.json`, and when it moved
  regenerates and opens a pull request from `chore/regenerate`.
- `.github/workflows/release.yml` tags the next patch release (`v*`) on the
  merge commit of a `chore/regenerate` pull request, so the Go module proxy
  picks it up without a manual step. If tag protection refuses the push, cut
  the tag by hand.
- Generated code is committed. Consumers get a reviewable release PR per
  API change, and git history documents every API delta.
- `README.md` and the CI workflows are hand-maintained; the generator is
  told to leave them alone via `.openapi-generator-ignore`.

## Regenerating locally

```bash
scripts/generate.sh
git diff   # review the API delta
```

`scripts/generate.sh` requires Docker and regenerates from the committed
snapshot. CI uses the same generator image, so local and CI output are
identical.

## Setup for the regeneration workflow

One repository secret, `BMC_BACKEND_SPEC_TOKEN`: a fine-grained personal
access token with **Contents: read** on `bare-metal-cloud/backend`. Without
it the watch job skips (it cannot read the private backend repo).

## Versioning

Go modules version by git tag. Merging a regeneration PR tags the next patch
release automatically (`release.yml`). To cut a minor or major version by
hand:

```bash
git tag -a v0.2.0 <merge-sha> -m "..." && git push origin v0.2.0
```

## Visibility

Public since 2026-10-11 — publishing to the Go module proxy (and npm, on
the TypeScript side) requires public content. The spec covers the
**entire** API surface, including internal and admin endpoints; filtering
the spec down to customer-facing endpoints remains an open option for a
future pass (the generation pipeline would drop non-public paths).

## Known spec quirks handled by the pipeline

`scripts/normalize-spec.py` papers over three swag (backend annotation
generator) defects; each becomes a no-op once the backend annotations are
fixed upstream:

1. Eight operations ship no `responses` section — backfilled as untyped 200.
2. The k8s-watch wildcard route declares `{proxyPath}` without defining it —
   backfilled as a required string path parameter.
3. No operation carries an `operationId` — all are assigned deterministic
   names (`postOrgsOrg_idVaultRotatePassword`-style), and the two operations
   tagged `vault` **and** `org` are trimmed to their first tag so types are
   not emitted twice.

## License

Apache-2.0 (matches the org's convention for developer-facing repositories).
