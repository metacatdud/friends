package execute

import "os"

func MkDir(path string) error {
	return os.MkdirAll(path, 0777)
}

func CheckDir(path string) error {
	dir, err := os.Open(path)
	defer func() {
		_ = dir.Close()
	}()

	if err == nil {
		return os.ErrExist
	}
	return nil
}
