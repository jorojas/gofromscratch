package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jorojas/gofromscratch/exercises"
)

var fileName string = "./files/txt/tabla.txt"

func SaveTable() {
	var text string = exercises.ShowMultiplicationTable()
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error creating file " + err.Error())
		return
	}

	fmt.Fprintln(file, text)
	file.Close()

}

func AddTable() {
	var text string = exercises.ShowMultiplicationTable()

	if !Append(fileName, text) {
		fmt.Println("Error appending content")
	}
}

func Append(fileN string, text string) bool {
	fileB, err := os.OpenFile(fileN, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error during file append " + err.Error())
		return false
	}

	_, err = fileB.WriteString(text)

	if err != nil {
		fmt.Println("Error during file writting " + err.Error())
		return false
	}

	fileB.Close()

	return true
}

func ReadFileIOUtil() {
	fileB, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error whilre reading file " + err.Error())
		return
	}

	fmt.Println(string(fileB))
}

func ReadFile() {
	fileB, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error whilre reading file " + err.Error())
		return
	}

	scanner := bufio.NewScanner(fileB)

	for scanner.Scan() {
		record := scanner.Text()

		fmt.Println("> " + record)
	}

	fileB.Close()
}
