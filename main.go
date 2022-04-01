package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"log"
)

type Student struct {
	Name   string
	Age    string
	Gender string
	Sport  []string
	Height string
	Weight string
}
type DetailWriter interface {
	WriteToFile(file string) error
}

type TJson struct {
	TJson []Student
}

type TYaml struct {
	TYaml []Student
}

func (a *TJson) WriteToFile(file string) error {
	b, err := json.MarshalIndent(a.TJson, "", "")
	if err != nil {
		return err
	}
	err1 := ioutil.WriteFile(file, b, 0644)
	if err1 != nil {
		return err1
	}
	return errors.New("no errors")

}

func (c *TYaml) WriteToFile(file string) error {
	d, err := yaml.Marshal(c.TYaml)
	if err != nil {
		return err
	}
	err1 := ioutil.WriteFile(file, d, 0644)
	if err1 != nil {
		return err1
	}
	return errors.New("no errors")
}

func CreateFile(item DetailWriter, file string) error {
	err := item.WriteToFile(file)
	return err
}

func main() {
	slice := make([]Student, 0)
	fmt.Println("Enter file name (same directory)")

	var fileName string
	fmt.Scan(&fileName)

	file, e := os.Open(fileName)
	if e != nil {
		fmt.Println("Error is = ", e)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		s := strings.Split(scanner.Text(), "[")

		//fmt.Println(s)
		p := strings.Split(s[0], ",")
		//fmt.Println(p)

		q := strings.Split(s[1], "]")
		//fmt.Println(q)

		q1 := strings.Split(q[0], ",")
		r := strings.Split(q[1], ",")
		//fmt.Println(r[1])

		var name Student
		name.Name, name.Age, name.Gender, name.Sport, name.Height, name.Weight = p[0], p[1], p[2], q1, r[1], r[2]
		//fmt.Println(name)
		slice = append(slice, name)

	}

	byte, err := json.MarshalIndent(&slice, "", "")
	if err != nil {
		fmt.Println(err)
		return

	}
	//fmt.Println(string(e))
	err = ioutil.WriteFile("test.json", byte, 0644)

	file.Close()

	for _, v := range slice {
		fmt.Println(v.Name, v.Age, v.Gender, v.Sport, v.Height, v.Weight)
		//fmt.Println(v)

	}
	var tea TYaml

	tea.TYaml = slice

	err = CreateFile(&tea, "test.yaml")

	var eat TJson

	eat.TJson = slice

	err = CreateFile(&eat, "test1.json")

	file, err2 := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err2 != nil {
		log.Fatal(err2)
	}

	log.SetOutput(file)

	log.Println(" created test.yaml")
	log.Println("created test1.json")

}
