package render

import (
	"github.com/Funnelish/liquid/parser"
)

// Config holds configuration information for parsing and rendering.
type Config struct {
	parser.Config
	grammar
	Cache           map[string][]byte
	StrictVariables bool
	TemplateStore   TemplateStore
}

type grammar struct {
	tags      map[string]TagCompiler
	blockDefs map[string]*blockSyntax
}

// NewConfig creates a new Settings.
//
// TemplateStore defaults to DeniedTemplateStore, which refuses every filesystem
// read, so {% include %} cannot read files off the server by default. This is
// deliberate: Liquid templates are frequently authored by untrusted users, and
// a filesystem-backed include is a local-file-disclosure primitive. Callers that
// genuinely need to include templates from disk opt in explicitly with
// Engine.RegisterTemplateStore(&FileTemplateStore{}).
func NewConfig() Config {
	g := grammar{
		tags:      map[string]TagCompiler{},
		blockDefs: map[string]*blockSyntax{},
	}
	return Config{
		Config:        parser.NewConfig(g),
		grammar:       g,
		Cache:         map[string][]byte{},
		TemplateStore: DeniedTemplateStore{},
	}
}
