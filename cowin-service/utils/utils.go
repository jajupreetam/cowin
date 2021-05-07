package utils

import (
	"bytes"
	slog "github.com/go-eden/slf4go"
	"github.com/pquerna/ffjson/ffjson"
	"io"
	"runtime"
	"runtime/debug"
	"sync"
)

func Convert(from interface{}, to interface{}) error {
	bs, err := ffjson.Marshal(from)
	if err != nil {
		return err
	}
	return ffjson.Unmarshal(bs, to)
}

func GetReader(r interface{}) io.Reader {
	requestBytes, err := ffjson.Marshal(r)
	if err != nil || requestBytes == nil {
		panic("Error while marshalling")
	}

	return bytes.NewReader(requestBytes)
}

func RecoverRoutine(wg *sync.WaitGroup) {
	wg.Done()
	if err := recover(); err != nil {
		slog.Error("Error in GoRoutine, Error:", err)
		debug.PrintStack()
		runtime.Goexit()
	}
}
