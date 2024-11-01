package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vbargl/go-roomba/pkg/roomba"
)

func Status(roombaClient *roomba.Roomba) {
	status := roombaClient.GetStatus(10000)
	jsonStatus, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}
	fmt.Println(string(jsonStatus))
}
