package tags

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/Funnelish/liquid/render"
)

func includeTag(source string) (func(io.Writer, render.Context) error, error) {
	return func(w io.Writer, ctx render.Context) error {
		// It might be more efficient to add a context interface to render bytes
		// to a writer. The status quo keeps the interface light at the expense of some overhead
		// here.
		value, err := ctx.EvaluateString(ctx.TagArgs())
		if err != nil {
			return err
		}
		rel, ok := value.(string)
		if !ok {
			return ctx.Errorf("include requires a string argument; got %v", value)
		}
		if !safeIncludePath(rel) {
			return ctx.Errorf("include: invalid template path %q", rel)
		}
		filename := filepath.Join(filepath.Dir(ctx.SourceFile()), rel)
		s, err := ctx.RenderFile(filename, map[string]any{})
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, s)
		return err
	}, nil
}

// safeIncludePath accepts only relative paths that stay below the including
// template: no absolute paths, no volume names, no ".." segments and no NUL
// bytes. The template store applies its own root check on top of this.
func safeIncludePath(rel string) bool {
	if rel == "" || strings.ContainsRune(rel, 0) {
		return false
	}
	if filepath.IsAbs(rel) || filepath.VolumeName(rel) != "" || strings.HasPrefix(rel, "/") || strings.HasPrefix(rel, "\\") {
		return false
	}
	for _, seg := range strings.FieldsFunc(rel, func(r rune) bool { return r == '/' || r == '\\' }) {
		if seg == ".." {
			return false
		}
	}
	return true
}
