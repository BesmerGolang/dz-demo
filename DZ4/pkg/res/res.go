package res

type ResOk struct {
	Data any `json:"data"`
}
type ResErr struct {
	Message string `json:"message"`
}
