package keyboard

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func DetectKeyboardDevice() ([]string, error) {

	data, err := os.ReadFile("/proc/bus/input/devices")
	if err != nil {
		return nil, fmt.Errorf("failure reading devices: %v", err)
	}

	content := string(data)

	blocks := strings.Split(content, "\n\n")

	var paths []string

	for _, block := range blocks {

		if strings.Contains(block, "Handlers=") &&
			strings.Contains(block, "kbd") &&
			strings.Contains(block, "event") &&
			strings.Contains(block, "leds") {

			re := regexp.MustCompile(`event[0-9]+`)
			match := re.FindString(block)

			if match != "" {
				paths = append(paths, fmt.Sprintf("/dev/input/%s", match))
			}
		}
	}

	fmt.Println(paths)
	if len(paths) > 0 {
		return paths, nil
	} else {
		return nil, fmt.Errorf("no keyboard found")
	}
}
