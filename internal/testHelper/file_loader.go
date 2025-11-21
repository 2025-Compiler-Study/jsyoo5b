package testHelper

import (
	"io/fs"
	"strings"
)

func LoadTestFile(filesystem fs.FS, trimExt string) (map[string]string, error) {
	testCases := make(map[string]string)

	err := fs.WalkDir(
		filesystem,
		".",
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			} else if d.IsDir() {
				return nil
			}

			data, err := fs.ReadFile(filesystem, path)
			if err != nil {
				return err
			}
			testName := strings.TrimSuffix(d.Name(), trimExt)
			testCases[testName] = string(data)
			return nil
		},
	)

	return testCases, err
}
