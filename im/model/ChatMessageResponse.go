package model

/**
 *
 * 聊天请求实体类
 *
**/

type ChatMessageResponse struct {
	BaseModel
	Id       int
	UserName string
	Message  string
}

//func main() {
//	chatMessageResponse := ChatMessageResponse{base: model.BaseModel{Status: "SUCCESS", Message: "no message", Data: nil}, Id: 5, UserName: "abin", Message: "hi"}
//}
