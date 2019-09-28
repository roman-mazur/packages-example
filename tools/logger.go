package tools

import (
	"fmt"
	"log"
	"os"
)

func NewLogger(name string) *log.Logger {
	return log.New(os.Stderr, fmt.Sprintf("[%s] ", name), log.LstdFlags)
}
