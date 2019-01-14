package main


/*
String和Byte Slice之间的转换
当你把一个字符串转换为一个byte slice（或者反之）时，你就得到了一个原始数据的完整拷贝。
这和其他语言中cast操作不同，也和新的slice变量指向原始byte slice使用的相同数组时的重新slice操作不同。
Go在[]byte到string和string到[]byte的转换中确实使用了一些优化来避免额外的分配（在todo列表中有更多的优化）。
第一个优化避免了当[]byte keys用于在map[string]集合中查询时的额外分配:m[string(key)]。
第二个优化避免了字符串转换为[]byte后在for range语句中的额外分配：for i,v := range []byte(str) {...}
*/
