package utils

import (
	"dousheng/global"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/policy"
	"go.uber.org/zap"
)

func CreatMinioBUcket(bucketName string) {
	// 检查存储桶是否已经存在。
	exists, err := global.MinioClient.BucketExists(bucketName)
	fmt.Println(exists)
	if err == nil && exists {
		fmt.Printf(" already own bucket%s\n", bucketName)
		return
	}
	location := "us-east-1"
	err = global.MinioClient.MakeBucket(bucketName, location)
	if err != nil {
		fmt.Println(err, exists)
		return
	}

	//设置bucket权限
	err = global.MinioClient.SetBucketPolicy(bucketName, policy.BucketPolicyReadWrite)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully created %s\n", bucketName)
}

// UploadFile 上传文件给minio指定的桶中
func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (ok bool, err error) {
	_, err = global.MinioClient.PutObject(bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return false, err
	}
	return true, err
}

//  GetFileUrl 获取文件url
func GetFileUrl(bucketName string, fileName string, expires time.Duration) string {
	//time.Second*24*60*60
	reqParams := make(url.Values)
	presignedURL, err := global.MinioClient.PresignedGetObject(bucketName, fileName, expires, reqParams)
	if err != nil {
		zap.L().Error(err.Error())
		return ""
	}
	return fmt.Sprintf("%s", presignedURL)
}
