### go

##### defer

在golang当中,defer代码块会在函数调用链表中增加一个函数调用.这个函数调用不是普通的函数调用,而是会在函数正常返回,也就是return之后添加一个函数调用.因此,defer通常用来释放函数内部变量.



###### 规则一 当defer被声明时,其参数就会被实时解析

######　规则二　defer执行顺序为先进后出

###### 规则三 defer可以读取有名返回值

