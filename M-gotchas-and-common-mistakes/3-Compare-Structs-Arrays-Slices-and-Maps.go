package main

/*
比较Structs, Arrays, Slices, and Maps
如果结构体中的各个元素都可以用,你可以使用等号来比较的话，那就可以使用 '==' 来比较结构体变量。
*/

//比较 结构体 例子：
/*
import "fmt"
type data struct {
	num int
	fp float32
	complex complex64
	str string
	char rune
	yes bool
	events <-chan string
	handler interface{}
	ref *byte
	raw [10]byte
}
func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",v1 == v2) //prints: v1 == v2: true
}
*/


//如果结构体中的元素无法比较，那使用等号将导致编译错误。注意数组仅在它们的数据元素可比较的情况下才可以比较。

/*
import "fmt"
type data struct {
	num int                //ok
	checks [10]func() bool //not comparable
	doit func() bool       //not comparable
	m map[string] string   //not comparable
	bytes []byte           //not comparable
}
func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",v1 == v2)
}
*/

//print::
//M-gotchas-and-common-mistakes\3-Compare-Structs-Arrays-Slices-and-Maps.go:44:
// invalid operation: v1 == v2 (struct containing [10]func() bool cannot be compared)

// Go确实提供了一些助手函数，用于比较那些无法使用等号比较的变量。
// 最常用的方法是使用reflect包中的DeepEqual()函数。

/*
import (
	"fmt"
	"reflect"
)
type data struct {
	num int                //ok
	checks [10]func() bool //not comparable
	doit func() bool       //not comparable
	m map[string] string   //not comparable
	bytes []byte           //not comparable
}
func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:",reflect.DeepEqual(v1,v2)) //prints: v1 == v2: true
	m1 := map[string]string{"one": "a","two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:",reflect.DeepEqual(m1, m2)) //prints: m1 == m2: true
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:",reflect.DeepEqual(s1, s2)) //prints: s1 == s2: true
}
*/

// 除了很慢（这个可能会也可能不会影响你的应用），DeepEqual()也有其他自身的技巧。


/*
import (
	"fmt"
	"reflect"
    "bytes"
)
func main() {
	var b1 []byte = nil
	b2 := []byte{}
	fmt.Println("b1 == b2:",reflect.DeepEqual(b1, b2)) //prints: b1 == b2: false
    fmt.Println("b1 == b2:",bytes.Equal(b1, b2)) //prints: b1 == b2: true
}
*/


// DeepEqual()不会认为空的slice与“nil”的slice相等。
// 这个行为与你使用bytes.Equal()函数的行为不同。
// bytes.Equal()认为“nil”和空的slice是相等的。

// 但是 DeepEqual()在比较slice时并不总是完美的。 比如：

import (
	"fmt"
	"reflect"
	"encoding/json"
)
func main() {
	var str string = "one"
	var in interface{} = "one"
	fmt.Println("str == in:",str == in,reflect.DeepEqual(str, in))
	//prints: str == in: true true
	v1 := []string{"one","two"}
	v2 := []interface{}{"one","two"}
	fmt.Println("v1 == v2:",reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: false (not ok)
	data := map[string]interface{}{
		"code": 200,
		"value": []string{"one","two"},
	}
	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("decode:", decoded)
	fmt.Println("data == decoded:",reflect.DeepEqual(data, decoded))
	//prints: data == decoded: false (not ok)
}


/*
如果你的byte slice（或者字符串）中包含文字数据，而当你要不区分大小写形式的值时（在使用==，bytes.Equal()，或者bytes.Compare()），你可能会尝试使用“bytes”和“string”包中的ToUpper()或者ToLower()函数。
对于英语文本，这么做是没问题的，但对于许多其他的语言来说就不行了。这时应该使用strings.EqualFold()和bytes.EqualFold()。

如果你的byte slice中包含需要验证用户数据的隐私信息（比如，加密哈希、tokens等），不要使用reflect.DeepEqual()、bytes.Equal()，或者bytes.Compare()，因为这些函数将会让你的应用易于被定时攻击。
为了避免泄露时间信息，使用'crypto/subtle'包中的函数（即，subtle.ConstantTimeCompare()）。
*/









