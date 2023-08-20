package errno

var (
	ErrRecordConnectFail = &Errno{HTTP: 500, Code: "MongoDB.ConnectError", Message: "Mongodb connect failed."}
)
