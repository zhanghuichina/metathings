package protobuf_helper

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func Now() timestamp.Timestamp {
	now := time.Now()
	return FromTime(now)
}

func FromTime(t time.Time) timestamp.Timestamp {
	return timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.UnixNano() % 1000000000),
	}
}

func ToTime(t timestamp.Timestamp) time.Time {
	return time.Unix(t.Seconds, int64(t.Nanos))
}
