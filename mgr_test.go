package gopyrun

import (
	_ "embed"
	"os"
	"testing"
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
	defer func() { os.Remove(data.Filename) }()

	ior := New(TestPyTemplate, PyBinPathForVenv(pathToVenvIor, true), data)
	t.Log(ior.Run())

}
