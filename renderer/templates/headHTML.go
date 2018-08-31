package templates

// HeadTemplate is the renderer/templates/head.tmpl template.
const HeadTemplate = `<head>
<meta charset="utf-8">
<title>{{.Title}}</title>

<style> @import url(css/main.css); </style>
<style> @import url(css/colors.css); </style>

<script>
	// wasm callbacks do not return values so use javascript to return true.
	window.onabort = function(e) {return true;}
	window.onbeforeunload = function(e) {return true;}
</script>
<script src="/wasm/wasm_exec.js"></script>
<script>

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

</script>

{{.UserHead}}

</head>
`
