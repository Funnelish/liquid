package render

import "errors"

// ErrTemplateStoreDisabled is returned by DeniedTemplateStore. {% include %}
// surfaces it as empty output (the renderer embeds render errors as placeholders
// rather than failing the whole template).
var ErrTemplateStoreDisabled = errors.New("liquid: filesystem template reads are disabled; call RegisterTemplateStore to enable {% include %}")

// DeniedTemplateStore is the default TemplateStore. It refuses every read, so a
// template cannot use {% include %} to read arbitrary files off the server's
// disk (e.g. {% include "../../../../proc/self/environ" %} to exfiltrate the
// process environment). Callers that need filesystem includes opt in with
// Engine.RegisterTemplateStore(&FileTemplateStore{}). Cache-backed includes
// (ParseTemplateAndCache / Config.Cache) still work regardless of the store.
type DeniedTemplateStore struct{}

// ReadTemplate always fails without touching the filesystem.
func (DeniedTemplateStore) ReadTemplate(string) ([]byte, error) {
	return nil, ErrTemplateStoreDisabled
}
