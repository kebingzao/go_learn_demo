package main

import (
	"fmt"
	"time"
)


func main() {

	//now := time.Now()
	//// 时间戳
	//secs := now.Unix()
	//nanos := now.UnixNano()
	//fmt.Println(now)
	//millis := nanos / 1000000
	//
	//fmt.Println(secs)
	//fmt.Println(millis)
	//fmt.Println(nanos)
	//
	//fmt.Println(time.Unix(secs, 0))
	//fmt.Println(time.Unix(0, nanos))

	// http://studygolang.com/articles/240

	p := fmt.Println

	now := time.Now()
	p(now)

	d := time.Duration(7200 * 1000 * 1000 * 1000)
	p(d)

	then := time.Date(
		2013, 1, 7, 20, 34, 58, 651387237, time.UTC)

	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	p(then.Date())
	p(then.ISOWeek())
	p("----------")
	p(now.UTC())
	p(now.Local())
	p(now.Location())
	p(now.Zone())
	p(now.Unix())
	p(time.Unix(now.Unix(), 0))
	p(now.UnixNano())
	p(time.Unix(0, now.UnixNano()))
	p(now.GobEncode())
	p(now.MarshalJSON())
	p(time.Since(now))
	p("----------")
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))

	p(d)
	p(d.Hours())
	p(d.Minutes())
	p(d.Seconds())
	p(d.Nanoseconds())
	p(then.Add(d))

}