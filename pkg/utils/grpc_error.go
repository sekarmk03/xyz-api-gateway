package utils

import (
	"log"
	"net/http"
	"google.golang.org/grpc/status"
)

func GetGrpcError(err error) *ErrorResponse {
	st, ok := status.FromError(err)
	
    if !ok {
        log.Println("ERROR: Failed to extract gRPC error response:", err)
        return NewErrorResponse(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Failed to extract gRPC error response")
    }

	errCode := GrpcToHttpCode(st.Code())
	status := grpcToHttpStatus[st.Code()]

	return NewErrorResponse(errCode, status, st.Message())
}