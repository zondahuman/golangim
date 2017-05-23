package main

import (
	"fmt"
	"golangim/im/model"
)

/**
 *
 * 基本状态基类
 *
**/

func main() {

	chatMessageResponse := new(model.ChatMessageResponse)
	//Set BaseModel
	chatMessageResponse.BaseModel.Status = "SUCCESS"
	chatMessageResponse.BaseModel.Message = "this is a message"
	chatMessageResponse.BaseModel.Data = "{'take' : 'lk'}"
	//Set Property
	chatMessageResponse.Id = 1
	chatMessageResponse.ResponseMessage = "haha"
	chatMessageResponse.UserName = "abin"

	fmt.Println("chatMessageResponse = %v ", chatMessageResponse)

	//	model.BaseModel: BaseModel{Status: "SUCCESS", Message: "no message", Data: nil}, Id: 5, UserName: "abin", Message: "hi"

}
