package ListHelper

import (
	"AdventOfCode/Helpers/MathHelper"
)

func GetListSum(inputList []int) (returnValue int) {
	for i := 0; i < len(inputList); i++ {
		returnValue += inputList[i]
	}
	return returnValue
}

// Gets diff of elements in same pos from 2 lists to 1 list
func GetListDiff(testList1 []int, testList2 []int) (returnList []int) {
	for i := 0; i < len(testList1); i++ {
		returnList = append(returnList, MathHelper.AbsDiffInt(testList1[i], testList2[i]))
	}

	return returnList
}

func GetTotalOccurenceInList(value int, targetList []int) int {
	occurenceCount := 0
	for i := 0; i < len(targetList); i++ {
		if targetList[i] == value {
			occurenceCount++
		}
	}
	return occurenceCount
}

func IsAscending(inputList []int) bool {
	previousValue := -1
	for i := 0; i < len(inputList); i++ {
		if i == 0 {
			previousValue = inputList[i]
			continue
		}
		if previousValue < inputList[i] {
			previousValue = inputList[i]
			continue
		}
		return false
	}
	return true
}

func IsDescending(inputList []int) bool {
	previousValue := -1
	for i := 0; i < len(inputList); i++ {
		if i == 0 {
			previousValue = inputList[i]
			continue
		}
		if previousValue > inputList[i] {
			previousValue = inputList[i]
			continue
		}
		return false
	}
	return true
}

func DiffersByAtleast(inputList []int, atleastValue int) bool {
	previousValue := -1
	for i := 0; i < len(inputList); i++ {
		if i == 0 {
			previousValue = inputList[i]
			continue
		}
		diff := MathHelper.AbsDiffInt(previousValue, inputList[i])
		if diff >= atleastValue {
			previousValue = inputList[i]
			continue
		}
		return false
	}
	return true
}

func DiffersByAtmost(inputList []int, atmostValue int) bool {
	previousValue := -1
	for i := 0; i < len(inputList); i++ {
		if i == 0 {
			previousValue = inputList[i]
			continue
		}
		diff := MathHelper.AbsDiffInt(previousValue, inputList[i])
		if diff <= atmostValue {
			previousValue = inputList[i]
			continue
		}
		return false
	}
	return true
}

func RemoveIndex(inputList []int, targetIndex int) []int {
	returnList := []int{}
	for i := 0; i < len(inputList); i++ {
		if targetIndex == i {
			continue
		}
		returnList = append(returnList, inputList[i])
	}
	return returnList
}
