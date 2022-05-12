package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containerIDCh := make(chan string, 1)
	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true}
	go func() {
		for {
			id := <-containerIDCh
			out, err := cli.ContainerLogs(ctx, id, options)
			if err != nil {
				fmt.Println("get log error", err)
				continue
			}
			res, err := ioutil.ReadAll(out)
			if err != nil {
				fmt.Println("read log out error", err)
				continue
			}
			fmt.Println("*****************")
			fmt.Println(string(res))
			out.Close()
		}
	}()

	args := filters.NewArgs(filters.Arg("event", "die"))
	msgCh, errCh := cli.Events(ctx, types.EventsOptions{Filters: args})
	for {
		select {
		case data := <-msgCh:
			// fmt.Println("*****************")
			// fmt.Printf("%#v\n", data)
			containerIDCh <- data.ID
		case err := <-errCh:
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
