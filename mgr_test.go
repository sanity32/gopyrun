package gopyrun

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"
	"time"
)

//go:embed testTemplate.py
var TestPyTemplate string

var pathToVenvIor = `d:\p\W\V16\ior\master\ior_venv`

type TestPyTemplateData struct {
	Greetings string
	Filename  string
}

func Test_Ior(t *testing.T) {
	data := TestPyTemplateData{"Hiho", "fancyfile.tmp"}
	// defer func() { os.Remove(data.Filename) }()

	ior := New(TestPyTemplate, PyBinPathForVenv(pathToVenvIor, true), data)
	h := ior.Handler().Start()
	go func() { h.Wait(); t.Log("wait done") }()

	h.Kill()

}

func f() (*bytes.Buffer, <-chan any) {
	ch := make(chan any)
	var b bytes.Buffer
	go func() {
		for i := 0; i < 5; i++ {
			b.WriteString("add... ")
			time.Sleep(time.Second)
			fmt.Println("i#", i)
		}
		fmt.Println("channel closing")
		ch <- 0
		fmt.Println("channel closed")
	}()

	return &b, ch
}

func Test_B(t *testing.T) {
	b, ch := f()
	t.Log(b)
	time.Sleep(time.Second)
	t.Log(b)
	time.Sleep(time.Second)
	t.Log(b)
	time.Sleep(time.Second)
	t.Log(b)
	time.Sleep(time.Second)
	<-ch
	// t.Log()
}
