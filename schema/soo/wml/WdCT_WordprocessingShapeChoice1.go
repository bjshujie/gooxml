// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/bjshujie/gooxml"
)

type WdCT_WordprocessingShapeChoice1 struct {
	Txbx       *WdCT_TextboxInfo
	LinkedTxbx *WdCT_LinkedTextboxInformation
}

func NewWdCT_WordprocessingShapeChoice1() *WdCT_WordprocessingShapeChoice1 {
	ret := &WdCT_WordprocessingShapeChoice1{}
	return ret
}

func (m *WdCT_WordprocessingShapeChoice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Txbx != nil {
		setxbx := xml.StartElement{Name: xml.Name{Local: "wp:txbx"}}
		e.EncodeElement(m.Txbx, setxbx)
	}
	if m.LinkedTxbx != nil {
		selinkedTxbx := xml.StartElement{Name: xml.Name{Local: "wp:linkedTxbx"}}
		e.EncodeElement(m.LinkedTxbx, selinkedTxbx)
	}
	return nil
}

func (m *WdCT_WordprocessingShapeChoice1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lWdCT_WordprocessingShapeChoice1:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "txbx"}:
				m.Txbx = NewWdCT_TextboxInfo()
				if err := d.DecodeElement(m.Txbx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "linkedTxbx"}:
				m.LinkedTxbx = NewWdCT_LinkedTextboxInformation()
				if err := d.DecodeElement(m.LinkedTxbx, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_WordprocessingShapeChoice1 %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WordprocessingShapeChoice1
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WordprocessingShapeChoice1 and its children
func (m *WdCT_WordprocessingShapeChoice1) Validate() error {
	return m.ValidateWithPath("WdCT_WordprocessingShapeChoice1")
}

// ValidateWithPath validates the WdCT_WordprocessingShapeChoice1 and its children, prefixing error messages with path
func (m *WdCT_WordprocessingShapeChoice1) ValidateWithPath(path string) error {
	if m.Txbx != nil {
		if err := m.Txbx.ValidateWithPath(path + "/Txbx"); err != nil {
			return err
		}
	}
	if m.LinkedTxbx != nil {
		if err := m.LinkedTxbx.ValidateWithPath(path + "/LinkedTxbx"); err != nil {
			return err
		}
	}
	return nil
}
