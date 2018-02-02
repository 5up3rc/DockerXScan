package check_syscall

import (
	"github.com/MXi4oyu/DockerXScan/probe/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
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

	fmt.Println("###########################################")
	fmt.Println("hostinfo")
	fmt.Println("###########################################")

	//host
	hinfo,_:=host.Info()
	fmt.Println(hinfo)

	fmt.Println("###########################################")
	fmt.Println("CUPinfo")
	fmt.Println("###########################################")
	//cpu
	cinfos,_:=cpu.Info()
	for _,cinfo :=range cinfos{

		fmt.Println(cinfo)
	}

	fmt.Println("###########################################")
	fmt.Println("Diskinfo")
	fmt.Println("###########################################")

	//disk
	dio,_:=disk.IOCounters()
	fmt.Println("###########################################")
	fmt.Println("disk.IOCounters")
	fmt.Println("###########################################")
	fmt.Println(dio)

	fmt.Println("###########################################")
	fmt.Println("disk.Partitions")
	fmt.Println("###########################################")
	dpartitions,_:=disk.Partitions(true)
	for _,dpartition := range dpartitions{
		fmt.Println(dpartition)
	}

	fmt.Println("###########################################")
	fmt.Println("disk.Usage")
	fmt.Println("###########################################")

	dusage,_:=disk.Usage("/")
	fmt.Println(dusage)

	fmt.Println("###########################################")
	fmt.Println("Meminfo")
	fmt.Println("###########################################")
	//mem
	mvirtual,_:=mem.VirtualMemory()
	fmt.Println(mvirtual)

	fmt.Println("###########################################")
	fmt.Println("net info")
	fmt.Println("###########################################")
	//net
	fmt.Println("###########################################")
	fmt.Println("net.IOCounters")
	fmt.Println("###########################################")
	nio,_:=net.IOCounters(true)
	fmt.Println(nio)
	fmt.Println("###########################################")
	fmt.Println("net.Interfaces")
	fmt.Println("###########################################")
	ninters,_:=net.Interfaces()
	for _,ninter :=range ninters{
		fmt.Println(ninter)
	}

	fmt.Println("###########################################")
	fmt.Println("processinfo")
	fmt.Println("###########################################")
	//process

	pocs,_:=process.Processes()

	for _,poc :=range pocs{
		fmt.Println("###########################################")
		fmt.Println("Process")
		fmt.Println("###########################################")

		ppid,_:=poc.Ppid()
		pname,_:=poc.Name()
		meminfo,_:=poc.MemoryInfo()
		cmdline,_:=poc.Cmdline()
		pconns,_:=poc.Connections()
		cpuaffi,_:=poc.CPUAffinity()
		cwd,_:=poc.Cwd()
		exe,_:=poc.Exe()
		pctime,_:=poc.CreateTime()
		gids,_:=poc.Gids()
		pioc,_:=poc.IOCounters()
		pnice,_:=poc.IOnice()
		pthreads,_:=poc.NumThreads()
		pusername,_:=poc.Username()

		fmt.Println("Ppid:",ppid)
		fmt.Println("Name:",pname)
		fmt.Println("MemoryInfo:",meminfo)
		fmt.Println("Cmdline:",cmdline)
		fmt.Println("Connections:",pconns)
		fmt.Println("CPUAffinity:",cpuaffi)
		fmt.Println("Cwd:",cwd)
		fmt.Println("Exe:",exe)
		fmt.Println("CreateTime",pctime)
		fmt.Println("Gids:",gids)
		fmt.Println("IOCounters:",pioc)
		fmt.Println("IOnice:",pnice)
		fmt.Println("NumThreads:",pthreads)
		fmt.Println("Username:",pusername)

	}

}
