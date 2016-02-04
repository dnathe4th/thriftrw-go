// Copyright (c) 2015 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gen

import "github.com/uber/thriftrw-go/compile"

// typedef generates code for the given typedef.
func (g *Generator) typedef(spec *compile.TypedefSpec) error {
	err := g.DeclareFromTemplate(
		`
		<$wire := import "github.com/uber/thriftrw-go/wire">

		// TODO add * to ToWire if target is not a reference type
		type <.Name> <typeReference .Target Required>

		<$v := newVar "v">
		func (<$v> <.Name>) ToWire() <$wire>.Value {
			return <toWire .Target $v>
		}

		<$w := newVar "w">
		func (<$v> *<.Name>) FromWire(<$w> <$wire>.Value) error {
			// TODO: Implement some sort of fromValue template function that
			// will tell us what to put here.
			return nil
		}
		`,
		spec,
	)
	// TODO(abg): To/FromWire.
	return wrapGenerateError(spec.Name, err)
}