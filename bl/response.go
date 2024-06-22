package bl

type Response struct {
	Codigo  int         `json:"codigo"`
	Mensaje string      `json:"mensaje"`
	Data    interface{} `json:"data"`
}
