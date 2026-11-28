// Package handtest carries hand-written SDK tests that must survive client
// regeneration. scripts/generate.sh wipes the generated layout (root *.go,
// api/, docs/, test/) but never touches this directory, so the proofs written
// for the org S3 connection lifecycle (BAR-6001, FR-Z-5) keep running after
// every automated regeneration.
//
// Every test exercises one lifecycle verb as a PAT caller: the client is
// configured with the BearerAuth scheme only (no cookie), mirroring how a
// personal access token authenticates. The lifecycle verbs under test are the
// ones BAR-5949 landed on the backend: list, versions, rename/repoint (PUT),
// rotate-credentials, and delete with the force query and its 409 consumers
// body. Per the BAR-5949 plan-review rule, the SDK sends no force parameter by
// default and merely surfaces the 409 and its consumers; beyond the API's own
// force query parameter it adds nothing.
package handtest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	openapiclient "github.com/bare-metal-cloud/client-go"
)

const (
	testOrgID  int32 = 42
	testConnID int32 = 7
	testPAT          = "bmc_pat_test_token"
)

// patContext returns a context that authenticates every request with the
// BearerAuth scheme carrying a personal access token, and nothing else.
func patContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, openapiclient.ContextAPIKeys, map[string]openapiclient.APIKey{
		"BearerAuth": {Key: testPAT, Prefix: "Bearer"},
	})
}

// newPATClient starts an httptest backend stub and returns an API client
// pointed at it, plus the recorder the handlers assert against.
func newPATClient(t *testing.T, handler http.HandlerFunc) (*openapiclient.APIClient, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	cfg := openapiclient.NewConfiguration()
	cfg.Servers = openapiclient.ServerConfigurations{{URL: server.URL}}
	return openapiclient.NewAPIClient(cfg), server
}

// assertPATRequest verifies the call went out as a PAT request: an
// Authorization: Bearer header carrying the token and no session cookie.
func assertPATRequest(t *testing.T, r *http.Request) {
	t.Helper()
	auth := r.Header.Get("Authorization")
	if auth != "Bearer "+testPAT {
		t.Errorf("Authorization header = %q, want %q", auth, "Bearer "+testPAT)
	}
	if cookie := r.Header.Get("Cookie"); cookie != "" {
		t.Errorf("Cookie header = %q, want none from a PAT caller", cookie)
	}
}

func TestListOrgS3ConnectionsWithPAT(t *testing.T) {
	client, _ := newPATClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/orgs/42/s3-connections" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		assertPATRequest(t, r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id":7,"name":"corner-storage"}]`))
	})

	resp, httpRes, err := client.OrgS3ConnectionsAPI.
		GetOrgsOrgIdS3Connections(patContext(context.Background()), testOrgID).Execute()
	if err != nil {
		t.Fatalf("list with PAT failed: %v", err)
	}
	if httpRes.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", httpRes.StatusCode)
	}
	if len(resp) != 1 || resp[0].GetName() != "corner-storage" {
		t.Errorf("unexpected list payload: %+v", resp)
	}
}

func TestGetOrgS3ConnectionVersionsWithPAT(t *testing.T) {
	client, _ := newPATClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/orgs/42/s3-connections/7/versions" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		assertPATRequest(t, r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"items":[{"version_number":2,"is_current":true}],"next_cursor":null}`))
	})

	versions, httpRes, err := client.SecretsAPI.
		GetOrgsOrgIdS3ConnectionsIdVersions(patContext(context.Background()), testOrgID, testConnID).Execute()
	if err != nil {
		t.Fatalf("versions with PAT failed: %v", err)
	}
	if httpRes.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", httpRes.StatusCode)
	}
	items, ok := versions["items"].([]interface{})
	if !ok {
		t.Fatalf("versions payload missing items array: %+v", versions)
	}
	if len(items) != 1 {
		t.Errorf("len(items) = %d, want 1 (backend sends {items, next_cursor}): %+v", len(items), versions)
	}
}

func TestRenameAndRepointOrgS3ConnectionWithPAT(t *testing.T) {
	newName := "renamed-connection"
	newEndpoint := "https://s3.us-east.example.com"
	client, _ := newPATClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut || r.URL.Path != "/orgs/42/s3-connections/7" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		assertPATRequest(t, r)
		var body openapiclient.HandlersUpdateOrgS3ConnectionRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("decode request body: %v", err)
			return
		}
		if body.GetName() != newName {
			t.Errorf("name in body = %q, want %q", body.GetName(), newName)
		}
		if body.GetEndpoint() != newEndpoint {
			t.Errorf("endpoint in body = %v, want %q", body.GetEndpoint(), newEndpoint)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":7,"name":"renamed-connection","endpoint":"https://s3.us-east.example.com"}`))
	})

	name, endpoint := newName, newEndpoint
	resp, httpRes, err := client.OrgS3ConnectionsAPI.
		PutOrgsOrgIdS3ConnectionsId(patContext(context.Background()), testOrgID, testConnID).
		Request(openapiclient.HandlersUpdateOrgS3ConnectionRequest{Name: &name, Endpoint: &endpoint}).
		Execute()
	if err != nil {
		t.Fatalf("rename/repoint with PAT failed: %v", err)
	}
	if httpRes.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", httpRes.StatusCode)
	}
	if resp.GetName() != newName {
		t.Errorf("response name = %q, want %q", resp.GetName(), newName)
	}
}

func TestRotateOrgS3ConnectionCredentialsWithPAT(t *testing.T) {
	accessKey := "AKIAEXAMPLE"
	secretKey := "secret-example"
	client, _ := newPATClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/orgs/42/s3-connections/7/rotate-credentials" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		assertPATRequest(t, r)
		var body openapiclient.HandlersRotateOrgS3ConnectionCredentialsRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("decode request body: %v", err)
			return
		}
		if body.GetAccessKeyId() != accessKey || body.GetSecretAccessKey() != secretKey {
			t.Errorf("rotate body = %q/%q, want %q/%q",
				body.GetAccessKeyId(), body.GetSecretAccessKey(), accessKey, secretKey)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":7,"has_vault_credentials":true}`))
	})

	resp, httpRes, err := client.OrgS3ConnectionsAPI.
		PostOrgsOrgIdS3ConnectionsIdRotateCredentials(patContext(context.Background()), testOrgID, testConnID).
		Request(openapiclient.HandlersRotateOrgS3ConnectionCredentialsRequest{AccessKeyId: &accessKey, SecretAccessKey: &secretKey}).
		Execute()
	if err != nil {
		t.Fatalf("rotate with PAT failed: %v", err)
	}
	if httpRes.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", httpRes.StatusCode)
	}
	if resp.GetHasVaultCredentials() != true {
		t.Errorf("has_vault_credentials = %v, want true", resp.GetHasVaultCredentials())
	}
}

// TestDeleteOrgS3ConnectionSurfaces409Consumers pins the BAR-5949 plan-review
// rule on the client: a plain delete carries NO force query parameter, and a
// 409 from the secret-reference guard surfaces to the caller with the
// consumers list intact — the SDK does not swallow it and adds nothing beyond
// the API's own force query.
func TestDeleteOrgS3ConnectionSurfaces409Consumers(t *testing.T) {
	client, _ := newPATClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/orgs/42/s3-connections/7" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		assertPATRequest(t, r)
		if force := r.URL.Query().Get("force"); force != "" {
			t.Errorf("force query = %q on a plain delete, want absent", force)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error":"s3 connection still in use","consumers":[{"id":3,"name":"nightly-backups","type":"backup_plan","scope":"org","link":"/api/v1/backup-plans/3"},{"id":9,"name":"pg-orders","type":"managed_database","scope":"org","link":"/api/v1/databases/9"}]}`))
	})

	_, httpRes, err := client.OrgS3ConnectionsAPI.
		DeleteOrgsOrgIdS3ConnectionsId(patContext(context.Background()), testOrgID, testConnID).Execute()
	if err == nil {
		t.Fatal("plain delete against a guarded connection must fail, got nil error")
	}
	if httpRes == nil || httpRes.StatusCode != http.StatusConflict {
		t.Fatalf("status = %v, want 409", httpRes)
	}
	apiErr, ok := err.(*openapiclient.GenericOpenAPIError)
	if !ok {
		t.Fatalf("error type %T, want *GenericOpenAPIError", err)
	}
	model, ok := apiErr.Model().(map[string]interface{})
	if !ok {
		t.Fatalf("409 model type %T, want map[string]interface{}", apiErr.Model())
	}
	consumers, ok := model["consumers"].([]interface{})
	if !ok {
		t.Fatalf("409 model missing consumers array: %+v", model)
	}
	if len(consumers) != 2 {
		t.Errorf("len(consumers) = %d, want 2: %+v", len(consumers), consumers)
	}
	first, ok := consumers[0].(map[string]interface{})
	if !ok {
		t.Fatalf("consumer[0] type %T, want object", consumers[0])
	}
	for _, key := range []string{"id", "name", "type", "scope", "link"} {
		if _, present := first[key]; !present {
			t.Errorf("consumer[0] missing %q key (S3ConnectionConsumer shape): %+v", key, first)
		}
	}
	if first["type"] != "backup_plan" {
		t.Errorf("consumer[0].type = %v, want backup_plan", first["type"])
	}
}

func TestForceDeleteOrgS3ConnectionWithPAT(t *testing.T) {
	client, _ := newPATClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete || r.URL.Path != "/orgs/42/s3-connections/7" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		assertPATRequest(t, r)
		if force := r.URL.Query().Get("force"); force != "true" {
			t.Errorf("force query = %q, want \"true\"", force)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	_, httpRes, err := client.OrgS3ConnectionsAPI.
		DeleteOrgsOrgIdS3ConnectionsId(patContext(context.Background()), testOrgID, testConnID).
		Force(true).Execute()
	if err != nil {
		t.Fatalf("force delete with PAT failed: %v", err)
	}
	if httpRes.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", httpRes.StatusCode)
	}
}
