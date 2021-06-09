package team

type Assignment struct {
	Reference string `json:"reference"`
	Summary   string `json:"summary"`
}

type Member struct {
	Name        string       `json:"name"`
	Assignments []Assignment `json:"assignments"`
}

type Team struct {
	Cycle   string   `json:"cycle"`
	Sprint  string   `json:"sprint"`
	Name    string   `json:"name"`
	Members []Member `json:"members"`
}
