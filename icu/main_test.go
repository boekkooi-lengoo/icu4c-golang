package icu

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestNewBreakIterator(t *testing.T) {
	t.Cleanup(Cleanup)
	t.Run("Invalid Break Type", func(t *testing.T) {
		bi, err := NewBreakIterator(-1, "en-us")
		assert.Error(t, err)
		if !assert.Nil(t, bi) {
			bi.Close()
		}
	})
}

func TestBreakIterator_SetText(t *testing.T) {
	t.Cleanup(Cleanup)
	t.Run("Text is not UTF-8", func(t *testing.T) {
		bi, err := NewBreakIterator(BreakSentence, "en-us")
		require.NoError(t, err)
		defer bi.Close()

		err = bi.SetText("a\xc5z")
		assert.ErrorContains(t, err, "text is not valid UTF-8")
	})
	t.Run("Set Text on closed iterator", func(t *testing.T) {
		bi, err := NewBreakIterator(BreakCharacter, "en-us")
		require.NoError(t, err)
		require.NotNil(t, bi)
		bi.Close()

		err = bi.SetText("Hello")
		assert.ErrorIs(t, err, ErrBreakIteratorIsClosed)
	})
	t.Run("Set Text 3 times", func(t *testing.T) {
		bi, err := NewBreakIterator(BreakCharacter, "en-us")
		require.NoError(t, err)
		defer bi.Close()

		err = bi.SetText("Welcome to Golang.")
		assert.Len(t, getBreaks(bi), 18)

		err = bi.SetText("This is a ICU wrapper.")
		assert.Len(t, getBreaks(bi), 22)

		err = bi.SetText("Welcome to Golang.")
		assert.Len(t, getBreaks(bi), 18)
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
				bi, err := NewBreakIterator(BreakSentence, "en-us")
				require.NoError(t, err)
				defer bi.Close()

				err = bi.SetText(testText)
				require.NoError(t, err)

				breaks := getBreaks(bi)
				assert.Equal(t, expectedBreaks, breaks)
			})
		}
	})
}

func FuzzBreakIterator_SetText(f *testing.F) {
	f.Cleanup(Cleanup)

	testcases := []string{"Hello, world", " ", "a\xc5z", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, text string) {
		bi, err := NewBreakIterator(BreakCharacter, "en-us")
		if err != nil {
			t.Error(err)
			return
		}

		err = bi.SetText(text)
		if err != nil {
			if !utf8.ValidString(text) && errors.Is(err, ErrTextIsNotUTF8) {
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

func BenchmarkBreakIterator_SetText(b *testing.B) {
	b.StopTimer()
	bi, err := NewBreakIterator(BreakCharacter, "en-us")
	require.NoError(b, err)
	b.Cleanup(Cleanup)
	b.Cleanup(bi.Close)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		txt := "Hello World" + strconv.Itoa(i)
		err = bi.SetText(txt)
		if err != nil {
			b.Error(txt, err)
			continue
		}

		if len(getBreaks(bi)) != utf8.RuneCountInString(txt) {
			b.Error("invalid character count")
		}
	}
}

func TestBreakIterator_Clone(t *testing.T) {
	t.Cleanup(Cleanup)

	t.Run("Clone and setText", func(t *testing.T) {
		bi, err := NewBreakIterator(BreakCharacter, "en-us")
		require.NoError(t, err)
		require.NotNil(t, bi)
		defer bi.Close()

		clone, err := bi.Clone()
		require.NoError(t, err)
		assert.NotSame(t, bi, clone)
		defer clone.Close()

		err = bi.SetText("Hello")
		require.NoError(t, err)

		err = clone.SetText("World!")
		require.NoError(t, err)

		assert.Len(t, getBreaks(bi), 5)
		assert.Len(t, getBreaks(clone), 6)
	})

	t.Run("Clone with Text", func(t *testing.T) {
		bi, err := NewBreakIterator(BreakCharacter, "en-us")
		require.NoError(t, err)
		require.NotNil(t, bi)
		defer bi.Close()

		err = bi.SetText("Hello World!")
		require.NoError(t, err)

		clone, err := bi.Clone()
		assert.ErrorIs(t, err, ErrBreakIteratorContainsText)
		if !assert.Nil(t, clone) {
			clone.Close()
		}
	})

	t.Run("Clone closed", func(t *testing.T) {
		bi, err := NewBreakIterator(BreakCharacter, "en-us")
		require.NoError(t, err)
		require.NotNil(t, bi)
		bi.Close()

		clone, err := bi.Clone()
		assert.ErrorIs(t, err, ErrBreakIteratorIsClosed)
		if !assert.Nil(t, clone) {
			clone.Close()
		}
	})
}

func getBreaks(bi *BreakIterator) []int32 {
	var breaks []int32
	for {
		pos := bi.Next()
		if pos == BreakDone {
			break
		}

		breaks = append(breaks, pos)
	}
	return breaks
}

func BenchmarkBreakIterator_New(b *testing.B) {
	b.Cleanup(Cleanup)
	for i := 0; i < b.N; i++ {
		bi, err := NewBreakIterator(BreakCharacter, "en-us")
		require.NoError(b, err)

		txt := "Hello World" + strconv.Itoa(i)
		err = bi.SetText(txt)
		if err != nil {
			b.Error(txt, err)
		} else if len(getBreaks(bi)) != utf8.RuneCountInString(txt) {
			b.Error("invalid character count")
		}
		bi.Close()
	}
}

func BenchmarkBreakIterator_Clone(b *testing.B) {
	b.Cleanup(Cleanup)
	bi, err := NewBreakIterator(BreakCharacter, "en-us")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		txt := "Hello World" + strconv.Itoa(i)
		err = bi.SetText(txt)
		if err != nil {
			b.Error(txt, err)
		} else if len(getBreaks(bi)) != utf8.RuneCountInString(txt) {
			b.Error("invalid character count")
		}
	}

	bi.Close()
}
