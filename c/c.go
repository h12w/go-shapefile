// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c

/*
#cgo LDFLAGS: -lshp
#include <shapefil.h>
#include <string.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

type (
	SHPHandle C.SHPHandle
	DBFHandle C.DBFHandle
	SHPObject struct {
		ShapeType    C.int
		ShapeId      C.int
		NParts       C.int
		PanPartStart *C.int
		PanPartType  *C.int
		NVertices    C.int
		PadfX        *C.double
		PadfY        *C.double
		PadfZ        *C.double
		PadfM        *C.double
		XMin         C.double
		YMin         C.double
		ZMin         C.double
		MMin         C.double
		XMax         C.double
		YMax         C.double
		ZMax         C.double
		MMax         C.double
	}
)

func SHPOpen(filename, mode string) SHPHandle {
	filename_, mode_ := C.CString(filename), C.CString(mode)
	defer C.free(unsafe.Pointer(filename_))
	defer C.free(unsafe.Pointer(mode_))
	return SHPHandle(C.SHPOpen(filename_, mode_))
}

func SHPClose(h SHPHandle) {
	C.SHPClose(h)
}

func SHPGetInfo(hSHP SHPHandle) (shapeType int, nEntries int, minBound, maxBound [4]C.double) {
	var shapeType_, nEntries_ C.int
	C.SHPGetInfo(hSHP,
		&nEntries_,
		&shapeType_,
		&minBound[0],
		&maxBound[0])
	shapeType, nEntries = int(shapeType_), int(nEntries_)
	return
}

func SHPReadObject(hSHP SHPHandle, iShape int) *SHPObject {
	return (*SHPObject)(unsafe.Pointer(C.SHPReadObject(hSHP, C.int(iShape))))
}

func SHPDestroyObject(o *SHPObject) {
	C.SHPDestroyObject((*C.SHPObject)(unsafe.Pointer(o)))
}

// ----------------------------------------------------------
// DBF

func DBFOpen(filename, mode string) DBFHandle {
	filename_, mode_ := C.CString(filename), C.CString(mode)
	defer C.free(unsafe.Pointer(filename_))
	defer C.free(unsafe.Pointer(mode_))
	return DBFHandle(C.DBFOpen(filename_, mode_))
}

func DBFClose(h DBFHandle) {
	C.DBFClose(h)
}

func DBFGetFieldCount(h DBFHandle) int {
	return int(C.DBFGetFieldCount(h))
}

func DBFGetRecordCount(h DBFHandle) int {
	return int(C.DBFGetRecordCount(h))
}

// ??
func DBFGetFieldIndex(h DBFHandle, fieldName string) int {
	fieldName_ := C.CString(fieldName)
	defer C.free(unsafe.Pointer(fieldName_))
	return int(C.DBFGetFieldIndex(h, fieldName_))
}

func DBFGetFieldInfo(h DBFHandle, fieldIndex int) (fieldName string, fieldType int, nWidth, nDecimals int) {
	var fieldName_ [12]C.char
	var nWidth_, nDecimals_ C.int
	fieldType = int(C.DBFGetFieldInfo(h, C.int(fieldIndex), &fieldName_[0], &nWidth_, &nDecimals_))
	fieldName = C.GoString(&fieldName_[0])
	nWidth = int(nWidth_)
	nDecimals = int(nDecimals_)
	return
}

func DBFReadIntegerAttribute(h DBFHandle, shapeIndex, fieldIndex int) int {
	return int(C.DBFReadIntegerAttribute(h, C.int(shapeIndex), C.int(fieldIndex)))
}

func DBFReadDoubleAttribute(h DBFHandle, shapeIndex, fieldIndex int) float64 {
	return float64(C.DBFReadDoubleAttribute(h, C.int(shapeIndex), C.int(fieldIndex)))
}

func DBFReadStringAttribute(h DBFHandle, shapeIndex, fieldIndex int) []byte {
	cstr := C.DBFReadStringAttribute(h, C.int(shapeIndex), C.int(fieldIndex))
	slen := int(C.strlen(cstr))
	bytes := make([]byte, slen)
	if slen > 0 {
		C.memcpy(unsafe.Pointer(&bytes[0]), unsafe.Pointer(cstr), C.size_t(slen))
	}
	return bytes
}
