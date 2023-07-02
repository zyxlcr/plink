package utils

import (
	"chatcser/config"
)

func CheckError(err error) {
	if err != nil {
		config.GVA_LOG.Error(err.Error())
	}
}
