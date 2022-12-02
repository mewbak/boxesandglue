package frontend

import (
	"fmt"
	"strings"

	"github.com/speedata/boxesandglue/backend/bag"
	"github.com/speedata/boxesandglue/backend/color"
	"github.com/speedata/boxesandglue/backend/document"
	"github.com/speedata/boxesandglue/backend/font"
	"github.com/speedata/boxesandglue/backend/lang"
	"github.com/speedata/boxesandglue/backend/node"
	"github.com/speedata/boxesandglue/pdfbackend/pdf"
	"github.com/speedata/textlayout/harfbuzz"
)

// SettingType represents a setting such as font weight or color.
type SettingType int

// FontWeight is the type which represents different font weights.
type FontWeight int

func (fw FontWeight) String() string {
	switch fw {
	case 100:
		return "Thin"
	case 200:
		return "Extra Light"
	case 300:
		return "Light"
	case 400:
		return "Normal"
	case 500:
		return "Medium"
	case 600:
		return "SemiBold"
	case 700:
		return "Bold"
	case 800:
		return "Ultra Bold"
	case 900:
		return "Black"
	default:
		return fmt.Sprintf("fontweight %d", fw)
	}
}

const (
	// FontWeight100 is commonly named “Thin”.
	FontWeight100 FontWeight = 100
	// FontWeight200 is commonly named “Extra Light”.
	FontWeight200 FontWeight = 200
	// FontWeight300 is commonly named “Light”.
	FontWeight300 FontWeight = 300
	// FontWeight400 is commonly named “Normal”.
	FontWeight400 FontWeight = 400
	// FontWeight500 is commonly named “Medium”.
	FontWeight500 FontWeight = 500
	// FontWeight600 is commonly named “Semi Bold”.
	FontWeight600 FontWeight = 600
	// FontWeight700 is commonly named Bold”.
	FontWeight700 FontWeight = 700
	// FontWeight800 is commonly named “Ultra Bold”.
	FontWeight800 FontWeight = 800
	// FontWeight900 is commonly named “Black”.
	FontWeight900 FontWeight = 900
)

// FontStyle is the type which represents different font styles such as italic or oblique.
type FontStyle int

func (fs FontStyle) String() string {
	switch fs {
	case FontStyleNormal:
		return "normal"
	case FontStyleItalic:
		return "italic"
	case FontStyleOblique:
		return "oblique"
	default:
		return "???"
	}
}

const (
	// FontStyleNormal is an upright font.
	FontStyleNormal FontStyle = iota
	// FontStyleItalic is an italicized font.
	FontStyleItalic
	// FontStyleOblique is an upright font tilted by an angle.
	FontStyleOblique
)

const (
	// SettingDummy is a no op.
	SettingDummy SettingType = iota
	// SettingColor sets a predefined color.
	SettingColor
	// SettingFontFamily selects a font family.
	SettingFontFamily
	// SettingFontWeight represents a font weight setting.
	SettingFontWeight
	// SettingHAlign sets the horizontal alignment of the paragraph.
	SettingHAlign
	// SettingHyperlink defines an external hyperlink.
	SettingHyperlink
	// SettingIndentLeft inserts a left margin
	SettingIndentLeft
	// SettingIndentLeftRows determines the number of rows to be indented (positive value), or the number of rows not indented (negative values). 0 means all rows.
	SettingIndentLeftRows
	// SettingLeading determines the distance between two base lines (line height).
	SettingLeading
	// SettingMarginBottom sets the bottom margin.
	SettingMarginBottom
	// SettingMarginLeft sets the left margin.
	SettingMarginLeft
	// SettingMarginRight sets the right margin.
	SettingMarginRight
	// SettingMarginTop sets the top margin.
	SettingMarginTop
	// SettingOpenTypeFeature allows the user to (de)select OpenType features such as ligatures.
	SettingOpenTypeFeature
	// SettingPreserveWhitespace makes a monospace paragraph with newlines.
	SettingPreserveWhitespace
	// SettingSize sets the font size.
	SettingSize
	// SettingStyle represents a font style such as italic or normal.
	SettingStyle
	// SettingYOffset shifts the glyph.
	SettingYOffset
)

func (st SettingType) String() string {
	var settingName string
	switch st {
	case SettingColor:
		settingName = "SettingColor"
	case SettingFontFamily:
		settingName = "SettingFontFamily"
	case SettingFontWeight:
		settingName = "SettingFontWeight"
	case SettingHAlign:
		settingName = "SettingHAlign"
	case SettingHyperlink:
		settingName = "SettingHyperlink"
	case SettingIndentLeft:
		settingName = "SettingIndentLeft"
	case SettingIndentLeftRows:
		settingName = "SettingIndentLeftRows"
	case SettingLeading:
		settingName = "SettingLeading"
	case SettingMarginBottom:
		settingName = "SettingMarginBottom"
	case SettingMarginLeft:
		settingName = "SettingMarginLeft"
	case SettingMarginRight:
		settingName = "SettingMarginRight"
	case SettingMarginTop:
		settingName = "SettingMarginTop"
	case SettingOpenTypeFeature:
		settingName = "SettingOpenTypeFeature"
	case SettingPreserveWhitespace:
		settingName = "SettingPreserveWhitespace"
	case SettingSize:
		settingName = "SettingSize"
	case SettingStyle:
		settingName = "SettingStyle"
	case SettingYOffset:
		settingName = "SettingYOffset"
	default:
		settingName = fmt.Sprintf("%d", st)
	}
	return fmt.Sprintf("%s", settingName)
}

// TypesettingSettings is a set of settings for text rendering.
type TypesettingSettings map[SettingType]any

// Text associates all items with the given settings. Items can be text
// (string), images, other instances of Text or nodes. Text behaves like a span
// in HTML or it just contains a collection of Go strings.
type Text struct {
	Settings TypesettingSettings
	Items    []any
}

// NewText returns an initialized text element.
func NewText() *Text {
	te := Text{}
	te.Settings = make(TypesettingSettings)
	return &te
}

func (ts *Text) String() string {
	ret := []string{}
	ret = append(ret, "Settings:")
	for k, v := range ts.Settings {
		ret = append(ret, fmt.Sprintf("%s:%v", k, v))
	}
	ret = append(ret, fmt.Sprintf("\nitems(len %d)", len(ts.Items)))
	for _, itm := range ts.Items {
		ret = append(ret, fmt.Sprintf("%s", itm))
	}
	return strings.Join(ret, " ")
}

type paragraph struct {
	alignment      HorizontalAlignment
	fontfamily     *FontFamily
	fontsize       bag.ScaledPoint
	hsize          bag.ScaledPoint
	indentLeft     bag.ScaledPoint
	indentLeftRows int
	language       *lang.Lang
	leading        bag.ScaledPoint
}

// TypesettingOption controls the formatting of the paragraph.
type TypesettingOption func(*paragraph)

// Leading sets the distance between two baselines in a paragraph.
func Leading(leading bag.ScaledPoint) TypesettingOption {
	return func(p *paragraph) {
		p.leading = leading
	}
}

// Language sets the default language for the whole paragraph (used for hyphenation).
func Language(language *lang.Lang) TypesettingOption {
	return func(p *paragraph) {
		p.language = language
	}
}

// FontSize sets the font size for the paragraph.
func FontSize(size bag.ScaledPoint) TypesettingOption {
	return func(p *paragraph) {
		p.fontsize = size
	}
}

// IndentLeft sets the left indent.
func IndentLeft(size bag.ScaledPoint, rows int) TypesettingOption {
	return func(p *paragraph) {
		p.indentLeft = size
		p.indentLeftRows = rows
	}
}

// Family sets the font family for the paragraph.
func Family(fam *FontFamily) TypesettingOption {
	return func(p *paragraph) {
		p.fontfamily = fam
	}
}

// HorizontalAlign sets the horizontal alignment for a paragraph.
func HorizontalAlign(a HorizontalAlignment) TypesettingOption {
	return func(p *paragraph) {
		p.alignment = a
	}
}

// FormatParagraph creates a rectangular text from the data stored in the
// Paragraph.
func (fe *Document) FormatParagraph(te *Text, hsize bag.ScaledPoint, opts ...TypesettingOption) (*node.VList, []*node.Breakpoint, error) {
	if len(te.Items) == 0 {
		g := node.NewGlue()
		g.Attributes = node.H{"origin": "empty list in FormatParagraph"}
		return node.Vpack(g), nil, nil
	}
	p := &paragraph{
		language: fe.Doc.DefaultLanguage,
		hsize:    hsize,
	}
	if ha, ok := te.Settings[SettingHAlign]; ok {
		if ha != nil {
			p.alignment = ha.(HorizontalAlignment)
		} else {
			p.alignment = HAlignDefault
		}
	}
	for _, opt := range opts {
		opt(p)
	}
	if p.fontsize != 0 {
		te.Settings[SettingSize] = p.fontsize
	}
	if p.fontfamily != nil {
		te.Settings[SettingFontFamily] = p.fontfamily
	}
	hlist, tail, err := fe.Mknodes(te)
	if err != nil {
		return nil, nil, err
	}

	Hyphenate(hlist, p.language)
	node.AppendLineEndAfter(hlist, tail)

	ls := node.NewLinebreakSettings()
	ls.HSize = p.hsize
	ls.Indent = p.indentLeft
	ls.IndentRows = p.indentLeftRows

	if p.leading == 0 {
		if l, ok := te.Settings[SettingLeading]; ok {
			ls.LineHeight = l.(bag.ScaledPoint)
		} else {
			ls.LineHeight = p.fontsize * 120 / 100
		}
	} else {
		ls.LineHeight = p.leading
	}
	if p.alignment == HAlignLeft || p.alignment == HAlignCenter {
		lg := node.NewGlue()
		lg.Stretch = bag.Factor
		lg.StretchOrder = 3
		lg.Subtype = node.GlueLineEnd
		ls.LineEndGlue = lg
	}
	if p.alignment == HAlignRight || p.alignment == HAlignCenter {
		lg := node.NewGlue()
		lg.Stretch = bag.Factor
		lg.StretchOrder = 3
		lg.Subtype = node.GlueLineStart
		ls.LineStartGlue = lg
	}
	vlist, info := node.Linebreak(hlist, ls)
	return vlist, info, nil
}

func parseHarfbuzzFontFeatures(featurelist any) []harfbuzz.Feature {
	fontfeatures := []harfbuzz.Feature{}
	switch t := featurelist.(type) {
	case string:
		for _, str := range strings.Split(t, ",") {
			f, err := harfbuzz.ParseFeature(str)
			if err != nil {
				bag.Logger.Errorf("cannot parse OpenType feature tag %q.", str)
			}
			fontfeatures = append(fontfeatures, f)
		}
	case []string:
		for _, single := range t {
			for _, str := range strings.Split(single, ",") {
				f, err := harfbuzz.ParseFeature(str)
				if err != nil {
					bag.Logger.Errorf("cannot parse OpenType feature tag %q.", str)
				}
				fontfeatures = append(fontfeatures, f)
			}

		}
	}
	return fontfeatures
}

func (fe *Document) BuildNodelistFromString(ts TypesettingSettings, str string) (node.Node, error) {
	fontweight := FontWeight400
	fontstyle := FontStyleNormal
	var fontfamily *FontFamily
	fontsize := 12 * bag.Factor
	var col *color.Color
	var hyperlink document.Hyperlink
	var hasHyperlink bool
	fontfeatures := make([]harfbuzz.Feature, 0, len(fe.DefaultFeatures))
	for _, f := range fe.DefaultFeatures {
		fontfeatures = append(fontfeatures, f)
	}
	preserveWhitespace := false
	yoffset := bag.ScaledPoint(0)
	var settingFontFeatures []harfbuzz.Feature
	for k, v := range ts {
		switch k {
		case SettingFontWeight:
			switch t := v.(type) {
			case int:
				fontweight = FontWeight(t)
			case FontWeight:
				fontweight = t
			}
		case SettingFontFamily:
			fontfamily = v.(*FontFamily)
		case SettingSize:
			fontsize = v.(bag.ScaledPoint)
		case SettingColor:
			switch t := v.(type) {
			case string:
				if c := fe.GetColor(t); c != nil {
					col = c
				}
			case *color.Color:
				col = t
			}
		case SettingHyperlink:
			hyperlink = v.(document.Hyperlink)
			hasHyperlink = true
		case SettingStyle:
			fontstyle = v.(FontStyle)
		case SettingOpenTypeFeature:
			settingFontFeatures = parseHarfbuzzFontFeatures(v)
		case SettingMarginTop, SettingMarginRight, SettingMarginBottom, SettingMarginLeft:
			// ignore
		case SettingHAlign, SettingLeading, SettingIndentLeft, SettingIndentLeftRows:
			// ignore
		case SettingPreserveWhitespace:
			preserveWhitespace = v.(bool)
		case SettingYOffset:
			yoffset = v.(bag.ScaledPoint)
		default:
			bag.Logger.DPanicf("Unknown setting %v", k)
		}
	}

	var fnt *font.Font
	var face *pdf.Face
	var fs *FontSource
	var err error
	if fs, err = fontfamily.GetFontSource(fontweight, fontstyle); err != nil {
		return nil, err
	}
	// fs.SizeAdjust is CSS size-adjust normalized so that 0 = 100% and negative = shrinking.
	if fs.SizeAdjust != 0 {
		fontsize = bag.ScaledPointFromFloat(fontsize.ToPT() * (1 - fs.SizeAdjust))
	}
	// First the font source default features should get applied, then the
	// features from the current settings.
	fontfeatures = append(fontfeatures, parseHarfbuzzFontFeatures(fs.FontFeatures)...)
	fontfeatures = append(fontfeatures, settingFontFeatures...)

	if face, err = fe.LoadFace(fs); err != nil {
		return nil, err
	}
	if fe.usedFonts[face] == nil {
		fe.usedFonts = make(map[*pdf.Face]map[bag.ScaledPoint]*font.Font)
	}
	if fe.usedFonts[face][fontsize] == nil {
		fe.usedFonts[face] = make(map[bag.ScaledPoint]*font.Font)
	}

	var found bool
	if fnt, found = fe.usedFonts[face][fontsize]; !found {
		fnt = fe.Doc.CreateFont(face, fontsize)
		fe.usedFonts[face][fontsize] = fnt
	}

	var head, cur node.Node
	var hyperlinkStart, hyperlinkStop *node.StartStop
	if hasHyperlink {
		hyperlinkStart = node.NewStartStop()
		hyperlinkStart.Action = node.ActionHyperlink
		hyperlinkStart.Value = &hyperlink
		if head != nil {
			head = node.InsertAfter(head, head, hyperlinkStart)
		}
		head = hyperlinkStart
	}

	if col != nil {
		colStart := node.NewStartStop()
		colStart.Position = node.PDFOutputPage
		colStart.Callback = func(n node.Node) string {
			return col.PDFStringNonStroking() + " "
		}
		if head != nil {
			head = node.InsertAfter(head, head, colStart)
		}
		head = colStart
	}
	cur = head
	var lastglue node.Node
	atoms := fnt.Shape(str, fontfeatures)
	for _, r := range atoms {
		if r.IsSpace {
			if preserveWhitespace {
				switch r.Components {
				case " ":
					g := node.NewRule()
					g.Width = fnt.Space
					head = node.InsertAfter(head, cur, g)
					cur = g
					lastglue = g
				case "\t":
					// tab size...
					for i := 0; i < 4; i++ {
						g := node.NewRule()
						g.Width = fnt.Space
						head = node.InsertAfter(head, cur, g)
						cur = g
						lastglue = g
					}
				case "\n":
					head, cur = node.AppendLineEndAfter(head, cur)
					lastglue = cur
				default:
					panic("unhandled whitespace type")
				}
			} else {
				if lastglue == nil {
					g := node.NewGlue()
					g.Width = fnt.Space
					g.Stretch = fnt.SpaceStretch
					g.Shrink = fnt.SpaceShrink
					head = node.InsertAfter(head, cur, g)
					cur = g
					lastglue = g
				}
			}
		} else {
			n := node.NewGlyph()
			n.Hyphenate = r.Hyphenate
			n.Codepoint = r.Codepoint
			n.Components = r.Components
			n.Font = fnt
			n.Width = r.Advance
			n.Height = r.Height
			n.Depth = r.Depth
			n.YOffset = yoffset
			head = node.InsertAfter(head, cur, n)
			cur = n
			lastglue = nil

			if r.Kernafter != 0 {
				k := node.NewKern()
				k.Kern = r.Kernafter
				head = node.InsertAfter(head, cur, k)
				cur = k
			}
		}
	}
	if col != nil {
		stop := node.NewStartStop()
		stop.Position = node.PDFOutputPage
		stop.Callback = func(n node.Node) string {
			return "0 0 0 RG 0 0 0 rg "
		}
		node.InsertAfter(head, cur, stop)
		cur = stop
	}
	if hasHyperlink {
		hyperlinkStop = node.NewStartStop()
		hyperlinkStop.StartNode = hyperlinkStart
		head = node.InsertAfter(head, cur, hyperlinkStop)
		cur = hyperlinkStop
	}

	return head, nil
}

// Mknodes creates a list of nodes which which can be formatted to a given
// width. The returned head and the tail are the beginning and the end of the
// node list.
func (fe *Document) Mknodes(ts *Text) (head node.Node, tail node.Node, err error) {
	if len(ts.Items) == 0 {
		return nil, nil, nil
	}
	var newSettings = make(TypesettingSettings)
	var nl, end node.Node
	for k, v := range ts.Settings {
		newSettings[k] = v
	}
	var hyperlinkStartNode *node.StartStop
	var hyperlinkDest string
	for _, itm := range ts.Items {
		switch t := itm.(type) {
		case string:
			if hyperlinkStartNode != nil {
				endHL := node.NewStartStop()
				endHL.Action = node.ActionNone
				endHL.StartNode = hyperlinkStartNode
				hyperlinkStartNode = nil
				node.InsertAfter(head, tail, endHL)
				tail = endHL
			}

			nl, err = fe.BuildNodelistFromString(newSettings, t)
			if err != nil {
				return nil, nil, err
			}
			if nl != nil {
				head = node.InsertAfter(head, tail, nl)
				tail = node.Tail(nl)
			}
		case *Text:
			if hyperlinkStartNode == nil {
				// we are within a hyperlink, so lets remove all startstop
				if hlSetting, ok := t.Settings[SettingHyperlink]; ok {
					hl := hlSetting.(document.Hyperlink)
					// insert a startstop with the hyperlink
					hyperlinkDest = hl.URI
					startHL := node.NewStartStop()
					startHL.Action = node.ActionHyperlink
					startHL.Value = &hl
					hyperlinkStartNode = startHL
					head = node.InsertAfter(head, tail, startHL)
					tail = startHL
				}
			} else {
				if hlSetting, ok := t.Settings[SettingHyperlink]; ok && hlSetting.(document.Hyperlink).URI == hyperlinkDest {
					// same destination
				} else {
					// probably no hyperlink, TODO: insert end startstop here?
				}
			}
			// copy current settings to the child if not already set.
			for k, v := range newSettings {
				if _, found := t.Settings[k]; !found {
					t.Settings[k] = v
				}
			}
			// we don't want to inherit hyperlinks
			delete(t.Settings, SettingHyperlink)
			nl, end, err = fe.Mknodes(t)
			if err != nil {
				return nil, nil, err
			}
			if nl != nil {
				head = node.InsertAfter(head, tail, nl)
				tail = end
			}
		case node.Node:
			head = node.InsertAfter(head, tail, t)
			tail = t
		default:
			bag.Logger.DPanicf("Mknodes: unknown item type %T", t)
		}
	}
	if hyperlinkStartNode != nil {
		endHL := node.NewStartStop()
		endHL.Action = node.ActionNone
		endHL.StartNode = hyperlinkStartNode
		node.InsertAfter(head, tail, endHL)
		tail = endHL
	}
	return head, tail, nil
}
