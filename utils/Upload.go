package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"personal-web-be/config"

	"github.com/minio/minio-go/v7"
)



func UploadFile(bucketName, objectName string, file multipart.File, fileSize int64, contentType string)  (string,error) {
	_ , err := config.MinioClient.PutObject(context.Background(),bucketName,objectName,file,fileSize,minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "" ,err 
	}
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", `inline; filename="`+objectName+`"`)

	baseUrl := config.MinioClient.EndpointURL().String()

	resultLink := fmt.Sprintf("%s/%s/%s", baseUrl,bucketName, objectName)
	
	return resultLink,nil
}