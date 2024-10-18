package match

import (
	"bufio"
	"context"
	"errors"
	"io"
	"iter"
	"os"
	"strings"
)

type Match func(ctx context.Context, line string, m map[string]struct{}) bool

func (m Match) ToLower() Match {
	return func(ctx context.Context, line string, s map[string]struct{}) bool {
		var lower string = strings.ToLower(line)
		return m(ctx, lower, s)
	}
}

func (m Match) ToSplited(line2splited func(string) []string) Match {
	return func(ctx context.Context, line string, s map[string]struct{}) bool {
		var splited []string = line2splited(line)
		for _, item := range splited {
			select {
			case <-ctx.Done():
				return false
			default:
			}

			var hit bool = m(ctx, item, s)
			if hit {
				return true
			}
		}
		return false
	}
}

func (m Match) WriteMatchAll(
	ctx context.Context,
	lines iter.Seq[string],
	s map[string]struct{},
	wtr func(string) error,
) error {
	for line := range lines {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		var matched bool = m(ctx, line, s)
		if !matched {
			continue
		}

		e := wtr(line)
		if nil != e {
			return e
		}
	}
	return nil
}

func (m Match) ReaderToWriter(
	ctx context.Context,
	s map[string]struct{},
	rdr io.Reader,
	wtr io.Writer,
) error {
	var bw *bufio.Writer = bufio.NewWriter(wtr)
	defer bw.Flush()

	var scanner *bufio.Scanner = bufio.NewScanner(rdr)
	var i iter.Seq[string] = func(yield func(string) bool) {
		for scanner.Scan() {
			var line string = scanner.Text()
			if !yield(line) {
				return
			}
		}
	}
	return m.WriteMatchAll(
		ctx,
		i,
		s,
		func(line string) error {
			_, e1 := bw.WriteString(line)
			e2 := bw.WriteByte('\n')
			return errors.Join(e1, e2)
		},
	)
}

func (m Match) StdinToStdout(
	ctx context.Context,
	s map[string]struct{},
) error {
	return m.ReaderToWriter(ctx, s, os.Stdin, os.Stdout)
}
