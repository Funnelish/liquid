package render

import (
	"github.com/Funnelish/liquid/parser"
)

// An Error is an error during template rendering.
type Error interface {
	Path() string
	LineNumber() int
	Cause() error
	Error() string
}

func renderErrorf(loc parser.Locatable, format string, a ...any) Error {
	return parser.Errorf(loc, format, a...)
}

func wrapRenderError(err error, loc parser.Locatable) Error {
	return parser.WrapError(err, loc)
}
