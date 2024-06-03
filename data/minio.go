package data

import (
	c "github.com/PLDao/gin-frame/config"
	"github.com/PLDao/gin-frame/internal/pkg/logger"
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/policy"
)

func initMinIO() {
	client, err := minio.New(
		c.Config.Minio.Endpoint+":"+c.Config.Minio.Port,
		c.Config.Minio.AccessKeyID,
		c.Config.Minio.SecretAccessKey,
		c.Config.Minio.UseSSL,
	)
	if err != nil {
		panic("Failed to connect to MinIO: " + err.Error())
	}

	logger.Logger.Sugar().Info(c.Config.Minio.Endpoint + ":" + c.Config.Minio.Port)
	logger.Logger.Sugar().Info(c.Config.Minio.AccessKeyID)
	logger.Logger.Sugar().Info(c.Config.Minio.SecretAccessKey)
	logger.Logger.Sugar().Info(c.Config.Minio.UseSSL)
	createMinoBuket(client, c.Config.Minio.BuckName)
	MinioClient = client

}

// create bucket
func createMinoBuket(client *minio.Client, bucketName string) {
	location := "us-east-1"
	err := client.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := client.BucketExists(bucketName)
		if err == nil && exists {
			logger.Logger.Info("MinIO bucket already exists")
		} else {
			logger.Logger.Error("Failed to create MinIO bucket: " + err.Error())
			return
		}
	}
	//
	err = client.SetBucketPolicy(bucketName, policy.BucketPolicyReadWrite)

	if err != nil {
		logger.Logger.Error("Failed to set bucket policy: " + err.Error())
		return
	}
	logger.Logger.Info("MinIO bucket created successfully")
}
