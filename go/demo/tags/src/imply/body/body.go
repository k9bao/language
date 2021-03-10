package body

type Config struct {
	Type     string `json:"type"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Result struct {
	Desc string `json:"desc"`
}

type Body interface {
	Output() chan *Result
	Close() error
}
