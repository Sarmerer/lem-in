package parser

func soreCheck(arr *[]string, usedIndexes *[]int, found *bool, sorePointer *string, index int, sore string) { //sore == start or end
	if !*found {
		if index < len(*arr)-1 {
			if _, _, valid := validRoom((*arr)[index+1], sorePointer); !valid {
				invalidInput(-1, "invalid "+sore+" room params")
			} else {
				*usedIndexes = append(*usedIndexes, index)
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
