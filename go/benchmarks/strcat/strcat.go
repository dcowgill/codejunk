package strcat

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/omakasecorp/samurai/pkg/crypt"
)

var left = []string{
	"appetizing",
	"baked",
	"bland",
	"blended",
	"boiled",
	"brined",
	"buttered",
	"char-broiled",
	"cheesy",
	"chilled",
	"chunked",
	"creamed",
	"creamy",
	"crispy",
	"crunchy",
	"cured",
	"deep-fried",
	"dressed",
	"drizzled",
}

var right = []string{
	"sauteed",
	"savory",
	"seared",
	"seasoned",
	"simmered",
	"sizzling",
	"smoked",
	"smothered",
	"spiced",
	"spicy",
	"steamed",
	"sticky",
	"stuffed",
	"sugared",
	"sweetened",
	"tangy",
	"tender",
	"tepid",
	"toasted",
	"tossed",
	"whipped",
	"zesty",
}

func CRandSprintf() string {
	lLen := int64(len(left))
	rLen := int64(len(right))
	return fmt.Sprintf("%s_%s", left[crypt.RandInt(lLen)], right[crypt.RandInt(rLen)])
}

func CRandConcat() string {
	lLen := int64(len(left))
	rLen := int64(len(right))
	return left[crypt.RandInt(lLen)] + "_" + right[crypt.RandInt(rLen)]
}

func RandSprintf() string {
	return fmt.Sprintf("%s_%s", left[rand.Intn(len(left))], right[rand.Intn(len(right))])
}

func RandConcat() string {
	return left[rand.Intn(len(left))] + "_" + right[rand.Intn(len(right))]
}

func RandBuffer() string {
	l := left[rand.Intn(len(left))]
	r := right[rand.Intn(len(right))]
	b := bytes.NewBuffer(make([]byte, 0, len(l)+len(r)+1))
	b.WriteString(l)
	b.WriteByte('_')
	b.WriteString(r)
	return b.String()
}
