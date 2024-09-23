package routes

import (
	"net/http"
	"strconv"
	"xyz-api-gateway/pkg/pb"
	"xyz-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CreateConsumerRequestBody struct {
	Nik                 string `json:"nik"`
	Fullname            string `json:"fullname"`
	LegalName           string `json:"legal_name"`
	BirthPlace          string `json:"birth_place"`
	BirthDate           string `json:"birth_date"`
	Salary              uint64 `json:"salary"`
	KtpBuffer           []byte `json:"ktp_buffer"`
	KtpFilename         string `json:"ktp_filename"`
	SelfiePhotoBuffer   []byte `json:"selfie_photo_buffer"`
	SelfiePhotoFilename string `json:"selfie_photo_filename"`
}

func CreateConsumer(ctx *gin.Context, c pb.ConsumerServiceClient) {
	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to parse multipart form")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	nik := ctx.PostForm("nik")
	fullname := ctx.PostForm("fullname")
	legalName := ctx.PostForm("legal_name")
	birthPlace := ctx.PostForm("birth_place")
	birthDate := ctx.PostForm("birth_date")
	salary := ctx.PostForm("salary")

	salaryInt, err := strconv.ParseUint(salary, 10, 64)

	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Failed to convert id to int")
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			errResp,
		)
		return
	}

	ktpFile, err := ctx.FormFile("ktp_file")
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "KTP file is required")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	ktpFileBytes, err := utils.FileToBytes(ktpFile)
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error", "Failed to read KTP file")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}

	selfieFile, err := ctx.FormFile("selfie_photo_file")
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusBadRequest, "Bad Request", "Selfie photo file is required")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResp)
		return
	}

	selfieFileBytes, err := utils.FileToBytes(selfieFile)
	if err != nil {
		errResp := utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error", "Failed to read selfie photo file")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errResp)
		return
	}

	res, err := c.CreateConsumer(ctx, &pb.ConsumerDataRequest{
		Nik:                 nik,
		Fullname:            fullname,
		LegalName:           legalName,
		BirthPlace:          birthPlace,
		BirthDate:           birthDate,
		Salary:              salaryInt,
		KtpBuffer:           ktpFileBytes,
		KtpFilename:         ktpFile.Filename,
		SelfiePhotoBuffer:   selfieFileBytes,
		SelfiePhotoFilename: selfieFile.Filename,
	})

	if err != nil {
		errResp := utils.GetGrpcError(err)
		ctx.AbortWithStatusJSON(
			errResp.Code,
			errResp,
		)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
