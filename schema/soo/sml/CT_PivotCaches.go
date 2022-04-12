// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/bjshujie/gooxml"
)

type CT_PivotCaches struct {
	// PivotCache
	PivotCache []*CT_PivotCache
}

func NewCT_PivotCaches() *CT_PivotCaches {
	ret := &CT_PivotCaches{}
	return ret
}

func (m *CT_PivotCaches) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sepivotCache := xml.StartElement{Name: xml.Name{Local: "ma:pivotCache"}}
	for _, c := range m.PivotCache {
		e.EncodeElement(c, sepivotCache)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PivotCaches) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PivotCaches:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotCache"}:
				tmp := NewCT_PivotCache()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PivotCache = append(m.PivotCache, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_PivotCaches %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PivotCaches
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PivotCaches and its children
func (m *CT_PivotCaches) Validate() error {
	return m.ValidateWithPath("CT_PivotCaches")
}

// ValidateWithPath validates the CT_PivotCaches and its children, prefixing error messages with path
func (m *CT_PivotCaches) ValidateWithPath(path string) error {
	for i, v := range m.PivotCache {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PivotCache[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
