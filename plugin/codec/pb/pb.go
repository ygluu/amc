package pb

//  Asynchronous Microservice Cluster Framework(AMC)
//          异步微服务集群框架
//        Author: Yigui Lu (卢益贵)
//         WX/QQ: 48092788

// Creation by: 2018-2020

import (
	"github.com/golang/protobuf/proto"
)

type Codec struct {
}

func (this *Codec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (this *Codec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func New() *Codec {
	return &Codec{}
}
