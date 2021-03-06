// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"runtime"

	"comail.io/go/colog"
	"comail.io/go/wincolog"
	"github.com/hawry/gote/cmd"
)

var logLevel string

func main() {
	colog.Register()
	if logLevel == "production" {
		colog.SetMinLevel(colog.LInfo)
		colog.SetFlags(log.Ltime)
	} else {
		colog.SetFlags(log.Lshortfile)
	}
	if runtime.GOOS == "windows" {
		colog.SetOutput(wincolog.Stdout())
	} else {
		if logLevel != "production" {
			colog.SetDefaultLevel(colog.LDebug)
		}
	}
	cmd.Execute()
}
