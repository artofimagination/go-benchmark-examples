package tests

// Command: go test -v -bench=. -benchmem ./tests/

import (
	"os"
	"runtime/trace"
	"testing"
)

type TestData struct {
	a, b, c string
	d, e, f int
	g, h, i float64
}

func copyByValue(test TestData) int {
	test2 := test
	return test2.d
}

func copyByPointer(test *TestData) int {
	test2 := *test
	return test2.d
}

type Class struct {
}

func (c Class) copyByValueValueReceiver(test TestData) int {
	test2 := test
	return test2.d
}

func (c Class) copyByPointerValueReceiver(test *TestData) int {
	test2 := *test
	return test2.d
}

func (c *Class) copyByValuePointerReceiver(test TestData) int {
	test2 := test
	return test2.d
}

func (c *Class) copyByPointerPointerReceiver(test *TestData) int {
	test2 := *test
	return test2.d
}

func BenchmarkPassValue(b *testing.B) {
	test := TestData{
		a: "test1", b: "test2", c: "test3",
		d: 3, e: 4, f: 5,
		g: 2.3, h: 4.5, i: 6.7,
	}

	f, err := os.Create("byValue.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = copyByValue(test)
	}

	trace.Stop()

	b.StopTimer()
}

func BenchmarkPassPointer(b *testing.B) {
	test := TestData{
		a: "test1", b: "test2", c: "test3",
		d: 3, e: 4, f: 5,
		g: 2.3, h: 4.5, i: 6.7,
	}

	f, err := os.Create("byPointer.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = copyByPointer(&test)
	}

	trace.Stop()

	b.StopTimer()
}

func BenchmarkPassValueValueReceiver(b *testing.B) {
	class := Class{}
	test := TestData{
		a: "test1", b: "test2", c: "test3",
		d: 3, e: 4, f: 5,
		g: 2.3, h: 4.5, i: 6.7,
	}

	f, err := os.Create("byValueVRec.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = class.copyByValueValueReceiver(test)
	}

	trace.Stop()

	b.StopTimer()
}

func BenchmarkPassPointerValueReceiver(b *testing.B) {
	class := Class{}
	test := TestData{
		a: "test1", b: "test2", c: "test3",
		d: 3, e: 4, f: 5,
		g: 2.3, h: 4.5, i: 6.7,
	}

	f, err := os.Create("byPointerVRec.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = class.copyByPointerValueReceiver(&test)
	}

	trace.Stop()

	b.StopTimer()
}

func BenchmarkPassValuePointerReceiver(b *testing.B) {
	class := Class{}
	test := TestData{
		a: "test1", b: "test2", c: "test3",
		d: 3, e: 4, f: 5,
		g: 2.3, h: 4.5, i: 6.7,
	}

	f, err := os.Create("byValuePtrRec.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = class.copyByValuePointerReceiver(test)
	}

	trace.Stop()

	b.StopTimer()
}

func BenchmarkPassPointerPointerReceiver(b *testing.B) {
	class := Class{}
	test := TestData{
		a: "test1", b: "test2", c: "test3",
		d: 3, e: 4, f: 5,
		g: 2.3, h: 4.5, i: 6.7,
	}

	f, err := os.Create("byPtrPtrRec.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		_ = class.copyByPointerPointerReceiver(&test)
	}

	trace.Stop()

	b.StopTimer()
}
