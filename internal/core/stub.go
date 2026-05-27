//go:build nocoretools

// Package core provides the built-in core tools for KIT's coding agent.
// This stub is compiled when the nocoretools build tag is set. All tool
// constructors return nil and all bundle functions return empty slices so
// the rest of the codebase compiles unchanged. The truncation helpers
// (truncate.go) are always available regardless of this tag.
package core

import (
	"context"
	"time"

	"charm.land/fantasy"
)

// ---------------------------------------------------------------------------
// Tool configuration (mirrors tools.go)
// ---------------------------------------------------------------------------

// ToolOption configures tool behavior.
type ToolOption func(*ToolConfig)

// ToolConfig holds configuration for tool construction.
type ToolConfig struct {
	WorkDir string
}

// WithWorkDir sets the working directory for file-based tools.
// If empty, os.Getwd() is used at execution time.
func WithWorkDir(dir string) ToolOption {
	return func(c *ToolConfig) {
		c.WorkDir = dir
	}
}

// ApplyOptions applies the given ToolOptions to a ToolConfig and returns it.
func ApplyOptions(opts []ToolOption) ToolConfig {
	var cfg ToolConfig
	for _, o := range opts {
		o(&cfg)
	}
	return cfg
}

// ---------------------------------------------------------------------------
// Tool bundle stubs (mirrors tools.go)
// ---------------------------------------------------------------------------

// CodingTools returns an empty slice when built with nocoretools.
func CodingTools(_ ...ToolOption) []fantasy.AgentTool { return nil }

// ReadOnlyTools returns an empty slice when built with nocoretools.
func ReadOnlyTools(_ ...ToolOption) []fantasy.AgentTool { return nil }

// SubagentTools returns an empty slice when built with nocoretools.
func SubagentTools(_ ...ToolOption) []fantasy.AgentTool { return nil }

// AllTools returns an empty slice when built with nocoretools.
func AllTools(_ ...ToolOption) []fantasy.AgentTool { return nil }

// ---------------------------------------------------------------------------
// Individual tool constructor stubs (mirrors bash.go, read.go, etc.)
// ---------------------------------------------------------------------------

// NewBashTool returns nil when built with nocoretools.
func NewBashTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// NewReadTool returns nil when built with nocoretools.
func NewReadTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// NewWriteTool returns nil when built with nocoretools.
func NewWriteTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// NewEditTool returns nil when built with nocoretools.
func NewEditTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// NewGrepTool returns nil when built with nocoretools.
func NewGrepTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// NewFindTool returns nil when built with nocoretools.
func NewFindTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// NewLsTool returns nil when built with nocoretools.
func NewLsTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// NewSubagentTool returns nil when built with nocoretools.
func NewSubagentTool(_ ...ToolOption) fantasy.AgentTool { return nil }

// ---------------------------------------------------------------------------
// bash.go exported types and context helpers
// ---------------------------------------------------------------------------

// ToolOutputCallback is the signature for streaming tool output.
type ToolOutputCallback func(toolCallID, toolName, chunk string, isStderr bool)

// PasswordPromptCallback is the signature for password prompts.
type PasswordPromptCallback func(prompt string) (password string, cancelled bool)

// ContextWithToolOutputCallback returns the context unchanged (no-op stub).
func ContextWithToolOutputCallback(ctx context.Context, _ ToolOutputCallback) context.Context {
	return ctx
}

// ContextWithPasswordPrompt returns the context unchanged (no-op stub).
func ContextWithPasswordPrompt(ctx context.Context, _ PasswordPromptCallback) context.Context {
	return ctx
}

// ---------------------------------------------------------------------------
// subagent.go exported types and context helpers
// ---------------------------------------------------------------------------

// SubagentSpawnResult carries the outcome of an in-process subagent spawn.
type SubagentSpawnResult struct {
	Response     string
	Error        error
	SessionID    string
	InputTokens  int64
	OutputTokens int64
	Elapsed      time.Duration
}

// SubagentSpawnFunc is a callback that spawns an in-process subagent.
type SubagentSpawnFunc func(ctx context.Context, toolCallID, prompt, model, systemPrompt string, timeout time.Duration) (*SubagentSpawnResult, error)

// WithSubagentSpawner returns the context unchanged (no-op stub).
func WithSubagentSpawner(ctx context.Context, _ SubagentSpawnFunc) context.Context {
	return ctx
}
