// Copyright (c) 2016 Uber Technologies, Inc.
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

package benchmarks

import (
	"io/ioutil"

	"github.com/rs/zerolog"
)

func (uu users) MarshalZerologArray(a *zerolog.Array) {
	for _, u := range uu {
		a.Object(u)
	}
}

func (u user) MarshalZerologObject(e *zerolog.Event) {
	e.Str("name", u.Name).
		Str("email", u.Email).
		Int64("createdAt", u.CreatedAt.UnixNano())
}

func (tA timeArray) MarshalZerologArray(a *zerolog.Array) {
	for _, t := range tA {
		a.Int64(t.Unix())
	}
}

func (tS stringArray) MarshalZerologArray(a *zerolog.Array) {
	for _, t := range tS {
		a.Str(t)
	}
}

func (tI intArray) MarshalZerologArray(a *zerolog.Array) {
	for _, t := range tI {
		a.Int(t)
	}
}

func newZerolog() zerolog.Logger {
	return zerolog.New(ioutil.Discard).With().Timestamp().Logger()
}

func newDisabledZerolog() zerolog.Logger {
	return newZerolog().Level(zerolog.Disabled)
}

func fakeZerologFields(e *zerolog.Event) *zerolog.Event {
	return e.
		Int("int", _tenInts[0]).
		Array("ints", intArray(_tenInts)).
		Str("string", _tenStrings[0]).
		Array("strings", stringArray(_tenStrings)).
		Time("time", _tenTimes[0]).
		Array("times", timeArray(_tenTimes)).
		Object("user1", _oneUser).
		Object("user2", _oneUser).
		Array("users", _tenUsers).
		Err(errExample)
}

func fakeZerologContext(c zerolog.Context) zerolog.Context {
	return c.
		Int("int", _tenInts[0]).
		Array("ints", intArray(_tenInts)).
		Str("string", _tenStrings[0]).
		Array("strings", stringArray(_tenStrings)).
		Time("time", _tenTimes[0]).
		Array("times", timeArray(_tenTimes)).
		Object("user1", _oneUser).
		Object("user2", _oneUser).
		Array("users", _tenUsers).
		Err(errExample)
}
