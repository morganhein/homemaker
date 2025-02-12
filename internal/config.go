/*
 * Copyright (c) 2015 Alex Yatskov <alex@foosoft.net>
 * Author: Alex Yatskov <alex@foosoft.net>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package internal

import (
	"os"
)

// Config stores the run configurtion
type Config struct {
	Tasks    map[string]task
	Macros   map[string]macro
	File     string
	SrcDir   string `mapstructure:"home-src"`
	DstDir   string `mapstructure:"home-dst"`
	Variants string //comma separated, FIFO priority list of variants

	Clobber     bool
	Force       bool
	Verbose     bool
	Nocmds      bool
	Nolinks     bool
	Notemplates bool
	Unlink      bool

	handled map[string]bool
}

func (c *Config) digest() {
	c.handled = make(map[string]bool)
	c.SrcDir = makeAbsPath(c.SrcDir)
	c.DstDir = makeAbsPath(c.DstDir)

	if c.Unlink {
		c.Nocmds = true
	}
}

func (c *Config) setEnv() {
	os.Setenv("HM_CONFIG", c.File)
	os.Setenv("HM_SRC", c.SrcDir)
	os.Setenv("HM_DEST", c.DstDir)
	os.Setenv("HM_VARIANT", c.Variants)
}
