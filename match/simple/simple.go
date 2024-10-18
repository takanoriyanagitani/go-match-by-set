package smatch

import (
	"context"
)

func SimpleMatch(_ context.Context, line string, m map[string]struct{}) bool {
	_, ok := m[line]
	return ok
}
