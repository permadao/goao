package schema

const (
	DataProtocol = "ao"
	Variant      = "ao.TN.1"
	TypeMessage  = "Message"
	TypeProcess  = "Process"
	SDK          = "goao"
)

type ResponseMu struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type ResponseCu struct {
}