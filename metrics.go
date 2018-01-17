package metrics

import (
	"encoding/json"
	"reflect"
	"strings"
	"fmt"
	"sync"
)

var events map[string]interface{}
var eventChan chan string
var mu sync.Mutex

func init() {
	events = make(map[string]interface{})
	eventChan = make(chan string, 1000)
	mu = sync.Mutex{}
	go func() {
		for {
			select {
			case e, ok := <-eventChan:
				if !ok {
					return
				}
				addEvent(e)
			}
		}
	}()
}

func addEvent(e string) {
	mu.Lock()
	defer mu.Unlock()

	split := strings.Split(e, ":")

	if len(split) == 1 {
		if events[split[0]] == nil {
			events[split[0]] = int(0)
		}
		if isInt(events[split[0]]) {
			i := events[split[0]].(int)
			i++
			events[split[0]] = i
			return
		}
		fmt.Println("unable to set int, already map - event: " + split[0])
		return
	}

	if events[split[0]] == nil {
		events[split[0]] = make(map[string]interface{})
	}
	if isInt(events[split[0]]) {
		fmt.Println("unable to set map, already int - event: " + split[0])
		return
	}

	entry := events[split[0]].(map[string]interface{})

	for key := 1; key < len(split); key++ {
		value := split[key]
		if key == len(split)-1 {
			if entry[value] == nil {
				entry[value] = int(0)
			}
			if isInt(entry[value]) {
				i := entry[value].(int)
				i++
				entry[value] = i
				return
			}
			fmt.Println("unable to set int, already map - event: " + value)
			return

		} else {
			if entry[value] == nil {
				entry[value] = make(map[string]interface{})
			}
			if isMapInterface(entry[value]) {
				entry = entry[value].(map[string]interface{})
				continue
			}
			fmt.Println("unable to set map, already int - event: " + value)
			return
		}
	}
}

func isMapInterface(obj interface{}) bool {
	if isInt(obj) {
		return false
	}
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Map {
		if t.Elem().Kind() == reflect.Interface {
			return true
		}
	}
	return false
}

func isInt(obj interface{}) bool {
	if reflect.TypeOf(obj).Kind() == reflect.Int {
		return true
	}
	return false
}

func Event(e string) {
	eventChan <- e
}

func ExportJson() ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()

	b, err := json.Marshal(events)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ClearEvents() {
	mu.Lock()
	defer mu.Unlock()
	events = make(map[string]interface{})
}
