package main

type BodyStruct struct {
	Password string      `json:"password"`
	Monitor  string      `json:"monitor"`
	List     []BotStruct `json:"list"`
}

type BotStruct struct {
	Name    string       `json:"name"`
	Status  string       `json:"status"`
	World   string       `json:"world"`
	Level   int32        `json:"level"`
	Captcha string       `json:"captcha"`
	X       int32        `json:"x"`
	Y       int32        `json:"y"`
	Profit  []ItemStruct `json:"profit"`
}

type ItemStruct struct {
	Id    int32 `json:"id"`
	Total int32 `json:"total"`
}

type BotDocumentStruct struct {
	Username   string		`json:"username"`
	LastUpdate int64        `json:"lastupdate"`
	Monitor    string       `json:"monitor"`
	Name       string       `json:"name"`
	Status     string       `json:"status"`
	World      string       `json:"world"`
	Level      int32        `json:"level"`
	Captcha    string       `json:"captcha"`
	X          int32        `json:"x"`
	Y          int32        `json:"y"`
	Profit []ItemStruct 	`json:"profit"`
}

type UserDocumentStruct struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Monitors []string `json:"monitors"`
}


type ResponseFindBot struct {
	List     []BotDocumentStruct `json:"list"`
}