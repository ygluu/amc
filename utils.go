package amc

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

var objectIdCounter uint32 = 0

var machineId = readMachineId()

func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}

// GUID returns a new unique ObjectId.
// 4byte 时间，
// 3byte 机器ID
// 2byte pid
// 3byte 自增ID
func GetGUID() string {
	b := make([]byte, 12)
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return hex.EncodeToString(b)
}

// func GetCurrentThreadId() int {
// 	var user32 *syscall.DLL
// 	var GetCurrentThreadId *syscall.Proc
// 	var err error

// 	user32, err = syscall.LoadDLL("Kernel32.dll")
// 	if err != nil {
// 		fmt.Printf("syscall.LoadDLL fail: %v\n", err.Error())
// 		return 0
// 	}
// 	GetCurrentThreadId, err = user32.FindProc("GetCurrentThreadId")
// 	if err != nil {
// 		fmt.Printf("user32.FindProc fail: %v\n", err.Error())
// 		return 0
// 	}

// 	var pid uintptr
// 	pid, _, err = GetCurrentThreadId.Call()

// 	return int(pid)
// }
