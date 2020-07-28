package parser

import (
	"fmt"
	"lem-in/config"
	"lem-in/types"
	"lem-in/utils"
)

//sore == start or end
//soreCheck function checks if start/end  has valid parameters,
//or if there is start/end dublicate.
func soreCheck(arr *[]string, usedIndexes *[]int, found *bool, sorePointer *types.Room, index int, sore string) {
	if !*found {
		if index < len(*arr)-1 {
			if x, y, valid := validRoom((*arr)[index+1], &sorePointer.Name); !valid {
				utils.InvalidInput(-1, fmt.Sprintf(config.ErrorSore, sore))
			} else {
				*usedIndexes = append(*usedIndexes, index)
				sorePointer.X = x
				sorePointer.Y = y
			}
			*found = true
		} else {
			utils.InvalidInput(-1, fmt.Sprintf(config.ErrorNoSoreCoords, sore))
		}
	} else {
		utils.InvalidInput(-1, fmt.Sprintf(config.ErrorAnotherSore, sore))
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
