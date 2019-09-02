package notjs

import "syscall/js"

// GET

// GetElementByID inovokes the document;s getElementById.
func (notjs *NotJS) GetElementByID(id string) js.Value {
	return notjs.document.Call("getElementById", id)
}

// GetElementsByTagName invokes the document's getElementsByTagName.
func (notjs *NotJS) GetElementsByTagName(tagName string) []js.Value {
	els := notjs.document.Call("getElementsByTagName", tagName)
	l := els.Length()
	tagNames := make([]js.Value, l, l)
	for i := 0; i < l; i++ {
		tagNames[i] = els.Index(i)
	}
	return tagNames
}

// VARIOUS CREATES

// CreateTextNode invokes the document's createElement.
func (notjs *NotJS) CreateTextNode(text string) js.Value {
	return notjs.document.Call("createTextNode", text)
}

// CreateElement invokes the document's createElement.
func (notjs *NotJS) CreateElement(tagName string) js.Value {
	return notjs.document.Call(createElementMethodName, tagName)
}

// CREATE FORM INPUT VARIATIONS.

// CreateElementINPUT creates a text input element.
func (notjs *NotJS) CreateElementINPUT() js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, "text")
	return input
}

// CreateElementCheckBoxInGroup creates a checkbox input element.
func (notjs *NotJS) CreateElementCheckBoxInGroup(group string) js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, checkboxTypeName)
	if len(group) > 0 {
		input.Set(groupAttributeName, group)
	}
	return input
}

// CreateElementRadioInGroup creates a radio input element
func (notjs *NotJS) CreateElementRadioInGroup(group string) js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, radioTypeName)
	input.Set(groupAttributeName, group)
	return input
}

// CreateElementCheckBox creates a checkbox input element.
func (notjs *NotJS) CreateElementCheckBox() js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, checkboxTypeName)
	return input
}

// CreateElementRadio creates a radio input element
func (notjs *NotJS) CreateElementRadio() js.Value {
	input := notjs.document.Call(createElementMethodName, inputTypeName)
	input.Set(typeAttributeName, radioTypeName)
	return input
}

// CreateElementA invokes the document's createElement.
func (notjs *NotJS) CreateElementA() js.Value {
	return notjs.document.Call(createElementMethodName, "a")
}

// CreateElementABBR invokes the document's createElement.
func (notjs *NotJS) CreateElementABBR() js.Value {
	return notjs.document.Call(createElementMethodName, "abbr")
}

// CreateElementADDRESS invokes the document's createElement.
func (notjs *NotJS) CreateElementADDRESS() js.Value {
	return notjs.document.Call(createElementMethodName, "address")
}

// CreateElementAREA invokes the document's createElement.
func (notjs *NotJS) CreateElementAREA() js.Value {
	return notjs.document.Call(createElementMethodName, "area")
}

// CreateElementARTICLE invokes the document's createElement.
func (notjs *NotJS) CreateElementARTICLE() js.Value {
	return notjs.document.Call(createElementMethodName, "article")
}

// CreateElementASIDE invokes the document's createElement.
func (notjs *NotJS) CreateElementASIDE() js.Value {
	return notjs.document.Call(createElementMethodName, "aside")
}

// CreateElementAUDIO invokes the document's createElement.
func (notjs *NotJS) CreateElementAUDIO() js.Value {
	return notjs.document.Call(createElementMethodName, "audio")
}

// CreateElementB invokes the document's createElement.
func (notjs *NotJS) CreateElementB() js.Value {
	return notjs.document.Call(createElementMethodName, "b")
}

// CreateElementBDI invokes the document's createElement.
func (notjs *NotJS) CreateElementBDI() js.Value {
	return notjs.document.Call(createElementMethodName, "bdi")
}

// CreateElementBDO invokes the document's createElement.
func (notjs *NotJS) CreateElementBDO() js.Value {
	return notjs.document.Call(createElementMethodName, "bdo")
}

// CreateElementBLOCKQUOTE invokes the document's createElement.
func (notjs *NotJS) CreateElementBLOCKQUOTE() js.Value {
	return notjs.document.Call(createElementMethodName, "blockquote")
}

// CreateElementBR invokes the document's createElement.
func (notjs *NotJS) CreateElementBR() js.Value {
	return notjs.document.Call(createElementMethodName, "br")
}

// CreateElementBUTTON invokes the document's createElement.
func (notjs *NotJS) CreateElementBUTTON() js.Value {
	return notjs.document.Call(createElementMethodName, "button")
}

// CreateElementCANVAS invokes the document's createElement.
func (notjs *NotJS) CreateElementCANVAS() js.Value {
	return notjs.document.Call(createElementMethodName, "canvas")
}

// CreateElementCAPTION invokes the document's createElement.
func (notjs *NotJS) CreateElementCAPTION() js.Value {
	return notjs.document.Call(createElementMethodName, "caption")
}

// CreateElementCITE invokes the document's createElement.
func (notjs *NotJS) CreateElementCITE() js.Value {
	return notjs.document.Call(createElementMethodName, "cite")
}

// CreateElementCODE invokes the document's createElement.
func (notjs *NotJS) CreateElementCODE() js.Value {
	return notjs.document.Call(createElementMethodName, "code")
}

// CreateElementCOL invokes the document's createElement.
func (notjs *NotJS) CreateElementCOL() js.Value {
	return notjs.document.Call(createElementMethodName, "col")
}

// CreateElementCOLGROUP invokes the document's createElement.
func (notjs *NotJS) CreateElementCOLGROUP() js.Value {
	return notjs.document.Call(createElementMethodName, "colgroup")
}

// CreateElementDATA invokes the document's createElement.
func (notjs *NotJS) CreateElementDATA() js.Value {
	return notjs.document.Call(createElementMethodName, "data")
}

// CreateElementDATALIST invokes the document's createElement.
func (notjs *NotJS) CreateElementDATALIST() js.Value {
	return notjs.document.Call(createElementMethodName, "datalist")
}

// CreateElementDD invokes the document's createElement.
func (notjs *NotJS) CreateElementDD() js.Value {
	return notjs.document.Call(createElementMethodName, "dd")
}

// CreateElementDEL invokes the document's createElement.
func (notjs *NotJS) CreateElementDEL() js.Value {
	return notjs.document.Call(createElementMethodName, "del")
}

// CreateElementDETAILS invokes the document's createElement.
func (notjs *NotJS) CreateElementDETAILS() js.Value {
	return notjs.document.Call(createElementMethodName, "details")
}

// CreateElementDFN invokes the document's createElement.
func (notjs *NotJS) CreateElementDFN() js.Value {
	return notjs.document.Call(createElementMethodName, "dfn")
}

// CreateElementDIALOG invokes the document's createElement.
func (notjs *NotJS) CreateElementDIALOG() js.Value {
	return notjs.document.Call(createElementMethodName, "dialog")
}

// CreateElementDIV invokes the document's createElement.
func (notjs *NotJS) CreateElementDIV() js.Value {
	return notjs.document.Call(createElementMethodName, "div")
}

// CreateElementDL invokes the document's createElement.
func (notjs *NotJS) CreateElementDL() js.Value {
	return notjs.document.Call(createElementMethodName, "dl")
}

// CreateElementDT invokes the document's createElement.
func (notjs *NotJS) CreateElementDT() js.Value {
	return notjs.document.Call(createElementMethodName, "dt")
}

// CreateElementELEMENT invokes the document's createElement.
func (notjs *NotJS) CreateElementELEMENT() js.Value {
	return notjs.document.Call(createElementMethodName, "element")
}

// CreateElementEM invokes the document's createElement.
func (notjs *NotJS) CreateElementEM() js.Value {
	return notjs.document.Call(createElementMethodName, "em")
}

// CreateElementEMBED invokes the document's createElement.
func (notjs *NotJS) CreateElementEMBED() js.Value {
	return notjs.document.Call(createElementMethodName, "embed")
}

// CreateElementFIELDSET invokes the document's createElement.
func (notjs *NotJS) CreateElementFIELDSET() js.Value {
	return notjs.document.Call(createElementMethodName, "fieldset")
}

// CreateElementFIGCAPTION invokes the document's createElement.
func (notjs *NotJS) CreateElementFIGCAPTION() js.Value {
	return notjs.document.Call(createElementMethodName, "figcaption")
}

// CreateElementFIGURE invokes the document's createElement.
func (notjs *NotJS) CreateElementFIGURE() js.Value {
	return notjs.document.Call(createElementMethodName, "figure")
}

// CreateElementFOOTER invokes the document's createElement.
func (notjs *NotJS) CreateElementFOOTER() js.Value {
	return notjs.document.Call(createElementMethodName, "footer")
}

// CreateElementH1 invokes the document's createElement.
func (notjs *NotJS) CreateElementH1() js.Value {
	return notjs.document.Call(createElementMethodName, "h1")
}

// CreateElementH2 invokes the document's createElement.
func (notjs *NotJS) CreateElementH2() js.Value {
	return notjs.document.Call(createElementMethodName, "h2")
}

// CreateElementH3 invokes the document's createElement.
func (notjs *NotJS) CreateElementH3() js.Value {
	return notjs.document.Call(createElementMethodName, "h3")
}

// CreateElementH4 invokes the document's createElement.
func (notjs *NotJS) CreateElementH4() js.Value {
	return notjs.document.Call(createElementMethodName, "h4")
}

// CreateElementH5 invokes the document's createElement.
func (notjs *NotJS) CreateElementH5() js.Value {
	return notjs.document.Call(createElementMethodName, "h5")
}

// CreateElementH6 invokes the document's createElement.
func (notjs *NotJS) CreateElementH6() js.Value {
	return notjs.document.Call(createElementMethodName, "h6")
}

// CreateElementHEADER invokes the document's createElement.
func (notjs *NotJS) CreateElementHEADER() js.Value {
	return notjs.document.Call(createElementMethodName, "header")
}

// CreateElementHGROUP invokes the document's createElement.
func (notjs *NotJS) CreateElementHGROUP() js.Value {
	return notjs.document.Call(createElementMethodName, "hgroup")
}

// CreateElementHR invokes the document's createElement.
func (notjs *NotJS) CreateElementHR() js.Value {
	return notjs.document.Call(createElementMethodName, "hr")
}

// CreateElementI invokes the document's createElement.
func (notjs *NotJS) CreateElementI() js.Value {
	return notjs.document.Call(createElementMethodName, "i")
}

// CreateElementIFRAME invokes the document's createElement.
func (notjs *NotJS) CreateElementIFRAME() js.Value {
	return notjs.document.Call(createElementMethodName, "iframe")
}

// CreateElementIMG invokes the document's createElement.
func (notjs *NotJS) CreateElementIMG() js.Value {
	return notjs.document.Call(createElementMethodName, "img")
}

// CreateElementINS invokes the document's createElement.
func (notjs *NotJS) CreateElementINS() js.Value {
	return notjs.document.Call(createElementMethodName, "ins")
}

// CreateElementKBD invokes the document's createElement.
func (notjs *NotJS) CreateElementKBD() js.Value {
	return notjs.document.Call(createElementMethodName, "kbd")
}

// CreateElementLABEL invokes the document's createElement.
func (notjs *NotJS) CreateElementLABEL() js.Value {
	return notjs.document.Call(createElementMethodName, "label")
}

// CreateElementLEGEND invokes the document's createElement.
func (notjs *NotJS) CreateElementLEGEND() js.Value {
	return notjs.document.Call(createElementMethodName, "legend")
}

// CreateElementLI invokes the document's createElement.
func (notjs *NotJS) CreateElementLI() js.Value {
	return notjs.document.Call(createElementMethodName, "li")
}

// CreateElementMAIN invokes the document's createElement.
func (notjs *NotJS) CreateElementMAIN() js.Value {
	return notjs.document.Call(createElementMethodName, "main")
}

// CreateElementMAP invokes the document's createElement.
func (notjs *NotJS) CreateElementMAP() js.Value {
	return notjs.document.Call(createElementMethodName, "map")
}

// CreateElementMARK invokes the document's createElement.
func (notjs *NotJS) CreateElementMARK() js.Value {
	return notjs.document.Call(createElementMethodName, "mark")
}

// CreateElementMENU invokes the document's createElement.
func (notjs *NotJS) CreateElementMENU() js.Value {
	return notjs.document.Call(createElementMethodName, "menu")
}

// CreateElementMETER invokes the document's createElement.
func (notjs *NotJS) CreateElementMETER() js.Value {
	return notjs.document.Call(createElementMethodName, "meter")
}

// CreateElementOBJECT invokes the document's createElement.
func (notjs *NotJS) CreateElementOBJECT() js.Value {
	return notjs.document.Call(createElementMethodName, "object")
}

// CreateElementOL invokes the document's createElement.
func (notjs *NotJS) CreateElementOL() js.Value {
	return notjs.document.Call(createElementMethodName, "ol")
}

// CreateElementOPTGROUP invokes the document's createElement.
func (notjs *NotJS) CreateElementOPTGROUP() js.Value {
	return notjs.document.Call(createElementMethodName, "optgroup")
}

// CreateElementOPTION invokes the document's createElement.
func (notjs *NotJS) CreateElementOPTION() js.Value {
	return notjs.document.Call(createElementMethodName, "option")
}

// CreateElementOUTPUT invokes the document's createElement.
func (notjs *NotJS) CreateElementOUTPUT() js.Value {
	return notjs.document.Call(createElementMethodName, "output")
}

// CreateElementP invokes the document's createElement.
func (notjs *NotJS) CreateElementP() js.Value {
	return notjs.document.Call(createElementMethodName, "p")
}

// CreateElementPARAM invokes the document's createElement.
func (notjs *NotJS) CreateElementPARAM() js.Value {
	return notjs.document.Call(createElementMethodName, "param")
}

// CreateElementPICTURE invokes the document's createElement.
func (notjs *NotJS) CreateElementPICTURE() js.Value {
	return notjs.document.Call(createElementMethodName, "picture")
}

// CreateElementPRE invokes the document's createElement.
func (notjs *NotJS) CreateElementPRE() js.Value {
	return notjs.document.Call(createElementMethodName, "pre")
}

// CreateElementPROGRESS invokes the document's createElement.
func (notjs *NotJS) CreateElementPROGRESS() js.Value {
	return notjs.document.Call(createElementMethodName, "progress")
}

// CreateElementQ invokes the document's createElement.
func (notjs *NotJS) CreateElementQ() js.Value {
	return notjs.document.Call(createElementMethodName, "q")
}

// CreateElementRB invokes the document's createElement.
func (notjs *NotJS) CreateElementRB() js.Value {
	return notjs.document.Call(createElementMethodName, "rb")
}

// CreateElementRP invokes the document's createElement.
func (notjs *NotJS) CreateElementRP() js.Value {
	return notjs.document.Call(createElementMethodName, "rp")
}

// CreateElementRT invokes the document's createElement.
func (notjs *NotJS) CreateElementRT() js.Value {
	return notjs.document.Call(createElementMethodName, "rt")
}

// CreateElementRTC invokes the document's createElement.
func (notjs *NotJS) CreateElementRTC() js.Value {
	return notjs.document.Call(createElementMethodName, "rtc")
}

// CreateElementRUBY invokes the document's createElement.
func (notjs *NotJS) CreateElementRUBY() js.Value {
	return notjs.document.Call(createElementMethodName, "ruby")
}

// CreateElementS invokes the document's createElement.
func (notjs *NotJS) CreateElementS() js.Value {
	return notjs.document.Call(createElementMethodName, "s")
}

// CreateElementSAMP invokes the document's createElement.
func (notjs *NotJS) CreateElementSAMP() js.Value {
	return notjs.document.Call(createElementMethodName, "samp")
}

// CreateElementSECTION invokes the document's createElement.
func (notjs *NotJS) CreateElementSECTION() js.Value {
	return notjs.document.Call(createElementMethodName, "section")
}

// CreateElementSELECT invokes the document's createElement.
func (notjs *NotJS) CreateElementSELECT() js.Value {
	return notjs.document.Call(createElementMethodName, "select")
}

// CreateElementSLOT invokes the document's createElement.
func (notjs *NotJS) CreateElementSLOT() js.Value {
	return notjs.document.Call(createElementMethodName, "slot")
}

// CreateElementSMALL invokes the document's createElement.
func (notjs *NotJS) CreateElementSMALL() js.Value {
	return notjs.document.Call(createElementMethodName, "small")
}

// CreateElementSOURCE invokes the document's createElement.
func (notjs *NotJS) CreateElementSOURCE() js.Value {
	return notjs.document.Call(createElementMethodName, "source")
}

// CreateElementSPAN invokes the document's createElement.
func (notjs *NotJS) CreateElementSPAN() js.Value {
	return notjs.document.Call(createElementMethodName, "span")
}

// CreateElementSTRONG invokes the document's createElement.
func (notjs *NotJS) CreateElementSTRONG() js.Value {
	return notjs.document.Call(createElementMethodName, "strong")
}

// CreateElementSTYLE invokes the document's createElement.
func (notjs *NotJS) CreateElementSTYLE() js.Value {
	return notjs.document.Call(createElementMethodName, "style")
}

// CreateElementSUB invokes the document's createElement.
func (notjs *NotJS) CreateElementSUB() js.Value {
	return notjs.document.Call(createElementMethodName, "sub")
}

// CreateElementSUMMARY invokes the document's createElement.
func (notjs *NotJS) CreateElementSUMMARY() js.Value {
	return notjs.document.Call(createElementMethodName, "summary")
}

// CreateElementSUP invokes the document's createElement.
func (notjs *NotJS) CreateElementSUP() js.Value {
	return notjs.document.Call(createElementMethodName, "sup")
}

// CreateElementTABLE invokes the document's createElement.
func (notjs *NotJS) CreateElementTABLE() js.Value {
	return notjs.document.Call(createElementMethodName, "table")
}

// CreateElementTBODY invokes the document's createElement.
func (notjs *NotJS) CreateElementTBODY() js.Value {
	return notjs.document.Call(createElementMethodName, "tbody")
}

// CreateElementTD invokes the document's createElement.
func (notjs *NotJS) CreateElementTD() js.Value {
	return notjs.document.Call(createElementMethodName, "td")
}

// CreateElementTEMPLATE invokes the document's createElement.
func (notjs *NotJS) CreateElementTEMPLATE() js.Value {
	return notjs.document.Call(createElementMethodName, "template")
}

// CreateElementTEXTAREA invokes the document's createElement.
func (notjs *NotJS) CreateElementTEXTAREA() js.Value {
	return notjs.document.Call(createElementMethodName, "textarea")
}

// CreateElementTFOOT invokes the document's createElement.
func (notjs *NotJS) CreateElementTFOOT() js.Value {
	return notjs.document.Call(createElementMethodName, "tfoot")
}

// CreateElementTH invokes the document's createElement.
func (notjs *NotJS) CreateElementTH() js.Value {
	return notjs.document.Call(createElementMethodName, "th")
}

// CreateElementTHEAD invokes the document's createElement.
func (notjs *NotJS) CreateElementTHEAD() js.Value {
	return notjs.document.Call(createElementMethodName, "thead")
}

// CreateElementTIME invokes the document's createElement.
func (notjs *NotJS) CreateElementTIME() js.Value {
	return notjs.document.Call(createElementMethodName, "time")
}

// CreateElementTR invokes the document's createElement.
func (notjs *NotJS) CreateElementTR() js.Value {
	return notjs.document.Call(createElementMethodName, "tr")
}

// CreateElementTRACK invokes the document's createElement.
func (notjs *NotJS) CreateElementTRACK() js.Value {
	return notjs.document.Call(createElementMethodName, "track")
}

// CreateElementTT invokes the document's createElement.
func (notjs *NotJS) CreateElementTT() js.Value {
	return notjs.document.Call(createElementMethodName, "tt")
}

// CreateElementU invokes the document's createElement.
func (notjs *NotJS) CreateElementU() js.Value {
	return notjs.document.Call(createElementMethodName, "u")
}

// CreateElementUL invokes the document's createElement.
func (notjs *NotJS) CreateElementUL() js.Value {
	return notjs.document.Call(createElementMethodName, "ul")
}

// CreateElementVAR invokes the document's createElement.
func (notjs *NotJS) CreateElementVAR() js.Value {
	return notjs.document.Call(createElementMethodName, "var")
}

// CreateElementVIDEO invokes the document's createElement.
func (notjs *NotJS) CreateElementVIDEO() js.Value {
	return notjs.document.Call(createElementMethodName, "video")
}

// CreateElementWBR invokes the document's createElement.
func (notjs *NotJS) CreateElementWBR() js.Value {
	return notjs.document.Call(createElementMethodName, "wbr")
}
