// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/bjshujie/gooxml/schema/soo/sml"
)

func TestCalcChainConstructor(t *testing.T) {
	v := sml.NewCalcChain()
	if v == nil {
		t.Errorf("sml.NewCalcChain must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CalcChain should validate: %s", err)
	}
}

func TestCalcChainMarshalUnmarshal(t *testing.T) {
	v := sml.NewCalcChain()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCalcChain()
	xml.Unmarshal(buf, v2)
}
