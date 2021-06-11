package node

import (
	"container/list"

	"github.com/speedata/texperiments/lang"
)

var (
	ids chan int
)

func genIntegerSequence(ids chan int) {
	i := int(0)
	for {
		ids <- i
		i++
	}
}

func init() {
	ids = make(chan int)
	go genIntegerSequence(ids)
}

type basenode struct {
	id int
}

// NewElement creates a list element from the node type. You must ensure that the val is a
// valid node type.
func NewElement(val interface{}) *list.Element {
	return &list.Element{Value: val}
}

// A Disc is a hyphenation point.
type Disc struct {
	basenode
}

// NewDisc creates an initialized Disc node
func NewDisc() *Disc {
	n := &Disc{}
	n.id = <-ids
	return n
}

// NewDiscWithContents creates an initialized Disc node with the given contents
func NewDiscWithContents(n *Disc) *Disc {
	n.id = <-ids
	return n
}

// IsDisc retuns the value of the element and true, if the element is a Disc node.
func IsDisc(elt *list.Element) (*Disc, bool) {
	Disc, ok := elt.Value.(*Disc)
	return Disc, ok
}

// Glyph nodes represents a single visible entity such as a letter or a ligature.
type Glyph struct {
	basenode
	GlyphID    int    // The font specific glyph id
	Components string // A codepoint can contain more than one rune, for example a fi ligature contains f + i
	Hyphenate  bool
}

// NewGlyph returns an initialized Glyph
func NewGlyph() *Glyph {
	n := &Glyph{}
	n.id = <-ids
	return n
}

// IsGlyph returns the value of the element and true, if the element is a Glyph node.
func IsGlyph(elt *list.Element) (*Glyph, bool) {
	n, ok := elt.Value.(*Glyph)
	return n, ok
}

// A Glue node has the value of a shrinking and stretching space
type Glue struct {
	basenode
	Width        float64
	Stretch      int
	Shrink       int
	StretchOrder int
	ShrinkOrder  int
}

// NewGlue creates an initialized Glue node
func NewGlue() *Glue {
	n := &Glue{}
	n.id = <-ids
	return n
}

// IsGlue retuns the value of the element and true, if the element is a Glue node.
func IsGlue(elt *list.Element) (*Glue, bool) {
	n, ok := elt.Value.(*Glue)
	return n, ok
}

// A HList is a horizontal list.
type HList struct {
	basenode
	List *list.Element
}

// NewHList creates an initialized HList node
func NewHList() *HList {
	n := &HList{}
	n.id = <-ids
	return n
}

// IsHList retuns the value of the element and true, if the element is a HList node.
func IsHList(elt *list.Element) (*HList, bool) {
	hlist, ok := elt.Value.(*HList)
	return hlist, ok
}

// A Lang is a node that sets the current language.
type Lang struct {
	basenode
	Lang *lang.Lang
}

// NewLang creates an initialized Lang node
func NewLang() *Lang {
	n := &Lang{}
	n.id = <-ids
	return n
}

// NewLangWithContents creates an initialized Lang node with the given contents
func NewLangWithContents(n *Lang) *Lang {
	n.id = <-ids
	return n
}

// IsLang retuns the value of the element and true, if the element is a Lang node.
func IsLang(elt *list.Element) (*Lang, bool) {
	lang, ok := elt.Value.(*Lang)
	return lang, ok
}