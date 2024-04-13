package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f := newFile("file.txt")
	fmt.Println("The file has been created successfully!")
	defer f.Close()

	f.write("Writting some bytes into the new created file")

	f.readAll()

	f.readByChunck(5)

	f.remove()
}

type File struct {
	*os.File
}

func newFile(name string) *File {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	return &File{file}
}

func (f *File) write(content string) {
	_, err := f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}

func (f *File) readByChunck(chunckSize int) error {
	reader := bufio.NewReader(f)
	buffer := make([]byte, chunckSize)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	return nil
}

func (f *File) readAll() {
	file, err := os.ReadFile(f.Name())
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
}

func (f *File) remove() {
	err := os.Remove(f.Name())
	if err != nil {
		panic(err)
	}
}
