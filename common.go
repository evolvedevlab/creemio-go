package creemio

// Mode represents the environment
type Mode string

const (
	ModeTest       Mode = "test"
	ModeProduction Mode = "prod"
	ModeSandbox    Mode = "sandbox"
)

type Pagination struct {
	TotalRecords int  `json:"total_records"`
	TotalPages   int  `json:"total_pages"`
	CurrentPage  int  `json:"current_page"`
	NextPage     int  `json:"next_page"`
	PreviousPage *int `json:"prev_page"`
}

type CustomField struct {
	Type     string    `json:"type"`
	Key      string    `json:"key"`
	Label    string    `json:"label"`
	Optional bool      `json:"optional"`
	Text     *TextSpec `json:"text,omitempty"`
}

type TextSpec struct {
	MaxLength int `json:"max_length"`
	MinLength int `json:"min_length"`
}
