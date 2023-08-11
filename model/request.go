package model

type GetRequest struct {
	Name string `json:"name"`
	//
	// add your data
	//
}

type AddRequest struct {
	// User string `json:"user" binding:"required"`
	//
	// add your data
	//
	User   string `json:"user" binding:"required"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	Url    string `json:"url"`
	Idrm   string `json:"idrm"`
	// UpdatedFields string `json:"updatedFields"`
}

type UpdateRequest struct {
	// ID   int`
	//
	// add your data
	//
	User   string `json:"user" binding:"required"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	Url    string `json:"url"`
	Idrm   string `json:"idrm" binding:"required"`
	// UpdatedFields string `json:"updatedFields"`
}

type DeleteRequest struct {
	// ID   int
	// User string `json:"user" binding:"required"`
	//
	// add your data
	//
	User   string `json:"user" binding:"required"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	Url    string `json:"url"`
	Idrm   string `json:"idrm" binding:"required"`
	// UpdatedFields string `json:"updatedFields"`
}
