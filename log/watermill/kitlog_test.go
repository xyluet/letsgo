package watermill_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ThreeDotsLabs/watermill"
	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/assert"

	letsgoassert "github.com/xyluet/letsgo/assert"
	letsgowatermill "github.com/xyluet/letsgo/log/watermill"
)

func TestKitLogger_All(t *testing.T) {
	var buf bytes.Buffer
	kitLogger := kitlog.NewJSONLogger(&buf)

	logger := letsgowatermill.NewKitLogger(kitLogger, letsgowatermill.WithTraceEnabled(true))
	logger.Trace("msg", watermill.LogFields{"key1": "value1"})
	logger.Debug("msg", watermill.LogFields{"key1": "value1"})
	logger.Info("msg", watermill.LogFields{"key1": "value1"})
	logger.Error("msg", errors.New("!"), watermill.LogFields{"key1": "value1"})
	logger.With(watermill.LogFields{"key2": "value2"}).Debug("msg", watermill.LogFields{"key1": "value1"})

	letsgoassert.JSONLEq(
		t,
		strings.Join([]string{
			`{"level":"debug","msg":"msg","key1":"value1"}`,
			`{"level":"debug","msg":"msg","key1":"value1"}`,
			`{"level":"info","msg":"msg","key1":"value1"}`,
			`{"level":"error","msg":"msg","err":"!","key1":"value1"}`,
			`{"level":"debug","msg":"msg","key1":"value1","key2":"value2"}`,
		}, "\n"),
		buf.String(),
	)
}

func TestKitLogger_With(t *testing.T) {
	var buf bytes.Buffer

	logger := letsgowatermill.NewKitLogger(kitlog.NewJSONLogger(&buf), letsgowatermill.WithTraceEnabled(true))
	logger.With(watermill.LogFields{"key1": "value1"}).Trace("msg", watermill.LogFields{"key2": "value2"})
	logger.With(watermill.LogFields{"key1": "value1"}).Trace("msg", watermill.LogFields{"key2": "value2"})

	letsgoassert.JSONLEq(
		t,
		strings.Join([]string{
			`{"level":"debug","msg":"msg","key1":"value1","key2":"value2"}`,
			`{"level":"debug","msg":"msg","key1":"value1","key2":"value2"}`,
		}, "\n"),
		buf.String(),
	)
}

func TestKitLogger_SkipTrace(t *testing.T) {
	var buf bytes.Buffer

	logger := letsgowatermill.NewKitLogger(kitlog.NewJSONLogger(&buf))
	logger.With(watermill.LogFields{"key1": "value1"}).Trace("msg", watermill.LogFields{"key2": "value2"})

	assert.Equal(t, "", buf.String())
}
