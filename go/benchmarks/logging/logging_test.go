package logging

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/omakasecorp/samurai/pkg/logmsg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	url     = "https://google.com/foo/bar/baz"
	attempt = 3
	backoff = 30 * time.Second

	zapLogger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(zapcore.AddSync(ioutil.Discard)),
		zap.NewAtomicLevel()))
)

func init() {
	log.SetOutput(ioutil.Discard)

}

func BenchmarkStdlibNoFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print("this is a test")
	}
}

func BenchmarkStdlib1Printf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Printf("this is a test url=%s", url)
	}
}

func BenchmarkStdlib3Printf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Printf("this is a test url=%s attempt=%d backoff=%s", url, attempt, backoff)
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(map[string]interface{}{
			"message": "this is a test",
			"url":     url,
			"attempt": attempt,
			"backoff": backoff,
		})
		log.Print(string(data))
	}
}

func BenchmarkLogmsg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print(logmsg.Error("this is a test").
			Set("url", url).
			Set("attempt", attempt).
			Set("backoff", backoff))
	}
}

func BenchmarkZapNoFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zapLogger.Error("this is a test")
	}
}

func BenchmarkZap1Field(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zapLogger.Error("this is a test", zap.String("url", url))
	}
}

func BenchmarkZap3Fields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		zapLogger.Error("this is a test",
			zap.String("url", url),
			zap.Int("attempt", attempt),
			zap.Duration("backoff", backoff))
	}
}
