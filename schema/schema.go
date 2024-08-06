package schema

const (
	DataProtocol = "ao"
	Variant      = "ao.TN.1"
	TypeMessage  = "Message"
	TypeProcess  = "Process"
	SDK          = "goao"

	DefaultModule       = "xT0ogTeagEGuySbKuUoo_NaWeeBv1fZ4MqgDdKVKY0U"
	DefaultSqliteModule = "sFNHeYzhHfP9vV9CPpqZMU-4Zzq_qKGKwlwMZozWi2Y"
	DefaultScheduler    = "_GQ33BkPtZrqxA84vM8Zk-N2aO0toNNu_C-l-rawrBA"
)

type ResponseMu struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type ResponseCu struct {
	Messages    []interface{} `json:"Messages"`
	Assignments []interface{} `json:"Assignments"`
	Spawns      []interface{} `json:"Spawns"`
	Output      interface{}   `json:"Output"`
	GasUsed     int64         `json:"GasUsed"`
}
