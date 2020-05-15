package main

import (
	"fmt"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"sync"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	curtime := time.Now().UnixNano()

	//单个测试
	/*pool := newPool()
	r := redsync.New([]redsync.Pool{pool})
	//配置redis锁
	mutex := r.NewMutex("test-mutex", redsync.SetExpiry(time.Duration(2)*time.Second),
		redsync.SetRetryDelay(time.Duration(10)*time.Millisecond))

	//获取锁
	if err := mutex.Lock(); err != nil {
		t.Fatalf("Expected err == nil, got %q", err)
		return
	}
	fmt.Println("add lock ....")
	str, _ := redis.String(DoRedisCmdByConn(pool, "GET", "name"))
	fmt.Println("before name is : ", str)
	//进行写操作
	DoRedisCmdByConn(pool, "SET", "name", "ywanbing")
	str, _ = redis.String(DoRedisCmdByConn(pool, "GET", "name"))
	fmt.Println("after name is : ", str)
	//释放锁
	mutex.Unlock()
	fmt.Println("del lock ....")*/

	//并发访问
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//这里模拟多个客户端获取redis的值
		go func(i int, wg *sync.WaitGroup) {
			//多个同时访问
			pool := newPool()
			r := redsync.New([]redsync.Pool{pool})
			//配置redis锁
			mutex := r.NewMutex("test-mutex", redsync.SetExpiry(time.Duration(2)*time.Second),
				redsync.SetRetryDelay(time.Duration(10)*time.Millisecond))

			//获取锁
			if err := mutex.Lock(); err != nil {
				t.Fatalf("Expected err == nil, got %q", err)
				return
			}

			//释放锁
			defer mutex.Unlock()
			fmt.Println(i, "add lock ....")

			str, _ := redis.String(DoRedisCmdByConn(pool, "GET", "name"))
			fmt.Println("before name is : ", str)
			//进行写操作
			_, _ = DoRedisCmdByConn(pool, "SET", "name", fmt.Sprintf("name%v", i))
			str, _ = redis.String(DoRedisCmdByConn(pool, "GET", "name"))
			fmt.Println("after name is : ", str)
			fmt.Println(i, "del lock ....")
			wg.Done()

		}(i, wg)
	}

	wg.Wait()
	fmt.Println(time.Now().UnixNano() - curtime)
}

//redis命令执行函数
func DoRedisCmdByConn(conn *redis.Pool, commandName string, args ...interface{}) (interface{}, error) {
	redisConn := conn.Get()
	defer redisConn.Close()
	//检查与redis的连接
	return redisConn.Do(commandName, args...)
}

//获取redis连接池
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: time.Duration(24) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.7.51:6379")
			if err != nil {
				panic(err.Error())
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return err
		},
	}
}
