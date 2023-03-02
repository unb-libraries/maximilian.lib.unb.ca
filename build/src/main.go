package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
)

func drupalUliCmd(actionID string) func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	return func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
		uri := request.Param("uri")
		namespace := request.Param("namespace")
		if namespace != "" {
			namespace := "prod"
		}
		userName := botCtx.Event().UserName
		fmt.Println(fmt.Sprintf("[drupal-uli] (%s) %s/%s", userName, uri, namespace))
		config := getK8sConfig()
		client := getK8sClient(config)
		drupalUli, errorMsg := executeDrupalUli(client, config, uri, namespace)
		if errorMsg != "" {
			response.Reply(errorMsg)
		} else {
			response.Reply(formatSlackUliString(drupalUli))
		}
	}
}

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.Command("drupal-uli {uri} {namespace}", &slacker.CommandDefinition{
		BlockID: "drupal-uli",
		Handler: drupalUliCmd("drupal-uli"),
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
