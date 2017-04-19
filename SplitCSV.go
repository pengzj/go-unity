package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"path/filepath"
	"encoding/csv"
	"io"
	"strings"
)

func main()  {
	fmt.Println("please input csv file")
	reader := bufio.NewReader(os.Stdin)

	filename, err := reader.ReadString('\n');
	if  err != nil {
		log.Fatal(err)
	}
	filename = strings.TrimSpace(filename)

	fmt.Println("filename: " , filename)

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		log.Fatal(filename + " does not exist")
	}
	if err != nil {
		log.Fatal(err)
	}


	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	r := csv.NewReader(fp)
	dir := filepath.Dir(filename)
	basename := filepath.Base(filename)
	ext := filepath.Ext(filename)
	var name = basename[0:len(basename) - len(ext)]


	const Num = 50000

	var w io.WriteCloser
	var writer *csv.Writer
	var index int = 0
//	var header []string
	for counter := 0;; counter++{
		record, err := r.Read()
		if counter == 0 {
		//	header = record
		}
		if err == io.EOF {
			writer.Flush()
			w.Close()
			break
		}
		if err != nil {
			log.Fatal(err)
		}


		index = 0
		if counter % Num == 0 {
			index = counter / Num + 1
		}

		if index > 0 {
			filename = name + fmt.Sprintf("%04d", index) + ext
			filename = filepath.Join(dir, filename)
			w, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				log.Fatal(err)
			}
			writer = csv.NewWriter(bufio.NewWriter(w))

			//write header
			if counter > 0 {
				//writer.Write(header)

			}
			fmt.Println("generate file: ", filename)
		}


		err = writer.Write(record)
		if err != nil {
			log.Fatal(err)
		}

		if (counter + 1) % Num == 0 {
			writer.Flush()
			w.Close()
		}
	}
	fmt.Println("success")
}
