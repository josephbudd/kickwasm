package types

// RendererToMainProcessGetContactsPageRecordsMatchStateCityParams are the GetContactsPage function parameters that the renderer sends to the main process.
type RendererToMainProcessGetContactsPageRecordsMatchStateCityParams struct {
	StateMatch  string
	CityMatch   string
	SortedIndex uint64
	PageSize    uint64
	State       uint64
}

// MainProcessToRendererGetContactsPageRecordsMatchStateCityParams are the GetContactsPage function parameters that the main process sends to the renderer.
type MainProcessToRendererGetContactsPageRecordsMatchStateCityParams struct {
	StateMatch   string
	CityMatch    string
	SortedIndex  uint64
	PageSize     uint64
	Records      []*ContactRecord
	RecordCount  uint64
	State        uint64
	Error        bool
	ErrorMessage string
}
