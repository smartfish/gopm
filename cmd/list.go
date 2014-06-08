// Copyright 2014 Unknown
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package cmd

import (
	"fmt"
	"sort"

	"github.com/Unknwon/com"
	"github.com/Unknwon/goconfig"
	"github.com/codegangsta/cli"

	"github.com/gpmgo/gopm/modules/doc"
	"github.com/gpmgo/gopm/modules/errors"
)

var CmdList = cli.Command{
	Name:  "list",
	Usage: "list all dependencies of current project",
	Description: `Command list lsit all dependencies of current Go project

gopm list

Make sure you run this command in the root path of a go project.`,
	Action: runList,
	Flags: []cli.Flag{
		cli.BoolFlag{"verbose, v", "show process details"},
	},
}

func verSuffix(gf *goconfig.ConfigFile, name string) string {
	val := gf.MustValue("deps", name)
	if len(val) > 0 {
		val = " @ " + val
	}
	return val
}

func runList(ctx *cli.Context) {
	if err := setup(ctx); err != nil {
		errors.SetError(err)
		return
	}

	gf, _, imports, err := genGopmfile()
	if err != nil {
		errors.SetError(err)
		return
	}

	list := make([]string, 0, len(imports))
	for _, name := range imports {
		if !com.IsSliceContainsStr(list, name) {
			list = append(list, name)
		}
	}
	sort.Strings(list)

	fmt.Printf("Dependency list(%d):\n", len(list))
	for _, name := range list {
		name = doc.GetRootPath(name)
		fmt.Printf("-> %s%s\n", name, verSuffix(gf, name))
	}
}
