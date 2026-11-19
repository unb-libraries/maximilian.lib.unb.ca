package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/slack-io/slacker"
)

func drupalUliCmd(ctx *slacker.CommandContext) {
	request := ctx.Request()
	uri := request.Param("uri")
	namespace := request.StringParam("namespace", "prod")
	uid := request.StringParam("uid", "1")
	fmt.Printf("[drupal-uli] (%s) %s/%s\n", uid, uri, namespace)
	config := getK8sConfig()
	client := getK8sClient(config)
	drupalUli, errorMsg := executeDrupalUli(client, config, uri, namespace, uid)
	if errorMsg != "" {
		ctx.Response().Reply(errorMsg)
	} else {
		ctx.Response().Reply(formatSlackUliString(drupalUli))
	}
}

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "drupal-uli {uri} {namespace} {uid}",
		Description: "Retrieves a Drupal site admin ULI link",
		Examples:    []string{"drupal-uli bnald.lib.unb.ca prod 1"},
		Handler:     drupalUliCmd,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
