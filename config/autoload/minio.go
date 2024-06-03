package autoload

type MinioConfig struct {
	Enable          bool   `mapstructure:"enable"`
	Endpoint        string `mapstructure:"endpoint"`
	Port            string `mapstructure:"port"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	UseSSL          bool   `mapstructure:"use_ssl"`
	BuckName        string `mapstructure:"bucket_name"`
}

var Minio = MinioConfig{
	Enable:          false,
	Endpoint:        "127.0.0.1",
	Port:            "9003",
	AccessKeyID:     "admin",
	SecretAccessKey: "admin1234",
	UseSSL:          false,
	BuckName:        "test",
}
