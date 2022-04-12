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

type CT_VolMain struct {
	// First String
	FirstAttr string
	// Topic
	Tp []*CT_VolTopic
}

func NewCT_VolMain() *CT_VolMain {
	ret := &CT_VolMain{}
	return ret
}

func (m *CT_VolMain) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "first"},
		Value: fmt.Sprintf("%v", m.FirstAttr)})
	e.EncodeToken(start)
	setp := xml.StartElement{Name: xml.Name{Local: "ma:tp"}}
	for _, c := range m.Tp {
		e.EncodeElement(c, setp)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_VolMain) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "first" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FirstAttr = parsed
			continue
		}
	}
lCT_VolMain:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tp"}:
				tmp := NewCT_VolTopic()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tp = append(m.Tp, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_VolMain %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_VolMain
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_VolMain and its children
func (m *CT_VolMain) Validate() error {
	return m.ValidateWithPath("CT_VolMain")
}

// ValidateWithPath validates the CT_VolMain and its children, prefixing error messages with path
func (m *CT_VolMain) ValidateWithPath(path string) error {
	for i, v := range m.Tp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tp[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
