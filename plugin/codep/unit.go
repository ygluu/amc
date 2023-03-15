package codep

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"unsafe"
)

func BufToUint32(datas []byte) uint32 {
	return uint32(datas[0]) + (uint32(datas[1]) << 8) + (uint32(datas[2]) << 16) + (uint32(datas[3]) << 24)
}

func Uint32ToBuf(value uint32, buf []byte) {
	buf[0] = byte(value)
	buf[1] = byte(value >> 8)
	buf[2] = byte(value >> 16)
	buf[3] = byte(value >> 24)
}

func UintptrToBuf(value uintptr, buf []byte) {
	l := (int)(unsafe.Sizeof(value))
	for i := 0; i < l; i++ {
		value >>= 8 * i
		buf[i] = byte(value)
	}
}

func BufToUintptr(buf []byte) (ret uintptr) {
	l := (int)(unsafe.Sizeof(ret))
	ret = 0
	for i := 0; i < l; i++ {
		ret = ret | ((uintptr)(buf[i]))<<(8*i)
	}
	return ret
}
