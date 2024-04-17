package main

import (
	"github.com/boekkooi-lengoo/icu4c-golang/icu"
	"log"
)

// https://github.com/unicode-org/icu/blob/617b094df3eb853a35f1227472178836ce625cff/docs/userguide/strings/utext.md?plain=1#L94
func main() {
	defer icu.Cleanup()

	bi, err := icu.NewBreakIterator(icu.BreakSentence, "de@ss=standard", "Well this is an n. Chr. hello world. This is a sentence.")
	if err != nil {
		log.Fatalln(err)
	}
	defer bi.Close()

	for {
		pos := bi.Next()
		if pos == icu.BreakDone {
			break
		}
		log.Println(pos)
	}
}
