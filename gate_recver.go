package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"container/list"
	"sync"
	"time"
)

type cdInfo struct {
	time          int64
	last          int64
	speedingCount int
}

type connCD struct {
	lastTimes     map[uint32]*cdInfo
	destConnOfSvc IConn
}

type setCdNode struct {
	msgid uint32
	time  uint32
}

// 客户端接受器，自带CD功能
type GateRecver struct {
	Recver
	cdTimes   map[uint32]int64
	mut       sync.Mutex
	setCDList list.List
	myMsgIds  []uint32
	cdObjs    map[IConn]*connCD
	snet      INet
	cnet      INet
	lb        ILB
}

func NewGateRecver(msgQueue *RecvQueue, snet, cnet INet, lb ILB) iThread {
	this := &GateRecver{}
	this.snet = snet
	this.cnet = cnet
	this.net = snet
	this.lb = lb

	if LB == nil {
		LB = lb
	}
	if recv_queue == nil {
		recv_queue = msgQueue
	}

	return this
}

func (this *GateRecver) SetMySvcMsg(msgIds []uint32) {
	this.myMsgIds = msgIds
}

func (this *GateRecver) isMyMsg(msgid uint32) bool {
	for _, id := range this.myMsgIds {
		if id == msgid {
			return true
		}
	}
	return false
}

func (this *GateRecver) SetCdTime(msgid uint32, time uint32) {
	this.mut.Lock()
	this.setCDList.PushBack(&setCdNode{msgid: msgid, time: time})
	this.mut.Unlock()
}

func (this *GateRecver) checkCd(msgid uint32, ccd *connCD) int {
	info := ccd.lastTimes[msgid]
	if info != nil {
		last := info.last
		now := time.Now().UnixMicro()
		info.last = now
		// 误差5%，即可以提前
		speeding := (now-last)-info.time >= -cdErrorValue
		if speeding {
			info.speedingCount++
		}
		if info.speedingCount >= cdSpeedingLimit {
			return 2
		} else {
			return 1
		}
	}

	info = &cdInfo{}
	ccd.lastTimes[msgid] = info

	info.last = time.Now().UnixMicro()
	info.time = this.cdTimes[msgid]
	if info.time == 0 {
		// 默认CD时间
		info.time = cdTimeDef
	}
	return 0
}

func (this *GateRecver) Run() {
	if this.setCDList.Len() > 0 {
		this.mut.Lock()
		for e := this.setCDList.Front(); e != nil; e = e.Next() {
			node := e.Value.(*setCdNode)
			this.cdTimes[node.msgid] = int64(node.time)
		}
		this.mut.Unlock()
	}

	conn, addr, msgid, flag, pack := this.net.RecvPack()
	if (conn == nil) || (msgid == 0) {
		return
	}

	// 断线
	if pack == nil {
		delete(this.cdObjs, conn)
		return
	}

	// 每个链接对应一个CD对象
	if this.cdObjs == nil {
		this.cdObjs = make(map[IConn]*connCD)
	}
	cdObj := this.cdObjs[conn]

	if cdObj == nil {
		cdObj = &connCD{
			lastTimes: make(map[uint32]*cdInfo),
		}
		this.cdObjs[conn] = cdObj
	} else {
		switch this.checkCd(msgid, cdObj) {
		case 0:
			{
			}
		case 1:
			{
				LogE("[amc]GateRecver.run => 服务请求过于频繁，host:%s, msgid:%d", addr, msgid)
				return
			}
		case 2:
			{
				LogE("[amc]GateRecver.run => 服务请求过于频繁，超过限定次数(%d)，host:%s, msgid:%d", cdSpeedingLimit, addr, msgid)
				onSpeedingLimit(addr, msgid)
				return
			}
		}
	}

	// 是发给网关的
	if this.isMyMsg(msgid) {
		this.UnmarshalAndEnqueue(conn, addr, msgid, flag, pack)
		return
	}

	// msgid对应的服务类型和cdObj.destConnOfSvc一致的，则发送到cdObj.destConnOfSvc指定的服务
	//   否则通过负载均衡发送到msgid对应的的服务之一
	destConn := this.lb.CheckGetConn(msgid, "", cdObj.destConnOfSvc)
	if destConn == nil {
		LogE("[amc]GateRecver.run => 服务尚未启动或客户端msgid错误：%d", msgid)
		return
	}
	this.snet.SendPack(destConn, msgid, conn.Uintptr(), pack)
}
