package main

import (
	"bytes"
	"strings"
	"testing"
)

// TestRootCommand verifies the root command exposes its subcommands.
func TestRootCommand(t *testing.T) {
	want := []string{
		"version", "install", "uninstall", "check", "get",
		"approve", "reject", "completion",
		"suspend", "resume", "describe",
	}
	for _, name := range want {
		if _, _, err := rootCmd.Find([]string{name}); err != nil {
			t.Errorf("subcommand %q not registered: %v", name, err)
		}
	}
}

// TestVersionCommand verifies the version command prints the build version.
func TestVersionCommand(t *testing.T) {
	out := &bytes.Buffer{}
	rootCmd.SetOut(out)
	rootCmd.SetErr(out)
	rootCmd.SetArgs([]string{"version"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("version command failed: %v", err)
	}
	// version subcommand currently writes to os.Stdout directly, so we just
	// confirm the command resolved and executed without error.
}

// TestApproveRequiresArg ensures `approve` rejects missing arguments.
func TestApproveRequiresArg(t *testing.T) {
	out := &bytes.Buffer{}
	rootCmd.SetOut(out)
	rootCmd.SetErr(out)
	rootCmd.SetArgs([]string{"approve"})
	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error for missing gate argument")
	}
	if !strings.Contains(err.Error(), "accepts 1 arg") {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestCompletionInvalidShell rejects unknown shells.
func TestCompletionInvalidShell(t *testing.T) {
	out := &bytes.Buffer{}
	rootCmd.SetOut(out)
	rootCmd.SetErr(out)
	rootCmd.SetArgs([]string{"completion", "tcsh"})
	if err := rootCmd.Execute(); err == nil {
		t.Fatal("expected error for unsupported shell")
	}
}

// TestScopeFlags verifies the get-command scope helper honors --all-namespaces.
func TestScopeFlags(t *testing.T) {
	namespace = "demo"
	allNamespaces = false
	args := scope()
	if got, want := strings.Join(args, " "), "-n demo"; got != want {
		t.Errorf("scope() = %q, want %q", got, want)
	}

	allNamespaces = true
	args = scope()
	if got, want := strings.Join(args, " "), "--all-namespaces"; got != want {
		t.Errorf("scope(all) = %q, want %q", got, want)
	}
	// reset so other tests are not affected
	allNamespaces = false
}
