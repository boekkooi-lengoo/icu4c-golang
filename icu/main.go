package icu

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var (
	ErrTextIsNotUTF8             = errors.New("text is not valid UTF-8")
	ErrBreakIteratorIsClosed     = errors.New("the BreakIterator is closed")
	ErrBreakIteratorContainsText = errors.New("operation not permitted because the BreakIterator contains text")
)

// Initialize will attempt to load some part of ICU's data, and is useful as a test for configuration or installation problems that leave the ICU data inaccessible.
// Use of this function is optional.
func Initialize() error {
	var errCode errorCode
	icuInit(&errCode)
	if errCode.IsFailure() {
		return errCode
	}
	return nil
}

type BreakIterator struct {
	textUnit *text
	iterator *breakIterator
}

func NewBreakIterator(breakType breakIteratorType, locale string) (*BreakIterator, error) {
	var errCode errorCode
	bi := &BreakIterator{}
	bi.iterator = breakOpen(breakType, locale, nil, 0, &errCode)
	if errCode.IsFailure() {
		bi.Close()
		return nil, errCode
	}
	return bi, nil
}

func (bi *BreakIterator) Clone() (*BreakIterator, error) {
	if bi.IsClosed() {
		return nil, ErrBreakIteratorIsClosed
	}
	if bi.textUnit != nil {
		return nil, ErrBreakIteratorContainsText
	}

	var errCode errorCode
	clone := &BreakIterator{}
	clone.iterator = breakClone(bi.iterator, &errCode)
	if errCode.IsFailure() {
		clone.Close()
		return nil, errCode
	}

	return clone, nil
}

func (bi *BreakIterator) SetText(txt string) error {
	if bi.IsClosed() {
		return ErrBreakIteratorIsClosed
	}
	if !utf8.ValidString(txt) {
		return ErrTextIsNotUTF8
	}

	var errCode errorCode
	textUnit := textOpenUTF8(nil, txt, int64(len(txt)), &errCode)
	if errCode.IsFailure() {
		freeText(textUnit)
		return errCode
	}

	breakSetText(bi.iterator, textUnit, &errCode)
	if errCode.IsFailure() {
		freeText(textUnit)
		return errCode
	}

	if bi.textUnit != nil {
		bi.closeTextUnit()
	}
	bi.textUnit = textUnit
	return nil
}

func (bi *BreakIterator) Next() int32 {
	return breakNext(bi.iterator)
}

func (bi *BreakIterator) Close() {
	bi.closeTextUnit()
	bi.closeBreakIterator()
}

func (bi *BreakIterator) IsClosed() bool {
	return bi.iterator == nil
}

func (bi *BreakIterator) closeTextUnit() {
	freeText(bi.textUnit)
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

func freeText(textUnit *text) {
	if textUnit != nil {
		textClose(textUnit).Free()
		textUnit.Free()
	}
}
