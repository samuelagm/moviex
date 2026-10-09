//go:build integration

package main

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestMainStartupEnvironment(t *testing.T) {
	empty := ""
	tests := []struct {
		name  string
		value *string
	}{
		{
			name: "unset",
		},
		{
			name:  "empty",
			value: &empty,
		},
	}

	buildContext, cancelBuild := context.WithTimeout(context.Background(), 90*time.Second)
	tempDir := t.TempDir()
	binaryPath := filepath.Join(tempDir, "moviex")
	build := exec.CommandContext(buildContext, "go", "build", "-o", binaryPath, ".")
	build.Env = sanitizedEnvironment(nil)
	buildOutput, buildErr := build.CombinedOutput()
	buildContextErr := buildContext.Err()
	cancelBuild()
	if errors.Is(buildContextErr, context.DeadlineExceeded) {
		t.Fatalf("building root binary exceeded deadline: %v", buildContextErr)
	}
	if buildErr != nil {
		t.Fatalf("building root binary: %v\n%s", buildErr, buildOutput)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			cmd := exec.CommandContext(ctx, binaryPath)
			cmd.Env = sanitizedEnvironment(test.value)
			output, err := cmd.CombinedOutput()
			if ctx.Err() != nil {
				t.Fatalf("root binary exceeded deadline: %v", ctx.Err())
			}

			var exitErr *exec.ExitError
			if !errors.As(err, &exitErr) {
				t.Fatalf("root binary error = %v, want exit status 1; output: %s", err, output)
			}
			if got := exitErr.ExitCode(); got != 1 {
				t.Fatalf("root binary exit status = %d, want 1; output: %s", got, output)
			}
			if got, want := string(output), "MOVIEX_ENV_PROBE is required\n"; got != want {
				t.Fatalf("root binary output = %q, want %q", got, want)
			}
		})
	}
}

func sanitizedEnvironment(probeValue *string) []string {
	environment := make([]string, 0, len(os.Environ())+1)
	for _, entry := range os.Environ() {
		key, _, _ := strings.Cut(entry, "=")
		if key == "MOVIEX_ENV_PROBE" || key == "ENV_1" || key == "SEC_1" {
			continue
		}
		environment = append(environment, entry)
	}
	if probeValue != nil {
		environment = append(environment, "MOVIEX_ENV_PROBE="+*probeValue)
	}
	return environment
}
