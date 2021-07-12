// https://www.youtube.com/watch?v=n2MLjGeK7qA
// http://algolab.valemak.com/tree-binary
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var strForApp = "\"CSV-Sorter\" is an app which sorts all lines alphabetically by the first value in each " +
	"line using the Tree Sort algorithm.\n" +
	"Flags:\n" +
	"\tflag -i > read file (-i=С:/docs/test.csv);\n" +
	"\tflag -h > data with header (-h);\n" +
	"\tflag -o > write data in file (-o=С:/docs/test.csv).\n" +
	"\tflag -r > reverse data (-r).\n"

//goland:noinspection GoPrintFunctions
func main() {
	fmt.Println(strForApp)

	/*Флаги.*/
	iFile := flag.String("i", "noValue", "using file for read")
	hHeader := flag.Bool("h", false, "using header")
	oFile := flag.String("o", "noValue", "using file for write")
	reverseOrder := flag.Bool("r", false, "reverse order")
	flag.Parse()

	/*Получаю лист csv строк для сортировки.*/
	linesArr := make([]string, 0)
	var err error
	if *iFile != "noValue" {
		linesArr, err = getListOfCsvStringsFromFile(iFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		linesArr, err = getListOfCsvStringsFromInput()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	/*Получаю ссылку на корневой объект Tree с распределенными данными.*/
	var header string
	var tree *Tree
	if *hHeader {
		header = linesArr[0]
		tree = binaryTreeSort(linesArr[1:])
	} else {
		tree = binaryTreeSort(linesArr)
	}

	/*Обратный сбор данных по дереву для печати.*/
	newLinesArr := make([]string, 0)
	tree.forEach(func(line string) {
		newLinesArr = append(newLinesArr, line)
	})

	/*Перевернуть список в случае наличия флага -r.*/
	if *reverseOrder {
		for i, j := 0, len(newLinesArr)-1; i < j; i, j = i+1, j-1 {
			newLinesArr[i], newLinesArr[j] = newLinesArr[j], newLinesArr[i]
		}
	}

	/*Добавить header.*/
	if *hHeader {
		newLinesArr = append([]string{header}, newLinesArr...)
	}

	/*Вывести результат в файл или консоль.*/
	if *oFile != "noValue" {
		err = ioutil.WriteFile(*oFile, []byte(strings.Join(newLinesArr, "\n")), 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Data written successfully!\nPass to data > [%s]\n", *oFile)
	} else {
		fmt.Println(strings.Join(newLinesArr, "\n"))
	}
}

func getListOfCsvStringsFromFile(filePath *string) ([]string, error) {
	bytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), "\n"), nil
}

func getListOfCsvStringsFromInput() ([]string, error) {
	newLinesArr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please, input csv string: ")
		scanner.Scan()
		text := scanner.Text()
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}

		if text != "" {
			newLinesArr = append(newLinesArr, text)
		} else {
			fmt.Println()
			break
		}
	}
	return newLinesArr, nil
}