# Ejemplo de API en golang 

Les comparto un ejemplo de un API utilizando golang como lenguaje de desarrollo utilizando varios frameworks interesantes :

- Gin Gonic (https://github.com/gin-gonic/gin) para desarrollar API rest.
- Gin Gonic JWT (https://github.com/appleboy/gin-jwt). Middleware para gestión de seguridad a través de JWT.
- Viper (https://github.com/spf13/viper) para gestión de archivos de configuración en yml.
- Mongo Driver (https://github.com/mongodb/mongo-go-driver) para conexión a base de datos MongoDb.
- AWS SNS (https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/sns/SnsPublish.go) para gestión de notificaciones a través de un modelo de publicador y suscriptores.


### ¿Cómo corro el ejemplo?
go run main.go

### ¿Cómo armo el ejecutable para deployarlo?

Si queremos que sea para entorno linux :
  env GOOS=linux GOARCH=amd64 go build main.go

### Packages necesarios

export GO111MODULE=on
go get github.com/aws/aws-sdk-go
