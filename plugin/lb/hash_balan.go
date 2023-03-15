package lb

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"strings"
)

type hashNode struct {
	index int
	hash  uint32
	infos []*SvcInfo
}

func (this *hashNode) get() *SvcInfo {
	ret := this.infos[this.index]
	this.index++
	if this.index >= len(this.infos) {
		this.index = 0
	}
	ret.hitCount++
	return ret
}

type nodes []*hashNode

func (s nodes) Len() int {
	return len(s)
}

func (s nodes) Less(i, j int) bool {
	return s[i].hash < s[j].hash
}

func (s nodes) Swap(i, j int) {
	s[i].hash, s[j].hash = s[j].hash, s[i].hash
}

type HashBalan struct {
	replicas   int
	index      int
	infos      []*SvcInfo
	infoCount  int
	nodeOfHash map[uint32]*hashNode
	hashNodes  nodes
	hashRing   nodes
}

func NewHashBalan(replicas int) *HashBalan {
	ret := &HashBalan{
		replicas:   replicas,
		nodeOfHash: make(map[uint32]*hashNode),
	}
	return ret
}

func getHash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(fmt.Sprintf("%x", md5.Sum([]byte(key)))))
}

func (this *HashBalan) getHash(i int, key string) uint32 {
	return getHash(strconv.FormatUint(uint64(i), 10) + key)
}
func (this *HashBalan) GetHash(key string) uint32 {
	return getHash(key)
}

func (this *HashBalan) Add(info *SvcInfo) {
	this.infoCount++
	this.infos = append(this.infos, info)
}

func (this *HashBalan) AddAddrs(name, addrs string) {
	arr := strings.Split(addrs, ";")
	for _, addr := range arr {
		info := &SvcInfo{
			name: name,
			addr: addr,
		}
		this.Add(info)
	}
}

func (this *HashBalan) Calc() {
	sumr := this.replicas * len(this.infos)
	sumw := 0
	for _, info := range this.infos {
		sumw += info.weight
	}

	for _, info := range this.infos {
		count := this.replicas
		if info.weight > 0 {
			count = int(sumr * info.weight / sumw)
		}

		for i := 0; i < count; i++ {
			hash := this.getHash(i, info.addr)
			node := this.nodeOfHash[hash]
			if node == nil {
				node = &hashNode{hash: hash}
				this.nodeOfHash[hash] = node
			}
			node.infos = append(node.infos, info)
		}
	}

	this.hashNodes = make(nodes, len(this.nodeOfHash))
	this.hashRing = make(nodes, len(this.nodeOfHash))

	count := 0
	// nodeOfHash本身已经是Map散列形态，所有对hashNodes进行轮询即可实现均衡和按权重分配
	for _, node := range this.nodeOfHash {
		this.hashNodes[count] = node
		this.hashRing[count] = node
		count++
	}

	sort.Sort(this.hashRing)
}

func (this *HashBalan) GetInfo() (name, addr string) {
	info := this.Get()
	if info == nil {
		return "", ""
	}
	return info.name, info.addr
}

// 权重及哈希均衡分配（轮询）
func (this *HashBalan) Get() *SvcInfo {
	node := this.hashNodes[this.index]
	this.index++
	if this.index >= len(this.hashNodes) {
		this.index = 0
	}
	return node.get()
}

// 一致性哈希
func (this *HashBalan) GetByKey(key string) *SvcInfo {
	return this.GetByHash(getHash(key))
}

// 一致性哈希
func (this *HashBalan) GetByHash(hash uint32) *SvcInfo {
	index := sort.Search(len(this.hashRing), func(i int) bool { return this.hashRing[i].hash >= hash })
	if index >= len(this.hashRing) {
		index = 0
	}
	return this.hashRing[index].get()
}

func HashBalanTest() {

	svcNum := 10
	getcount := 10000
	replicas := 10000

	fmt.Println("哈希均衡")
	hb := NewHashBalan(replicas)
	for i := 0; i < 10; i++ {
		info := &SvcInfo{
			addr: fmt.Sprintf("Names%d", i),
		}
		hb.Add(info)
	}
	hb.Calc()

	for i := 0; i < getcount; i++ {
		hb.Get()
	}

	for _, info := range hb.infos {
		fmt.Printf("Name:%s, Weight:%d, HitCount:%d, HitRate:%0.1f%%\r\n",
			info.addr, info.weight, info.hitCount, float32(info.hitCount)/float32(getcount)*100)
	}

	fmt.Println("")

	fmt.Println("权重哈希")
	hb = NewHashBalan(replicas)
	for i := 0; i < svcNum; i++ {
		info := &SvcInfo{
			weight: 1 + i,
			addr:   fmt.Sprintf("Names%d", i),
		}
		hb.Add(info)
	}
	hb.Calc()

	for i := 0; i < getcount; i++ {
		hb.Get()
	}

	for _, info := range hb.infos {
		fmt.Printf("Name:%s, Weight:%d, HitCount:%d, HitRate:%0.1f%%\r\n",
			info.addr, info.weight, info.hitCount, float32(info.hitCount)/float32(getcount)*100)
	}
	fmt.Println("")

	fmt.Println("一致性哈希")
	hb = NewHashBalan(replicas)
	for i := 0; i < svcNum; i++ {
		info := &SvcInfo{
			addr: fmt.Sprintf("Names%d", i),
		}
		hb.Add(info)
	}
	hb.Calc()

	for i := 0; i < getcount; i++ {
		hb.GetByKey(fmt.Sprintf("key%d", i))
	}

	for _, info := range hb.infos {
		fmt.Printf("Name:%s, Weight:%d, HitCount:%d, HitRate:%0.1f%%\r\n",
			info.addr, info.weight, info.hitCount, float32(info.hitCount)/float32(getcount)*100)
	}

	fmt.Println("")
}
