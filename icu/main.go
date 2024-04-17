package icu

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type BreakIterator struct {
	textUnit *text
	iterator *breakIterator
}

func NewBreakIterator(breakType breakIteratorType, locale, text string) (*BreakIterator, error) {
	if !utf8.ValidString(text) {
		return nil, errors.New("text is not valid UTF-8")
	}

	var errCode errorCode
	bi := &BreakIterator{}
	bi.textUnit = textOpenUTF8(bi.textUnit, text, int64(len(text)), &errCode)
	if errCode.IsFailure() {
		bi.closeTextUnit()
		return nil, errCode
	}

	bi.iterator = breakOpen(breakType, locale, nil, 0, &errCode)
	if errCode.IsFailure() {
		bi.Close()
		return nil, errCode
	}

	breakSetText(bi.iterator, bi.textUnit, &errCode)
	if errCode.IsFailure() {
		bi.Close()
		return nil, errCode
	}

	return bi, nil
}

func (bi *BreakIterator) Next() int32 {
	return breakNext(bi.iterator)
}

func (bi *BreakIterator) Close() {
	bi.closeTextUnit()
	bi.closeBreakIterator()
}

func (bi *BreakIterator) closeTextUnit() {
	if bi.textUnit != nil {
		textClose(bi.textUnit).Free()
		bi.textUnit.Free()
	}
	bi.textUnit = nil
}

func (bi *BreakIterator) closeBreakIterator() {
	if bi.iterator != nil {
		breakClose(bi.iterator)
		// Calling free causes `free(): double free detected in tcache 2`
		//bi.iterator.Free()
	}
	bi.iterator = nil
}

// IsSuccess returns true if the error code indicate success (mimics U_SUCCESS)
// See https://github.com/unicode-org/icu/blob/release-74-1/icu4c/source/common/unicode/utypes.h#L700
func (c errorCode) IsSuccess() bool {
	return c <= 0
}

// IsFailure returns true if the error code indicate a failure (mimics U_FAILURE)
// See https://github.com/unicode-org/icu/blob/release-74-1/icu4c/source/common/unicode/utypes.h#L706
func (c errorCode) IsFailure() bool {
	return c > 0
}

func (c errorCode) Error() string {
	return fmt.Sprintf("ICU error (%d) %s", c, errorName(c))
}
