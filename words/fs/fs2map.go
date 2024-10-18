package fsmap

import (
	"bufio"
	"io/fs"
	"os"

	sw "github.com/takanoriyanagitani/go-match-by-set/words/std/scanner"
)

func FsToMap(fsys fs.FS, path2dict string) map[string]struct{} {
	f, e := fsys.Open(path2dict)
	if nil != e {
		return map[string]struct{}{}
	}
	defer f.Close()

	var s *bufio.Scanner = bufio.NewScanner(f)
	return sw.ScannerToMap(s)
}

func FsToMapDefault(path2dict string) map[string]struct{} {
	return FsToMap(
		os.DirFS("/"),
		path2dict,
	)
}
