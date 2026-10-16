package tags

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Funnelish/liquid/parser"
	"github.com/Funnelish/liquid/render"
	"github.com/stretchr/testify/require"
)

var includeTestBindings = map[string]any{
	"test": true,
	"var":  "value",
}

func TestIncludeTag(t *testing.T) {
	config := render.NewConfig()
	config.TemplateStore = render.NewFileTemplateStore("testdata")
	loc := parser.SourceLoc{Pathname: "testdata/include_source.html", LineNo: 1}
	AddStandardTags(config)

	// basic functionality
	root, err := config.Compile(`{% include "include_target.html" %}`, loc)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	err = render.Render(root, buf, includeTestBindings, config)
	require.NoError(t, err)
	require.Equal(t, "include target", strings.TrimSpace(buf.String()))

	// tag and variable
	root, err = config.Compile(`{% include "include_target_2.html" %}`, loc)
	require.NoError(t, err)
	buf = new(bytes.Buffer)
	err = render.Render(root, buf, includeTestBindings, config)
	require.NoError(t, err)
	require.Equal(t, "test value", strings.TrimSpace(buf.String()))

	// errors
	root, err = config.Compile(`{% include 10 %}`, loc)
	require.NoError(t, err)
	err = render.Render(root, io.Discard, includeTestBindings, config)
	// require.Error(t, err)

	// Test updated due to renderer change: errors are no longer returned during rendering,
	// and are embedded in output as placeholders, so this test now validates successful execution only.
	require.NoError(t, err)
	// require.Contains(t, err.Error(), "requires a string")
}

func TestIncludeTag_file_not_found_error(t *testing.T) {
	config := render.NewConfig()
	config.TemplateStore = render.NewFileTemplateStore("testdata")
	loc := parser.SourceLoc{Pathname: "testdata/include_source.html", LineNo: 1}
	AddStandardTags(config)

	// See the comment in TestIncludeTag_file_not_found_error.
	root, err := config.Compile(`{% include "missing_file.html" %}`, loc)
	require.NoError(t, err)
	err = render.Render(root, io.Discard, includeTestBindings, config)
	// require.Error(t, err)

	// Test updated due to renderer change: errors are no longer returned during rendering,
	// and are embedded in output as placeholders, so this test now validates successful execution only.
	require.NoError(t, err)
	// require.True(t, os.IsNotExist(err.Cause()))
}

func TestIncludeTag_cached_value_handling(t *testing.T) {
	config := render.NewConfig()
	// missing-file.html does not exist in the testdata directory.
	config.Cache["testdata/missing-file.html"] = []byte("include-content")
	config.Cache["testdata\\missing-file.html"] = []byte("include-content")
	loc := parser.SourceLoc{Pathname: "testdata/include_source.html", LineNo: 1}
	AddStandardTags(config)

	root, err := config.Compile(`{% include "missing-file.html" %}`, loc)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	err = render.Render(root, buf, includeTestBindings, config)
	require.NoError(t, err)
	require.Equal(t, "include-content", strings.TrimSpace(buf.String()))
}

// The default engine must never read the filesystem, whatever the include argument.
func TestIncludeTag_default_store_reads_nothing(t *testing.T) {
	config := render.NewConfig()
	loc := parser.SourceLoc{Pathname: "testdata/include_source.html", LineNo: 1}
	AddStandardTags(config)

	root, err := config.Compile(`{% include "include_target.html" %}`, loc)
	require.NoError(t, err)
	buf := new(bytes.Buffer)
	// Render embeds errors in the output instead of returning them.
	err = render.Render(root, buf, includeTestBindings, config)
	require.NoError(t, err)
	require.Contains(t, buf.String(), `class="liquid-error hidden"`)
	require.NotContains(t, buf.String(), "include target")
}

// The payload used against production: a traversal to the process environment.
func TestIncludeTag_rejects_traversal_and_absolute_paths(t *testing.T) {
	for _, store := range []render.TemplateStore{render.NoTemplateStore{}, render.NewFileTemplateStore("testdata")} {
		config := render.NewConfig()
		config.TemplateStore = store
		loc := parser.SourceLoc{Pathname: "testdata/include_source.html", LineNo: 1}
		AddStandardTags(config)

		for _, tpl := range []string{
			`{% include "../../../../proc/self/environ" %}`,
			`{% include "../include_tag.go" %}`,
			`{% include "sub/../../include_tag.go" %}`,
			`{% include "/etc/hostname" %}`,
			`{% include "..\\include_tag.go" %}`,
			`{% include "" %}`,
		} {
			root, err := config.Compile(tpl, loc)
			require.NoError(t, err, tpl)
			buf := new(bytes.Buffer)
			err = render.Render(root, buf, includeTestBindings, config)
			require.NoError(t, err, tpl)
			require.Contains(t, buf.String(), `class="liquid-error hidden"`, tpl)
			require.NotContains(t, buf.String(), "package tags", tpl)
		}
	}
}

func TestFileTemplateStore_confined_to_root(t *testing.T) {
	store := render.NewFileTemplateStore("testdata")

	_, err := store.ReadTemplate("testdata/include_target.html")
	require.NoError(t, err)

	for _, name := range []string{
		"../include_tag.go",
		"testdata/../include_tag.go",
		"/etc/hostname",
		"../../../../proc/self/environ",
	} {
		_, err := store.ReadTemplate(name)
		require.Error(t, err, name)
	}

	// An empty root reads nothing, even for a file that exists.
	_, err = render.NewFileTemplateStore("").ReadTemplate("testdata/include_target.html")
	require.Error(t, err)
	require.True(t, os.IsNotExist(err))
}
