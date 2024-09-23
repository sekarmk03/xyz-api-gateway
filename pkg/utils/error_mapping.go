package utils

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var grpcToHttpCode = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Canceled:           http.StatusRequestTimeout,
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.Unauthenticated:    http.StatusUnauthorized,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.AlreadyExists:      http.StatusConflict,
}

var grpcToHttpStatus = map[codes.Code]string{
	codes.OK:                 "OK",
	codes.Canceled:           "REQUEST_TIMEOUT",
	codes.Unknown:            "INTERNAL_SERVER_ERROR",
	codes.InvalidArgument:    "BAD_REQUEST",
	codes.DeadlineExceeded:   "GATEWAY_TIMEOUT",
	codes.NotFound:           "NOT_FOUND",
	codes.PermissionDenied:   "FORBIDDEN",
	codes.Unauthenticated:    "UNAUTHORIZED",
	codes.ResourceExhausted:  "TOO_MANY_REQUESTS",
	codes.FailedPrecondition: "PRECONDITION_FAILED",
	codes.Aborted:            "CONFLICT",
	codes.OutOfRange:         "BAD_REQUEST",
	codes.Unimplemented:      "NOT_IMPLEMENTED",
	codes.Internal:           "INTERNAL_SERVER_ERROR",
	codes.Unavailable:        "SERVICE_UNAVAILABLE",
	codes.DataLoss:           "INTERNAL_SERVER_ERROR",
	codes.AlreadyExists:      "CONFLICT",
}

func GrpcToHttpCode(grpcCode codes.Code) int {
	httpCode, ok := grpcToHttpCode[grpcCode]
	if !ok {
		httpCode = http.StatusInternalServerError
	}

	return httpCode
}

func GrpcToHttpStatus(grpcCode codes.Code) string {
	httpStatus, ok := grpcToHttpStatus[grpcCode]
	if !ok {
		httpStatus = "INTERNAL_SERVER_ERROR"
	}

	return httpStatus
}
