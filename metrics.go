package metrics

import (
	"encoding/json"
	"reflect"
	"strings"
	"fmt"
	"sync"
)

var events map[string]interface{}
var eventChan chan interface{}
var mu sync.Mutex

type eventValue struct {
	e string
	v float64
}

func init() {
	events = make(map[string]interface{})
	eventChan = make(chan interface{}, 1000)
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

func addEvent(e interface{}) {
	mu.Lock()
	defer mu.Unlock()

	isStruct := reflect.TypeOf(e).Kind() == reflect.Struct

	var split []string

	if isStruct {
		split = strings.Split(e.(eventValue).e, ":")
	} else {
		split = strings.Split(e.(string), ":")
	}

	if len(split) == 1 {
		if isStruct {
			if events[split[0]] == nil {
				events[split[0]] = make([]float64, 0)
			}
			if isSliceFloat(events[split[0]]) {
				events[split[0]] = append(events[split[0]].([]float64), e.(eventValue).v)
				return
			}
		} else {
			if events[split[0]] == nil {
				events[split[0]] = int(0)
			}
			if isInt(events[split[0]]) {
				i := events[split[0]].(int)
				i++
				events[split[0]] = i
				return
			}
		}
		fmt.Println("unable to set int or slice, already map - event: " + split[0])
		return
	}

	if events[split[0]] == nil {
		events[split[0]] = make(map[string]interface{})
	}
	if isInt(events[split[0]]) || isSliceFloat(events[split[0]]) {
		fmt.Println("unable to set map, already int or slice - event: " + split[0])
		return
	}

	entry := events[split[0]].(map[string]interface{})

	for key := 1; key < len(split); key++ {
		value := split[key]
		if key == len(split)-1 {
			if isStruct {
				if entry[value] == nil {
					entry[value] = make([]float64, 0)
				}
				if isSliceFloat(entry[value]) {
					entry[value] = append(entry[value].([]float64), e.(eventValue).v)
					return
				}
			} else {
				if entry[value] == nil {
					entry[value] = int(0)
				}
				if isInt(entry[value]) {
					i := entry[value].(int)
					i++
					entry[value] = i
					return
				}
			}
			fmt.Println("unable to set int or slice, already map - event: " + value)
			return
		} else {
			if entry[value] == nil {
				entry[value] = make(map[string]interface{})
			}
			if isMapInterface(entry[value]) {
				entry = entry[value].(map[string]interface{})
				continue
			}
			fmt.Println("unable to set map, already int or slice - event: " + value)
			return
		}
	}
}

func isMapInterface(obj interface{}) bool {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Map {
		return t.Elem().Kind() == reflect.Interface
	}
	return false
}

func isSliceFloat(obj interface{}) bool {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Slice {
		return t.Elem().Kind() == reflect.Float64
	}
	return false
}

func isInt(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Int
}

//Event adds +1 to event. Use ":" as separator to add depth to events. Example: ("Hello:World:love:kittens")
//Cannot overwrite event that is already a value. Don't do this: ("Hello:World"),("Hello:World:Dog")
func Event(e string) {
	eventChan <- e
}

//EventValue adds a value to event, will be saved in a slice of values. Example: ("User:Score", 134)
//can be used to calculate average or total sum
func EventValue(e string, value float64) {
	eventChan <- eventValue{
		e: e,
		v: value,
	}
}

func ExportJson() ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()
	return json.Marshal(events)
}

func ClearEvents() {
	mu.Lock()
	defer mu.Unlock()
	events = make(map[string]interface{})
}
