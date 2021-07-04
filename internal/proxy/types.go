package proxy

type Response struct {
	Code    int               `json:"code"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}
