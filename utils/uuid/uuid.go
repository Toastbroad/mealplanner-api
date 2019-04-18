package uuid

import (
	"log"
	"os/exec"
	"strings"
)

func GenerateUUID() string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(string(uuid), "\n")
}
