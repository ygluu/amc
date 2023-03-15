package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"reflect"
)

// 数据通讯上下文
type MsgCtx struct {
}

func (this *MsgCtx) GetConn() IConn {
	return nil
}

func (this *MsgCtx) GetFlag() uintptr {
	return 0
}

func (this *MsgCtx) GetAddr() string {
	return ""
}

// 空数据通讯上下文
var NullCtx = &MsgCtx{}

// 进程内部消息：P2P
type Msg struct {
}

func (this *Msg) Reset()         {}
func (this *Msg) String() string { return "" }
func (this *Msg) ProtoMessage()  {}

// 进程内部消息：广播
type MsgB struct {
	Msg
}

func (this *MsgB) GetBroadcast() bool {
	return true
}

func isMsgObj(typ reflect.Type) bool {
	_, ret1 := typ.MethodByName("Reset")
	_, ret2 := typ.MethodByName("String")
	_, ret3 := typ.MethodByName("ProtoMessage")
	return ret1 && ret2 && ret3
}

func isBroadcastMsg(typ reflect.Type) bool {
	_, ret1 := typ.MethodByName("GetBroadcast")
	return ret1
}

type msgInfo struct {
	name        string
	id          uint32
	isProto     bool
	isBroadcast bool
}

var msgInfoList = make([]*msgInfo, 0)
var msgTypeOfId = make(map[uint32]reflect.Type)
var msgIdOfType = make(map[reflect.Type]uint32)
var msgInfoOfType = make(map[reflect.Type]*msgInfo)
var msgIdOfName = make(map[string]uint32)
var msgNameOfId = make(map[uint32]string)

func isMyMsg(msgid uint32) bool {
	return msgTypeOfId[msgid] != nil
}

func ProtoMsgIds() []uint32 {
	ret := make([]uint32, len(msgInfoList))
	count := 0
	for _, info := range msgInfoList {
		if info.isProto {
			ret[count] = info.id
			count++
		}
	}
	return ret[:count]
}

func NewMsgOfId(id uint32) IMsg {
	typ := msgTypeOfId[id]
	if typ == nil {
		return nil
	}
	return reflect.New(typ.Elem()).Elem().Addr().Interface().(IMsg)
}

func MsgIdOfName(name string) uint32 {
	return msgIdOfName[name]
}

func MsgIdOfMsg(msg IMsg) uint32 {
	return msgIdOfType[reflect.TypeOf(msg)]
}

func MsgInfoOfMsg(msg IMsg) (uint32, bool) {
	info := msgInfoOfType[reflect.TypeOf(msg)]
	if info == nil {
		return 0, false
	}
	return info.id, info.isProto
}

func MsgNameOfId(id uint32) string {
	return msgNameOfId[id]
}

func doRegisterMsg(msg IMsg, isProto bool) uint32 {
	t := reflect.TypeOf(msg)

	msgName := t.Elem().Name()
	msgId := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%x", md5.Sum([]byte(msgName)))))
	msgType := reflect.TypeOf(msg)

	if msgIdOfName[msgName] != 0 {
		Panic("[amc]doRegisterMsg => 消息类型已经注册：" + msgName)
	}

	if msgTypeOfId[msgId] != nil {
		Panic("[amc]doRegisterMsg => 消息类型名称HASH值碰撞，请改名：" + msgName)
	}

	msgTypeOfId[msgId] = msgType
	msgIdOfType[msgType] = msgId
	msgIdOfName[msgName] = msgId
	msgNameOfId[msgId] = msgName

	info := &msgInfo{
		name:        msgName,
		id:          msgId,
		isProto:     isProto,
		isBroadcast: isBroadcastMsg(msgType),
	}
	msgInfoOfType[msgType] = info
	msgInfoList = append(msgInfoList, info)

	return msgId
}

func registerMsg(msg IMsg) uint32 {
	return doRegisterMsg(msg, false)
}

func RegisterProtoMsg(msg IMsg) uint32 {
	return doRegisterMsg(msg, true)
}
