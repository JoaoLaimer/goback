package keyboard

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type InputEvent struct {
	TimeSec  int64
	TimeUSec int64
	Type     uint16
	Code     uint16
	Value    int32
}

var shiftIsPressed bool = false

func Setup(keyChan chan<- string) {
	devicePath, err := DetectKeyboardDevice()
	if err != nil {
		log.Fatalf("Error Detecting device: %v", err)
	}

	var wg sync.WaitGroup

	for _, path := range devicePath {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			readKeys(path, keyChan)
		}(path)
	}

	wg.Wait()

}

func readKeys(path string, keyChan chan<- string) {

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error tracking device (are you root?): %v", err)
	}
	defer f.Close()

	eventSize := int(24)
	buffer := make([]byte, eventSize)

	for {
		_, err := f.Read(buffer)
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		var event InputEvent
		err = binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &event)
		if err != nil {
			log.Println("Parsing Error: ", err)
			break
		}
		//fmt.Println(event.Code, " ", event.Type, " ", event.Value)

		if (event.Code == 42 || event.Code == 54) && event.Value == 1 {
			shiftIsPressed = true
		} else if (event.Code == 42 || event.Code == 54) && event.Value == 0 {
			shiftIsPressed = false
		}

		if event.Value == 1 && event.Type == 1 {
			keyName := mapKeycode(event.Code, CapsLockStatus(int(f.Fd())), shiftIsPressed)

			fmt.Printf("%s (code %d)\n", keyName, event.Code)

			select {
			case keyChan <- keyName:
			default:
			}
		}
	}
}

func mapKeycode(code uint16, caps bool, shift bool) string {
	keyMap := map[uint16]string{
		1: "ESC",
		2: "1", 3: "2", 4: "3", 5: "4", 6: "5", 7: "6", 8: "7", 9: "8", 10: "9", 11: "0", 12: "-", 13: "=", 14: "BACKSPACE",
		15: "TAB", 26: "[", 27: "]", 43: "\\",
		58: "CAPS", 39: ";", 40: "'", 28: "\n",
		86: "\\", 51: ",", 52: ".", 53: "/",
		29: "LCTRL", 125: "SUPER", 56: "LALT", 57: " ", 97: "RCTRL", 105: "LARROW", 108: "DARROW", 103: "UARROW", 106: "RARROW",
	}

	symbolsMap := map[uint16]string{
		2: "!", 3: "@", 4: "#", 5: "$", 6: "%", 7: "^", 8: "&", 9: "*", 10: "(", 11: ")", 12: "_", 13: "+",
		26: "{", 27: "}", 43: "|",
		39: ":", 40: "\"",
		86: "?", 51: "<", 52: ">", 53: "?",
		42: "LSHIFT", 54: "RSHIFT",
	}

	lettersMap := map[uint16]string{
		16: "Q", 17: "W", 18: "E", 19: "R", 20: "T", 21: "Y", 22: "U", 23: "I", 24: "O", 25: "P",
		30: "A", 31: "S", 32: "D", 33: "F", 34: "G", 35: "H", 36: "J", 37: "K", 38: "L",
		44: "Z", 45: "X", 46: "C", 47: "V", 48: "B", 49: "N", 50: "M",
	}
	if name, ok := lettersMap[code]; ok {
		if caps != shift {
			return strings.ToUpper(name)
		} else {
			return strings.ToLower(name)
		}
	} else if shift {
		if name, ok := symbolsMap[code]; ok {
			return name
		}
	} else if name, ok := keyMap[code]; ok {
		return name
	}

	return "."
}
