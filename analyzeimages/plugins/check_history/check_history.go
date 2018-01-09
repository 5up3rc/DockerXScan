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

	fmt.Println("history_config_file:",config)

	//解析配置文件
	mhf,err:=os.Open(path + "/"+config)
	if err != nil {
		return err
	}

	defer mhf.Close()


	type configItem struct {
		Architecture string
		Config []string

	}




	return nil


}
