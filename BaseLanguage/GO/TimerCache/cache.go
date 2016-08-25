package timecache

import (
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	cacheGroup = make(map[string]*CacheTable)
)

func Cache(table string) *CacheTable {
	t, ok := cacheGroup[table]
	if !ok {
		t = &CacheTable{
			name:          table,
			pool:          make(map[interface{}]*CacheItem),
			cleanInterval: 0,
		}
		cacheGroup[table] = t
	}

	for i := 0; i < 20000; i++ {
		key := "key_" + strconv.Itoa(i)
		val := "val_" + strconv.Itoa(i)
		t.Add(key, val, 3*time.Second)
	}

	return t
}

//Define cacheItem
type CacheItem struct {
	value    interface{}
	accessOn time.Time
	lifeSpan time.Duration
}

func CreateCacheItem(val interface{}, duration time.Duration) CacheItem {
	nowTime := time.Now()
	return CacheItem{
		value:    val,
		lifeSpan: duration,
		accessOn: nowTime,
	}
}

//Define cacheTable
type CacheTable struct {
	name string
	pool map[interface{}]*CacheItem

	cleanTime     *time.Timer
	cleanInterval time.Duration
	sync.RWMutex
}

func (c *CacheTable) Add(key interface{}, val interface{}, lifeSpan time.Duration) {
	item := CreateCacheItem(val, lifeSpan)

	c.Lock()
	c.pool[key] = &item
	//log.Println("[AddTime]" + key.(string) + " " + time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"))
	c.Unlock()
	if 0 == c.cleanInterval || lifeSpan < c.cleanInterval {
		c.Clean()
	}
}

//Clean the Exp Object
func (c *CacheTable) Clean() {
	if nil != c.cleanTime {
		c.cleanTime.Stop()
	}

	//Copy map
	items := c.pool

	//Get min duration
	minDuration := 0 * time.Second
	clean_st := time.Now()
	for key, item := range items {
		nowTime := time.Now()
		if item.lifeSpan <= nowTime.Sub(item.accessOn) {
			c.Delete(key)
		} else {
			if 0 == minDuration || item.lifeSpan-nowTime.Sub(item.accessOn) < minDuration {
				minDuration = item.lifeSpan - nowTime.Sub(item.accessOn)
				//log.Println("innnn %f", minDuration.Seconds())
			}
		}
	}
	if minDuration <= 0 {
		return
	}

	log.Println("LoopCost %f", time.Now().Sub(clean_st).Seconds())
	c.cleanInterval = minDuration
	//Check if need do clean
	c.cleanTime = time.AfterFunc(minDuration, func() {
		go c.Clean()
	})
}

func (c *CacheTable) Delete(key interface{}) {
	//log.Println("[DeleteTime] %s %f", key.(string), time.Now().Sub(c.pool[key].accessOn).Seconds()*1000)
	delete(c.pool, key)
}
