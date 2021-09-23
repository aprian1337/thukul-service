package salaries

type Request struct {
	Minimal float64 `json:"minimal"`
	Maximal float64 `json:"maximal"`
}
