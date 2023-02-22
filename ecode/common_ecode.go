package ecode

func init() {
	Load()
}

var (
	OK        = New(0, "成功") // 正确
	ErrString = New(1, "无效错误码")

	// 定义错误码
	// 通用错误码 1000 - 1099
	ErrBind                 = New(1000, "请求结构不合法")
	ErrValidation           = New(1001, "请求内容校验失败")
	ErrToken                = New(1002, "操作鉴权失败")
	ErrParameter            = New(1003, "无效参数")
	ErrSupport              = New(1004, "方法未支持")
	ErrImplement            = New(1005, "方法未实现")
	ErrPassword             = New(1006, "无效密码")
	ErrOnlyAdmin            = New(1007, "管理员才能操作")
	ErrAdminPassword        = New(1008, "管理员密码错误")
	ErrTimeFormat           = New(1009, "时间格式错误, 合法的格式:2020-03-01 10:30:45")
	ErrWebSocketConnection  = New(1010, "WebSocket连接异常")
	ErrWebSocketMessageType = New(1011, "WebSocket消息类型错误")
	ErrWebSocketWrite       = New(1012, "WebSocket发送数据异常")
	ErrWebSocketHeartbeat   = New(1013, "WebSocket心跳超时")

	//Database  & Redis 1100 - 1199
	ErrDatabase = New(1100, "DB error.")
	ErrDBDelete = New(1101, "DB delete error.")
	ErrDBCreate = New(1102, "DB create error.")
	ErrDBGet    = New(1103, "DB get error.")
	ErrDBUpdate = New(1104, "DB update error.")
)
