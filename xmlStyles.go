// Copyright 2016 - 2022 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.15 or later.

package excelize

import (
	"encoding/xml"
	"sync"
)

// xlsxStyleSheet is the root element of the Styles part.
type xlsxStyleSheet struct {
	sync.Mutex
	XMLName      xml.Name          `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main styleSheet"`
	NumFmts      *xlsxNumFmts      `xml:"numFmts"`
	Fonts        *xlsxFonts        `xml:"fonts"`
	Fills        *xlsxFills        `xml:"fills"`
	Borders      *xlsxBorders      `xml:"borders"`
	CellStyleXfs *xlsxCellStyleXfs `xml:"cellStyleXfs"`
	CellXfs      *xlsxCellXfs      `xml:"cellXfs"`
	CellStyles   *xlsxCellStyles   `xml:"cellStyles"`
	Dxfs         *xlsxDxfs         `xml:"dxfs"`
	TableStyles  *xlsxTableStyles  `xml:"tableStyles"`
	Colors       *xlsxStyleColors  `xml:"colors"`
	ExtLst       *xlsxExtLst       `xml:"extLst"`
}

// xlsxAlignment formatting information pertaining to text alignment in cells.
// There are a variety of choices for how text is aligned both horizontally and
// vertically, as well as indentation settings, and so on.
type xlsxAlignment struct {
	Horizontal      string `xml:"horizontal,attr,omitempty" json:"horizontal,omitempty"`
	Indent          int    `xml:"indent,attr,omitempty" json:"indent,omitempty"`
	JustifyLastLine bool   `xml:"justifyLastLine,attr,omitempty" json:"justify_last_line,omitempty"`
	ReadingOrder    uint64 `xml:"readingOrder,attr,omitempty" json:"reading_order,omitempty"`
	RelativeIndent  int    `xml:"relativeIndent,attr,omitempty" json:"relative_indent,omitempty"`
	ShrinkToFit     bool   `xml:"shrinkToFit,attr,omitempty" json:"shrink_to_fit,omitempty"`
	TextRotation    int    `xml:"textRotation,attr,omitempty" json:"text_rotation,omitempty"`
	Vertical        string `xml:"vertical,attr,omitempty" json:"vertical,omitempty"`
	WrapText        bool   `xml:"wrapText,attr,omitempty" json:"wrap_text,omitempty"`
}

// xlsxProtection (Protection Properties) contains protection properties
// associated with the cell. Each cell has protection properties that can be
// set. The cell protection properties do not take effect unless the sheet has
// been protected.
type xlsxProtection struct {
	Hidden *bool `xml:"hidden,attr" json:"hidden,omitempty"`
	Locked *bool `xml:"locked,attr" json:"locked,omitempty"`
}

// xlsxLine expresses a single set of cell border.
type xlsxLine struct {
	Style string     `xml:"style,attr,omitempty" json:"style,omitempty"`
	Color *xlsxColor `xml:"color" json:"color,omitempty"`
}

// xlsxColor is a common mapping used for both the fgColor and bgColor elements.
// Foreground color of the cell fill pattern. Cell fill patterns operate with
// two colors: a background color and a foreground color. These combine together
// to make a patterned cell fill. Background color of the cell fill pattern.
// Cell fill patterns operate with two colors: a background color and a
// foreground color. These combine together to make a patterned cell fill.
type xlsxColor struct {
	Auto    bool    `xml:"auto,attr,omitempty" json:"auto,omitempty"`
	RGB     string  `xml:"rgb,attr,omitempty" json:"rgb,omitempty"`
	Indexed int     `xml:"indexed,attr,omitempty" json:"indexed,omitempty"`
	Theme   *int    `xml:"theme,attr" json:"theme,omitempty"`
	Tint    float64 `xml:"tint,attr,omitempty" json:"tint,omitempty"`
}

// xlsxFonts directly maps the font element. This element contains all font
// definitions for this workbook.
type xlsxFonts struct {
	Count int         `xml:"count,attr"`
	Font  []*xlsxFont `xml:"font"`
}

// xlsxFont directly maps the font element. This element defines the
// properties for one of the fonts used in this workbook.
type xlsxFont struct {
	B        *attrValBool   `xml:"b" json:"b,omitempty"`
	I        *attrValBool   `xml:"i" json:"i,omitempty"`
	Strike   *attrValBool   `xml:"strike" json:"strike,omitempty"`
	Outline  *attrValBool   `xml:"outline" json:"outline,omitempty"`
	Shadow   *attrValBool   `xml:"shadow" json:"shadow,omitempty"`
	Condense *attrValBool   `xml:"condense" json:"condense,omitempty"`
	Extend   *attrValBool   `xml:"extend" json:"extend,omitempty"`
	U        *attrValString `xml:"u" json:"u,omitempty"`
	Sz       *attrValFloat  `xml:"sz" json:"sz,omitempty"`
	Color    *xlsxColor     `xml:"color" json:"color,omitempty"`
	Name     *attrValString `xml:"name" json:"name,omitempty"`
	Family   *attrValInt    `xml:"family" json:"family,omitempty"`
	Charset  *attrValInt    `xml:"charset" json:"charset,omitempty"`
	Scheme   *attrValString `xml:"scheme" json:"scheme,omitempty"`
}

// xlsxFills directly maps the fills element. This element defines the cell
// fills portion of the Styles part, consisting of a sequence of fill records. A
// cell fill consists of a background color, foreground color, and pattern to be
// applied across the cell.
type xlsxFills struct {
	Count int         `xml:"count,attr"`
	Fill  []*xlsxFill `xml:"fill"`
}

// xlsxFill directly maps the fill element. This element specifies fill
// formatting.
type xlsxFill struct {
	PatternFill  *xlsxPatternFill  `xml:"patternFill" json:"pattern_fill,omitempty"`
	GradientFill *xlsxGradientFill `xml:"gradientFill" json:"gradient_fill,omitempty"`
}

// xlsxPatternFill is used to specify cell fill information for pattern and
// solid color cell fills. For solid cell fills (no pattern), fgColor is used.
// For cell fills with patterns specified, then the cell fill color is
// specified by the bgColor element.
type xlsxPatternFill struct {
	PatternType string     `xml:"patternType,attr,omitempty" json:"pattern_type,omitempty"`
	FgColor     *xlsxColor `xml:"fgColor" json:"fg_color,omitempty"`
	BgColor     *xlsxColor `xml:"bgColor" json:"bg_color,omitempty"`
}

// xlsxGradientFill defines a gradient-style cell fill. Gradient cell fills can
// use one or two colors as the end points of color interpolation.
type xlsxGradientFill struct {
	Bottom float64                 `xml:"bottom,attr,omitempty"`
	Degree float64                 `xml:"degree,attr,omitempty"`
	Left   float64                 `xml:"left,attr,omitempty"`
	Right  float64                 `xml:"right,attr,omitempty"`
	Top    float64                 `xml:"top,attr,omitempty"`
	Type   string                  `xml:"type,attr,omitempty"`
	Stop   []*xlsxGradientFillStop `xml:"stop"`
}

// xlsxGradientFillStop directly maps the stop element.
type xlsxGradientFillStop struct {
	Position float64   `xml:"position,attr"`
	Color    xlsxColor `xml:"color,omitempty"`
}

// xlsxBorders directly maps the borders element. This element contains borders
// formatting information, specifying all border definitions for all cells in
// the workbook.
type xlsxBorders struct {
	Count  int           `xml:"count,attr"`
	Border []*xlsxBorder `xml:"border"`
}

// xlsxBorder directly maps the border element. Expresses a single set of cell
// border formats (left, right, top, bottom, diagonal). Color is optional. When
// missing, 'automatic' is implied.
type xlsxBorder struct {
	DiagonalDown bool     `xml:"diagonalDown,attr,omitempty" json:"diagonal_down,omitempty"`
	DiagonalUp   bool     `xml:"diagonalUp,attr,omitempty" json:"diagonal_up,omitempty"`
	Outline      bool     `xml:"outline,attr,omitempty" json:"outline,omitempty"`
	Left         xlsxLine `xml:"left,omitempty" json:"left"`
	Right        xlsxLine `xml:"right,omitempty" json:"right"`
	Top          xlsxLine `xml:"top,omitempty" json:"top"`
	Bottom       xlsxLine `xml:"bottom,omitempty" json:"bottom"`
	Diagonal     xlsxLine `xml:"diagonal,omitempty" json:"diagonal"`
}

// xlsxCellStyles directly maps the cellStyles element. This element contains
// the named cell styles, consisting of a sequence of named style records. A
// named cell style is a collection of direct or themed formatting (e.g., cell
// border, cell fill, and font type/size/style) grouped together into a single
// named style, and can be applied to a cell.
type xlsxCellStyles struct {
	XMLName   xml.Name         `xml:"cellStyles"`
	Count     int              `xml:"count,attr"`
	CellStyle []*xlsxCellStyle `xml:"cellStyle"`
}

// xlsxCellStyle directly maps the cellStyle element. This element represents
// the name and related formatting records for a named cell style in this
// workbook.
type xlsxCellStyle struct {
	XMLName       xml.Name `xml:"cellStyle"`
	Name          string   `xml:"name,attr"`
	XfID          int      `xml:"xfId,attr"`
	BuiltInID     *int     `xml:"builtinId,attr"`
	ILevel        *int     `xml:"iLevel,attr"`
	Hidden        *bool    `xml:"hidden,attr"`
	CustomBuiltIn *bool    `xml:"customBuiltin,attr"`
}

// xlsxCellStyleXfs directly maps the cellStyleXfs element. This element
// contains the master formatting records (xf's) which define the formatting for
// all named cell styles in this workbook. Master formatting records reference
// individual elements of formatting (e.g., number format, font definitions,
// cell fills, etc.) by specifying a zero-based index into those collections.
// Master formatting records also specify whether to apply or ignore particular
// aspects of formatting.
type xlsxCellStyleXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

// xlsxXf directly maps the xf element. A single xf element describes all of the
// formatting for a cell.
type xlsxXf struct {
	Lang              *string         `xml:"lang,attr" json:"lang,omitempty"`
	NumFmtID          *int            `xml:"numFmtId,attr" json:"num_fmt_id,omitempty"`
	FontID            *int            `xml:"fontId,attr" json:"font_id,omitempty"`
	FillID            *int            `xml:"fillId,attr" json:"fill_id,omitempty"`
	BorderID          *int            `xml:"borderId,attr" json:"border_id,omitempty"`
	XfID              *int            `xml:"xfId,attr" json:"xf_id,omitempty"`
	QuotePrefix       *bool           `xml:"quotePrefix,attr" json:"quote_prefix,omitempty"`
	PivotButton       *bool           `xml:"pivotButton,attr" json:"pivot_button,omitempty"`
	ApplyNumberFormat *bool           `xml:"applyNumberFormat,attr" json:"apply_number_format,omitempty"`
	ApplyFont         *bool           `xml:"applyFont,attr" json:"apply_font,omitempty"`
	ApplyFill         *bool           `xml:"applyFill,attr" json:"apply_fill,omitempty"`
	ApplyBorder       *bool           `xml:"applyBorder,attr" json:"apply_border,omitempty"`
	ApplyAlignment    *bool           `xml:"applyAlignment,attr" json:"apply_alignment,omitempty"`
	ApplyProtection   *bool           `xml:"applyProtection,attr" json:"apply_protection,omitempty"`
	Alignment         *xlsxAlignment  `xml:"alignment" json:"alignment,omitempty"`
	Protection        *xlsxProtection `xml:"protection" json:"protection,omitempty"`
}

// xlsxCellXfs directly maps the cellXfs element. This element contains the
// master formatting records (xf) which define the formatting applied to cells
// in this workbook. These records are the starting point for determining the
// formatting for a cell. Cells in the Sheet Part reference the xf records by
// zero-based index.
type xlsxCellXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

// xlsxDxfs directly maps the dxfs element. This element contains the master
// differential formatting records (dxf's) which define formatting for all non-
// cell formatting in this workbook. Whereas xf records fully specify a
// particular aspect of formatting (e.g., cell borders) by referencing those
// formatting definitions elsewhere in the Styles part, dxf records specify
// incremental (or differential) aspects of formatting directly inline within
// the dxf element. The dxf formatting is to be applied on top of or in addition
// to any formatting already present on the object using the dxf record.
type xlsxDxfs struct {
	Count int        `xml:"count,attr"`
	Dxfs  []*xlsxDxf `xml:"dxf"`
}

// xlsxDxf directly maps the dxf element. A single dxf record, expressing
// incremental formatting to be applied.
type xlsxDxf struct {
	Dxf string `xml:",innerxml"`
}

// dxf directly maps the dxf element.
type dxf struct {
	Font       *xlsxFont       `xml:"font"`
	NumFmt     *xlsxNumFmt     `xml:"numFmt"`
	Fill       *xlsxFill       `xml:"fill"`
	Alignment  *xlsxAlignment  `xml:"alignment"`
	Border     *xlsxBorder     `xml:"border"`
	Protection *xlsxProtection `xml:"protection"`
	ExtLst     *xlsxExt        `xml:"extLst"`
}

// xlsxTableStyles directly maps the tableStyles element. This element
// represents a collection of Table style definitions for Table styles and
// PivotTable styles used in this workbook. It consists of a sequence of
// tableStyle records, each defining a single Table style.
type xlsxTableStyles struct {
	Count             int               `xml:"count,attr"`
	DefaultPivotStyle string            `xml:"defaultPivotStyle,attr"`
	DefaultTableStyle string            `xml:"defaultTableStyle,attr"`
	TableStyles       []*xlsxTableStyle `xml:"tableStyle"`
}

// xlsxTableStyle directly maps the tableStyle element. This element represents
// a single table style definition that indicates how a spreadsheet application
// should format and display a table.
type xlsxTableStyle struct {
	Name              string `xml:"name,attr,omitempty"`
	Pivot             int    `xml:"pivot,attr"`
	Count             int    `xml:"count,attr,omitempty"`
	Table             bool   `xml:"table,attr,omitempty"`
	TableStyleElement string `xml:",innerxml"`
}

// xlsxNumFmts directly maps the numFmts element. This element defines the
// number formats in this workbook, consisting of a sequence of numFmt records,
// where each numFmt record defines a particular number format, indicating how
// to format and render the numeric value of a cell.
type xlsxNumFmts struct {
	Count  int           `xml:"count,attr"`
	NumFmt []*xlsxNumFmt `xml:"numFmt"`
}

// xlsxNumFmt directly maps the numFmt element. This element specifies number
// format properties which indicate how to format and render the numeric value
// of a cell.
type xlsxNumFmt struct {
	NumFmtID   int    `xml:"numFmtId,attr" json:"num_fmt_id,omitempty"`
	FormatCode string `xml:"formatCode,attr,omitempty" json:"format_code,omitempty"`
}

// xlsxStyleColors directly maps the colors element. Color information
// associated with this stylesheet. This collection is written whenever the
// legacy color palette has been modified (backwards compatibility settings) or
// a custom color has been selected while using this workbook.
type xlsxStyleColors struct {
	Color string `xml:",innerxml"`
}

// Alignment directly maps the alignment settings of the cells.
type Alignment struct {
	Horizontal      string `json:"horizontal"`
	Indent          int    `json:"indent"`
	JustifyLastLine bool   `json:"justify_last_line"`
	ReadingOrder    uint64 `json:"reading_order"`
	RelativeIndent  int    `json:"relative_indent"`
	ShrinkToFit     bool   `json:"shrink_to_fit"`
	TextRotation    int    `json:"text_rotation"`
	Vertical        string `json:"vertical"`
	WrapText        bool   `json:"wrap_text"`
}

// Border directly maps the border settings of the cells.
type Border struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}

// Font directly maps the font settings of the fonts.
type Font struct {
	Bold      bool    `json:"bold"`
	Italic    bool    `json:"italic"`
	Underline string  `json:"underline"`
	Family    string  `json:"family"`
	Size      float64 `json:"size"`
	Strike    bool    `json:"strike"`
	Color     string  `json:"color"`
	VertAlign string  `json:"vertAlign"`
}

// Fill directly maps the fill settings of the cells.
type Fill struct {
	Type    string   `json:"type"`
	Pattern int      `json:"pattern"`
	Color   []string `json:"color"`
	Shading int      `json:"shading"`
}

// Protection directly maps the protection settings of the cells.
type Protection struct {
	Hidden bool `json:"hidden"`
	Locked bool `json:"locked"`
}

// Style directly maps the style settings of the cells.
type Style struct {
	Border        []Border    `json:"border"`
	Fill          Fill        `json:"fill"`
	Font          *Font       `json:"font"`
	Alignment     *Alignment  `json:"alignment"`
	Protection    *Protection `json:"protection"`
	NumFmt        int         `json:"number_format"`
	DecimalPlaces int         `json:"decimal_places"`
	CustomNumFmt  *string     `json:"custom_number_format"`
	Lang          string      `json:"lang"`
	NegRed        bool        `json:"negred"`
}

type StyleOutput struct {
	Lang       string          `json:"lang,omitempty"`
	NumFmt     *xlsxNumFmt     `json:"num_fmt,omitempty"`
	Font       *xlsxFont       `json:"fonts,omitempty"`
	Border     *xlsxBorder     `json:"borders,omitempty"`
	Fill       *xlsxFill       `json:"fills,omitempty"`
	Alignment  *xlsxAlignment  `json:"alignment,omitempty"`
	Protection *xlsxProtection `json:"protection,omitempty"`
}
