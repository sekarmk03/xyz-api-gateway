package consumer_limit

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/server"
)

type ConsumerLimitServiceClient struct {
	Client pb.ConsumerLimitServiceClient
}

func InitServiceClient(c *config.Config) pb.ConsumerLimitServiceClient {
	cc := server.InitGRPCConn(c.ConsumerSvcUrl, false, "")

	return pb.NewConsumerLimitServiceClient(cc)
}
