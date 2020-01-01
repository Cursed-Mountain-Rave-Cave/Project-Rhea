package web

//Message ss
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//Login ss
type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
