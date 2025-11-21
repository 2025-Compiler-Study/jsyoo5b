package parser

import (
	"io/fs"
	"strings"
)

func loadTestFiles(filesystem fs.FS, ext string) (map[string]string, error) {
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

			data, err := testLexerSuccessFs.ReadFile(path)
			if err != nil {
				return err
			}
			testName := strings.TrimSuffix(d.Name(), ext)
			testCases[testName] = string(data)
			return nil
		},
	)

	return testCases, err
}
