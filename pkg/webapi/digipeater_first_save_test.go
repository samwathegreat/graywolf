package webapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chrissnell/graywolf/pkg/configstore"
	"github.com/chrissnell/graywolf/pkg/webapi/dto"
)

func TestDigipeaterFirstSaveInherit(t *testing.T) {
	for _, body := range []string{
		`{"enabled":true,"dedupe_window_seconds":30,"my_call":""}`,
		`{"enabled":true,"dedupe_window_seconds":30}`,
		`{"enabled":false,"dedupe_window_seconds":30,"my_call":""}`,
	} {
		t.Run(body, func(t *testing.T) {
			srv, _ := newTestServer(t)
			ctx := context.Background()
			if err := srv.store.UpsertStationConfig(ctx, configstore.StationConfig{Callsign: "W1AW-9"}); err != nil {
				t.Fatal(err)
			}
			mux := http.NewServeMux()
			srv.RegisterRoutes(mux)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/api/digipeater", strings.NewReader(body)))
			if rec.Code != http.StatusOK {
				t.Fatalf("PUT: %d %s", rec.Code, rec.Body.String())
			}
			var response dto.DigipeaterConfigResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if response.MyCall != "" {
				t.Fatalf("PUT changed inherited override to %q", response.MyCall)
			}
			get := httptest.NewRecorder()
			mux.ServeHTTP(get, httptest.NewRequest(http.MethodGet, "/api/digipeater", nil))
			if get.Code != http.StatusOK {
				t.Fatalf("GET: %d %s", get.Code, get.Body.String())
			}
			if err := json.Unmarshal(get.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if response.MyCall != "" {
				t.Fatalf("GET changed inherited override to %q", response.MyCall)
			}
		})
	}
}

func TestDigipeaterRejectsEnabledPlaceholderOverride(t *testing.T) {
	for _, override := range []string{"N0CALL", "n0call-7"} {
		t.Run(override, func(t *testing.T) {
			srv, _ := newTestServer(t)
			ctx := context.Background()
			if err := srv.store.UpsertStationConfig(ctx, configstore.StationConfig{Callsign: "W1AW-9"}); err != nil {
				t.Fatal(err)
			}
			mux := http.NewServeMux()
			srv.RegisterRoutes(mux)
			body := `{"enabled":true,"dedupe_window_seconds":30,"my_call":"` + override + `"}`
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/api/digipeater", strings.NewReader(body)))
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("PUT: %d %s", rec.Code, rec.Body.String())
			}
			cfg, err := srv.store.GetDigipeaterConfig(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if cfg != nil {
				t.Fatalf("rejected request changed config: %+v", cfg)
			}
		})
	}
}

func TestDigipeaterStoredPlaceholderCanBeClearedExplicitly(t *testing.T) {
	srv, _ := newTestServer(t)
	ctx := context.Background()
	if err := srv.store.UpsertStationConfig(ctx, configstore.StationConfig{Callsign: "W1AW-9"}); err != nil {
		t.Fatal(err)
	}
	if err := srv.store.UpsertDigipeaterConfig(ctx, &configstore.DigipeaterConfig{MyCall: "N0CALL", DedupeWindowSeconds: 30}); err != nil {
		t.Fatal(err)
	}
	mux := http.NewServeMux()
	srv.RegisterRoutes(mux)
	for _, step := range []struct {
		body   string
		status int
	}{
		{`{"enabled":true,"dedupe_window_seconds":30}`, http.StatusBadRequest},
		{`{"enabled":true,"dedupe_window_seconds":30,"my_call":""}`, http.StatusOK},
	} {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(http.MethodPut, "/api/digipeater", strings.NewReader(step.body)))
		if rec.Code != step.status {
			t.Fatalf("PUT %s: %d %s", step.body, rec.Code, rec.Body.String())
		}
	}
	got, err := srv.store.GetDigipeaterConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil || !got.Enabled || got.MyCall != "" {
		t.Fatalf("cleared config: %+v", got)
	}
}
