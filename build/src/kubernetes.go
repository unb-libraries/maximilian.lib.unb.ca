package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"os"
	"strings"
)

func getK8sConfig() *rest.Config {
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	if _, err := os.Stat(kubeconfig); errors.Is(err, os.ErrNotExist) {
		return getInClusterConfig()
	} else {
		return getFileKubeConfig(kubeconfig)
	}
}

func getK8sClient(config *rest.Config) corev1client.CoreV1Interface {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset.CoreV1()
}

func getFileKubeConfig(kubeconfig string) *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func getInClusterConfig() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return config
}

func getPodNameFromUri(client corev1client.CoreV1Interface, podUri string, podNamespace string) (string, string) {
	labelSelector := fmt.Sprintf("app.kubernetes.io/instance=%s", strings.ReplaceAll(podUri, ".", "-"))

	listOptions := metaV1.ListOptions{
		LabelSelector: labelSelector,
		Limit:         1,
	}
	pods, err := client.Pods(podNamespace).List(context.TODO(), listOptions)

	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		return pod.GetName(), ""
	}

	labelSelector = fmt.Sprintf("instance=%s", podUri)
	pods, err = client.Pods(podNamespace).List(context.TODO(), listOptions)

	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		return pod.GetName(), ""
	}

	return "", fmt.Sprintf("Error! No pods exist matching [uri=%s, namespace=%s]", podUri, podNamespace)
}

func executeRemotePodCommand(client corev1client.CoreV1Interface, config *rest.Config, podName string, podNamespace string, command string) (string, string, error) {
	buf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}
	request := client.RESTClient().
		Post().
		Namespace(podNamespace).
		Resource("pods").
		Name(podName).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Command: []string{"/bin/sh", "-c", command},
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
			TTY:     true,
		}, scheme.ParameterCodec)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", request.URL())
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: buf,
		Stderr: errBuf,
	})
	if err != nil {
		return "", "", fmt.Errorf("%w Failed executing command %s on %v/%v", err, command, podNamespace, podName)
	}

	return buf.String(), errBuf.String(), nil
}
