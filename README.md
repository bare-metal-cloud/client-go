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

Status: early scaffold, generated from `bare-metal-cloud/backend` at tag
**v1.2.0**. The client regenerates automatically whenever the backend
publishes a new release tag; the release process (module tags, GOPROXY
visibility) is still being decided.

## Install

```bash
go get github.com/bare-metal-cloud/client-go
```

(Once the first version tag is cut; until then, pin a commit.)

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
- `spec/source-state.json` records which backend release tag the client
  currently mirrors, and the hash of the spec at that tag.
- `scripts/generate.sh` normalizes the snapshot and regenerates the client
  using `openapi-generator` in Docker (no local Java needed).
- `.github/workflows/generate.yml` keeps the client in lockstep with
  released backend API versions. Every hour (and on any
  `repository_dispatch` of type `api-spec-updated`, e.g. the backend's own
  notify workflow pinging at tag time) it compares the latest `v*` tag on
  `bare-metal-cloud/backend` against `spec/source-state.json`; when the tag
  or its spec moved, it fetches the spec **at that tag**, regenerates, and
  opens a pull request. The tag contract means the SDK always mirrors a
  released backend, not a moving branch.
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
it the tag-watch job skips (it cannot read the private backend repo).

## Versioning

Go modules version by git tag. Cut a release after merging a regeneration
PR:

```bash
git tag v0.1.0 && git push origin v0.1.0
```

## Visibility

Private for now (ruled 2026-10-11). The spec covers the **entire** API
surface, including internal and admin endpoints — if this repository ever
goes public, first decide whether to filter the spec down to
customer-facing endpoints (the generation pipeline would drop non-public
paths). Flipping visibility is a one-command change when that day comes.

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
