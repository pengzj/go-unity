package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
)

func main() {
	xlsxDir := "/path/config/data/xlsx_dev"	
	filepath.Walk(xlsxDir, func(myPath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if strings.Contains(myPath, ".svn") {
			return nil
		}
		
		//如果是目录，过滤.svn
		if info.IsDir() {
			return nil
		}
		
		if filepath.Ext(info.Name()) != ".xlsx" {
			return nil
		}

		fmt.Printf("current path is %s \t", myPath)
		fmt.Println(info.Name(), info.Size(), info.IsDir())
		
		return nil
	})
}
