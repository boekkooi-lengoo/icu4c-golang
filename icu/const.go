// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package icu

/*
#cgo pkg-config: icu-i18n icu-io icu-uc
#cgo LDFLAGS: -licuuc -licudata
#include <unicode/utypes.h>
#include <unicode/utext.h>
#include <unicode/ubrk.h>
#include <unicode/uclean.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// BreakH as defined in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L11

	// BreakTypedefUbreakIterator as defined in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L26

	// BreakDone as defined in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L133
	BreakDone = ((int32)(-1))
	// BreakSafecloneBuffersize as defined in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L346
	BreakSafecloneBuffersize = 1
)

// breakIteratorType as declared in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L128
type breakIteratorType int32

// breakIteratorType enumeration from https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L128
const (
	BreakCharacter breakIteratorType = iota
	BreakWord      breakIteratorType = 1
	BreakLine      breakIteratorType = 2
	BreakSentence  breakIteratorType = 3
	BreakTitle     breakIteratorType = 4
	BreakCount     breakIteratorType = 5
)

// wordBreak as declared in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L170
type wordBreak int32

// wordBreak enumeration from https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L170
const (
	BreakWordNone        wordBreak = iota
	BreakWordNoneLimit   wordBreak = 100
	BreakWordNumber      wordBreak = 100
	BreakWordNumberLimit wordBreak = 200
	BreakWordLetter      wordBreak = 200
	BreakWordLetterLimit wordBreak = 300
	BreakWordKana        wordBreak = 300
	BreakWordKanaLimit   wordBreak = 400
	BreakWordIdeo        wordBreak = 400
	BreakWordIdeoLimit   wordBreak = 500
)

// lineBreakTag as declared in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L193
type lineBreakTag int32

// lineBreakTag enumeration from https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L193
const (
	BreakLineSoft      lineBreakTag = iota
	BreakLineSoftLimit lineBreakTag = 100
	BreakLineHard      lineBreakTag = 100
	BreakLineHardLimit lineBreakTag = 200
)

// sentenceBreakTag as declared in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L224
type sentenceBreakTag int32

// sentenceBreakTag enumeration from https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/ubrk.h#L224
const (
	BreakSentenceTerm      sentenceBreakTag = iota
	BreakSentenceTermLimit sentenceBreakTag = 100
	BreakSentenceSep       sentenceBreakTag = 100
	BreakSentenceSepLimit  sentenceBreakTag = 200
)

// wordBreakValues as declared in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/uchar.h#L2355
type wordBreakValues int32

// wordBreakValues enumeration from https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/uchar.h#L2355
const (
	WordBreakOther             wordBreakValues = iota
	WordBreakAletter           wordBreakValues = 1
	WordBreakFormat            wordBreakValues = 2
	WordBreakKatakana          wordBreakValues = 3
	WordBreakMidletter         wordBreakValues = 4
	WordBreakMidnum            wordBreakValues = 5
	WordBreakNumeric           wordBreakValues = 6
	WordBreakExtendnumlet      wordBreakValues = 7
	WordBreakCr                wordBreakValues = 8
	WordBreakExtend            wordBreakValues = 9
	WordBreakLf                wordBreakValues = 10
	WordBreakMidnumlet         wordBreakValues = 11
	WordBreakNewline           wordBreakValues = 12
	WordBreakRegionalIndicator wordBreakValues = 13
	WordBreakHebrewLetter      wordBreakValues = 14
	WordBreakSingleQuote       wordBreakValues = 15
	WordBreakDoubleQuote       wordBreakValues = 16
	WordBreakEBase             wordBreakValues = 17
	WordBreakEBaseGaz          wordBreakValues = 18
	WordBreakEModifier         wordBreakValues = 19
	WordBreakGlueAfterZwj      wordBreakValues = 20
	WordBreakZwj               wordBreakValues = 21
	WordBreakWsegspace         wordBreakValues = 22
	WordBreakCount             wordBreakValues = 23
)

// errorCode as declared in https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/utypes.h#L689
type errorCode int32

// errorCode enumeration from https://github.com/unicode-org/icu/blob/release-74-2/icu4c/source/common/unicode/utypes.h#L689
const (
	ZeroError                  errorCode = 0
	BreakInternalError         errorCode = 66048
	BreakErrorStart            errorCode = 66048
	BreakHexDigitsExpected     errorCode = 66049
	BreakSemicolonExpected     errorCode = 66050
	BreakRuleSyntax            errorCode = 66051
	BreakUnclosedSet           errorCode = 66052
	BreakAssignError           errorCode = 66053
	BreakVariableRedfinition   errorCode = 66054
	BreakMismatchedParen       errorCode = 66055
	BreakNewLineInQuotedString errorCode = 66056
	BreakUndefinedVariable     errorCode = 66057
	BreakInitError             errorCode = 66058
	BreakRuleEmptySet          errorCode = 66059
	BreakUnrecognizedOption    errorCode = 66060
	BreakMalformedRuleTag      errorCode = 66061
	BreakErrorLimit            errorCode = 66062
)

const ()

const ()

const ()

const ()

const ()
