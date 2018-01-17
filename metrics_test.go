package metrics

import (
	"time"
	"testing"
)

func TestEvent(t *testing.T) {

	ClearEvents()

	Event("test")

	time.Sleep(time.Millisecond * 100)

	b, err := ExportJson()
	if err != nil {
		panic(err)
	}
	if string(b) != string(`{"test":1}`) {
		t.Fail()
	}
}

func TestMultipleOfOneEvent(t *testing.T) {
	ClearEvents()

	Event("test")
	Event("test")
	Event("test")

	time.Sleep(time.Millisecond * 100)

	b, err := ExportJson()
	if err != nil {
		panic(err)
	}

	if string(b) != string(`{"test":3}`) {
		t.Fail()
	}
}

func TestMultipleDepthEvent(t *testing.T) {
	ClearEvents()

	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")

	time.Sleep(time.Millisecond * 100)

	b, err := ExportJson()
	if err != nil {
		panic(err)
	}

	if string(b) != string(`{"depth1":{"depth2":{"depth3":{"depth":{"4":{"depth5":{"depth6":{"depth7":{"depth8":{"depth9":{"depth10":1}}}}}}}}}}}`) {
		t.Fail()
	}
}

func TestMultipleDepthOfOneEvent(t *testing.T) {
	ClearEvents()

	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")
	Event("depth1:depth2:depth3:depth:4:depth5:depth6:depth7:depth8:depth9:depth10")

	time.Sleep(time.Millisecond * 100)

	b, err := ExportJson()
	if err != nil {
		panic(err)
	}

	if string(b) != string(`{"depth1":{"depth2":{"depth3":{"depth":{"4":{"depth5":{"depth6":{"depth7":{"depth8":{"depth9":{"depth10":5}}}}}}}}}}}`) {
		t.Fail()
	}
}
