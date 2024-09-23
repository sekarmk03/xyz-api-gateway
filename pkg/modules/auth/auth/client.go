package auth

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/server"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc := server.InitGRPCConn(c.AuthSvcUrl, false, "")

	return pb.NewAuthServiceClient(cc)
}