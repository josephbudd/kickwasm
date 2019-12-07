package renderer

import (
	"fmt"

	"github.com/josephbudd/kickwasm/pkg/paths"
	"github.com/josephbudd/kickwasm/pkg/project"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func buildIndexHTMLNode(appPaths paths.ApplicationPathsI, builder *project.Builder, addLocations bool) *html.Node {
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
	head := buildHeadNode(appPaths, builder)
	htm.AppendChild(head)
	// body
	body := &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	}
	htm.AppendChild(body)
	// master view
	mainMasterView := builder.ToHTMLNode(mainMasterViewID, addLocations)
	body.AppendChild(mainMasterView)
	// modal view
	modal := buildModalNode(builder)
	body.AppendChild(modal)
	// black screen behind printing
	black := buildBlackNode(builder)
	body.AppendChild(black)
	return doc
}

func buildHeadNode(appPaths paths.ApplicationPathsI, builder *project.Builder) (head *html.Node) {
	fileNames := appPaths.GetFileNames()
	folderNames := appPaths.GetFolderNames()

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
		// Data: "@import url(css/main.css);",
		Data: fmt.Sprintf("@import url(%s/%s);", folderNames.CSS, fileNames.MainDotCSS),
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
		// Data: "@import url(css/colors.css);",
		Data: fmt.Sprintf("@import url(%s/%s);", folderNames.CSS, fileNames.ColorsDotCSS),
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
		// Data: "@import url(mycss/Usercontent.css);",
		Data: fmt.Sprintf("@import url(%s/%s);", folderNames.MyCSS, fileNames.UserContentDotCSS),
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
			// {Key: "src", Val: "/wasm/wasm_exec.js"},
			{Key: "src", Val: fmt.Sprintf("/%s/%s", folderNames.WASM, fileNames.WasmExecJS)},
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

func buildBlackNode(builder *project.Builder) (black *html.Node) {
	black = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
		Data:     "div",
		Attr: []html.Attribute{
			{Key: "id", Val: "blackMasterView"},
			{Key: "class", Val: builder.Classes.UnSeen},
		},
	}
	return
}
