package main

type HttpNotFound struct {
	Code int32 `json:code`
	Status string `json:status`
}

type HttpUnauthorized struct {
	Code int32 `json:code`
	Status string `json:status`
}

type HttpBadRequest struct {
	Code int32 `json:code`
	Status string `json:status`	
}

type HttpOK struct {
	Code int32 `json:code`
	Status string `json:status`	
}