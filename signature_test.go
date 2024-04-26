package polysdk

import (
	"testing"
	"time"

	"github.com/quanxiang-cloud/go-polysdk/internal/polysign"
)

func TestSignature(t *testing.T) {
	println(time.Now().Format(polysign.PingTimestampFmt))
}
