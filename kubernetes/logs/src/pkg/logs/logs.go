package logs

import (
	"bufio"
	"encoding/base64"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubernetes/logs/pkg/prompt"
	"log"
	"strconv"
	"strings"
)

type Inputs struct {
	Namespace 		string
	PodPartName 	string
	Kubeconfig 		string
}

func (in Inputs) Run() {
	kubeConfigBytes, _ := base64.StdEncoding.DecodeString(in.Kubeconfig)
	clientConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigBytes)
	if err != nil {
		log.Fatalln("Failed to load config. Verify if you set credential kubeconfig in base64 format.")
	}

	config, err := clientConfig.ClientConfig()
	if err != nil {
		log.Fatalln("Failed to load config. Verify if you set credential kubeconfig in base64 format.")
	}

	// create the client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create kubernetes client.")
	}

	pods, err := client.CoreV1().Pods(in.Namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list pods for namespace: %s.\n", in.Namespace)
	}
	if len(pods.Items) == 0 {
		log.Fatalf("No result pods to namespace: %s.\n", in.Namespace)
	}
	var podsFilter []v1.Pod
	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, in.PodPartName) {
			podsFilter = append(podsFilter, pod)
		}
	}
	if len(podsFilter) == 0 {
		log.Fatalf("No result pods to name: %s and namespace: %s.\n", in.PodPartName, in.Namespace)
	}
	var items []string
	for i, pod := range podsFilter {
		items = append(items, fmt.Sprint(i, " - ", pod.Name))
	}
	itemSelect , _ := prompt.List("Select Pod: ", items)
	ind, _ := strconv.Atoi(strings.Split(itemSelect, " - ")[0])
	podSelect := podsFilter[ind]

	var containersItems []string
	for _, container := range podSelect.Spec.Containers {
		containersItems = append(containersItems, container.Name)
	}
	containerSelect , _ := prompt.List("Select container: ", containersItems)

	podLogOpts := v1.PodLogOptions{Follow:true, Container:containerSelect}

	req := client.CoreV1().Pods(podSelect.Namespace).GetLogs(podSelect.Name, &podLogOpts)
	podLogs, err := req.Stream()
	scanner := bufio.NewScanner(podLogs)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
}
