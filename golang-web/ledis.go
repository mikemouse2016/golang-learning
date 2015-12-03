package main

import (
	//"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
	"log"
	"time"
)

type ResourceConn struct {
	redis.Conn
}

func (r ResourceConn) Close() {
	r.Conn.Close()
}

func main() {
	p := pools.NewResourcePool(func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", "211.152.55.72:6380")
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

	b, err := redis.Ints(c.Do("SMEMBERS", "to_match_user_ids"))

	log.Println(len(b))
	log.Println(b[len(b)-1])
	/*
		go func() {
			for i := 800000; i < 100; i += 1 {
				a := fmt.Sprintf("to_match_user_ids:%d", b[i])
				log.Println(a)
				//b, _ := redis.Ints(c.Do("GET", a))
				//log.Println(b)
			}
		}()
		select {}
	*/
}

/*
func demo1(step int, max int, ) {
	var a string
	for i := 0; i < max; i += step {
		fmt.Fprintf(a, "to_match_user_ids:%d", i)
		b, err := redis.Ints(c.Do("GET", a))
		log.Println(len(b))
	}
}
*/
