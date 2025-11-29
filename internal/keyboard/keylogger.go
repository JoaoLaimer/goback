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

		if event.Value == 1 && event.Type == 1 {
			keyName := mapKeycode(event.Code, CapsLockStatus(int(f.Fd())))
			fmt.Printf("%s (code %d)\n", keyName, event.Code)

			select {
			case keyChan <- keyName:
			default:
			}
		}
	}
}

func mapKeycode(code uint16, caps bool) string {
	keyMap := map[uint16]string{
		1: "ESC",
		2: "1", 3: "2", 4: "3", 5: "4", 6: "5", 7: "6", 8: "7", 9: "8", 10: "9", 11: "0",
		28: "\n", 57: " ",
		16: "Q", 17: "W", 18: "E", 19: "R", 20: "T", 21: "Y", 22: "U", 23: "I", 24: "O", 25: "P",
		30: "A", 31: "S", 32: "D", 33: "F", 34: "G", 35: "H", 36: "J", 37: "K", 38: "L",
		44: "Z", 45: "X", 46: "C", 47: "V", 48: "B", 49: "N", 50: "M",
		42: "SHIFT",
	}
	if name, ok := keyMap[code]; ok {

		if caps {
			return strings.ToUpper(name)
		} else {
			return strings.ToLower(name)
		}
	}
	return "."
}
