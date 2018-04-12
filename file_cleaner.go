package downloader

import "os"

type FileCleaner struct {
	paths []string
}

func (fc *FileCleaner) AddPath(path string) *FileCleaner {
	fc.paths = append(fc.paths, path)

	return fc
}

func (fc *FileCleaner) Clean() error {
	for _, path := range fc.paths {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}
