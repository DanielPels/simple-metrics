package metrics

import (
	"time"
	"testing"
	"fmt"
)

const sleepTime = 50

func TestEvent(t *testing.T) {
	expectedResult := string(`{"test":1}`)

	ClearEvents()

	Event("test")

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestEvent: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestEvent invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestMultipleOfOneEvent(t *testing.T) {
	expectedResult := string(`{"test":3}`)

	ClearEvents()

	Event("test")
	Event("test")
	Event("test")

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestMultipleOfOneEvent: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestMultipleOfOneEvent invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestMultipleDepthEvent(t *testing.T) {
	expectedResult := string(`{"depth1":{"depth2":{"depth3":{"depth":{"4":{"depth5":{"depth6":{"depth7":{"depth8":{"depth9":{"depth10":1}}}}}}}}}}}`)

	ClearEvents()

	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestMultipleDepthEvent: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestMultipleDepthEvent invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestMultipleDepthOfOneEvent(t *testing.T) {
	expectedResult := string(`{"depth1":{"depth2":{"depth3":{"depth":{"4":{"depth5":{"depth6":{"depth7":{"depth8":{"depth9":{"depth10":5}}}}}}}}}}}`)

	ClearEvents()

	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestMultipleDepthOfOneEvent: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestMultipleDepthOfOneEvent invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestEventValue(t *testing.T) {
	expectedResult := string(`{"test":[20.38]}`)

	ClearEvents()

	EventValue("test", 20.38)

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestEventValue: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestEventValue invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestMultipleOfOneEventValue(t *testing.T) {
	expectedResult := string(`{"test":[42.38,-0.38,32847.38]}`)

	ClearEvents()

	EventValue("test", 42.38)
	EventValue("test", -0.38)
	EventValue("test", 32847.38)

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestMultipleOfOneEventValue: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestMultipleOfOneEventValue invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestMultipleDepthOfOneEventValue(t *testing.T) {
	expectedResult := string(`{"dept1":{"depth2":{"depth3":[947.82]}}}`)

	ClearEvents()

	EventValue("dept1:depth2:depth3", 947.82)

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestMultipleDepthOfOneEventValue: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestMultipleDepthOfOneEventValue invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestMultipleDepthOfMultipleEventValue(t *testing.T) {
	expectedResult := string(`{"dept1":{"depth2":{"depth3":[2987.7,-9283,387.2]}}}`)

	ClearEvents()

	EventValue("dept1:depth2:depth3", 2987.7)
	EventValue("dept1:depth2:depth3", -9283)
	EventValue("dept1:depth2:depth3", 387.2)

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestMultipleDepthOfMultipleEventValue: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestMultipleDepthOfMultipleEventValue invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestInvalidEventString(t *testing.T) {
	expectedResult := string(`{}`)

	ClearEvents()

	Event("::")
	Event("hello:")
	Event(":hello")
	Event("depth0:depth1::depth2")

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestInvalidEventString: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestInvalidEventString invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}

func TestFinal(t *testing.T) {
	expectedResult := string(`{"depth1":{"depth2":2,"other":2},"hello":1,"life":[42,-180],"test":2,"valueDepth1":{"other1":[872,-48],"valueDepth2":[-99,273]}}`)

	ClearEvents()

	Event("hello")
	Event("test")
	Event("test")
	Event("depth1:depth2")
	Event("depth1:depth2")
	Event("depth1:other")
	Event("depth1:other")

	EventValue("life", 42)
	EventValue("life", -180)
	EventValue("valueDepth1:valueDepth2", -99)
	EventValue("valueDepth1:valueDepth2", 273)
	EventValue("valueDepth1:other1", 872)
	EventValue("valueDepth1:other1", -48)

	Event("::")

	time.Sleep(time.Millisecond * sleepTime)

	b, err := ExportJson()
	if err != nil {
		t.Error("TestMultipleDepthOfMultipleEventValue: export json")
		return
	}

	if string(b) != expectedResult {
		t.Error("TestMultipleDepthOfMultipleEventValue invalid return. Got: " + string(b) + " Should be: " + expectedResult)
	}
}
