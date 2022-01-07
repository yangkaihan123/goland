**strconv和strings的用法**
-
    1.前缀和后缀
    HasPrefix 判断字符串 s 是否以prefix开头
    strings.HasPrefix(s,prefix string) bool
    HasSuffix 判断字符串是否以suffix结尾
    strings.HasSuffix(s,suffix string) bool
    
    2.字符串包含关系
    Contains 判断字符串s是否包含substr
    格式： strings.Contains(s, substr string) bool
    
    3.判断子字符串或字符在父字符串中出现的位置（索引）
    Index返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str
    strings.Index(s, str string) int
    LastIndex返回字符串str在字符串s中最后出现位置的索引（str的第一个字符的索引），-1 表示字符串s不包含str
    strings.LastIndex(s,str string) int
    如果字符串是非ASCII编码的字符，可以使用IndexRune函数来对字符串定位
    strings.IndexRune(s string, r rune) int

    4.字符串替换
    Replace 用于将字符串 str 中的前n个字符串 old 替换为字符串 new ，并返回一个新字符串，如果n = -1，则替换所有字符串old为字符串new
    strings.Replace(str, old, new, n) string
    
    5.统计字符串出现的次数
    Count 用于计算字符串str在字符串s中出现的非重叠次数
    strings.Count(s,str string) int

    6.重复字符串
    Repeat用于重复count次字符串s并返回一个新的字符串
    strings.Repeat(s,count int) string
    
    7.修改字符串大小写
    Tolower 将字符串中的Unicode字符全部转换为相应的小写字符
    strings.Tolower(s) string
    ToUpper 将字符串中的Unicode字符全部转换为相应的大写字符
    strings.ToUpper(s) string

    8.修剪字符串
    strings.TrimSpace(s)来剔除字符串开头和结尾的空白符号
    strings.Trim(s,"s")来把开头和结尾的s去掉，这个函数第二个参数可以包含任何字符
    strings.TrimsLeft(s,"s")只剔除开头
    strings.TrimsRight(s,"s")只剔除结尾

    9.分割字符串
    strings.Fields(s)利用空白作为分隔符将字符串分割为若干块，返回为一个切片
    如果字符串只包含空白符号，返回一个长度为0的切片
    strings.Split(s,sep)自定义分割符号对字符串分割，返回一个切片
    
    10.拼接slice到字符串
    Join 用于将元素类型为string的slice使用分割符号来拼接组成一个字符串
    strings.Join(s1 []string,sep string) string

    11.从字符串中读取内容
    strings.NewReader(str)用于生成一个Reader并读取字符串中的内容，返回的是指向这个Reader的指针
        。 Read() 从[] byte中读取内容
        。 ReadByte() 和 ReadRune() 从字符串中读取下一个byte或Rune

    12.字符串和其他类型的转换
    通过strconv包实现
    该包包含了一些变量用于获取程序运行的操作系统平台下 int 类型所占的位数，如：strconv.IntSize。
    任何类型 T 转换为字符串总是成功的。
    针对从数字类型转换到字符串，Go 提供了以下函数：
        。 strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数
        。 strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，
        其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64
    
    将字符串转换为其它类型 tp 并不总是可能的，可能会在运行时抛出错误 parsing "…": invalid argument
    针对从字符串类型转换为数字类型，Go 提供了以下函数：
        。strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型
        。strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型
    利用多返回值的特性，这些函数会返回 2 个值，第 1 个是转换后的结果（如果转换成功），第 2 个是可能出现的错误，因此，我们一般使用以下形式来进行从字符串到其它类型的转换