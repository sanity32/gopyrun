package gopyrun

import (
	"bytes"
	"os/exec"
	"time"
)

func NewHandler(cmd *exec.Cmd) *Handler {
	c := Handler{
		start: time.Now(),
		cmd:   cmd,
	}
	cmd.Stdout = &c.Stdout
	cmd.Stderr = &c.Stderr
	return &c
}

type Handler struct {
	Stdout bytes.Buffer
	Stderr bytes.Buffer
	cmd    *exec.Cmd
	start  time.Time
}

func (h Handler) Passed() time.Duration {
	return time.Since(h.start)
}

func (h *Handler) Wait() error {
	return h.cmd.Wait()
}

func (h *Handler) Start() *Handler {
	h.cmd.Start()
	return h
}

func (h *Handler) Run() error {
	return h.cmd.Run()
}

func (h *Handler) Kill() error {
	if p := h.cmd.Process; p != nil {
		return p.Kill()
	}
	return nil
}
