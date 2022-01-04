**字符串**  
与其他的差异  
string的零值是一个空字符串  
1.string是数据类型，不是引用或指针类型  
2.string是只读的byte slice, len() 函数可以查到string包含的byte数  
byte数和字符串的多少没关系  
3.string的byte数组可以存放任何数据    
Unicode 和 UTF8  
1.Unicode是一种字符集（code point）   
2.UTF8 是unicode的存储实现 (转换为字节序列的规则)  
就是UTF8是unicode字符集转换成机器可以识别的语言的规则    

编码与存储   
字符： ”中“    
unicode： 0x4E2D    
utf-8： 0xE4B8AD   
string/[]byte： [0xE4,0xB8,0xAD]   string的切片里放的是[0xE4,0xB8,0xAD]    
常用的字符串函数： string包（https://golang.org/pkg/strings)   
                strconv包（https://golang.org/pkg/strconv)    