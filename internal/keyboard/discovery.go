package keyboard

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func DetectKeyboardDevice() (string, error) {

	data, err := os.ReadFile("/proc/bus/input/devices")
	if err != nil {
		return "", fmt.Errorf("Failure reading devices: %v", err)
	}

	content := string(data)

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {

		if strings.Contains(block, "Handlers=") &&
			strings.Contains(block, "kbd") &&
			strings.Contains(block, "event") &&
			strings.Contains(block, "leds") {

			re := regexp.MustCompile(`event[0-9]+`)
			match := re.FindString(block)

			if match != "" {
				fullPath := fmt.Sprintf("/dev/input/%s", match)
				return fullPath, nil
			}
		}
	}
	return "", fmt.Errorf("No keyboard found")
}
