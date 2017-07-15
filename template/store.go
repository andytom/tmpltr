package template

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Store is a directory that contains multiple templates.
type Store struct {
	dir string
}

// OpenStore returns a new Store for the passed dir. It will also create the
// dir if it doesn't exist.
func OpenStore(dir string) (Store, error) {
	ts := Store{dir}

	err := os.MkdirAll(dir, 0755)

	return ts, err
}

func (ts *Store) getPath(key string) string {
	return path.Join(ts.dir, key)
}

// List lists all the templates that are in a Store. If there is a template
// that cannot be parsed it will return an error.
func (ts *Store) List() (map[string]*Template, error) {
	templates := map[string]*Template{}

	files, err := ioutil.ReadDir(ts.dir)
	if err != nil {
		return templates, err
	}

	for _, file := range files {
		if file.IsDir() {
			path := ts.getPath(file.Name())

			t, err := New(path)
			if err != nil {
				return templates, err
			}

			templates[file.Name()] = &t

		}
	}

	return templates, nil
}

// Create copies a template into the Store and stores it under the passed key.
// If there are any errors this might leave a partially created template in the
// store.
func (ts *Store) Create(key string, t *Template) error {
	dest := ts.getPath(key)

	walker := func(path string, info os.FileInfo, err error) error {
		// First handle any incoming error by returning it to the
		// caller rather than trying to do something fancy here.
		if err != nil {
			return err
		}

		// Build our target path from the template baseDir and the
		// relative path to this path.
		relPath, err := filepath.Rel(t.baseDir, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dest, relPath)

		// If this is a dir we just create the dir and set the mode.
		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		// Create/truncate the target file and set it's mode before
		// we copy the content across.e
		dst, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		err = dst.Chmod(info.Mode())
		if err != nil {
			return err
		}

		src, err := os.Open(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(dst, src)
		return err
	}

	return filepath.Walk(t.baseDir, walker)
}

// Delete removes a template from the store
func (ts *Store) Delete(key string) error {
	path := ts.getPath(key)

	return os.RemoveAll(path)
}

// Get returns the template that is store under the passed key. If the template
// cannot be loaded it will return an error
func (ts *Store) Get(key string) (Template, error) {
	path := ts.getPath(key)

	return New(path)
}

// Replace is a helper method that calls Delete with the passed key and then
// Create with the passed key and template.
func (ts *Store) Replace(key string, t *Template) error {
	err := ts.Delete(key)
	if err != nil {
		return err
	}

	return ts.Create(key, t)
}
