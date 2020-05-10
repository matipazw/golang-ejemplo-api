# api-messages

## Ejemplo de API en golang utilizando :

- Gin Gonic (https://github.com/gin-gonic/gin) para desarrollar API rest.
- Gin Gonic JWT (https://github.com/appleboy/gin-jwt). Middleware para gestión de seguridad a través de JWT.
- Viper (https://github.com/spf13/viper) para gestión de archivos de configuración en yml.
- Mongo Driver (https://github.com/mongodb/mongo-go-driver) para conexión a base de datos MongoDb.
- AWS SNS (https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/sns/SnsPublish.go) para gestión de notificaciones a través de un modelo de publicador y suscriptores.


##running
go run main.go

## deployment
env GOOS=linux GOARCH=amd64 go build main.go

## packages necesarios

export GO111MODULE=on
go get github.com/aws/aws-sdk-go