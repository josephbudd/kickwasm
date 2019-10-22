package templates

// VSCodeWorkSpaceJSON is the rp.code-workspace file.
const VSCodeWorkSpaceJSON = `{
	"folders": [
		{
			"path": "."
		},
		{
			"path": "../site"
		},
		{
			"path": "../domain"
		},
	],
	"settings": {
		"go":{
			"toolsEnvVars":{
				"GOARCH":"wasm",
				"GOOS":"js",
			},
			"testEnvVars": {
				"GOARCH":"wasm",
				"GOOS":"js",
			},
			"installDependenciesWhenBuilding": false,
		},
	},
}`
