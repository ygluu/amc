package lb

import (
	"testing"
)

func HashBalanTest(t *testing.T) {
	svcNum := 10
	getcount := 10000
	replicas := 10000

	t.Println("哈希均衡")
	hb := NewHashBalan(replicas)
	for i := 0; i < 10; i++ {
		info := &svcInfo{
			addr: t.Sprintf("Names%d", i),
		}
		hb.Add(info)
	}
	hb.Calc()

	for i := 0; i < getcount; i++ {
		hb.Get()
	}

	for _, info := range hb.infos {
		t.Logf("Name:%s, Weight:%d, HitCount:%d, HitRate:%0.1f%%\r\n",
			info.addr, info.weight, info.hitCount, float32(info.hitCount)/float32(getcount)*100)
	}

	t.Log("")

	t.Println("权重哈希")
	hb = NewHashBalan(replicas)
	for i := 0; i < svcNum; i++ {
		info := &svcInfo{
			weight: 1 + i,
			addr:   t.Sprintf("Names%d", i),
		}
		hb.Add(info)
	}
	hb.Calc()

	for i := 0; i < getcount; i++ {
		hb.Get()
	}

	for _, info := range hb.infos {
		t.Logf("Name:%s, Weight:%d, HitCount:%d, HitRate:%0.1f%%\r\n",
			info.addr, info.weight, info.hitCount, float32(info.hitCount)/float32(getcount)*100)
	}
	t.Log("")

	t.Println("一致性哈希")
	hb = NewHashBalan(replicas)
	for i := 0; i < svcNum; i++ {
		info := &svcInfo{
			addr: t.Sprintf("Names%d", i),
		}
		hb.Add(info)
	}
	hb.Calc()

	for i := 0; i < svcNum; i++ {
		for j := 0; j < getcount; j++ {
			hb.GetByKey(t.Sprintf("key%d", i))
		}
	}

	for _, info := range hb.infos {
		t.Logf("Name:%s, Weight:%d, HitCount:%d, HitRate:%0.1f%%\r\n",
			info.addr, info.weight, info.hitCount, float32(info.hitCount)/float32(getcount)*100)
	}

	t.Log("")
}
