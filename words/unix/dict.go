package unixdict

import (
	"os"

	fw "github.com/takanoriyanagitani/go-match-by-set/words/fs"
)

const WordsPath string = "usr/share/dict/words"

func DefaultMap() map[string]struct{} {
	return fw.FsToMap(
		os.DirFS("/"),
		WordsPath,
	)
}
