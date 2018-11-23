package types

// RendererToMainProcessGetContactsPageStatesParams are the GetContactsPage function parameters that the renderer sends to the main process.
type RendererToMainProcessGetContactsPageStatesParams struct {
	SortedIndex uint64
	PageSize    uint64
	State       uint64
}

// MainProcessToRendererGetContactsPageStatesParams are the GetContactsPage function parameters that the main process sends to the renderer.
type MainProcessToRendererGetContactsPageStatesParams struct {
	SortedIndex  uint64
	PageSize     uint64
	Records      []*ContactRecord
	RecordCount  uint64
	State        uint64
	Error        bool
	ErrorMessage string
}
