package FileHelper

import (
	"bufio"
	"fmt"
	"os"
)

func FileToStringList(inputFile string) (returnList []string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		returnList = append(returnList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return returnList
}
