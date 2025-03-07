// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/base/servicehub/examples/dependency/dependency"
)

type config struct {
	Name string `file:"name" default:"recallsong"`
}

type provider struct {
	Cfg *config
	Log logs.Logger
}

func (p *provider) Init(ctx servicehub.Context) error {
	dep1 := ctx.Service("example-dependency@label").(dependency.Interface)
	fmt.Println(dep1.Hello(p.Cfg.Name))

	dep2 := ctx.Service("example-dependency").(dependency.Interface)
	fmt.Println(dep2.Hello(p.Cfg.Name))
	return nil
}

func init() {
	servicehub.Register("hello-provider", &servicehub.Spec{
		Services:     []string{"hello"},
		Dependencies: []string{"example-dependency@label", "example-dependency"},
		Description:  "hello for example",
		ConfigFunc:   func() interface{} { return &config{} },
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}

func main() {
	hub := servicehub.New()
	hub.Run("examples", "", os.Args...)
}
