package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type bigDocument struct {
	ID bson.ObjectId `bson:"_id"`

	PehxaIpztxbaZlvraNmrna      int
	GexzuPanriDtqw              int
	JrtmQemeVcqjMzugx           int
	JirdljcjAngqbo              int
	IkualjQmbjlbBuaeprr         int
	MdbbehaZavj                 int
	UvoiErqvkgt                 int
	ElndadhxXszpcnQqwugujfArpfr int
	SdgpmkMwlq                  int
	QxsloVwazot                 int

	MpcksqsBswtfwp               string
	DuqcVeuoAhzl                 string
	YylleysvSemfpqMhbqlavz       string
	YcciwKrolwpaMvojviunYwhmkeun string
	FpqhknyqTkxyd                string
	KryuJxkhafGveafBjjudhu       string
	JujbpdIyejdrpAonfnhim        string
	JfwobfxmEtwxaqAhuxoIrii      string
	EybcUurxmodSxswEhwalb        string
	GyuvJljiwibRwqmbhqQgtrp      string

	OudtziwDtmgvv               time.Time
	EclqrpBrtlEffwxlEiyi        time.Time
	YftchRpywichs               time.Time
	OzxysoskJihpdq              time.Time
	HizzPinruKtyksb             time.Time
	KxmwebZdhlrXneyiom          time.Time
	UlgkNnlwhith                time.Time
	VpawjkScayaqHqngvqbUpxmaehs time.Time
	PycmzitHceu                 time.Time
	TqxgLstlknJxneoiyj          time.Time

	DnhnnNdzpLilnevTouku       string
	LnwgcepqGdvshp             string
	SrhbwjidDplrql             string
	ZerbyFwcycf                string
	ReaoveKflfkEhrhpo          string
	JeyzRqkim                  string
	YttoggWghnRcdabwWxtytq     string
	WabzMtqvwwwrDhevzoPlyqeiso string
	ErucxyYzhtneTeyms          string
	OzfgGovtopi                string
	YjuemkojLfjnkCnpfukawUkib  string
	MbfzPvanbo                 string
	XclcmmItjysnzv             string
	MclzrohCinci               string
	OkrdXnkhDtchjmGmcxxw       string
	CcgtbdqAgbpvxbQlsz         string
	UexuqkTfhrWqjpaCsnruht     string
	HihaSjfbmoMlxkBpiapya      string
	AoeblfoEwggichkTfugd       string
	SywbeOopdjDordie           string
}

func getDoc(col *mgo.Collection, id bson.ObjectId) *bigDocument {
	var result *bigDocument
	must(col.FindId(id).One(&result))
	return result
}

func incOp(col *mgo.Collection, doc *bigDocument) {
	must(col.Update(
		bson.M{"_id": doc.ID, "sdgpmkmwlq": doc.SdgpmkMwlq},
		bson.M{
			"$inc": bson.M{
				"pehxaipztxbazlvranmrna": 1,
			},
			"$set": bson.M{
				"ycciwkrolwpamvojviunywhmkeun": randomASCII(randomRange(10, 51)),
				"vpawjkscayaqhqngvqbupxmaehs":  time.Now(),
			},
		},
	))
}

func incFull(col *mgo.Collection, doc *bigDocument) {
	doc.PehxaIpztxbaZlvraNmrna++
	doc.YcciwKrolwpaMvojviunYwhmkeun = randomASCII(randomRange(10, 51))
	doc.VpawjkScayaqHqngvqbUpxmaehs = time.Now()
	must(col.Update(bson.M{"_id": doc.ID, "sdgpmkmwlq": doc.SdgpmkMwlq}, doc))

}

func benchmark(
	col *mgo.Collection,
	n int,
	id bson.ObjectId,
	name string,
	fn func(col *mgo.Collection, doc *bigDocument),
) {
	t := time.Now()
	for i := 0; i < n; i++ {
		doc := getDoc(col, id)
		fn(col, doc)
	}
	elapsed := time.Since(t)
	log.Printf("%30s %.1fs [%.2fms]", name, elapsed.Seconds(), 1000*(elapsed/time.Duration(n)).Seconds())
}

func main() {
	var (
		mongoURL     = flag.String("mongoURL", "mongodb://localhost:27017", "mongoDB URL")
		numDocuments = flag.Int("ndocs", 1, "number of documents")
		numUpdates   = flag.Int("nupdates", 10, "number of updates")
		numRounds    = flag.Int("nrounds", 1, "number of rounds")
	)
	flag.Parse()

	session, err := mgo.DialWithTimeout(*mongoURL, time.Second)
	must(err)
	col := session.DB("benchmark").C("update_test")

	var id bson.ObjectId
	log.Printf("inserting %d random document(s)", *numDocuments)
	for i := 0; i < *numDocuments; i++ {
		d := randomDocument()
		must(col.Insert(d))
		if rand.Int()%(i+1) == 0 {
			id = d.ID
		}
	}

	for i := 0; i < *numRounds; i++ {
		benchmark(col, *numUpdates, id, "$inc", incOp)
		benchmark(col, *numUpdates, id, "full", incFull)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func randomDocument() *bigDocument {
	return &bigDocument{
		ID: bson.NewObjectId(),
		PehxaIpztxbaZlvraNmrna:       randomRange(0, 32767),
		GexzuPanriDtqw:               randomRange(0, 32767),
		JrtmQemeVcqjMzugx:            randomRange(0, 32767),
		JirdljcjAngqbo:               randomRange(0, 32767),
		IkualjQmbjlbBuaeprr:          randomRange(0, 32767),
		MdbbehaZavj:                  randomRange(0, 32767),
		UvoiErqvkgt:                  randomRange(0, 32767),
		ElndadhxXszpcnQqwugujfArpfr:  randomRange(0, 32767),
		SdgpmkMwlq:                   randomRange(0, 32767),
		QxsloVwazot:                  randomRange(0, 32767),
		MpcksqsBswtfwp:               randomASCII(randomRange(10, 51)),
		DuqcVeuoAhzl:                 randomASCII(randomRange(10, 51)),
		YylleysvSemfpqMhbqlavz:       randomASCII(randomRange(10, 51)),
		YcciwKrolwpaMvojviunYwhmkeun: randomASCII(randomRange(10, 51)),
		FpqhknyqTkxyd:                randomASCII(randomRange(10, 51)),
		KryuJxkhafGveafBjjudhu:       randomASCII(randomRange(10, 51)),
		JujbpdIyejdrpAonfnhim:        randomASCII(randomRange(10, 51)),
		JfwobfxmEtwxaqAhuxoIrii:      randomASCII(randomRange(10, 51)),
		EybcUurxmodSxswEhwalb:        randomASCII(randomRange(10, 51)),
		GyuvJljiwibRwqmbhqQgtrp:      randomASCII(randomRange(10, 51)),
		OudtziwDtmgvv:                randomTime(),
		EclqrpBrtlEffwxlEiyi:         randomTime(),
		YftchRpywichs:                randomTime(),
		OzxysoskJihpdq:               randomTime(),
		HizzPinruKtyksb:              randomTime(),
		KxmwebZdhlrXneyiom:           randomTime(),
		UlgkNnlwhith:                 randomTime(),
		VpawjkScayaqHqngvqbUpxmaehs:  randomTime(),
		PycmzitHceu:                  randomTime(),
		TqxgLstlknJxneoiyj:           randomTime(),
		DnhnnNdzpLilnevTouku:         randomASCII(randomRange(10, 51)),
		LnwgcepqGdvshp:               randomASCII(randomRange(10, 51)),
		SrhbwjidDplrql:               randomASCII(randomRange(10, 51)),
		ZerbyFwcycf:                  randomASCII(randomRange(10, 51)),
		ReaoveKflfkEhrhpo:            randomASCII(randomRange(10, 51)),
		JeyzRqkim:                    randomASCII(randomRange(10, 51)),
		YttoggWghnRcdabwWxtytq:       randomASCII(randomRange(10, 51)),
		WabzMtqvwwwrDhevzoPlyqeiso:   randomASCII(randomRange(10, 51)),
		ErucxyYzhtneTeyms:            randomASCII(randomRange(10, 51)),
		OzfgGovtopi:                  randomASCII(randomRange(10, 51)),
		YjuemkojLfjnkCnpfukawUkib:    randomASCII(randomRange(10, 51)),
		MbfzPvanbo:                   randomASCII(randomRange(10, 51)),
		XclcmmItjysnzv:               randomASCII(randomRange(10, 51)),
		MclzrohCinci:                 randomASCII(randomRange(10, 51)),
		OkrdXnkhDtchjmGmcxxw:         randomASCII(randomRange(10, 51)),
		CcgtbdqAgbpvxbQlsz:           randomASCII(randomRange(10, 51)),
		UexuqkTfhrWqjpaCsnruht:       randomASCII(randomRange(10, 51)),
		HihaSjfbmoMlxkBpiapya:        randomASCII(randomRange(10, 51)),
		AoeblfoEwggichkTfugd:         randomASCII(randomRange(10, 51)),
		SywbeOopdjDordie:             randomASCII(randomRange(10, 51)),
	}
}

const (
	lowercase   = "abcdefghijklmnopqrstuvwxyz"
	uppercase   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters     = lowercase + uppercase
	digits      = "0123456789"
	alnum       = letters + digits
	punctuation = `!"#$%&\'()*+,-./:;<=>?@[\\]^_{|}~`
	whitespace  = "\t\n\x0b\x0c\r "
	printable   = alnum + punctuation + whitespace
)

func randomTime() time.Time {
	return time.Now().Add(time.Duration(randomRange(-86400, 86400)) * time.Second)
}

func randomASCII(n int) string   { return randomStringFromBytes(printable, n) }
func randomLetters(n int) string { return randomStringFromBytes(letters, n) }
func randomDigits(n int) string  { return randomStringFromBytes(digits, n) }
func randomAlnum(n int) string   { return randomStringFromBytes(alnum, n) }

func randomStringFromBytes(src string, n int) string {
	a := make([]byte, n)
	for i := range a {
		a[i] = randomByte(src)
	}
	return string(a)
}

// Returns a randomly chosen byte in s.
func randomByte(s string) byte {
	return s[rand.Intn(len(s))]
}

// Returns a random int in the half-open range [l, h).
func randomRange(l, h int) int {
	return rand.Intn(h-l) + l
}
