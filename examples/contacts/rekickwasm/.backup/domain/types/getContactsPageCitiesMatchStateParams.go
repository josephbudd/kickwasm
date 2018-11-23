package types

// RendererToMainProcessGetContactsPageCitiesMatchStateParams are the GetContactsPage function parameters that the renderer sends to the main process.
type RendererToMainProcessGetContactsPageCitiesMatchStateParams struct {
	StateMatch  string
	SortedIndex uint64
	PageSize    uint64
	State       uint64
}

// MainProcessToRendererGetContactsPageCitiesMatchStateParams are the GetContactsPage function parameters that the main process sends to the renderer.
type MainProcessToRendererGetContactsPageCitiesMatchStateParams struct {
	StateMatch   string
	SortedIndex  uint64
	PageSize     uint64
	Records      []*ContactRecord
	RecordCount  uint64
	State        uint64
	Error        bool
	ErrorMessage string
}
