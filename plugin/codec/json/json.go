package codec

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"encoding/json"
)

type Codec struct {
}

func (this *Codec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (this *Codec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func New() *Codec {
	return &Codec{}
}
