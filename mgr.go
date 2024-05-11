package gopyrun

import (
	"bytes"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"
)

const (
	DEFAULT_BIN_PATH_LINUX   = "bin/python"
	DEFAULT_BIN_PATH_WINDOWS = `Scripts\python.exe`
)

func NewIorVenv(tpl string) *Mgr {
	r := Mgr{
		Template: tpl,
		VenvPath: `d:\p\W\V16\ior\master\ior_venv`,
		BinPath:  `Scripts\python.exe`,
	}
	return &r
}

func New(tmplte string, venv string, binpath string, data any) *Mgr {
	return &Mgr{
		Template: tmplte,
		VenvPath: venv,
		BinPath:  binpath, //relative to venv
		Data:     data,
	}
}

func NewWindows(tpl string, venv string, data any) *Mgr {
	return New(tpl, venv, DEFAULT_BIN_PATH_WINDOWS, data)
}

func NewLinux(tpl string, venv string, data any) *Mgr {
	return New(tpl, venv, DEFAULT_BIN_PATH_WINDOWS, data)
}

type Mgr struct {
	Template string
	VenvPath string
	BinPath  string
	Data     any
}

func (sc Mgr) binPath() string {
	return path.Join(sc.VenvPath, sc.BinPath)
}

func (sc Mgr) parsedTpl() (r bytes.Buffer, err error) {
	if t, err := template.New("scr").Parse(sc.Template); err != nil {
		return r, err
	} else {
		return r, t.Execute(&r, sc.Data)
	}
}

func (sc Mgr) stdizeLines() string {
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

func (sc *Mgr) Run() (duration time.Duration, stdout bytes.Buffer, err error) {
	start := time.Now()
	c := exec.Command(sc.binPath(), "-c", sc.stdizeLines())
	c.Stdout = &stdout
	c.Stderr = &stdout
	err = c.Run()
	return time.Since(start), stdout, err
}
