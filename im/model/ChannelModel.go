package model

/**
 *
 * 聊天请求实体类
 *
**/

type ChannelHeader struct {
	ClientMsgId string //客户端生成的消息Id
	ServerMsgId string //服务端生成的消息Id
	MsgContent  string //消息内容
	EquipmentId string //设备号
	Channel     string //渠道IOS，ANDROID，PAD
	SessionId   string //
	MessageType string //消息类型，1：文本，2：语音，音频，3：图片

}
