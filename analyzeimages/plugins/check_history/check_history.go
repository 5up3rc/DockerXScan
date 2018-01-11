package check_history

import (
	"os"
	"encoding/json"
	"fmt"
)

func Check(path string)(error) {

	mf,err:= os.Open(path + "/manifest.json")
	if err != nil {
		return err
	}
	defer mf.Close()

	type manifestItem struct {
		Config   string
		RepoTags []string
		Layers   []string
	}

	var manifest []manifestItem
	if err = json.NewDecoder(mf).Decode(&manifest); err != nil {
		return  err
	} else if len(manifest) != 1 {
		return  err
	}


	config:= manifest[0].Config

	fmt.Println("########################################")
	fmt.Println("########################################")

	fmt.Println("history_config_file:",config)

	//解析配置文件
	mhf,err:=os.Open(path + "/"+config)
	if err != nil {
		return err
	}

	defer mhf.Close()


	type configItem struct {
		Architecture string
		Config map[string]interface{}
		Container string
		Container_config map[string]interface{}
		Created string
		Docker_version string
		History []map[string]interface{}
		Osv string
		Rootfs map[string]interface{}

	}

	var cfg configItem

	if err = json.NewDecoder(mhf).Decode(&cfg); err != nil {
		return  err
	}

	mhistory:=cfg.History
	fmt.Println("########################################")
	fmt.Println("########################################")
	fmt.Println("Historys：")

	for i,_:= range mhistory{
		created:=mhistory[i]["created"]
		fmt.Println(created)
		createdby:=mhistory[i]["created_by"]
		fmt.Println(createdby)
	}




	return nil


}
