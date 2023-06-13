package common

import "log"

// @dev handle error exeptions
func HandleException(err error) {
	if err != nil {
		log.Panic(err)
	}
}