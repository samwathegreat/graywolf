package configstore

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/chrissnell/graywolf/pkg/callsign"
)

// The first insert must preserve the same empty-override semantics as updates.
// The old N0CALL GORM default silently disabled a freshly enabled digipeater.
func TestDigipeaterFirstSaveInheritsAcrossReopen(t *testing.T) {
	ctx := context.Background()
	path := filepath.Join(t.TempDir(), "config.db")
	s, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if s != nil {
			_ = s.Close()
		}
	}()
	if err := s.UpsertStationConfig(ctx, StationConfig{Callsign: "W1AW-9"}); err != nil {
		t.Fatal(err)
	}
	c := &DigipeaterConfig{Enabled: true, DedupeWindowSeconds: 30}
	if err := s.UpsertDigipeaterConfig(ctx, c); err != nil {
		t.Fatal(err)
	}
	if c.MyCall != "" {
		t.Fatalf("first save changed inherited override to %q", c.MyCall)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
	s = nil
	s, err = Open(path)
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.GetDigipeaterConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil || got.MyCall != "" || !got.Enabled {
		t.Fatalf("reopened config: %+v", got)
	}
	resolved, err := callsign.Resolve(got.MyCall, "W1AW-9")
	if err != nil || resolved != "W1AW-9" {
		t.Fatalf("runtime resolution: %q, %v", resolved, err)
	}
}

func TestDigipeaterOverridePreserved(t *testing.T) {
	for _, override := range []string{"MTNTOP-1", "N0CALL"} {
		t.Run(override, func(t *testing.T) {
			s := newTestStore(t)
			ctx := context.Background()
			if err := s.UpsertStationConfig(ctx, StationConfig{Callsign: "W1AW-9"}); err != nil {
				t.Fatal(err)
			}
			c := &DigipeaterConfig{MyCall: override, DedupeWindowSeconds: 30}
			if err := s.UpsertDigipeaterConfig(ctx, c); err != nil {
				t.Fatal(err)
			}
			if err := s.Migrate(); err != nil {
				t.Fatal(err)
			}
			got, err := s.GetDigipeaterConfig(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if got == nil || got.MyCall != override {
				t.Fatalf("override changed: %+v", got)
			}
		})
	}
}
