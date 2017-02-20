package main
// 实现一个解析并格式化命令行输入的时间集合的例子
import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

type interval []time.Duration

//实现String接口
func (i *interval) String() string {
	return fmt.Sprintf("%v", *i)
}

//实现Set接口,Set接口决定了如何解析flag的值
func (i *interval) Set(value string) error {
	fmt.Println("flag origin value", value);
	//此处决定命令行是否可以设置多次-deltaT
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		fmt.Println(duration);
		if err != nil {
			return err
		}
		// 格式化之后，放到数组
		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	// 解析命令行参数
	flag.Parse()
	fmt.Println(intervalFlag)
}