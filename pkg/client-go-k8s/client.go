package clientgok8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	k8sClient kubernetes.Interface
}

func (c *Client) restClient() {
	// https://blog.csdn.net/a1369760658/article/details/135752515
	// 最基础的客户端，提供与APIServer通信的最基本封装，可以向APIServer发送 Restful 风格请求。
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}
	pods := v1.PodList{}
	err = restClient.Get().Resource("pods").Do(context.TODO()).Into(&pods)
	if err != nil {
		panic(err)
	}
	// 打印pod名称
	for _, pod := range pods.Items {
		println(pod.Name)
	}
}
