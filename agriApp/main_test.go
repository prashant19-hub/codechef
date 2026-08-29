package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultDBURL(t *testing.T) {
	if got := defaultDBURL(); got == "" {
		t.Fatal("defaultDBURL should not be empty")
	}
	if got := defaultDBURL(); got != "postgres://postgres:postgres@localhost:5433/agriapp?sslmode=disable" {
		t.Fatalf("unexpected postgres DSN: %q", got)
	}
}

func TestTemplatePath(t *testing.T) {
	got := templatePath()
	if got == "" {
		t.Fatal("template path should not be empty")
	}

	wantSuffix := filepath.Join("template", "index.html")
	if !strings.HasSuffix(got, wantSuffix) {
		t.Fatalf("unexpected template path: got %q want suffix %q", got, wantSuffix)
	}
}
