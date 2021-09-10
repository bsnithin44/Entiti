package api

type Error struct {
	Code           string      `json:"code,omitempty"`
	Description    string      `json:"description,omitempty"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
}
type Request struct {
	Uri         string      `json:"uri,omitempty"`
	QueryString string      `json:"queryString,omitempty"`
	Body        interface{} `json:"body,omitempty"`
}
type Data struct {
	Id string `json:"id,omitempty"`
}

type CreatedResponse struct {
	Error   `json:"error,omitempty"`
	Data    `json:"data,omitempty"`
	Request `json:"request,omitempty"`
}
type Response struct {
	Error   `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Request `json:"request,omitempty"`
}

type MyError struct {
	Code           string      `json:"code,omitempty"`
	Description    string      `json:"description,omitempty"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
	StatusCode     int         `json:"status_code,omitempty"`
}

func (e1 *MyError) Error() string {
	return e1.Description

}
