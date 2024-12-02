package Day2

import (
	"AdventOfCode/Helpers/ConvertHelper"
	"AdventOfCode/Helpers/FileHelper"
	"AdventOfCode/Helpers/ListHelper"
	"fmt"
)

//-------------------- Challange 1 --------------------

func RunChallenge1() {

	fmt.Println("\nCalculating Challange 1..")
	inputText := FileHelper.FileToStringList("Inputs/Day2/Input")
	inputList := ConvertHelper.StringSpliceToIntSpliceList(inputText)

	result := CalculateChallenge1(inputList)

	fmt.Println("Hohoho! Answer to Challenge 1 is: ", result)
}

func CalculateChallenge1(testList [][]int) (safeReports int) {
	safeReports = 0
	for i := 0; i < len(testList); i++ {
		if ListHelper.IsAscending(testList[i]) || ListHelper.IsDescending(testList[i]) {
			if ListHelper.DiffersByAtleast(testList[i], 1) && ListHelper.DiffersByAtmost(testList[i], 3) {
				safeReports++
			}
		}
	}
	return safeReports
}

//-------------------- Challange 2 --------------------

func RunChallenge2() {
	fmt.Println("\nCalculating Challange 2..")
	inputText := FileHelper.FileToStringList("Inputs/Day2/Input")
	inputList := ConvertHelper.StringSpliceToIntSpliceList(inputText)

	result := CalculateChallenge2(inputList)

	fmt.Println("Hohoho! Answer to Challenge 2 is: ", result)
}

func CalculateChallenge2(inputList [][]int) (safeReports int) {
	safeReports = 0
	for i := 0; i < len(inputList); i++ {
		if IsSafe(inputList[i]) {
			safeReports++
		}
	}
	return safeReports
}

func IsSafe(inputList []int) bool {
	for i := 0; i < len(inputList); i++ {
		currentList := ListHelper.RemoveIndex(inputList, i)
		if ListHelper.IsAscending(currentList) || ListHelper.IsDescending(currentList) {
			if ListHelper.DiffersByAtleast(currentList, 1) && ListHelper.DiffersByAtmost(currentList, 3) {
				return true
			}
		}
	}
	return false
}
