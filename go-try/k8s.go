package main

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "./kube/config")
}
