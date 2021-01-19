package main

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	/*ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   false,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)*/
	file, err := ioutil.ReadFile("/home/jzd/buxybox_name_1.28.3.tar")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	imageNameOrSha256, isTagged, err := extraImageName(&file)
	if err := imageTag(imageNameOrSha256, "busybox:1.99"); err != nil {
		fmt.Println(err.Error())
	}
	if err = imageRemove("busybox:1.99"); err != nil {
		fmt.Println(err.Error())
	}
	if isTagged {
		if err = imageRemove(imageNameOrSha256); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func extraImageName(file *[]byte) (string, bool, error) {
	fmt.Println("reading image tar file.")
	reader := bytes.NewReader(*file)
	tr := tar.NewReader(reader)
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			fmt.Println(err.Error())
			return "", false, err
		}
		if hdr.Name != "manifest.json" {
			continue
		}
		// 读取文件信息
		var info []map[string]interface{}
		body, err := ioutil.ReadAll(tr)
		if err != nil {
			fmt.Println(err.Error())
			return "", false, err
		}
		if err := yaml.Unmarshal(body, &info); err != nil {
			fmt.Println(err.Error())
			return "", false, err
		}
		//get repoTags from manifest
		tags := info[0]["RepoTags"]
		//if docker save with {imageID}
		if tags == nil {
			name := info[0]["Config"].(string)
			return strings.Split(name, ".")[0], false, nil
		} else {
			//if docker save with {imageName}
			name := tags.([]interface{})[0]
			return name.(string), true, nil
		}
	}
	return "", false, errors.New(" can not found manifest.json in tar")
}

//source: docker image source name
//target: hoarbor tag images name
func imageTag(source string, target string) error {
	ctx := context.Background()
	//cli, err := client.NewClientWithOpts(client.FromEnv)
	cli, err := newDockerClient()
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	if err := cli.ImageTag(ctx, source, target); err != nil {
		return err
	}
	return nil
}

//img: docker image name
func imageRemove(img string) error {
	ctx := context.Background()
	//cli, err := client.NewClientWithOpts(client.FromEnv)
	cli, err := newDockerClient()
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	_, err = cli.ImageRemove(ctx, img, types.ImageRemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

func newDockerClient() (*client.Client, error) {
	/*hbinfo := hbclient.NewHarborInfo()
	if hbinfo.DockerTcp != "" {
		return client.NewClientWithOpts(client.WithHost(hbinfo.DockerTcp), client.WithVersion(""), client.WithHTTPClient(nil), client.WithHTTPHeaders(nil))

	}*/
	return client.NewClientWithOpts(client.FromEnv)
}
