package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

type IModule interface {
	Queue() *RecvQueue
	Name() string
	Init()
	Ready()
	Final()
	Type() moduleType
}

// 消息类接口
type IMsg interface {
	Reset()
	String() string
	ProtoMessage()
}

// 编码器接口（消息序列化和反序列化）
type ICodec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// 封包器接口
type ICodep interface {
	CalcPackLen(dataLen uint32) uint32
	// 封包
	Enpack(msgid uint32, flag uintptr, src []byte, dest []byte) error
	// 解包
	Depack(srcBuf *[]byte, srclen *uint32) (msgid uint32, flag uintptr, pack []byte, err error)
}

// 数据通讯上下文接口
type ICtx interface {
	// 获取网络链接，如果接收数据则表示发送方，如果发送数据则表示接受方(0为广播给所有服务)
	GetConn() IConn
	// 获取该数据包的flag
	//    如果GetConn是客户端则是该客户端对应的接入网关链接
	GetFlag() uintptr
	GetAddr() string
}

// 网络链接
type IConn interface {
	IsClose() bool
	Uintptr() uintptr
}

// 网络接口
type INet interface {
	// 服务端监听，可服务端、客户端同时
	ListenByIp(ip string, beginPort int, endPort int) (string, error)
	ListenByAddr(addr string) (err error)
	// 客户端拨号链接，可服务端、客户端同时
	Dial(addr string) (conn IConn, err error)
	// 发送一个数据包，Net内部自行封包
	SendPack(conn IConn, msgid uint32, flag uintptr, pack []byte) (err error)
	// 接收一个Net内部已经解包好的数据包
	RecvPack() (conn IConn, addr string, msgid uint32, flag uintptr, pack []byte)
	Close()
	Disconn(conn IConn)
}

// 服务发现接口
type ISD interface {
	Watch(onWatch func(key string, weight, flag int))
	ParseName(key string) (svcname, addr string)
	MsgIds(etcdKey string) string
}

// 一致性哈希(Consistent Hash)
type ICHash interface {
	Get(key string) string
	GetBy(hash uint32) string
	Add(key, value string)
	Del(key string)
}

// 负载均衡哈希
type ILB interface {
	// 获取可以对应的哈希值
	GetHash(key string) uint32

	// 获取msgid对于的服务链接（负载均衡）
	GetConn(msgid uint32) IConn
	// 获取msgid对应的所以服务链接
	GetConns(msgid uint32) []IConn

	// 获取msgid对于的服务链接
	//   如果响应msgid的服务和checkSvcName相同，则返回DefConn，
	//   否则返回其中一个服务链接（负载均衡），等同于GetSvcConn
	CheckGetConn(msgid uint32, checkSvcName string, Defconn IConn) IConn

	// 根据哈希Key获取指定服务链接（一致性哈希）,key转为hash值
	GetConnByKey(msgid uint32, key string) IConn
	// 根据哈希值获取指定服务链接（一致性哈希）,可作为预先算好的hash值使用以提高效率
	GetConnByHash(msgid, hash uint32) IConn
	// 获取对应的链接
	GetConnByAddr(addr string) IConn

	// 获取msgid对于的服务信息（服务名和地址，负载均衡）
	GetInfo(msgid uint32) (name, addr string)
	// 获取key对应服务信息（服务名和地址，一致性哈希）
	GetInfoByKey(msgid uint32, key string) (name, addr string)
	// 获取hash对应服务信息（服务名和地址，一致性哈希）
	GetInfoByHash(msgid, hash uint32) (name, addr string)

	// 增加一个服务链接
	AddConn(addr string, conn IConn)
}
