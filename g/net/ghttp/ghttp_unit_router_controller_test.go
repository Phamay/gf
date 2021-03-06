// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp_test

import (
    "fmt"
    "github.com/gogf/gf/g"
    "github.com/gogf/gf/g/frame/gmvc"
    "github.com/gogf/gf/g/net/ghttp"
    "github.com/gogf/gf/g/test/gtest"
    "testing"
    "time"
)

// 控制器
type ControllerBasic struct {
    gmvc.Controller
}

func (c *ControllerBasic) Init(r *ghttp.Request) {
    c.Controller.Init(r)
    c.Response.Write("1")
}

func (c *ControllerBasic) Shut() {
    c.Response.Write("2")
}

func (c *ControllerBasic) Index() {
    c.Response.Write("Controller Index")
}

func (c *ControllerBasic) Show() {
    c.Response.Write("Controller Show")
}


// 控制器注册测试
func Test_Router_Controller(t *testing.T) {
    p := ports.PopRand()
    s := g.Server(p)
    s.BindController("/", new(ControllerBasic))
    s.SetPort(p)
    s.SetDumpRouteMap(false)
    s.Start()
    defer s.Shutdown()

    // 等待启动完成
    time.Sleep(time.Second)
    gtest.Case(t, func() {
        client := ghttp.NewClient()
        client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
        
        gtest.Assert(client.GetContent("/"),            "1Controller Index2")
        gtest.Assert(client.GetContent("/init"),        "Not Found")
        gtest.Assert(client.GetContent("/shut"),        "Not Found")
        gtest.Assert(client.GetContent("/index"),       "1Controller Index2")
        gtest.Assert(client.GetContent("/show"),        "1Controller Show2")
        gtest.Assert(client.GetContent("/none-exist"),  "Not Found")
    })
}
