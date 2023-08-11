package model

type Data struct {
	ID     int
	Name   string
	Type   string
	Detail string
	Url    string
	//
	// input your data
	//
}

type DataResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	Url    string `json:"url"`
	//
	// input your data
	//
}
