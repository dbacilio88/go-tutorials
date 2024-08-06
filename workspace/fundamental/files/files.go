package files

import (
	"bufio"
	"fmt"
	"github.com/dbacilio88/go/pkg/fundamental/bucles"

	"os"
	"time"
)

/**
*
* files
* <p>
* files file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author cbaciliod
* @author dbacilio88@outlook.es
* @since 8/07/2024
*
 */

var fileName = "./files/txt/table.txt"

func Execute() {
	readFile()
	writeFile()
	addWriteFile()
	DeleteFile()
}

func ValidateFileExist() bool {
	stat, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	fmt.Println(stat)
	return true
}

func DeleteFile() {

	if !ValidateFileExist() {
		fmt.Println("file does not exist")
		return
	}

	err := os.Remove(fileName)

	if err != nil {
		fmt.Println("Error deleting file ", err.Error())
		return
	}
}

func readFile() {

	bytes, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error to read file ", err.Error())
		return
	}

	scanner := bufio.NewScanner(bytes)

	for scanner.Scan() {
		fmt.Println("> ", scanner.Text())
	}

	err = bytes.Close()

	if err != nil {
		fmt.Println("error to close file ", err.Error())
		return
	}
}

func addWriteFile() {

	text := bucles.Iteraciones()

	if !Append(fileName, text) {
		fmt.Println("error to concat text")
	}
}

func Append(fileName string, text string) bool {

	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("error to open file: ", err.Error())
		return false
	}

	r, err := file.WriteString(text)

	if err != nil {
		fmt.Println("error to write string: ", err.Error())
		return false
	}

	err = file.Close()

	if err != nil {
		fmt.Println("error to close file: ", err.Error())
		return false
	}

	return r != 0

}

func writeFile() {

	start := time.Now()

	text := bucles.Iteraciones()

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("error create file: ", err.Error())
	}

	_, err = fmt.Fprintln(file, text)

	err = file.Close()

	if err != nil {
		fmt.Println("error close file: ", err.Error())
	}

	endTime := time.Now()

	duration := endTime.Sub(start)

	fmt.Println("elapsed time:", duration.String())
}
