扩展与复用
-
不支持LSP不支持重载，没有继承

这里继续查看其他文档看看...

Go 不⽀持继承，但可以通过复合的⽅式来复⽤
###匿名类型嵌入
它不是继承，如果我们把“ 内部 struct ”看作⽗类，把“外部 struct”  看作⼦类， 会发现如下问题：   
1.  不⽀持⼦类替换
2.  ⼦类并不是真正继承了⽗类的⽅法   
        • ⽗类的定义的⽅法⽆法访问⼦类的数据和⽅法