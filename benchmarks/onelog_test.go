package benchmarks

import (
	"io/ioutil"
	"time"

	"github.com/francoispqt/onelog"
	"go.uber.org/zap/zapcore"
)

func (u *user) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", u.Name)
	enc.AddString("email", u.Email)
	enc.AddInt64("createdAt", u.CreatedAt.UnixNano())
	return nil
}

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
	logger.Hook(func(enc onelog.Entry) {
		enc.Int64("time", time.Now().UnixNano())
	})
	return logger
}

func newDisabledOnelog() *onelog.Logger {
	return onelog.New(ioutil.Discard, onelog.DEBUG)
}

func getOnelogFields() func(onelog.Entry) {
	return func(enc onelog.Entry) {
		enc.Int("int", _tenInts[0])
		enc.Array("ints", intArray(_tenInts))
		enc.String("string", _tenStrings[0])
		enc.Array("strings", stringArray(_tenStrings))
		enc.Int64("time", _tenTimes[0].Unix())
		enc.Array("times", timeArray(_tenTimes))
		enc.Object("user1", _oneUser)
		enc.Object("user2", _oneUser)
		enc.Array("users", _tenUsers)
		enc.Err("err", errExample)
	}
}
