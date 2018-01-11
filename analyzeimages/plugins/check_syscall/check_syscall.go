package check_syscall

import (
	"github.com/MXi4oyu/DockerXScan/probe/docker"
	"fmt"
)

func Check()  {
	fmt.Println("###########################################")
	fmt.Println("ListAllContainers")
	fmt.Println("###########################################")
	docker.ListAllContainers()
	fmt.Println("###########################################")
	fmt.Println("ListNetworks")
	fmt.Println("###########################################")
	docker.ListNetworks()
}