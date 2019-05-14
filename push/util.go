package push

import (
	"log"
	"strconv"
)

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Print(err)
	}

	return i
}
