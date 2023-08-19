package status

const (
	OK   int = 200
	FAIL int = 400
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
