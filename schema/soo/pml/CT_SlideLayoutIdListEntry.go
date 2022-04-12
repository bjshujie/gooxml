// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/bjshujie/gooxml"
)

type CT_SlideLayoutIdListEntry struct {
	// ID Tag
	IdAttr  *uint32
	RIdAttr string
	ExtLst  *CT_ExtensionList
}

func NewCT_SlideLayoutIdListEntry() *CT_SlideLayoutIdListEntry {
	ret := &CT_SlideLayoutIdListEntry{}
	return ret
}

func (m *CT_SlideLayoutIdListEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.RIdAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideLayoutIdListEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RIdAttr = parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IdAttr = &pt
			continue
		}
	}
lCT_SlideLayoutIdListEntry:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SlideLayoutIdListEntry %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideLayoutIdListEntry
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideLayoutIdListEntry and its children
func (m *CT_SlideLayoutIdListEntry) Validate() error {
	return m.ValidateWithPath("CT_SlideLayoutIdListEntry")
}

// ValidateWithPath validates the CT_SlideLayoutIdListEntry and its children, prefixing error messages with path
func (m *CT_SlideLayoutIdListEntry) ValidateWithPath(path string) error {
	if m.IdAttr != nil {
		if *m.IdAttr < 2147483648 {
			return fmt.Errorf("%s/m.IdAttr must be >= 2147483648 (have %v)", path, *m.IdAttr)
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
