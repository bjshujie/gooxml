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

func TestCT_RevisionConflictConstructor(t *testing.T) {
	v := sml.NewCT_RevisionConflict()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionConflict must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionConflict should validate: %s", err)
	}
}

func TestCT_RevisionConflictMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionConflict()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionConflict()
	xml.Unmarshal(buf, v2)
}
