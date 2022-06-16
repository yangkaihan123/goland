package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.40"))
	defer cli.Close()
	log(err)
	fmt.Println("请输入镜像包路径：")
	reader := bufio.NewReader(os.Stdin)
	path, _ := reader.ReadString(':')
	loadimage(cli, path)
}

func listimages(cli *client.Client) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	log(err)

	for _, image := range images {
		fmt.Println(image)
	}
}

func loadimage(cli *client.Client, path string) error {
	file, err := os.Open(path)
	log(err)
	defer file.Close()
	_, err = cli.ImageLoad(context.TODO(), file, true)
	return err
}

func log(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}
}