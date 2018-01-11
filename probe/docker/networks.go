package docker


import (

	"fmt"

	"github.com/fsouza/go-dockerclient"
)


func ListNetworks()  {

	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	nts,err:= client.ListNetworks()
	if err != nil {
		panic(err)
	}
	for _, nt := range nts {

		info,_:=client.NetworkInfo(nt.ID)

		fmt.Println("ID: ",info.ID)
		fmt.Println("Name: ",info.Name)
		fmt.Println("Labels: ",info.Labels)
		fmt.Println("Containers: ",info.Containers)
		fmt.Println("Driver: ",info.Driver)
		fmt.Println("Internal: ",info.Internal)
		fmt.Println("IPAM: ",info.IPAM)
		fmt.Println("Options: ",info.Options)
		fmt.Println("Scope: ",info.Scope)

	}

}