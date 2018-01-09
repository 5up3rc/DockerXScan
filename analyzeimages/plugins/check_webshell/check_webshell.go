package check_webshell

import (

	"fmt"
	//"github.com/MXi4oyu/Utils/walkpath"
	"github.com/MXi4oyu/riskdetect/webshell"

)



func Check(check_path string)  {

	//check_files,_:=walkpath.WalkDir(check_path,"")

	//for _,f :=range check_files {

		//fmt.Println(f)
	//}

	phpshells:=webshell.Ssdeep("/gopath/src/github.com/MXi4oyu/riskdetect/libs/php.ssdeep",check_path,".php")
	jspshells:=webshell.Ssdeep("/gopath/src/github.com/MXi4oyu/riskdetect/libs/jsp.ssdeep",check_path,".jsp")
	aspshells:=webshell.Ssdeep("/gopath/src/github.com/MXi4oyu/riskdetect/libs/asp.ssdeep",check_path,".asp")


	phpl:=len(phpshells)
	for i:=0;i<phpl;i++{
		fmt.Println(phpshells[i]["type"],"----",phpshells[i]["path"],"----",phpshells[i]["level"],"----",phpshells[i]["like"])
	}

	jspl:=len(jspshells)
	for i:=0;i<jspl;i++{
		fmt.Println(jspshells[i]["type"],"----",jspshells[i]["path"],"----",jspshells[i]["level"],"----",jspshells[i]["like"])
	}

	aspl:=len(aspshells)
	for i:=0;i<aspl;i++{
		fmt.Println(aspshells[i]["type"],"----",aspshells[i]["path"],"----",aspshells[i]["level"],"----",aspshells[i]["like"])
	}

	//将检测结果提交到服务器



}