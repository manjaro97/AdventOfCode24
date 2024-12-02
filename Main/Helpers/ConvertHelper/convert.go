package ConvertHelper

import (
	"strconv"
	"strings"
)

func StringToList(input string) (returnList []string) {
	returnList = strings.Fields(input)
	return returnList
}

func StringToInt(input string) (returnValue int) {
	returnValue, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return returnValue
}

// Convert a List of strings into two lists of numbers. The strings contain a pair of numbers that are inserted into the number lists
func StringToTwoSplices(inputText []string) (returnList1 []int, returnList2 []int) {
	for i := 0; i < len(inputText); i++ {
		stringList := StringToList(inputText[i])
		value1 := StringToInt(stringList[0])
		value2 := StringToInt(stringList[1])

		returnList1 = append(returnList1, value1)
		returnList2 = append(returnList2, value2)
	}

	return returnList1, returnList2
}

// Convert a list of strings into a list of numbers
func StringSpliceToIntSplice(inputList []string) (returnList []int) {
	for i := 0; i < len(inputList); i++ {
		convertedValue := StringToInt(inputList[i])
		returnList = append(returnList, convertedValue)
	}
	return returnList
}

// Convert each 'string' from a 'List of Strings' into a 'list of numbers'. Return value will be a 'list of a list of numbers'
func StringSpliceToIntSpliceList(inputList []string) (returnList [][]int) {
	for i := 0; i < len(inputList); i++ {
		stringList := StringToList(inputList[i])
		intList := StringSpliceToIntSplice(stringList)
		returnList = append(returnList, intList)
	}
	return returnList
}
