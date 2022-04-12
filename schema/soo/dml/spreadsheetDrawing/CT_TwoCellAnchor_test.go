// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/bjshujie/gooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_TwoCellAnchorConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_TwoCellAnchor()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_TwoCellAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_TwoCellAnchor should validate: %s", err)
	}
}

func TestCT_TwoCellAnchorMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_TwoCellAnchor()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_TwoCellAnchor()
	xml.Unmarshal(buf, v2)
}
