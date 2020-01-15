package api

// Record Struct (Model)
type Record struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	Start       string `json:"start"`
	End         string `json:"end"`
	BreakTime   string `json:"breakTime"`
	Description string `json:"description"`
}
