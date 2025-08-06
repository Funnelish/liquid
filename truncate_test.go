package liquid

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTruncateBug(t *testing.T) {
	src := `{{ blogPost.excerpt | truncate: 90, "-STOP-" }}`
	data := map[string]any{
		"blogPost": map[string]any{
			"excerpt": `6 Ridiculously Easy Ways to Increase Your Funnel Conversion Rate
In the e-commerce world, your funnel conversion rate is everything… No conversions = No sales.
Let's dive into 6 actionable steps that you can take to increase your funnel conversion rates – today! FYI – These are explicit changes to make to your approach.`,
		},
	}

	engine := NewEngine()
	tpl, err := engine.ParseTemplate([]byte(src))
	require.NoError(t, err)

	var buf bytes.Buffer
	out, err := tpl.Render(data)
	require.NoError(t, err)
	buf.Write(out)
	result := buf.String()

	t.Logf("Rendered output:\n%s", result)

	require.Contains(t, result, "-STOP-", "truncate should append '-STOP-'")
	require.NotContains(t, result, "Let's dive", "truncate should cut text after '-STOP-'")
}
