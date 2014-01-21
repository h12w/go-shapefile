// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c

import (
	"fmt"
	"unsafe"
)

import "C"

func GetFloat(p *C.double, i int) float64 {
	base, offset := uintptr(unsafe.Pointer(p)), unsafe.Sizeof(*p)*uintptr(i)
	return float64(*(*C.double)(unsafe.Pointer(base + offset)))
}

func GetInt(p *C.int, i int) int {
	base, offset := uintptr(unsafe.Pointer(p)), unsafe.Sizeof(*p)*uintptr(i)
	return int(*(*C.int)(unsafe.Pointer(base + offset)))
}

func p(v ...interface{}) {
	fmt.Println(v...)
}
