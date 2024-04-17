package icu

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestNewBreakIterator(t *testing.T) {
	t.Cleanup(Cleanup)
	t.Run("Invalid Break Type", func(t *testing.T) {
		bi, err := NewBreakIterator(-1, "en-us", "Unknown Locale provided")
		assert.Error(t, err)
		if !assert.Nil(t, bi) {
			bi.Close()
		}
	})
	t.Run("Text is not UTF-8", func(t *testing.T) {
		bi, err := NewBreakIterator(BreakSentence, "en-us", "a\xc5z")
		assert.ErrorContains(t, err, "text is not valid UTF-8")
		if !assert.Nil(t, bi) {
			bi.Close()
		}
	})

	t.Run("Sentences", func(t *testing.T) {
		testcases := []struct {
			text   string
			breaks []int32
		}{
			{
				text:   "1",
				breaks: []int32{1},
			},
			{
				text:   "\x001",
				breaks: []int32{2},
			},
			{
				text:   "Welcome to Golang. This is a ICU wrapper.",
				breaks: []int32{19, 41},
			},
			{
				text:   "Welcome to Golang. This is a ICU wrapper😜.",
				breaks: []int32{19, 45},
			},
		}
		t.Parallel()
		for i, testcase := range testcases {
			var testText = testcase.text
			var expectedBreaks = testcase.breaks
			t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
				bi, err := NewBreakIterator(BreakSentence, "en-us", testText)
				require.NoError(t, err)
				defer bi.Close()

				var breaks []int32
				for {
					pos := bi.Next()
					if pos == BreakDone {
						break
					}

					breaks = append(breaks, pos)
				}
				assert.Equal(t, expectedBreaks, breaks)
			})
		}
	})
}

func FuzzNewBreakIterator(f *testing.F) {
	f.Cleanup(Cleanup)

	testcases := []string{"Hello, world", " ", "a\xc5z", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, text string) {
		bi, err := NewBreakIterator(BreakCharacter, "en-us", text)
		if err != nil {
			if !utf8.ValidString(text) && err.Error() == "text is not valid UTF-8" {
				t.Skip("not a UTF-8 string")
			} else {
				t.Error("unexpected error", err)
			}
			return
		}
		if bi == nil {
			t.Error("bi should not be nil")
			return
		}
		defer bi.Close()

		sentences := strings.Builder{}
		var prev int32
		for {
			pos := bi.Next()
			if pos == BreakDone {
				break
			}

			str := text[prev:pos]
			if !utf8.ValidString(str) {
				t.Errorf("invalid UTF-8 string result %q", str)
			}
			sentences.WriteString(str)

			prev = pos
		}

		if res := sentences.String(); res != text {
			t.Errorf("sentences not matching text.\ntext:\n%q\n\nsentences:\n%q", text, res)
		}
	})
}
