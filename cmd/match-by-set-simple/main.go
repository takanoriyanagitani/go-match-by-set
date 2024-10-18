package main

import (
	"context"
	"log"
	"os"

	m "github.com/takanoriyanagitani/go-match-by-set/match"
	sm "github.com/takanoriyanagitani/go-match-by-set/match/simple"

	fm "github.com/takanoriyanagitani/go-match-by-set/words/fs"
	um "github.com/takanoriyanagitani/go-match-by-set/words/unix"
)

var WordsPath string = os.Getenv("ENV_WORDS_PATH")

var match m.Match = sm.SimpleMatch

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
