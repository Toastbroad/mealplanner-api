package uuid

import (
	"log"
	"os/exec"
)

func GenerateUUID() []byte {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return uuid
}
