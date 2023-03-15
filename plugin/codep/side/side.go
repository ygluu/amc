package scodep

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"errors"
	"fmt"
	"unsafe"

	"lib/amc"
	"lib/amc/plugin/codep"
)

var MaxPackLen = uint32(5 * 1024 * 1024)

const PackLenSize = uint32(4)
const PackCmdSize = uint32(4)

type coder struct {
}

func New() amc.ICodep {
	return &coder{}
}

func (this *coder) CalcPackLen(dataLen uint32) uint32 {
	flag := uintptr(0)
	return dataLen + PackLenSize*2 + PackCmdSize + uint32(unsafe.Sizeof(flag))
}

func (this *coder) Enpack(cmd uint32, flag uintptr, src []byte, dest []byte) error {
	srcLen := uint32(len(src))
	if srcLen > MaxPackLen {
		return errors.New(fmt.Sprintf("sidep.coder.EncodePack=>封包时遇到包长度超限，上限：%d，实际：%d", MaxPackLen, srcLen))
	}

	destLen := uint32(len(dest))
	flagSize := uint32(unsafe.Sizeof(flag))
	packLen := PackCmdSize + flagSize + srcLen
	sumLen := PackLenSize*2 + packLen

	if sumLen > destLen {
		return errors.New(fmt.Sprintf("sidep.coder.EncodePack=>封包时遇到目标区域长度不足，需要：%d，实际：%d", sumLen, destLen))
	}

	codep.Uint32ToBuf(packLen, dest[:4])
	codep.Uint32ToBuf(cmd, dest[4:8])
	codep.UintptrToBuf(flag, dest[8:8+flagSize])
	copy(dest[8+flagSize:], src)
	codep.Uint32ToBuf(packLen, dest[sumLen-4:sumLen])

	return nil
}

func (this *coder) Depack(srcBuf *[]byte, srclen *uint32) (cmd uint32, flag uintptr, pack []byte, err error) {
	sLen := *srclen
	sBuf := *srcBuf

	flagSize := uint32(unsafe.Sizeof(flag))
	if sLen <= PackLenSize*2+PackCmdSize+flagSize {
		return 0, 0, nil, nil
	}

	packLen := codep.BufToUint32(sBuf[:4])
	if packLen-PackCmdSize-flagSize > MaxPackLen {
		return 0, 0, nil, errors.New(fmt.Sprintf("sidep.coder.DecodePack=>解包时遇到包长度超限，上限：%d，实际：%d",
			MaxPackLen, packLen-PackCmdSize-flagSize))
	}

	sumLen := packLen + PackLenSize*2
	if sumLen > sLen {
		return 0, 0, nil, nil
	}

	if packLen != codep.BufToUint32(sBuf[sumLen-4:sumLen]) {
		return 0, 0, nil, errors.New("sidep.coder.DecodePack=>解包时遇到无效封包格式")
	}

	cmd = codep.BufToUint32(sBuf[4:8])
	if cmd == 0 {
		return 0, 0, nil, errors.New("sidep.coder.DecodePack=>解包时遇到命令为0")
	}

	flag = codep.BufToUintptr(sBuf[8 : 8+flagSize])
	pack = sBuf[8+flagSize : sumLen-4]

	sLen -= sumLen
	if sLen > 0 {
		copy(sBuf, sBuf[sumLen:sumLen+sLen])
		*srclen = sLen
	} else {
		*srclen = 0
		*srcBuf = nil
	}

	err = nil
	return
}
