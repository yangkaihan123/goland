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