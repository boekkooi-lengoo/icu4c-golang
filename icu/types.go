// WARNING: This file has automatically been generated on Wed, 17 Apr 2024 17:23:44 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package icu

/*
#cgo pkg-config: icu-i18n icu-io icu-uc
#cgo LDFLAGS: -licuuc -licudata
#include <unicode/utypes.h>
#include <unicode/utext.h>
#include <unicode/ubrk.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// breakIterator as declared in unicode/ubrk.h:31
type breakIterator C.UBreakIterator

// text as declared in unicode/utext.h:153
type text struct {
	magic               uint32
	flags               int32
	providerProperties  int32
	sizeOfStruct        int32
	chunkNativeLimit    int64
	extraSize           int32
	nativeIndexingLimit int32
	chunkNativeStart    int64
	chunkOffset         int32
	chunkLength         int32
	chunkContents       []uint16
	pExtra              unsafe.Pointer
	context             unsafe.Pointer
	p                   unsafe.Pointer
	q                   unsafe.Pointer
	r                   unsafe.Pointer
	privP               unsafe.Pointer
	a                   int64
	b                   int32
	c                   int32
	privA               int64
	privB               int32
	privC               int32
	ref8c2c6043         *C.UText
	allocs8c2c6043      interface{}
}
