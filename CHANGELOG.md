# Changelog

Written for the product owner and customers.

## Unreleased

- The Go SDK now covers the managed-database query API from the backend: send a database query and read its result rows and outcome. Try it: call the databases query operation on a database id and read the response and outcome fields (new `DatabasequeryQueryResponse` and `DatabasequeryOutcome` types).

## v0.2.2 (2026-11-22)

- The Go SDK can now fetch a managed database's deployment status and its installed Postgres extensions — the two operations that were missing from the SDK. Try it: call `GetServicesDatabasesIdStatus` or `GetServicesDatabasesIdExtensions` on a database id and read the deployment status or extension list.
- The Go SDK now follows the backend's dev-release branch within about an hour, and each regeneration is tagged as a new version automatically. Try it: after the next regeneration PR from chore/regenerate merges, open github.com/bare-metal-cloud/client-go/tags and check for the new patch tag.
