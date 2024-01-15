package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	//"io"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// Refer to https://docs.docker.com/engine/api/sdk/examples/
func execute(data map[string]string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {

		// Error in creation of docker client
		return "", err
	}

	id, lang, timelimit, memorylimit :=
		data["id"], data["lang"], data["timelimit"], data["memorylimit"]
	
	image_name := lang_image_map[lang]
	fmt.Println(id,lang , timelimit , memorylimit , image_name)
	// Image Pull
	// reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	// if err != nil {
	//     panic(err)
	// }
	// io.Copy(os.Stdout, reader)

	// Path to the directory where the
	// files are mounted (in the host)
	location, _ := os.Getwd()
	// Modified to work on windows as well
	location = filepath.Join(filepath.Dir(location), "interface", bind_mnt_dir)
	// Create int variable for memory limit
	var mlimit int
	mlimit, err = strconv.Atoi(memorylimit[:len(memorylimit)-2])
	if err != nil {
		return "", err
	}

	//Container creation
	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: image_name,
			Cmd:   []string{id, lang_extension_map[lang], bind_mnt_dir, timelimit},
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: location,
					Target: "/home/" + unp_user + "/" + bind_mnt_dir,
				},
			},
			Resources: container.Resources{Memory: int64(mlimit * 1e6)},
		},
		nil,
		nil,
		"")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return "", err
	}
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	case <-statusCh:
	}
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", err
	}
	defer out.Close()

	std_output, std_err := new(strings.Builder), new(strings.Builder)
	_, err = stdcopy.StdCopy(std_output, std_err, out)

	if err != nil {
		return "", err
	}

	return std_output.String(), err
}
