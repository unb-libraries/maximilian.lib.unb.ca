package main

import (
	"fmt"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

func executeDrupalUli(client corev1client.CoreV1Interface, config *rest.Config, podUri string, podNamespace string, userUid string) (string, string) {
	podName, errorMsg := getPodNameFromUri(client, podUri, podNamespace)
	if errorMsg != "" {
		return "", errorMsg
	}
	output, errorMsg, err := executeRemotePodCommand(client, config, podName, podNamespace, "/scripts/drupalUli.sh " + userUid)
	if err != nil {
		fmt.Println(errorMsg)
		panic(err.Error())
	}
	return output, ""
}
