package common

import (
	"github.com/iris-contrib/go.uuid"
	"log"
)

func GenUUID() string {
	if uuid, err := uuid.NewV4(); err != nil {
		log.Println("common genuuid  error: ", err)
		return ""
	} else {
		return uuid.String()
	}

}
