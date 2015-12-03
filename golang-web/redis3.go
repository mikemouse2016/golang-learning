package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
	"log"
	"strconv"
	"time"
)

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

func main() {
	p := pools.NewResourcePool(func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", ":6380")
		return ResourceConn{c}, err
	}, 1, 2, time.Minute)
	defer p.Close()
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Put(r)
	c := r.(ResourceConn)
	n, err := c.Do("INFO")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("info=%s", n)
	for i := 0; i < 3000000; i++ {
		c.Do("setex", "abcde"+strconv.Itoa(i), 600, "{wegewgreFrgregergergergerger2334 就 非常wefWEFfewffasaddefdefdefefoeirgvjrepjbvpreobjvperove[kverpv[erpkver[kvepv[rp"+strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println(time.Now())
	/*
	   log.Println(time.Now())
	   for i :=0;i<100000;i++{
	       c.Do("sadd", "demo3", i)
	   }
	*/

	/*
	   n, err = c.Do("get", "a")
	   for i :=0;i<10000;i++{
	       c.Do("hset", "demo4", i, i)
	   }
	   log.Println(time.Now())
	   for i :=0;i<10000;i++{
	       c.Do("hget", "demo4", i)
	   }
	   log.Println(time.Now())
	*/
}
