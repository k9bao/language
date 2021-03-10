package test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

//http://docscn.studygolang.com/pkg/time/
const TimeFormat = "2006-01-02 15:04:05"

func AfterDemo() {
	time.AfterFunc(time.Millisecond,
		func() {
			fmt.Println("243")
			if true {
				fmt.Println("23333")
			}
		},
	)
}

//2006-01-02 15:04:05 07 --- 年 月 日 时 分 秒 时区
func FmortDemo() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339Nano))
	fmt.Println(now.Format("2006-Jan-2"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("15:04:05.000"))
}

func ConstructDemo() {
	t1, _ := time.Parse("2006-01-02 15:04:05", "2013-08-11 11:18:46")
	fmt.Println(t1)

	t2, _ := time.Parse("2006-01-02:15", "2019-05-24:13")
	fmt.Println(t2)

	t3 := time.Unix(1362984425, 0)
	fmt.Println(t3)
}

func DurationDemo() {
	t1 := time.Now()
	time.Sleep(time.Second)
	fmt.Println(time.Since(t1))
	fmt.Println(time.Since(t1) / time.Second)
	fmt.Println(time.Since(t1) / time.Millisecond)
	fmt.Println(time.Since(t1) / time.Microsecond)
}

func TickDemo() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	count := 0
	start := time.Now()
EXIT:
	for {
		select {
		case <-tick.C:
			log.Println("tick is coming", time.Since(start))
			count++
		default:
			//log.Println("sleep")
			if count%3 == 0 {
				time.Sleep(2 * time.Second)
			}
			if time.Since(start) > time.Second*15 {
				break EXIT
			}
		}
	}
}

func TestTime(t *testing.T) {
	// now := time.Now()
	// time.Sleep(time.Second)
	// log.Println(int(time.Since(now))) //1000,708,900 微妙

	t3399, _ := time.Parse(time.RFC3339, "2020-03-12T00:00:00Z")
	fmt.Println(t3399.UnixNano() / 1000000)

	// AfterDemo()
	//FmortDemo()
	// ConstructDemo()
	// DurationDemo()
	//TickDemo()
}
