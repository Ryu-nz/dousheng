package Init

import (
	"dousheng/global"
	"log"

	"github.com/fatih/color"
	"github.com/minio/minio-go"
)

func MinIO() {
	minioInfo := global.Settings.MinioInfo
	// 初使化 minio client对象。false是关闭https证书校验
	minioClient, err := minio.New(minioInfo.Endpoint, minioInfo.ID, minioInfo.Key, false)
	if err != nil {
		log.Fatalln(err)
	}
	//客户端注册到全局变量中
	global.MinioClient = minioClient
	color.Green("MinIO connection setup")
}
