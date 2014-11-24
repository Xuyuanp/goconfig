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
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSection(t *testing.T) {
	var section *Section
	convey.Convey(`Given a new Section{Name: "Test", Fields: {"ip": "192.168.1.1", "port": 8080}}`, t, func() {
		section = &Section{
			Name: "Test",
			Fields: map[string]string{
				"ip":      "192.168.1.1",
				"port":    "8080",
				"id_list": "[1, 2, 3, 4]",
			},
		}
		convey.Convey(`GetString("ip")`, func() {
			value, err := section.GetString("ip")
			convey.Convey(`value should be "192.168.1.1"`, func() {
				convey.So(value, convey.ShouldEqual, "192.168.1.1")
			})
			convey.Convey(`and error should be nil`, func() {
				convey.So(err, convey.ShouldBeNil)
			})
		})
		convey.Convey(`GetInt("port")`, func() {
			value, err := section.GetInt("port")
			convey.Convey(`value should be 8080`, func() {
				convey.So(value, convey.ShouldEqual, 8080)
			})
			convey.Convey(`and error should be nil`, func() {
				convey.So(err, convey.ShouldBeNil)
			})
		})
		convey.Convey(`GetStrings("id_list")`, func() {
			value, err := section.GetStrings("id_list")
			convey.Convey(`value should be ["1", "2", "3", "4"]`, func() {
				convey.So(value[0], convey.ShouldEqual, "1")
				convey.So(value[1], convey.ShouldEqual, "2")
				convey.So(value[2], convey.ShouldEqual, "3")
				convey.So(value[3], convey.ShouldEqual, "4")
			})
			convey.Convey(`and error should be nil`, func() {
				convey.So(err, convey.ShouldBeNil)
			})
		})
		convey.Convey(`GetInts("id_list")`, func() {
			value, err := section.GetInts("id_list")
			convey.Convey(`value should be [1, 2, 3, 4]`, func() {
				convey.So(value[0], convey.ShouldEqual, 1)
				convey.So(value[1], convey.ShouldEqual, 2)
				convey.So(value[2], convey.ShouldEqual, 3)
				convey.So(value[3], convey.ShouldEqual, 4)
			})
			convey.Convey(`and error should be nil`, func() {
				convey.So(err, convey.ShouldBeNil)
			})
		})
	})
}
