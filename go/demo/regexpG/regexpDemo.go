package regexpG

import (
	"bytes"
	"fmt"
	"regexp"
)

func PrintSlice2(id string, text [][]string) {
	fmt.Println(id, "len=", len(text))
	for _, v := range text {
		PrintSlice1(id, v)
	}
}

func PrintSlice1(id string, text []string) {
	fmt.Print(id, "len=", len(text), "---")
	for i, v := range text {
		fmt.Print(i, ":", v, "---")
	}
	fmt.Println("")
}

func DemoRegexp() {
	//说明:所有的函数,去掉String字段就是对slice的操作,返回值也是slice
	match, _ := regexp.MatchString("H(.*)d!", "Hello World!") //是否匹配字符串,.匹配任意一个字符 ，*匹配零个或多个 ，优先匹配更多(贪婪)
	fmt.Println("01:", match)                                 //true

	//通过`Compile`来使用一个优化过的正则对象,
	r, _ := regexp.Compile("(h|H)[a-z][a-z][a-z](d|o)")
	checkStr := "Hello!hello"
	fmt.Println("02:", r.MatchString(checkStr)) //true

	fmt.Println("03:", r.FindString(checkStr))      //第一次匹配的子串,贪婪匹配
	fmt.Println("04:", r.FindStringIndex(checkStr)) //同上,返回索引,[开始位置,结束位置(不包含结束位置)]

	PrintSlice1("05:", r.FindAllString(checkStr, -1)) //返回所有正则匹配的字符，不仅仅是第一个,返回列表

	PrintSlice1("06:", r.FindStringSubmatch(checkStr))             //在FindString的基础上,返回局部匹配内容.所谓局部匹配就是通过匹配得到的内容
	fmt.Println("07:", r.FindStringSubmatchIndex(checkStr))        //同上,返回对应索引
	PrintSlice2("08:", r.FindAllStringSubmatch(checkStr, -1))      //类似FindStringSubmatch,只不过返回的是全部.双层列表
	fmt.Println("09:", r.FindAllStringSubmatchIndex(checkStr, -1)) //同上,返回对应索引

	// 为这个方法提供一个正整数参数来限制匹配数量
	res, _ := regexp.Compile("H([a-z]+)d!")
	PrintSlice1("10:", res.FindAllString("Hello World! Held! Hellowrld! world", -1)) //[Held! Hellowrld!]
	PrintSlice1("11:", res.FindAllString("Hello World! Held! Hellowrld! world", 2))  //[Held! Hellowrld!]

	fmt.Println(r.FindAllString("Hello World! Held! world", 2)) //[Hello World! Held!]
	//注意上面两个不同，第二参数是一最大子串为单位计算。

	// regexp包也可以用来将字符串的一部分替换为其他的值
	fmt.Println(r.ReplaceAllString("Hello World! Held! world", "html")) //html world

	// `Func`变量可以让你将所有匹配的字符串都经过该函数处理
	// 转变为所需要的值
	in := []byte("Hello World! Held! world")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	// 在 b 中查找 reg 中编译好的正则表达式，并返回第一个匹配的位置
	// {起始位置, 结束位置}
	b := bytes.NewReader([]byte("Hello World!"))
	reg := regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindReaderIndex(b)) //[0 5]

	// 在 字符串 中查找 r 中编译好的正则表达式，并返回所有匹配的位置
	// {{起始位置, 结束位置}, {起始位置, 结束位置}, ...}
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项

	fmt.Println(r.FindAllIndex([]byte("Hello World!"), -1)) //[[0 12]]
	//同上
	fmt.Println(r.FindAllStringIndex("Hello World!", -1)) //[[0 12]]

	// 在 s 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
	// 同时返回子表达式匹配的内容
	// {
	//     {完整匹配项, 子匹配项, 子匹配项, ...},
	//     {完整匹配项, 子匹配项, 子匹配项, ...},
	//     ...
	// }
	// 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
	reg = regexp.MustCompile(`(\w)(\w)+`)                      //[[Hello H o] [World W d]]
	fmt.Println(reg.FindAllStringSubmatch("Hello World!", -1)) //[[Hello H o] [World W d]]

	// 将 template 的内容经过处理后，追加到 dst 的尾部。
	// template 中要有 $1、$2、${name1}、${name2} 这样的“分组引用符”
	// match 是由 FindSubmatchIndex 方法返回的结果，里面存放了各个分组的位置信息
	// 如果 template 中有“分组引用符”，则以 match 为标准，
	// 在 src 中取出相应的子串，替换掉 template 中的 $1、$2 等引用符号。
	reg = regexp.MustCompile(`(\w+),(\w+)`)
	src := []byte("Golang,World!")           // 源文本
	dst := []byte("Say: ")                   // 目标文本
	template := []byte("Hello $1, Hello $2") // 模板
	m := reg.FindSubmatchIndex(src)          // 解析源文本
	// 填写模板，并将模板追加到目标文本中
	fmt.Printf("%q", reg.Expand(dst, template, src, m))
	// "Say: Hello Golang, Hello World"

	// LiteralPrefix 返回所有匹配项都共同拥有的前缀（去除可变元素）
	// prefix：共同拥有的前缀
	// complete：如果 prefix 就是正则表达式本身，则返回 true，否则返回 false
	reg = regexp.MustCompile(`Hello[\w\s]+`)
	fmt.Println(reg.LiteralPrefix())
	// Hello false
	reg = regexp.MustCompile(`Hello`)
	fmt.Println(reg.LiteralPrefix())
	// Hello true

	text := `Hello World! hello world`
	// 正则标记“非贪婪模式”(?U)
	reg = regexp.MustCompile(`(?U)H[\w\s]+o`)
	fmt.Printf("%q\n", reg.FindString(text)) // Hello
	// 切换到“贪婪模式”
	reg.Longest()
	fmt.Printf("%q\n", reg.FindString(text)) // Hello Wo

	// 统计正则表达式中的分组个数（不包括“非捕获的分组”）
	fmt.Println(r.NumSubexp()) //1

	//返回 r 中的“正则表达式”字符串
	fmt.Printf("%s\n", r.String())

	// 在 字符串 中搜索匹配项，并以匹配项为分割符，将 字符串 分割成多个子串
	// 最多分割出 n 个子串，第 n 个子串不再进行分割
	// 如果 n < 0，则分割所有子串
	// 返回分割后的子串列表
	fmt.Printf("%q\n", r.Split("Hello World! Helld! hello", -1)) //["" " hello"]

	// 在 字符串 中搜索匹配项，并替换为 repl 指定的内容
	// 如果 rep 中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
	// 全部替换，并返回替换后的结果
	s := "Hello World, hello!"
	reg = regexp.MustCompile(`(Hell|h)o`)
	rep := "${1}"
	fmt.Printf("%q\n", reg.ReplaceAllLiteralString(s, rep)) //"${1} World, hello!"

	// 在 字符串 中搜索匹配项，然后将匹配的内容经过 repl 处理后，替换 字符串 中的匹配项
	// 如果 repb 的返回值中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
	// 全部替换，并返回替换后的结果
	ss := []byte("Hello World!")
	reg = regexp.MustCompile("(H)ello")
	repb := []byte("$0$1")
	fmt.Printf("%s\n", reg.ReplaceAll(ss, repb))
	// HelloH World!

	fmt.Printf("%s\n", reg.ReplaceAllFunc(ss,
		func(b []byte) []byte {
			rst := []byte{}
			rst = append(rst, b...)
			rst = append(rst, "$1"...)
			return rst
		}))
	// Hello$1 World!

}
