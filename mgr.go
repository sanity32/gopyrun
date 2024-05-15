package gopyrun

import (
	"bytes"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

const (
	DEFAULT_BIN_PATH_LINUX   = "bin/python"
	DEFAULT_BIN_PATH_WINDOWS = `Scripts\python.exe`
)

func PyBinPathForVenv(venv string, forWindows bool) string {
	m := map[bool]string{true: DEFAULT_BIN_PATH_WINDOWS, false: DEFAULT_BIN_PATH_LINUX}
	return path.Join(venv, m[forWindows])
}

func New(tmplte string, pyBinPath string, data any) *Launcher {
	return &Launcher{
		Template:  tmplte,
		PyBinPath: pyBinPath,
		Data:      data,
	}
}

type Launcher struct {
	Template  string
	PyBinPath string
	Data      any
	Dir       string
}

func (sc Launcher) parsedTpl() (r bytes.Buffer, err error) {
	if t, err := template.New("scr").Parse(sc.Template); err != nil {
		return r, err
	} else {
		return r, t.Execute(&r, sc.Data)
	}
}

func (sc Launcher) stdizeLines() string {
	s, err := sc.parsedTpl()
	if err != nil {
		return ""
	}
	a := strings.Split(s.String(), "\n")
	r := []string{}
	for _, line := range a {
		line = strings.TrimPrefix(line, " ")
		if len(line) < 1 || line[0] == '#' {
			continue
		}
		r = append(r, ""+line)
	}
	return strings.Join(r, "\n")
}

func (sc *Launcher) Handler() *Handler {
	c := exec.Command(sc.PyBinPath, "-c", sc.stdizeLines())
	if d := sc.Dir; d != "" {
		c.Dir = d
	}
	return NewHandler(c)
}
