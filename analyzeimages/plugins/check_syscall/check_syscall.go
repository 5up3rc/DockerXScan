package check_syscall

import (
	"github.com/MXi4oyu/DockerXScan/probe/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
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

	//host
	hinfo,_:=host.Info()
	fmt.Println(hinfo.BootTime,hinfo.HostID,hinfo.Hostname,hinfo.KernelVersion,hinfo.OS,hinfo.Platform,hinfo.PlatformFamily,hinfo.PlatformVersion,hinfo.Procs,hinfo.Uptime,hinfo.VirtualizationRole,hinfo.VirtualizationSystem)

	//cpu
	cinfo,_:=cpu.Info()
	fmt.Println(cinfo)

	//disk
	dio,_:=disk.IOCounters()
	fmt.Println(dio)
	dpartition,_:=disk.Partitions(true)
	fmt.Println(dpartition)
	dusage,_:=disk.Usage("/")
	fmt.Println(dusage)

	//load
	lavg,_:=load.Avg()
	fmt.Println(lavg)
	lmisc,_:=load.Misc()
	fmt.Println(lmisc)

	//mem
	mvirtual,_:=mem.VirtualMemory()
	fmt.Println(mvirtual)

	//net
	nio,_:=net.IOCounters(true)
	fmt.Println(nio)
	ninter,_:=net.Interfaces()
	fmt.Println(ninter)

	//process

	pocs,_:=process.Processes()

	for _,poc :=range pocs{

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

		fmt.Println(ppid,pname,meminfo,cmdline,pconns,cpuaffi,cwd,exe,pctime,gids,pioc,pnice,pthreads,pusername)

	}

}