**基本的数据类型**
-
    bool  
    string  
    int int8 int16 int32 int64  
    uint uint8 uint16 uint32 uint64 uintptr  
    byte //alias for uint8  
    rune //alias for int32,represents a Unicode code point  
    float32 float64  
    complex64 complex128  
###**差异**
    1.Go不允许隐式转换   
    2.别名类型和原有类型也不可以隐式转换  
###**类型的预定义值**
    1.math.MaxInt64   
    2.math.MaxFloat64  
    3.math.MaxUint32  
###**差异**
    1.不支持指针运算  
    2.string是值类型，其初始化类型为空字符串，不是nil  
###**总结**
    不支持隐式转换，不支持指针预算，字符类型是值类型，初始化值是空字符串不是nil  
