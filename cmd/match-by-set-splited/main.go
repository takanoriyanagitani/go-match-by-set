package main

import (
	"context"
	"log"
	"os"
	"strings"

	m "github.com/takanoriyanagitani/go-match-by-set/match"
	sm "github.com/takanoriyanagitani/go-match-by-set/match/simple"

	fm "github.com/takanoriyanagitani/go-match-by-set/words/fs"
	um "github.com/takanoriyanagitani/go-match-by-set/words/unix"
)

func getenvOrAlt(key, alt string) string {
	val, found := os.LookupEnv(key)
	switch found {
	case true:
		return val
	default:
		return alt
	}
}

var WordsPath string = os.Getenv("ENV_WORDS_PATH")

var Separator string = getenvOrAlt("ENV_SPLIT_CHAR", " ")

var match m.Match = m.Match(sm.SimpleMatch).ToSplited(
	func(line string) []string { return strings.Split(line, Separator) },
)

func onMap(s map[string]struct{}) error {
	return match.StdinToStdout(
		context.Background(),
		s,
	)
}

func sub() error {
	switch len(WordsPath) {
	case 0:
		return onMap(um.DefaultMap())
	default:
		return onMap(fm.FsToMapDefault(WordsPath))
	}
}

func main() {
	e := sub()
	if nil != e {
		log.Printf("%v\n", e)
	}
}
