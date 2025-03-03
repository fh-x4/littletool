package demo

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/fh-x4/littletool/component/logger"
)

func GenDemoCallback(data string) (*demoTimer, error) {
	r := strconv.Itoa(rand.Intn(2 << 18))
	key := strings.ReplaceAll(uuid.New().String(), "-", "")
	key += strings.Repeat("0", 7-len(r)) + r
	key += strconv.FormatInt(time.Now().UnixMilli(), 10)

	return &demoTimer{
		key:  key,
		Data: data,
	}, nil
}

type demoTimer struct {
	key  string
	Data string
}

func (h *demoTimer) Key() string {
	return h.key
}
func (h *demoTimer) Call() error {
	logger.GetLogger().Infof("demo timer called, key=%s, data=%s, ts=%d", h.key, string(h.Data), time.Now().Unix())
	return nil
}
