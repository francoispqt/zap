package benchmarks

import (
	"io/ioutil"
	"time"

	"github.com/francoispqt/onelog"
)

func (u *user) MarshalObject(enc *onelog.Encoder) {
	enc.AddStringKey("name", u.Name)
	enc.AddStringKey("email", u.Email)
	enc.AddInt64Key("createdAt", u.CreatedAt.UnixNano())
}

func (uu users) MarshalArray(enc *onelog.Encoder) {
	for i := range uu {
		enc.AddObject(uu[i])
	}
}
func (uu users) IsNil() bool {
	return len(uu) == 0
}

func (u *user) IsNil() bool {
	return u == nil
}

type timeArray []time.Time

func (tA timeArray) MarshalArray(enc *onelog.Encoder) {
	for _, t := range tA {
		enc.AddInt64(t.Unix())
	}
}

func (tA timeArray) IsNil() bool {
	return len(tA) == 0
}

type stringArray []string

func (tS stringArray) MarshalArray(enc *onelog.Encoder) {
	for _, t := range tS {
		enc.AddString(t)
	}
}

func (tS stringArray) IsNil() bool {
	return len(tS) == 0
}

type intArray []int

func (tI intArray) MarshalArray(enc *onelog.Encoder) {
	for _, t := range tI {
		enc.AddInt(t)
	}
}

func (tI intArray) IsNil() bool {
	return len(tI) == 0
}

func newOnelogLogger() *onelog.Logger {
	logger := onelog.New(ioutil.Discard, onelog.INFO)
	logger.Hook(func(e onelog.Entry) {
		e.Int64("time", time.Now().UnixNano())
	})
	return logger
}

func newDisabledOnelog() *onelog.Logger {
	return onelog.New(ioutil.Discard, onelog.DEBUG)
}

func getOnelogFields() func(onelog.Entry) {
	return func(e onelog.Entry) {
		e.Int("int", _tenInts[0])
		e.Array("ints", intArray(_tenInts))
		e.String("string", _tenStrings[0])
		e.Array("strings", stringArray(_tenStrings))
		e.Int64("time", _tenTimes[0].Unix())
		e.Array("times", timeArray(_tenTimes))
		e.Object("user1", _oneUser)
		e.Object("user2", _oneUser)
		e.Array("users", _tenUsers)
		e.String("err", errExample.Error())
	}
}
