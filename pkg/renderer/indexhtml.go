package renderer

import (
	"fmt"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func buildIndexHTMLNode(builder *project.Builder, addLocations bool) *html.Node {
	// document
	doc := &html.Node{
		Type: html.DocumentNode,
		Data: "!DOCTYPE html",
		/*
			Attr: []html.Attribute{
				html.Attribute{Key: "html"},
			},
		*/
	}
	// html
	htm := &html.Node{
		Type:     html.ElementNode,
		Data:     "html",
		DataAtom: atom.Html,
	}
	doc.AppendChild(htm)
	// head
	head := buildHeadNode(builder)
	htm.AppendChild(head)
	// body
	body := &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	}
	htm.AppendChild(body)
	// master view
	tabsMasterView := builder.ToHTMLNode(tabsMasterViewID, addLocations)
	body.AppendChild(tabsMasterView)
	// modal view
	modal := buildModalNode(builder)
	body.AppendChild(modal)
	// closer view
	closer := buildCloserNode(builder)
	body.AppendChild(closer)
	return doc
}

func buildHeadNode(builder *project.Builder) (head *html.Node) {
	head = &html.Node{
		Type:     html.ElementNode,
		Data:     "head",
		DataAtom: atom.Head,
	}
	meta := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Meta,
		Data:     "meta",
		Attr: []html.Attribute{
			{Key: "charset", Val: "utf-8"},
		},
	}
	head.AppendChild(meta)
	title := &html.Node{
		Type:     html.ElementNode,
		Data:     "title",
		DataAtom: atom.Title,
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: builder.Title,
	}
	title.AppendChild(textNode)
	head.AppendChild(title)
	// styles
	style := &html.Node{
		Type:     html.ElementNode,
		Data:     "style",
		DataAtom: atom.Style,
	}
	textNode = &html.Node{
		Type: html.TextNode,
		Data: "@import url(css/main.css);",
	}
	style.AppendChild(textNode)
	head.AppendChild(style)
	style = &html.Node{
		Type:     html.ElementNode,
		Data:     "style",
		DataAtom: atom.Style,
	}
	textNode = &html.Node{
		Type: html.TextNode,
		Data: "@import url(css/colors.css);",
	}
	style.AppendChild(textNode)
	head.AppendChild(style)
	// scripts
	script := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Script,
		Data:     "script",
	}
	textNode = &html.Node{
		Type: html.TextNode,
		Data: `
		// wasm callbacks do not return values so use javascript to return true.
		window.onabort = function(e) {return true;}
		window.onbeforeunload = function(e) {return true;}
	`,
	}
	script.AppendChild(textNode)
	head.AppendChild(script)
	script = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Script,
		Data:     "script",
		Attr: []html.Attribute{
			{Key: "src", Val: "/wasm/wasm_exec.js"},
		},
	}
	head.AppendChild(script)
	script = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Script,
		Data:     "script",
	}
	textNode = &html.Node{
		Type: html.TextNode,
		Data: `
		if (!WebAssembly.instantiateStreaming) { // polyfill
			WebAssembly.instantiateStreaming = async (resp, importObject) => {
				const source = await (await resp).arrayBuffer();
				return await WebAssembly.instantiate(source, importObject);
			};
		}
	
		(
			async function() {
				const go = new Go()
				const { instance } = await WebAssembly.instantiateStreaming(fetch("/wasm/app.wasm"), go.importObject)
				go.run(instance)
			}
		)()
		`,
	}
	script.AppendChild(textNode)
	head.AppendChild(script)
	// head template reference
	fileNames := paths.GetFileNames()
	textNode = &html.Node{
		Type: html.TextNode,
		Data: fmt.Sprintf(`{{template "%s"}}`, fileNames.HeadDotTMPL),
	}
	head.AppendChild(textNode)
	return
}

func buildModalNode(builder *project.Builder) (modal *html.Node) {
	// modal view
	modal = &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
		Attr: []html.Attribute{
			{Key: "id", Val: "modalInformationMasterView"},
			{Key: "class", Val: builder.Classes.UnSeen},
		},
	}
	center := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
		Attr: []html.Attribute{
			{Key: "id", Val: "modalInformationMasterView-center"},
			{Key: "class", Val: builder.Classes.ModalUserContent},
		},
	}
	modal.AppendChild(center)
	h1 := &html.Node{
		Type:     html.ElementNode,
		Data:     "h1",
		DataAtom: atom.H1,
		Attr: []html.Attribute{
			{Key: "id", Val: "modalInformationMasterView-h1"},
			{Key: "class", Val: builder.Classes.PanelHeading},
		},
	}
	center.AppendChild(h1)
	message := &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div,
		Attr: []html.Attribute{
			{Key: "id", Val: "modalInformationMasterView-message"},
		},
	}
	center.AppendChild(message)
	p := &html.Node{
		Type:     html.ElementNode,
		Data:     "p",
		DataAtom: atom.P,
	}
	center.AppendChild(p)
	button := &html.Node{
		Type:     html.ElementNode,
		Data:     "button",
		DataAtom: atom.Button,
		Attr: []html.Attribute{
			{Key: "id", Val: "modalInformationMasterView-close"},
		},
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: "Close",
	}
	button.AppendChild(textNode)
	p.AppendChild(button)
	return
}

func buildCloserNode(builder *project.Builder) (closer *html.Node) {
	closer = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: "closerMasterView"},
			{Key: "class", Val: builder.Classes.UnSeen},
		},
	}
	center := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: "closerMasterView-center"},
			{Key: "class", Val: builder.Classes.CloserUserContent},
		},
	}
	closer.AppendChild(center)
	h1 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.H1,
		Data:     "h1",
	}
	textNode := &html.Node{
		Type: html.TextNode,
		Data: "Close this application.",
	}
	h1.AppendChild(textNode)
	center.AppendChild(h1)
	// Are you certain paragraph
	p := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.P,
		Data:     "p",
	}
	textNode = &html.Node{
		Type: html.TextNode,
		Data: "Are you certain that you want to close this application?",
	}
	p.AppendChild(textNode)
	center.AppendChild(p)
	// buttons paragraph
	p = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.P,
		Data:     "p",
	}
	center.AppendChild(p)
	button := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "button",
		Attr: []html.Attribute{
			{Key: "id", Val: "closerMasterView-cancel"},
		},
	}
	textNode = &html.Node{
		Type: html.TextNode,
		Data: "No! Do not close this application.",
	}
	button.AppendChild(textNode)
	p.AppendChild(button)
	button = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "button",
		Attr: []html.Attribute{
			{Key: "id", Val: "closerMasterView-close"},
		},
	}
	textNode = &html.Node{
		Type: html.TextNode,
		Data: "Yes, close this application.",
	}
	button.AppendChild(textNode)
	p.AppendChild(button)
	return
}
