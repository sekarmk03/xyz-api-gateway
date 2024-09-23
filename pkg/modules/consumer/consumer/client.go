package consumer

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/server"
)

type ConsumerServiceClient struct {
	Client pb.ConsumerServiceClient
}

func InitServiceClient(c *config.Config) pb.ConsumerServiceClient {
	cc := server.InitGRPCConn(c.ConsumerSvcUrl, false, "")

	return pb.NewConsumerServiceClient(cc)
}
