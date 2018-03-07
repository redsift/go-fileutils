package fileutils

import (
	"os"
	"io/ioutil"
	"path/filepath"
)

func AtomicWrite(filename string, data []byte, perm os.FileMode) error {
	tname, err := func() (string, error){
		f, err := ioutil.TempFile(filepath.Dir(filename), "._")
		if err != nil {
			return "", err
		}
		defer f.Close()

		_, err = f.Write(data)
		if err != nil {
			return "", err
		}
		return f.Name(), nil
	}()

	if err != nil {
		return err
	}

	err = os.Rename(tname, filename)
	if err != nil {
		return err
	}

	return os.Chmod(filename, perm)
}