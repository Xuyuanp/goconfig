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
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Default = New()

// Configuration struct
type Configuration struct {
	Fields   map[string]string
	Sections map[string]*Section
}

// New methods return an empty Configuration.
func New() *Configuration {
	return &Configuration{
		Fields:   map[string]string{},
		Sections: map[string]*Section{},
	}
}

// GetString method returns string value.
// The key may be a simple string or <section_name>.<key>
func (c *Configuration) GetString(key string) (string, error) {
	keys := strings.SplitN(key, ".", 2)
	if len(keys) == 2 {
		if section, ok := c.Sections[keys[0]]; ok {
			return section.GetString(keys[1])
		}
	}
	if value, ok := c.Fields[keys[0]]; ok {
		return value, nil
	}
	return "", fmt.Errorf("Invalid key: %s", key)
}

// GetInt method returns integer value.
// The key may be a simple string or <section_name>.<key>
func (c *Configuration) GetInt(key string) (int, error) {
	keys := strings.SplitN(key, ".", 2)
	if len(keys) == 2 {
		if section, ok := c.Sections[keys[0]]; ok {
			return section.GetInt(keys[1])
		}
	}
	if value, ok := c.Fields[keys[0]]; ok {
		return strconv.Atoi(value)
	}
	return 0, fmt.Errorf("Invalid key: %s", key)
}

// LoadFile method loads config from a file.
func (c *Configuration) LoadFile(fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()
	return c.Load(file)
}

var configHolderRegexp = regexp.MustCompile(`\$\([\w_][\w\d_]*[\.[\w_][\w\d_]*]?\)`)
var envHolderRegexp = regexp.MustCompile(`\$\{[\w_][\w\d_]*[\.[\w_][\w\d_]*]?\}`)
var keyRegexp = regexp.MustCompile(`^[\w_][\w\d_]*$`)
var sectionRegexp = regexp.MustCompile(`^\[[\w_][\w\d_]*\]$`)

// Load method loads config from an io.Reader. An error will be returned
// if loading failed.
func (c *Configuration) Load(rd io.Reader) (err error) {
	reader := bufio.NewReader(rd)
	var currentSection *Section
	linno := 0
	var line string
	for {
		line, err = readLine(reader)
		if err != nil {
			break
		}
		line = strings.Trim(line, " \t")
		linno++
		// Ignore empty line.
		if line == "" {
			continue
		}
		// Ignore comment.
		if strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.SplitN(line, "=", 2)
		// Parse new section.
		if len(fields) == 1 {
			field := strings.Trim(fields[0], " \t")
			if sectionRegexp.MatchString(field) {
				sectionName := field[1 : len(field)-1]
				section := &Section{
					Name:   sectionName,
					Fields: make(map[string]string),
				}
				c.Sections[sectionName] = section
				currentSection = section
			} else {
				return fmt.Errorf("Invalid line: %d %s", linno, line)
			}
		} else {
			key := strings.Trim(fields[0], " \t")
			value := strings.Trim(fields[1], " \t")
			if !keyRegexp.MatchString(key) {
				return fmt.Errorf("Invalid key: %d %s", linno, line)
			}
			// Fill system environment var placeholder.
			value = envHolderRegexp.ReplaceAllStringFunc(value, func(s string) string {
				k := s[2 : len(s)-1]
				ss := os.Getenv(k)
				return ss
			})
			// Fill local config var placeholder.
			value = configHolderRegexp.ReplaceAllStringFunc(value, func(s string) string {
				k := s[2 : len(s)-1]
				ss, e := c.GetString(k)
				if e != nil {
					err = fmt.Errorf("Invalid key: %d %s", linno, k)
				}
				return ss
			})
			if err != nil {
				return
			}
			if currentSection == nil {
				c.Fields[key] = value
			} else {
				currentSection.Fields[key] = value
			}
		}
	}
	return nil
}

func readLine(rd *bufio.Reader) (line string, err error) {
	var l []byte
	var isPrefix bool
	for l, isPrefix, err = rd.ReadLine(); err == nil; {
		line += string(l)
		if !isPrefix {
			return
		}
	}
	return
}

func LoadFile(fname string) error {
	return Default.LoadFile(fname)
}

func Load(rd io.Reader) error {
	return Default.Load(rd)
}

func GetString(key string) (string, error) {
	return Default.GetString(key)
}

func GetInt(key string) (int, error) {
	return Default.GetInt(key)
}
