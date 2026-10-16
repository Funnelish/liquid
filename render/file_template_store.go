package render

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// TemplateStoreDisabledError is returned by the default template store. It
// satisfies os.IsNotExist so that templates registered through the cache
// (Engine.ParseTemplateAndCache) are still found by {% include %}.
func templateStoreDisabledError(name string) error {
	return &fs.PathError{Op: "include", Path: name, Err: fs.ErrNotExist}
}

// NoTemplateStore is the default store. It never reads the filesystem, so a
// template cannot include arbitrary files unless the host application opts in
// by registering a FileTemplateStore or its own TemplateStore.
type NoTemplateStore struct{}

func (NoTemplateStore) ReadTemplate(name string) ([]byte, error) {
	return nil, templateStoreDisabledError(name)
}

// FileTemplateStore reads included templates from the filesystem, confined to
// Root. Names are resolved the way os.ReadFile would resolve them, then any
// path that lands outside Root, including through symlinks, is rejected. A
// FileTemplateStore with an empty Root reads nothing.
type FileTemplateStore struct {
	Root string
}

// NewFileTemplateStore returns a store confined to root.
func NewFileTemplateStore(root string) *FileTemplateStore {
	return &FileTemplateStore{Root: root}
}

func (tl *FileTemplateStore) ReadTemplate(filename string) ([]byte, error) {
	if tl == nil || tl.Root == "" {
		return nil, templateStoreDisabledError(filename)
	}
	resolved, err := tl.resolve(filename)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(resolved)
}

// resolve returns the absolute path of filename if and only if it lies within
// the store's root.
func (tl *FileTemplateStore) resolve(filename string) (string, error) {
	if strings.ContainsRune(filename, 0) {
		return "", fmt.Errorf("include: invalid template path %q", filename)
	}
	root, err := filepath.Abs(tl.Root)
	if err != nil {
		return "", err
	}
	if r, err := filepath.EvalSymlinks(root); err == nil {
		root = r
	}
	// The include tag has already joined the name onto the including template's
	// directory, so the name is resolved as given (relative to the working
	// directory, like os.ReadFile) and then required to sit inside Root.
	candidate, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	if !within(root, candidate) {
		return "", fmt.Errorf("include: template path %q is outside the template root", filename)
	}
	// Re-check after resolving symlinks, so a link inside root cannot point outside it.
	if real, err := filepath.EvalSymlinks(candidate); err == nil {
		if !within(root, real) {
			return "", fmt.Errorf("include: template path %q is outside the template root", filename)
		}
		candidate = real
	} else if !os.IsNotExist(err) {
		return "", err
	}
	return candidate, nil
}

func within(root, path string) bool {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}
