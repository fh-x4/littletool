package implement

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type base struct {
	key     string
	types   string
	startTs int64
}

func (b *base) GenerateBaseInfo(types string) {
	r := strconv.Itoa(rand.Intn(2 << 18))
	key := strings.ReplaceAll(uuid.New().String(), "-", "")
	key += strings.Repeat("0", 7-len(r)) + r
	key += strconv.FormatInt(time.Now().UnixMilli(), 10)
	b.key = key
	b.startTs = time.Now().UnixMilli()
	b.types = types
}
