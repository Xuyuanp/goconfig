/*
 * Copyright 2014 Xuyuan Pang <xuyuanp # gmail dot com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goconfig

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var config = `
# Lines start with '#' would be ignored as comments
foo 		= bar
a 			= 1
# Using ${key} to get the system environment var
workspace 	= ${PWD}

[dev]
ip 		= 127.0.0.1
port 	= 8080
# Using $(key) to get the local config var. The key must be defined above this line.
url 	= $(dev.ip):$(dev.port)
foo 	= dev$(foo)
`

func TestConfiguration(t *testing.T) {
	convey.Convey("Given a new Configuration", t, func() {
		convey.Convey(fmt.Sprintf("Given a config: %s", config), func() {
			rd := bytes.NewReader([]byte(config))
			convey.Convey("Load()", func() {
				err := Load(rd)
				convey.Convey("error should be nil", func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
			convey.Convey(`GetString("foo")`, func() {
				value, err := GetString("foo")
				convey.Convey(`value should be "bar"`, func() {
					convey.So(value, convey.ShouldEqual, "bar")
				})
				convey.Convey(`error should be nil`, func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
			convey.Convey(`GetInt("a")`, func() {
				value, err := GetInt("a")
				convey.Convey(`value should be 1`, func() {
					convey.So(value, convey.ShouldEqual, 1)
				})
				convey.Convey(`error should be nil`, func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
			convey.Convey(`GetString("workspace")`, func() {
				value, err := GetString("workspace")
				convey.Convey(`value should be os.Getenv("PWD")`, func() {
					convey.So(value, convey.ShouldEqual, os.Getenv("PWD"))
				})
				convey.Convey(`error should be nil`, func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
			convey.Convey(`GetString("dev.ip")`, func() {
				value, err := GetString("dev.ip")
				convey.Convey(`value should be "127.0.0.1"`, func() {
					convey.So(value, convey.ShouldEqual, "127.0.0.1")
				})
				convey.Convey(`error should be nil`, func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
			convey.Convey(`GetInt("dev.port")`, func() {
				value, err := GetInt("dev.port")
				convey.Convey(`value should be 8080`, func() {
					convey.So(value, convey.ShouldEqual, 8080)
				})
				convey.Convey(`error should be nil`, func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
			convey.Convey(`GetString("dev.url")`, func() {
				value, err := GetString("dev.url")
				convey.Convey(`value should be "127.0.0.1:8080"`, func() {
					convey.So(value, convey.ShouldEqual, "127.0.0.1:8080")
				})
				convey.Convey(`error should be nil`, func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
			convey.Convey(`GetString("dev.foo")`, func() {
				value, err := GetString("dev.foo")
				convey.Convey(`value should be "devbar"`, func() {
					convey.So(value, convey.ShouldEqual, "devbar")
				})
				convey.Convey(`error should be nil`, func() {
					convey.So(err, convey.ShouldBeNil)
				})
			})
		})
	})
}
