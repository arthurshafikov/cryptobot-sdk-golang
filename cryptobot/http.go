package cryptobot

type Response struct {
	Ok     bool          `json:"ok"`
	Result any           `json:"result"`
	Error  ResponseError `json:"error"`
}

type ResponseError struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}
