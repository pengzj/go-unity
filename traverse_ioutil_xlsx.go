package main
import (
	"io/ioutil"
	"fmt"
)

func listFile(p string) ([]string, error) {
	var fullpath string
	var fileList []string
	files , err := ioutil.ReadDir(p)
	if err != nil {
		fmt.Println(err)
	}
	
	for _, fi := range files {
		if fi.IsDir() {
			continue
		}
		fullpath = p + "/" + fi.Name()		
		fileList = append(fileList, fullpath)
	}

	return fileList , nil
}

func main() {
		xlsxDir := "/path/config/data/xlsx_dev"
		fileList, _ := listFile(xlsxDir)
		for _, file := range fileList {
			fmt.Println(file)
		}
}
