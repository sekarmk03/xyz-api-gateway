package transaction

import (
	"xyz-api-gateway/pkg/config"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/server"
)

type TransactionServiceClient struct {
	Client pb.TransactionServiceClient
}

func InitServiceClient(c *config.Config) pb.TransactionServiceClient {
	cc := server.InitGRPCConn(c.TransactionSvcUrl, false, "")

	return pb.NewTransactionServiceClient(cc)
}
