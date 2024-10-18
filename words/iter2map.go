package words

import (
	"iter"
	"maps"
)

func IteratorToMap(i iter.Seq2[string, struct{}]) map[string]struct{} {
	return maps.Collect(i)
}
