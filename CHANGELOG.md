# Changelog

Written for the product owner and customers.

## Unreleased

## v0.2.5 (2026-11-26)

- The Go SDK's backup and alert builders now accept a personal access token: backup status, backup runs, backup plan, list alerts, acknowledge alert and resolve alert send the PAT alongside the cookie session. Try it: create a PAT, set it on the client, and read a cluster's backup status or acknowledge and resolve a missed-backup alert from a terminal with no browser session.
- `GET /alerts` returns a typed response envelope, and the SDK picks up the backend's storage-overview response, agent re-enroll command response and usage-measurement update models.

## v0.2.4 (2026-11-26)

- The Go SDK now builds the backend's alert types (alert events, alert rules, storage alert status), app restore/clone and dump export/restore request types, the durability section and stack-slot onboarding/review payloads, and accepts `mysql` as a managed-database type. Try it: build a restore or dump-export request with the new builders and send it as before.

## v0.2.3 (2026-11-24)

- The Go SDK now covers the managed-database query API from the backend: send a database query and read its result rows and outcome. Try it: call the databases query operation on a database id and read the response and outcome fields (new `DatabasequeryQueryResponse` and `DatabasequeryOutcome` types).

## v0.2.2 (2026-11-22)

- The Go SDK can now fetch a managed database's deployment status and its installed Postgres extensions — the two operations that were missing from the SDK. Try it: call `GetServicesDatabasesIdStatus` or `GetServicesDatabasesIdExtensions` on a database id and read the deployment status or extension list.
- The Go SDK now follows the backend's dev-release branch within about an hour, and each regeneration is tagged as a new version automatically. Try it: after the next regeneration PR from chore/regenerate merges, open github.com/bare-metal-cloud/client-go/tags and check for the new patch tag.
