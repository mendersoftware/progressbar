// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package progressbar

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"strings"
)

func TestProgressTTY(t *testing.T) {
	b := &bytes.Buffer{}
	p := &Bar{
		Renderer: &TTYRenderer{
			Out:            b,
			ProgressMarker: ".",
			terminalWidth:  80,
		},
		Size: 100,
	}
	p.Tick(10)
	assert.Equal(t,
		"\r.......                                                                  -  10 %",
		b.String())

	b.Reset()
	p.Tick(10)
	assert.Equal(t,
		"\r..............                                                           -  20 %",
		b.String())
	b.Reset()
	p.Tick(10)
	assert.Equal(t,
		"\r.....................                                                    -  30 %",
		b.String())
	b.Reset()
	p.Tick(50)
	assert.Equal(t,
		"\r.........................................................                -  80 %",
		b.String())
	b.Reset()
	p.Tick(20)
	assert.Equal(t,
		"\r........................................................................ - 100 %\n",
		b.String())
	b.Reset()
	p.Tick(10)
	assert.Equal(t,
		"\r........................................................................ - 110 %",
		b.String())
}

func TestProgressNoTTY(t *testing.T) {
	b := &bytes.Buffer{}
	p := &Bar{
		Renderer: &NoTTYRenderer{
			Out:            b,
			ProgressMarker: ".",
			terminalWidth:  100,
		},
		Size: 100,
	}
	p.Tick(10)
	header := "0%" + strings.Repeat(" ", 45) + "50%" + strings.Repeat(" ", 45) + "100%\n" +
		"|" + strings.Repeat("-", 48) + "|" + strings.Repeat("-", 48) + "|\n"
	assert.Equal(t,
		header+strings.Repeat(".", 10),
		b.String())
	p.Tick(10)
	assert.Equal(t,
		header+strings.Repeat(".", 20),
		b.String())
	p.Tick(10)
	assert.Equal(t,
		header+strings.Repeat(".", 30),
		b.String())
	p.Tick(50)
	assert.Equal(t,
		header+strings.Repeat(".", 80),
		b.String())
	p.Tick(20)
	assert.Equal(t,
		header+strings.Repeat(".", 100),
		b.String())
	p.Tick(10)
	assert.Equal(t,
		header+strings.Repeat(".", 100),
		b.String())
}
