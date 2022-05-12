package plan

// Pass is a plan to pass
type Pass string

func (pass Pass) Score() int {
	return 1
}

func (pass Pass) Submit(request RequestFunc) {
	request("pass", map[string]any{
		"id": string(pass),
	})
}

func (pass *Pass) String() string {
	return "pass"
}
