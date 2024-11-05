package infrastructure

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
)

type IOAdapter struct {
	r      io.Reader
	w      io.Writer
	logger *slog.Logger
}

func NewIOAdapter(r io.Reader, w io.Writer, logger *slog.Logger) *IOAdapter {
	return &IOAdapter{
		r:      r,
		w:      w,
		logger: logger,
	}
}

func (a *IOAdapter) Input() (string, error) {
	reader := bufio.NewReader(a.r)
	return reader.ReadString('\n')
}

func (a *IOAdapter) Output(content string) {
	_, err := fmt.Fprint(a.w, content)
	if err != nil {
		a.logger.Error(err.Error())
	}
}
