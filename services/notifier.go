package services

import (
	"api-messages/config"
	"api-messages/models"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/sns"
)

var Notify = func(message *models.Message) error {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(config.Provider().Aws.AccessKeyId, config.Provider().Aws.SecretAccessKey, ""),
		},
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	subject := GetMessageSubject(message.Client)
	topicArn := config.Provider().Aws.Sns.MessagesTopicArn
	snsMessage := struct {
		IdClient string
		UserName string
		Title    string
		Message  string
		To       string
		ToName   string
	}{
		message.Client.Id.Hex(),
		message.Sender.UserName,
		message.Title,
		message.Text,
		"",
		"",
	}

	client := sns.New(sess)

	for _, item := range FilterToList(message.To) {

		snsMessage.To = item.Email
		snsMessage.ToName = item.Name
		snsMessageText, _ := json.Marshal(&snsMessage)
		input := &sns.PublishInput{
			Subject:  aws.String(subject),
			TopicArn: aws.String(topicArn),
			Message:  aws.String(string(snsMessageText)),
		}

		_, err := client.Publish(input)

		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func FilterToList(list []models.To) (ret []models.To) {

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	test := func(t models.To) bool { return t.Email != "" && re.MatchString(t.Email) }

	for _, item := range list {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}

func GetMessageSubject(client models.Client) string {

	now := time.Now()
	return "messages-" + now.Format("2006-01-02 15:04:05") + "-" + client.Id.Hex()
}
