// +build js, wasm

package dom

import (
	"syscall/js"

	"github.com/josephbudd/kickwasm/examples/spawnwidgets/rendererprocess/api/markup"
)

var (
	document js.Value
)

const (
	createElementMethodName = "createElement"

	inputTypeName      = "input"
	checkboxTypeName   = "checkbox"
	radioTypeName      = "radio"
	typeAttributeName  = "type"
	groupAttributeName = "group"
)

func init() {
	g := js.Global()
	document = g.Get("document")
}

type DOM struct {
	panelUniqueID uint64
}

// NewDOM constructs a new DOM.
func NewDOM(panelUniqueID uint64) (d *DOM) {
	d = &DOM{
		panelUniqueID: panelUniqueID,
	}
	return
}

// ElementByID finds an element in the document with a matching id.
// Returns the element or nil.
func (d *DOM) ElementByID(id string) (el *markup.Element) {
	e := document.Call("getElementById", id)
	if e == js.Null() {
		return
	}
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// ElementsByTagName finds all the elements in the document with a matching tag name.
func (d *DOM) ElementsByTagName(tagName string) (els []*markup.Element) {
	ee := document.Call("getElementsByTagName", tagName)
	l := ee.Length()
	els = make([]*markup.Element, l)
	for i := 0; i < l; i++ {
		els[i] = markup.NewElement(ee.Index(i), d.panelUniqueID)
	}
	return
}

// NewElementFromJSValue creates a new markup element using the js.Value.
func (d *DOM) NewElementFromJSValue(jsv js.Value) (el *markup.Element) {
	el = markup.NewElement(jsv, d.panelUniqueID)
	return
}

// New creates an element when no such func exists here. Ex: document.New("myelement")
func (d *DOM) New(tagName string) (el *markup.Element) {
	e := document.Call(createElementMethodName, tagName)
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// Constructors for HTML 5 Elements.

// NewText creates a text node.
func (d *DOM) NewText(text string) (el *markup.Element) {
	e := document.Call("createTextNode", text)
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// These are shortcuts for input types.

// NewINPUT creates a text input element.
func (d *DOM) NewINPUT() (el *markup.Element) {
	e := document.Call(createElementMethodName, inputTypeName)
	e.Set(typeAttributeName, "text")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCheckBoxInGroup creates a checkbox input element.
func (d *DOM) NewCheckBoxInGroup(group string) (el *markup.Element) {
	e := document.Call(createElementMethodName, inputTypeName)
	e.Set(typeAttributeName, checkboxTypeName)
	if len(group) > 0 {
		e.Set(groupAttributeName, group)
	}
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewRadioInGroup creates a radio input element
func (d *DOM) NewRadioInGroup(group string) (el *markup.Element) {
	e := document.Call(createElementMethodName, inputTypeName)
	e.Set(typeAttributeName, radioTypeName)
	e.Set(groupAttributeName, group)
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCheckBox creates a checkbox input element.
func (d *DOM) NewCheckBox() (el *markup.Element) {
	e := document.Call(createElementMethodName, inputTypeName)
	e.Set(typeAttributeName, checkboxTypeName)
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewRadio creates a radio input element
func (d *DOM) NewRadio() (el *markup.Element) {
	e := document.Call(createElementMethodName, inputTypeName)
	e.Set(typeAttributeName, radioTypeName)
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// Tag specific constructors.

// NewA creates a new HTML A element.
func (d *DOM) NewA() (el *markup.Element) {
	e := document.Call(createElementMethodName, "a")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewABBR creates a new HTML ABBR element.
func (d *DOM) NewABBR() (el *markup.Element) {
	e := document.Call(createElementMethodName, "abbr")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewADDRESS creates a new HTML ADDRESS element.
func (d *DOM) NewADDRESS() (el *markup.Element) {
	e := document.Call(createElementMethodName, "address")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewAREA creates a new HTML AREA element.
func (d *DOM) NewAREA() (el *markup.Element) {
	e := document.Call(createElementMethodName, "area")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewARTICLE creates a new HTML ARTICLE element.
func (d *DOM) NewARTICLE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "article")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewASIDE creates a new HTML ASIDE element.
func (d *DOM) NewASIDE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "aside")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewAUDIO creates a new HTML AUDIO element.
func (d *DOM) NewAUDIO() (el *markup.Element) {
	e := document.Call(createElementMethodName, "audio")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewB creates a new HTML B element.
func (d *DOM) NewB() (el *markup.Element) {
	e := document.Call(createElementMethodName, "b")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewBDI creates a new HTML BDI element.
func (d *DOM) NewBDI() (el *markup.Element) {
	e := document.Call(createElementMethodName, "bdi")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewBDO creates a new HTML BDO element.
func (d *DOM) NewBDO() (el *markup.Element) {
	e := document.Call(createElementMethodName, "bdo")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewBLOCKQUOTE creates a new HTML BLOCKQUOTE element.
func (d *DOM) NewBLOCKQUOTE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "blockquote")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewBODY creates a new HTML BODY element.
func (d *DOM) NewBODY() (el *markup.Element) {
	e := document.Call(createElementMethodName, "body")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewBR creates a new HTML BR element.
func (d *DOM) NewBR() (el *markup.Element) {
	e := document.Call(createElementMethodName, "br")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewBUTTON creates a new HTML BUTTON element.
func (d *DOM) NewBUTTON() (el *markup.Element) {
	e := document.Call(createElementMethodName, "button")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCANVAS creates a new HTML CANVAS element.
func (d *DOM) NewCANVAS() (el *markup.Element) {
	e := document.Call(createElementMethodName, "canvas")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCAPTION creates a new HTML CAPTION element.
func (d *DOM) NewCAPTION() (el *markup.Element) {
	e := document.Call(createElementMethodName, "caption")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCITE creates a new HTML CITE element.
func (d *DOM) NewCITE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "cite")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCODE creates a new HTML CODE element.
func (d *DOM) NewCODE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "code")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCOL creates a new HTML COL element.
func (d *DOM) NewCOL() (el *markup.Element) {
	e := document.Call(createElementMethodName, "col")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewCOLGROUP creates a new HTML COLGROUP element.
func (d *DOM) NewCOLGROUP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "colgroup")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDATA creates a new HTML DATA element.
func (d *DOM) NewDATA() (el *markup.Element) {
	e := document.Call(createElementMethodName, "data")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDATALIST creates a new HTML DATALIST element.
func (d *DOM) NewDATALIST() (el *markup.Element) {
	e := document.Call(createElementMethodName, "datalist")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDD creates a new HTML DD element.
func (d *DOM) NewDD() (el *markup.Element) {
	e := document.Call(createElementMethodName, "dd")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDEL creates a new HTML DEL element.
func (d *DOM) NewDEL() (el *markup.Element) {
	e := document.Call(createElementMethodName, "del")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDETAILS creates a new HTML DETAILS element.
func (d *DOM) NewDETAILS() (el *markup.Element) {
	e := document.Call(createElementMethodName, "details")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDFN creates a new HTML DFN element.
func (d *DOM) NewDFN() (el *markup.Element) {
	e := document.Call(createElementMethodName, "dfn")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDIALOG creates a new HTML DIALOG element.
func (d *DOM) NewDIALOG() (el *markup.Element) {
	e := document.Call(createElementMethodName, "dialog")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDIV creates a new HTML DIV element.
func (d *DOM) NewDIV() (el *markup.Element) {
	e := document.Call(createElementMethodName, "div")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDL creates a new HTML DL element.
func (d *DOM) NewDL() (el *markup.Element) {
	e := document.Call(createElementMethodName, "dl")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewDT creates a new HTML DT element.
func (d *DOM) NewDT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "dt")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewELEMENT creates a new HTML ELEMENT element.
func (d *DOM) NewELEMENT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "element")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewEM creates a new HTML EM element.
func (d *DOM) NewEM() (el *markup.Element) {
	e := document.Call(createElementMethodName, "em")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewEMBED creates a new HTML EMBED element.
func (d *DOM) NewEMBED() (el *markup.Element) {
	e := document.Call(createElementMethodName, "embed")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewFIELDSET creates a new HTML FIELDSET element.
func (d *DOM) NewFIELDSET() (el *markup.Element) {
	e := document.Call(createElementMethodName, "fieldset")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewFIGCAPTION creates a new HTML FIGCAPTION element.
func (d *DOM) NewFIGCAPTION() (el *markup.Element) {
	e := document.Call(createElementMethodName, "figcaption")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewFIGURE creates a new HTML FIGURE element.
func (d *DOM) NewFIGURE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "figure")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewFOOTER creates a new HTML FOOTER element.
func (d *DOM) NewFOOTER() (el *markup.Element) {
	e := document.Call(createElementMethodName, "footer")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewH1 creates a new HTML H1 element.
func (d *DOM) NewH1() (el *markup.Element) {
	e := document.Call(createElementMethodName, "h1")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewH2 creates a new HTML H2 element.
func (d *DOM) NewH2() (el *markup.Element) {
	e := document.Call(createElementMethodName, "h2")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewH3 creates a new HTML H3 element.
func (d *DOM) NewH3() (el *markup.Element) {
	e := document.Call(createElementMethodName, "h3")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewH4 creates a new HTML H4 element.
func (d *DOM) NewH4() (el *markup.Element) {
	e := document.Call(createElementMethodName, "h4")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewH5 creates a new HTML H5 element.
func (d *DOM) NewH5() (el *markup.Element) {
	e := document.Call(createElementMethodName, "h5")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewH6 creates a new HTML H6 element.
func (d *DOM) NewH6() (el *markup.Element) {
	e := document.Call(createElementMethodName, "h6")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewHEAD creates a new HTML HEAD element.
func (d *DOM) NewHEAD() (el *markup.Element) {
	e := document.Call(createElementMethodName, "head")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewHEADER creates a new HTML HEADER element.
func (d *DOM) NewHEADER() (el *markup.Element) {
	e := document.Call(createElementMethodName, "header")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewHGROUP creates a new HTML HGROUP element.
func (d *DOM) NewHGROUP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "hgroup")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewHR creates a new HTML HR element.
func (d *DOM) NewHR() (el *markup.Element) {
	e := document.Call(createElementMethodName, "hr")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewHTML creates a new HTML HTML element.
func (d *DOM) NewHTML() (el *markup.Element) {
	e := document.Call(createElementMethodName, "html")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewI creates a new HTML I element.
func (d *DOM) NewI() (el *markup.Element) {
	e := document.Call(createElementMethodName, "i")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewIFRAME creates a new HTML IFRAME element.
func (d *DOM) NewIFRAME() (el *markup.Element) {
	e := document.Call(createElementMethodName, "iframe")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewIMG creates a new HTML IMG element.
func (d *DOM) NewIMG() (el *markup.Element) {
	e := document.Call(createElementMethodName, "img")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewINS creates a new HTML INS element.
func (d *DOM) NewINS() (el *markup.Element) {
	e := document.Call(createElementMethodName, "ins")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewKBD creates a new HTML KBD element.
func (d *DOM) NewKBD() (el *markup.Element) {
	e := document.Call(createElementMethodName, "kbd")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewLABEL creates a new HTML LABEL element.
func (d *DOM) NewLABEL() (el *markup.Element) {
	e := document.Call(createElementMethodName, "label")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewLEGEND creates a new HTML LEGEND element.
func (d *DOM) NewLEGEND() (el *markup.Element) {
	e := document.Call(createElementMethodName, "legend")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewLI creates a new HTML LI element.
func (d *DOM) NewLI() (el *markup.Element) {
	e := document.Call(createElementMethodName, "li")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewMAIN creates a new HTML MAIN element.
func (d *DOM) NewMAIN() (el *markup.Element) {
	e := document.Call(createElementMethodName, "main")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewMAP creates a new HTML MAP element.
func (d *DOM) NewMAP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "map")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewMARK creates a new HTML MARK element.
func (d *DOM) NewMARK() (el *markup.Element) {
	e := document.Call(createElementMethodName, "mark")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewMENU creates a new HTML MENU element.
func (d *DOM) NewMENU() (el *markup.Element) {
	e := document.Call(createElementMethodName, "menu")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewMETER creates a new HTML METER element.
func (d *DOM) NewMETER() (el *markup.Element) {
	e := document.Call(createElementMethodName, "meter")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewOBJECT creates a new HTML OBJECT element.
func (d *DOM) NewOBJECT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "object")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewOL creates a new HTML OL element.
func (d *DOM) NewOL() (el *markup.Element) {
	e := document.Call(createElementMethodName, "ol")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewOPTGROUP creates a new HTML OPTGROUP element.
func (d *DOM) NewOPTGROUP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "optgroup")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewOPTION creates a new HTML OPTION element.
func (d *DOM) NewOPTION() (el *markup.Element) {
	e := document.Call(createElementMethodName, "option")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewOUTPUT creates a new HTML OUTPUT element.
func (d *DOM) NewOUTPUT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "output")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewP creates a new HTML P element.
func (d *DOM) NewP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "p")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewPARAM creates a new HTML PARAM element.
func (d *DOM) NewPARAM() (el *markup.Element) {
	e := document.Call(createElementMethodName, "param")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewPICTURE creates a new HTML PICTURE element.
func (d *DOM) NewPICTURE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "picture")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewPRE creates a new HTML PRE element.
func (d *DOM) NewPRE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "pre")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewPROGRESS creates a new HTML PROGRESS element.
func (d *DOM) NewPROGRESS() (el *markup.Element) {
	e := document.Call(createElementMethodName, "progress")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewQ creates a new HTML Q element.
func (d *DOM) NewQ() (el *markup.Element) {
	e := document.Call(createElementMethodName, "q")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewRB creates a new HTML RB element.
func (d *DOM) NewRB() (el *markup.Element) {
	e := document.Call(createElementMethodName, "rb")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewRP creates a new HTML RP element.
func (d *DOM) NewRP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "rp")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewRT creates a new HTML RT element.
func (d *DOM) NewRT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "rt")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewRTC creates a new HTML RTC element.
func (d *DOM) NewRTC() (el *markup.Element) {
	e := document.Call(createElementMethodName, "rtc")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewRUBY creates a new HTML RUBY element.
func (d *DOM) NewRUBY() (el *markup.Element) {
	e := document.Call(createElementMethodName, "ruby")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewS creates a new HTML S element.
func (d *DOM) NewS() (el *markup.Element) {
	e := document.Call(createElementMethodName, "s")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSAMP creates a new HTML SAMP element.
func (d *DOM) NewSAMP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "samp")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSCRIPT creates a new HTML SCRIPT element.
func (d *DOM) NewSCRIPT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "script")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSECTION creates a new HTML SECTION element.
func (d *DOM) NewSECTION() (el *markup.Element) {
	e := document.Call(createElementMethodName, "section")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSELECT creates a new HTML SELECT element.
func (d *DOM) NewSELECT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "select")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSLOT creates a new HTML SLOT element.
func (d *DOM) NewSLOT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "slot")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSMALL creates a new HTML SMALL element.
func (d *DOM) NewSMALL() (el *markup.Element) {
	e := document.Call(createElementMethodName, "small")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSOURCE creates a new HTML SOURCE element.
func (d *DOM) NewSOURCE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "source")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSPAN creates a new HTML SPAN element.
func (d *DOM) NewSPAN() (el *markup.Element) {
	e := document.Call(createElementMethodName, "span")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSTRONG creates a new HTML STRONG element.
func (d *DOM) NewSTRONG() (el *markup.Element) {
	e := document.Call(createElementMethodName, "strong")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSTYLE creates a new HTML STYLE element.
func (d *DOM) NewSTYLE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "style")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSUB creates a new HTML SUB element.
func (d *DOM) NewSUB() (el *markup.Element) {
	e := document.Call(createElementMethodName, "sub")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSUMMARY creates a new HTML SUMMARY element.
func (d *DOM) NewSUMMARY() (el *markup.Element) {
	e := document.Call(createElementMethodName, "summary")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewSUP creates a new HTML SUP element.
func (d *DOM) NewSUP() (el *markup.Element) {
	e := document.Call(createElementMethodName, "sup")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTABLE creates a new HTML TABLE element.
func (d *DOM) NewTABLE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "table")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTBODY creates a new HTML TBODY element.
func (d *DOM) NewTBODY() (el *markup.Element) {
	e := document.Call(createElementMethodName, "tbody")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTD creates a new HTML TD element.
func (d *DOM) NewTD() (el *markup.Element) {
	e := document.Call(createElementMethodName, "td")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTEMPLATE creates a new HTML TEMPLATE element.
func (d *DOM) NewTEMPLATE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "template")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTEXTAREA creates a new HTML TEXTAREA element.
func (d *DOM) NewTEXTAREA() (el *markup.Element) {
	e := document.Call(createElementMethodName, "textarea")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTFOOT creates a new HTML TFOOT element.
func (d *DOM) NewTFOOT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "tfoot")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTH creates a new HTML TH element.
func (d *DOM) NewTH() (el *markup.Element) {
	e := document.Call(createElementMethodName, "th")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTHEAD creates a new HTML THEAD element.
func (d *DOM) NewTHEAD() (el *markup.Element) {
	e := document.Call(createElementMethodName, "thead")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTIME creates a new HTML TIME element.
func (d *DOM) NewTIME() (el *markup.Element) {
	e := document.Call(createElementMethodName, "time")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTITLE creates a new HTML TITLE element.
func (d *DOM) NewTITLE() (el *markup.Element) {
	e := document.Call(createElementMethodName, "title")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTR creates a new HTML TR element.
func (d *DOM) NewTR() (el *markup.Element) {
	e := document.Call(createElementMethodName, "tr")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTRACK creates a new HTML TRACK element.
func (d *DOM) NewTRACK() (el *markup.Element) {
	e := document.Call(createElementMethodName, "track")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewTT creates a new HTML TT element.
func (d *DOM) NewTT() (el *markup.Element) {
	e := document.Call(createElementMethodName, "tt")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewU creates a new HTML U element.
func (d *DOM) NewU() (el *markup.Element) {
	e := document.Call(createElementMethodName, "u")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewUL creates a new HTML UL element.
func (d *DOM) NewUL() (el *markup.Element) {
	e := document.Call(createElementMethodName, "ul")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewVAR creates a new HTML VAR element.
func (d *DOM) NewVAR() (el *markup.Element) {
	e := document.Call(createElementMethodName, "var")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewVIDEO creates a new HTML VIDEO element.
func (d *DOM) NewVIDEO() (el *markup.Element) {
	e := document.Call(createElementMethodName, "video")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}

// NewWBR creates a new HTML WBR element.
func (d *DOM) NewWBR() (el *markup.Element) {
	e := document.Call(createElementMethodName, "wbr")
	el = markup.NewElement(e, d.panelUniqueID)
	return
}
