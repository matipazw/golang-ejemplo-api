package config

type Aws struct {
	AccessKeyId     string
	SecretAccessKey string
	Sns             AwsSns
}
