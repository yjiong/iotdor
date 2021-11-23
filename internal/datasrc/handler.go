package datasrc

// DSrcer defines the interface of a handler data
type DSrcer interface {
	Close() error                                     // 断开mqtt连接
	SendData(topic string, paylosd interface{}) error // 发送上行数据
	DataDownChan() chan DataDownPayload               // 返回订阅到的消息数据channel
	IsConnected() bool
}
