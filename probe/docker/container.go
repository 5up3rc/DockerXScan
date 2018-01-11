package docker


import (

	"fmt"

	"github.com/fsouza/go-dockerclient"
)


func ListAllContainers()  {

	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	cts,err:= client.ListContainers(docker.ListContainersOptions{All:false})
	if err != nil {
		panic(err)
	}
	for _, ct := range cts {
		fmt.Println("ID: ",ct.ID)
		fmt.Println("Names: ",ct.Names)
		fmt.Println("Image: ",ct.Image)
		fmt.Println("Created: ",ct.Created)
		fmt.Println("Ports: ",ct.Ports)
		fmt.Println("Labels: ",ct.Labels)
		fmt.Println("Command: ",ct.Command)
		fmt.Println("Mounts: ",ct.Mounts)
		fmt.Println("Networks: ",ct.Networks)
		fmt.Println("SizeRootFs: ",ct.SizeRootFs)
		fmt.Println("SizeRw: ",ct.SizeRw)
		fmt.Println("State: ",ct.State)
		fmt.Println("Status: ",ct.Status)


	}
}