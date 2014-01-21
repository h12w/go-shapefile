// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c

import (
	"code.google.com/p/mahonia"
	"testing"
)

const (
	String = iota
	Integer
	Double
	Logical
	Invalid
)

func Test_quick(t *testing.T) {
	dbfHandle := DBFOpen("../map/bou2_4p.dbf", "rb")
	defer DBFClose(dbfHandle)

	decoder := mahonia.NewDecoder("gbk")

	for i := 0; i < DBFGetRecordCount(dbfHandle); i++ {
		for j := 0; j < DBFGetFieldCount(dbfHandle); j++ {
			name, type_, _, _ := DBFGetFieldInfo(dbfHandle, j)
			switch type_ {
			case String:
				p(name, decoder.ConvertString(string(DBFReadStringAttribute(dbfHandle, i, j))))
			case Integer, Logical:
				p(name, DBFReadIntegerAttribute(dbfHandle, i, j))
			case Double:
				p(name, DBFReadDoubleAttribute(dbfHandle, i, j))
			}
		}
	}
}
