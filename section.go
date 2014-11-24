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
	"fmt"

	"strconv"
	"strings"
)

// Section struct is considered as a unite of series of config.
type Section struct {
	Name   string
	Fields map[string]string
}

// GetString method returns string value.
func (s *Section) GetString(key string) (string, error) {
	if value, ok := s.Fields[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("Unknown key: %s", key)
}

// GetInt method returns integer value.
func (s *Section) GetInt(key string) (int, error) {
	if value, ok := s.Fields[key]; ok {
		return strconv.Atoi(value)
	}
	return 0, fmt.Errorf("Unknown key: %s", key)
}

// GetStrings method returns []string value.
// The value should start with [ and end with ].
func (s *Section) GetStrings(key string) ([]string, error) {
	if value, ok := s.Fields[key]; ok {
		if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
			value = value[1 : len(value)-1]
			values := strings.Split(value, ",")
			for index, v := range values {
				values[index] = strings.Trim(v, " \t")
			}
			return values, nil
		}
		return nil, fmt.Errorf("Wrong type: %s", value)
	}
	return nil, fmt.Errorf("Unknown key: %s", key)
}

// GetStrings method returns []int value.
// The value should start with [ and end with ].
func (s *Section) GetInts(key string) ([]int, error) {
	if value, ok := s.Fields[key]; ok {
		if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
			value = value[1 : len(value)-1]
			values := strings.Split(value, ",")
			ints := make([]int, len(values))
			for index, v := range values {
				v = strings.Trim(v, " \t")
				i, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				ints[index] = i
			}
			return ints, nil
		}
		return nil, fmt.Errorf("Wrong type: %s", value)
	}
	return nil, fmt.Errorf("Unknown key: %s", key)
}
