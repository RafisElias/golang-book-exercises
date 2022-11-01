package main

import (
	"encoding/gob"
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
)

func main() {
	//	ReadFileOs()
	//	ReadFileIoutil()
	//
	//	OpenDir()
	//	OpenDirWalk()

	//	kids := []ordenacao.Person{
	//		{"Jill", 9},
	//		{"Jack", 10},
	//	}
	//	sort.Sort(ordenacao.ByName(kids))
	//	sort.Sort(ordenacao.ByAge(kids))

}

func ReadFileOs() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// obtém o tamanho do arquivo
	stat, err := file.Stat()
	if err != nil {
		return
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func ReadFileIoutil() {
	file, err := os.ReadFile("teste.txt")
	if err != nil {
		if os.IsNotExist(err) {
			CreateFile("teste.txt")
			return
		}
		fmt.Println(err)
		return
	}
	str := string(file)
	fmt.Println(str)
}

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.WriteString("Olá")
}

func OpenDir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func OpenDirWalk() {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path, info.Size())
		return nil
	})
}
