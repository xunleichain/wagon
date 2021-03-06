// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validate

type Logger interface {
	Printf(string, ...interface{})
	Println(string, ...interface{})
}

var logger Logger

func init() {
	logger = NoopLogger{}
}

func SetLogger(l Logger) {
	logger = l
}

type NoopLogger struct{}

func (l NoopLogger) Printf(fmt string, v ...interface{})  {}
func (l NoopLogger) Println(fmt string, v ...interface{}) {}

/*
import (
	"io/ioutil"
	"log"
	"os"
)

var PrintDebugInfo = false

var logger *log.Logger

func init() {
	w := ioutil.Discard

	if PrintDebugInfo {
		w = os.Stderr
	}

	logger = log.New(w, "", log.Lshortfile)
	log.SetFlags(log.Lshortfile)
}
*/
