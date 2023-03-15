package tcp

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"container/list"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync/atomic"
	"time"
	"unsafe"

	"lib/amc"
)

var DelConnTime = int64(10 * 60 * 1000)

const perReadLen = uint32(4096)

type rNode struct {
	conn  amc.IConn
	msgid uint32
	flag  uintptr
	pack  []byte
	addr  string
}

type conn struct {
	c        net.Conn
	codep    amc.ICodep
	saveBuf  []byte
	saveLen  uint32
	isDis    bool
	addr     string
	isListen bool
}

func (this *conn) IsClose() bool {
	return this.isDis
}

func (this *conn) Uintptr() uintptr {
	return uintptr(unsafe.Pointer(this))
}

func (this *conn) close() {
	this.isDis = true
	if this.c != nil {
		this.c.Close()
		this.c = nil
	}
}

func (this *conn) sendPack(msgid uint32, flag uintptr, pack []byte) (err error) {
	dataLen := this.codep.CalcPackLen(uint32(len(pack)))
	dataBuf := make([]byte, dataLen)
	err = this.codep.Enpack(msgid, flag, pack, dataBuf)
	if err != nil {
		return
	}
	sumRLen := uint32(0)
	for sumRLen < dataLen {
		wLen, err := this.c.Write(dataBuf[sumRLen:dataLen])
		if err != nil {
			break
		}
		sumRLen += uint32(wLen)
	}
	return nil
}

func (this *conn) recvPack() (msgid uint32, flag uintptr, pack []byte, err error) {

	for {
		if this.saveLen > 12 {
			msgid, flag, pack, err = this.codep.Depack(&this.saveBuf, &this.saveLen)
			if err == nil {
				return
			}
		}

		var newBuf []byte

		if this.saveBuf != nil {
			oldLen := this.saveLen
			newLen := oldLen + perReadLen
			newBuf := make([]byte, newLen)
			this.saveBuf = newBuf
			copy(newBuf, this.saveBuf)
			newBuf = newBuf[oldLen:]
		} else {
			newBuf := make([]byte, perReadLen)
			this.saveBuf = newBuf
			this.saveLen = 0
		}

		newLen, err := this.c.Read(newBuf)
		if (err != nil) || (newLen == 0) {
			return 0, 0, nil, err
		}

		this.saveLen += uint32(newLen)
	}

	return 0, 0, nil, nil
}

type delNode struct {
	delTime int64
	conn    *conn
}

type tcp struct {
	codep         amc.ICodep
	rChan         chan *rNode
	connsOfAddr   map[string]*conn
	connsOfHandle map[amc.IConn]*conn
	delList       list.List
	cCount        int32
	lister        net.Listener
	tasks         amc.Tasks
	isClose       bool
}

func New(codep amc.ICodep) amc.INet {
	return &tcp{
		codep:         codep,
		rChan:         make(chan *rNode, 100),
		connsOfAddr:   make(map[string]*conn),
		connsOfHandle: make(map[amc.IConn]*conn),
	}
}

func IsIPv6(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return ip != nil && strings.Contains(ipAddr, ":")
}

func (this *tcp) ListenByIp(ip string, beginPort int, endPort int) (addr string, err error) {
	rand.Seed(time.Now().UnixNano())

	doListen := func(port int) bool {
		addr = fmt.Sprintf("%s:%d", ip, port)
		if IsIPv6(ip) {
			addr = fmt.Sprintf("[%s]:%d", ip, port)
		}
		err = this.ListenByAddr(addr)
		return err == nil
	}

	for i := 0; i < 20; i++ {
		port := beginPort + rand.Intn(endPort-beginPort)
		if doListen(port) {
			return
		}
	}

	for port := beginPort; port <= endPort; port++ {
		if doListen(port) {
			return
		}
	}

	return "", errors.New(fmt.Sprintf("地址端口范围已经被占用：%s(%d~%d)", ip, beginPort, endPort))
}

func (this *tcp) ListenByAddr(addr string) (err error) {
	this.Close()
	lister, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}

	this.lister = lister
	go this.goAccept()

	return
}

type addConnTask struct {
	tcp  *tcp
	conn *conn
}

func (this *addConnTask) DoExec() {
	this.tcp.connsOfHandle[this.conn] = this.conn
	this.tcp.connsOfAddr[this.conn.addr] = this.conn
}

func (this *tcp) postNullRecv() {
	//this.rChan <- &rNode{conn: (*conn)(uintptr(unsafe.Pointer(1)))}
}

func (this *tcp) addToList(c net.Conn) (ret *conn) {
	addr := c.RemoteAddr().String()
	ret = &conn{}
	ret.c = c
	ret.codep = this.codep
	ret.addr = addr

	task := &addConnTask{tcp: this, conn: ret}
	this.tasks.Add(task)
	this.postNullRecv()
	return
}

func (this *tcp) goAccept() {
	for {
		c, err := this.lister.Accept()
		if err != nil {
			break
		}

		conn := this.addToList(c)
		go this.goRecv(conn)
	}
	this.lister = nil
	time.Sleep(time.Duration(time.Millisecond * 100))
}

func (this *tcp) Dial(addr string) (ret amc.IConn, err error) {
	if this.connsOfAddr != nil {
		ret := this.connsOfAddr[addr]
		if ret != nil {
			return nil, nil
		}
	}

	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	connObj := this.addToList(c)
	ret = connObj
	go this.goRecv(connObj)
	return
}

func (this *tcp) goRecv(conn *conn) {
	defer atomic.AddInt32(&this.cCount, -1)
	atomic.AddInt32(&this.cCount, 1)

	for {
		msgid, flag, pack, err := conn.recvPack()
		if err == nil {
			this.rChan <- &rNode{conn: conn, addr: conn.c.RemoteAddr().String(), msgid: msgid, flag: flag, pack: pack}
			continue
		}

		this.rChan <- &rNode{conn: conn, addr: conn.c.RemoteAddr().String(), msgid: 0, flag: 0, pack: nil}

		conn.isDis = true
		conn.close()
		this.disconn(conn, false)
		return
	}
}

func (this *tcp) Close() {
	this.isClose = true

	if this.lister != nil {
		this.lister.Close()

		for {
			if this.lister == nil {
				break
			}
			time.Sleep(time.Duration(time.Millisecond * 100))
		}
	}

	if this.connsOfHandle != nil {
		for _, conn := range this.connsOfHandle {
			conn.c.Close()
		}
	}
	this.connsOfHandle = nil
	this.connsOfAddr = nil

	for {
		if this.cCount == 0 {
			break
		}
		time.Sleep(time.Duration(time.Millisecond * 100))
	}

	this.rChan <- &rNode{}
}

type delConnTask struct {
	tcp   *tcp
	conn  amc.IConn
	isDel bool
}

func (this *delConnTask) DoExec() {
	conn := this.tcp.connsOfHandle[this.conn]
	if conn == nil {
		return
	}

	conn.close()
	delete(this.tcp.connsOfHandle, this.conn)
	delete(this.tcp.connsOfAddr, conn.addr)

	if this.isDel && conn.isListen {
		del := &delNode{
			conn:    conn,
			delTime: time.Now().UnixMilli(),
		}
		this.tcp.delList.PushBack(del)
	}
}

func (this *tcp) disconn(conn amc.IConn, isDel bool) {
	task := &delConnTask{
		tcp:   this,
		conn:  conn,
		isDel: isDel,
	}
	this.tasks.Add(task)
	this.postNullRecv()
}

func (this *tcp) Disconn(conn amc.IConn) {
	this.disconn(conn, true)
}

func (this *tcp) procAsynTasks() {
	this.tasks.Exec()

	iter := this.delList.Front()
	if iter == nil {
		return
	}

	dNode := iter.Value.(*delNode)
	if time.Now().UnixMilli()-dNode.delTime >= DelConnTime {
		this.delList.Remove(iter)
		dNode.conn = nil
		dNode = nil
	}
}

func (this *tcp) RecvPack() (conn amc.IConn, addr string, msgid uint32, flag uintptr, pack []byte) {
	if this.isClose {
		return nil, "", 0, 0, nil
	}
	for {
		this.procAsynTasks()
		node := <-this.rChan
		if node.conn == nil {
			return nil, "", 0, 0, nil
		}
		if node.pack == nil {
			continue
		}
		return node.conn, node.addr, node.msgid, node.flag, node.pack
	}
}

func (this *tcp) reconn(conn *conn) (err error) {
	c, err := net.Dial("tcp", conn.addr)
	if err != nil {
		return
	}

	conn.c = c
	go this.goRecv(conn)
	return
}

func (this *tcp) SendPack(conn amc.IConn, msgid uint32, flag uintptr, pack []byte) (err error) {
	connObj := this.connsOfHandle[conn]
	if connObj == nil {
		return errors.New(fmt.Sprintf("tcp.SendPack=>无效Handle:%d", conn))
	}

	if connObj.isDis && (!connObj.isListen) {
		err = this.reconn(connObj)
		if err != nil {
			return
		}
	}

	err = connObj.sendPack(msgid, flag, pack)
	if err != nil {
		connObj.isDis = true
		connObj.close()
	}
	return
}
