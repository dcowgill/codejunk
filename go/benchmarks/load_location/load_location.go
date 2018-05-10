package load_location

import (
	"log"
	"sync"
	"time"
)

const tz = "America/New_York"

func LoadLocation() *time.Location {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Fatal(err)
	}
	return loc
}

var tzCache struct {
	sync.Mutex
	cache map[string]*time.Location
}

func LoadLocationCached() *time.Location {
	tzCache.Lock()
	defer tzCache.Unlock()
	if tzCache.cache == nil {
		tzCache.cache = make(map[string]*time.Location)
	}
	location, ok := tzCache.cache[tz]
	if !ok {
		var err error
		location, err = time.LoadLocation(tz)
		if err != nil {
			log.Fatal(err)
		}
		tzCache.cache[tz] = location
	}
	return location
}
