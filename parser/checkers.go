package parser

import "lem-in/types"

//sore == start or end
//This function checks if start/end  has valid parameters,
//or if there is start/end dublicate.
func soreCheck(arr *[]string, usedIndexes *[]int, found *bool, sorePointer *types.Room, index int, sore string) {
	if !*found {
		if index < len(*arr)-1 {
			if x, y, valid := validRoom((*arr)[index+1], &sorePointer.Name); !valid {
				invalidInput(-1, "invalid "+sore+" room params")
			} else {
				*usedIndexes = append(*usedIndexes, index)
				sorePointer.X = x
				sorePointer.Y = y
			}
			*found = true
		} else {
			invalidInput(-1, "no "+sore+" room coords")
		}
	} else {
		invalidInput(-1, "another "+sore+" declaration")
	}
}

func indexIsFree(index int, usedIndexes *[]int) bool {
	for _, idx := range *usedIndexes {
		if idx == index {
			return false
		}
	}
	return true
}
