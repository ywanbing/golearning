package test

import (
	"sync"
	"sync/atomic"
	"testing"
)

/*
请教个问题 如果50个协程去访问全局变量并且进行 +1的操作
最终这个全局变量会不会是50呢?

目前我了解的三种方式：
1 原子操作
2 chan 通道
3 加锁
4 是使用单线程执行，目的是为了性能对比

go test -test.bench="." -test.benchmem
测试结果：
goos: windows
goarch: amd64

BenchmarkAtomic-12                 42664             28082 ns/op              24 B/op          2 allocs/op
BenchmarkRWMutex-12                42214             28452 ns/op              56 B/op          3 allocs/op
BenchmarkChannel-12                24367             48412 ns/op             121 B/op          3 allocs/op
BenchmarkSingleThread-12         1501362               801 ns/op              16 B/op          1 allocs/op

*/

func TestAtomic(t *testing.T) {
	atomicAchieve()
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomicAchieve()
	}
}

func atomicAchieve() {
	i := int64(0)
	wg := new(sync.WaitGroup)
	for n := 0; n < 50; n++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			atomic.AddInt64(&i, 1)
			wg.Done()
		}(wg)
	}
	wg.Wait()
	//println(i)
}

func TestRWMutex(t *testing.T) {
	RWMutexAchieve()
}

func BenchmarkRWMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RWMutexAchieve()
	}
}

func RWMutexAchieve() {
	i := int64(0)
	wg := new(sync.WaitGroup)
	rw := new(sync.RWMutex)
	for n := 0; n < 50; n++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			rw.RLock()
			defer rw.RUnlock()
			i++
			wg.Done()
		}(wg)
	}
	wg.Wait()
	//println(i)
}

func TestChannel(t *testing.T) {
	ChannelAchieve()
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChannelAchieve()
	}
}

func ChannelAchieve() {
	i := int64(0)
	wg := new(sync.WaitGroup)
	addNum := make(chan struct{}, 1)
	for n := 0; n < 50; n++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			addNum <- struct{}{}
			i++
			<-addNum
			wg.Done()
		}(wg)
	}
	wg.Wait()
	//println(i)
}

func TestSingleThread(t *testing.T) {
	SingleThreadAchieve()
}

func BenchmarkSingleThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleThreadAchieve()
	}
}

//测试一下单线程计算,
//加上waitGroup是为了满足上面的压测
func SingleThreadAchieve() {
	i := int64(0)
	wg := new(sync.WaitGroup)
	for n := 0; n < 50; n++ {
		wg.Add(1)
		i++
		wg.Done()
	}
	wg.Wait()
	//println(i)
}
