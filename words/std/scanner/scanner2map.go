package scan2map

import (
	"bufio"
	"iter"

	mw "github.com/takanoriyanagitani/go-match-by-set/words"
)

func ScannerToIter(s *bufio.Scanner) iter.Seq2[string, struct{}] {
	return func(yield func(string, struct{}) bool) {
		for s.Scan() {
			if !yield(s.Text(), struct{}{}) {
				return
			}
		}
	}
}

func ScannerToMap(s *bufio.Scanner) map[string]struct{} {
	var i iter.Seq2[string, struct{}] = ScannerToIter(s)
	return mw.IteratorToMap(i)
}
