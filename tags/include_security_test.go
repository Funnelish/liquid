package tags

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/Funnelish/liquid/parser"
	"github.com/Funnelish/liquid/render"
	"github.com/stretchr/testify/require"
)

// TestIncludeTag_default_store_denies_filesystem is the regression guard for the
// local-file-disclosure fix: with the default config, {% include "<path>" %}
// must not read a real file off disk. Only an explicit FileTemplateStore (opt in)
// may do so.
func TestIncludeTag_default_store_denies_filesystem(t *testing.T) {
	secret := filepath.Join(t.TempDir(), "secret.txt")
	require.NoError(t, os.WriteFile(secret, []byte("TOP_SECRET_DSN"), 0o600))

	config := render.NewConfig() // default: DeniedTemplateStore
	AddStandardTags(config)
	loc := parser.SourceLoc{Pathname: "page.html", LineNo: 1}

	root, err := config.Compile(`{% include "`+secret+`" %}`, loc)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	err = render.Render(root, buf, map[string]any{}, config)
	require.NoError(t, err)
	// The file must not be read. The renderer embeds the denial as a hidden
	// placeholder span, so the secret never reaches the output.
	require.NotContains(t, buf.String(), "TOP_SECRET_DSN",
		"default engine leaked a file off disk via {%% include %%}")
	require.Contains(t, buf.String(), "filesystem template reads are disabled")
}

// TestIncludeTag_default_store_ignores_cache guards that Config.Cache is not a
// bypass: with the default (denied) store, {% include %} must not resolve a
// template even when one is present in the cache.
func TestIncludeTag_default_store_ignores_cache(t *testing.T) {
	config := render.NewConfig() // default: DeniedTemplateStore
	config.Cache["secret.html"] = []byte("TOP_SECRET_FROM_CACHE")
	AddStandardTags(config)
	loc := parser.SourceLoc{Pathname: "page.html", LineNo: 1}

	root, err := config.Compile(`{% include "secret.html" %}`, loc)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	err = render.Render(root, buf, map[string]any{}, config)
	require.NoError(t, err)
	require.NotContains(t, buf.String(), "TOP_SECRET_FROM_CACHE",
		"default engine resolved an include from the cache")
}
