package Day1

import (
	"AdventOfCode/Helpers/ConvertHelper"
	"AdventOfCode/Helpers/FileHelper"
	"AdventOfCode/Helpers/ListHelper"
	"AdventOfCode/Helpers/SortHelper"
	"fmt"
)

//-------------------- Challange 1 --------------------

func RunChallenge1() {
	fmt.Println("\nCalculating Challange 1..")
	inputText := FileHelper.FileToStringList("Inputs/Day1/Input")
	inputList1, inputList2 := ConvertHelper.StringToTwoSplices(inputText)

	result := CalculateChallenge1(inputList1, inputList2)

	fmt.Println("Hohoho! Answer to Challenge 1 is: ", result)
}

func CalculateChallenge1(testList1 []int, testList2 []int) (sum int) {
	sortedList1 := SortHelper.SortList(testList1)
	sortedList2 := SortHelper.SortList(testList2)
	returnList := ListHelper.GetListDiff(sortedList1, sortedList2)
	sum = ListHelper.GetListSum(returnList)
	return sum
}

//-------------------- Challange 2 --------------------

func RunChallenge2() {
	fmt.Println("\nCalculating Challange 2..")
	inputText := FileHelper.FileToStringList("Inputs/Day1/Input")
	inputList1, inputList2 := ConvertHelper.StringToTwoSplices(inputText)

	result := CalculateChallenge2(inputList1, inputList2)

	fmt.Println("Hohoho! Answer to Challenge 2 is: ", result)
}

func CalculateChallenge2(testList1 []int, testList2 []int) (similarityScore int) {
	similarityScore = 0
	for i := 0; i < len(testList1); i++ {
		occurenceCount := ListHelper.GetTotalOccurenceInList(testList1[i], testList2)
		similarityScore += occurenceCount * testList1[i]
	}
	return similarityScore
}
