# go 学习





## go 安装(vscode)

### 1) 前往go中文网下载go安装包(windows版本)

```
https://studygolang.com/dl
```

### 2) 安装完成后 配置环境变量

```
# 在windows系统属性里
# 增加系统变量
GOROOT  设置go的安装目录
GOPATH  设置go的工作目录并同时修改用户变量(go的工作目录包含src,pkg,bin等目录sec放置代码文件)
```

```go
# 设置vscode开发环境
# 下载拓展包go
ctrl+shift+p 打开命令输入 go:install/update tools 全部安装
ctrl+shift+p 打开命令输入 launch启动命令中心
输入 settings.json
"[go]": {
        "editor.insertSpaces": false,
        "editor.formatOnSave": true,
        "editor.formatOnType": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true
        }
    }
```

### 3)第一个go代码

```go
package main                  # 引进一个包
import "fmt"                  # 导入一个系统包fmt用来输出
func main(){                  # func(函数)
  fmt.Println("你好")         # 打印输出你好其中ln为换行
}

# 运行代码
# 到代码所在文件夹
# cmd输入go run 文件.go

```

### 注意事项:

```go
在非工作目录运行go文件出错,需要在windows命令行更改参数
go env
go env -w GO111MODULE=off
```





## 基础语法学习

### 注释

1) #### 单行注释

```golang
package main
// 单行注释
import "fmt"

func main(){
  fmt.Println("")
}
```

2) #### 多行注释

```golang
package main
/*
多行注释
多行注释
*/
import "fmt"

func main(){
  fmt.Println("")
}
```

### 变量

#### 变量的定义

Go语言是静态类型语言,就是所有的类型我们都需要明确的去定义.这里先不管其他类型,先了解string,我们用它来表示字符串

在Go语言中,声明一个变量一般是使用==var==关键字:

```golang
var name type
```

- 第一个 var 是声明变量的关键字,是固定的写法

- 第二个 name 是变量的名字

- 第三个 type 是代表变量的类型

举个例子:

```golang
// 定义一个字符串变量 name
var name string

// 定义一个数字类型变量 age
var age int
```

在声明变量时将变量的类型放在变量名称之后.这样的好处就是可以避免像C语言中那样含糊不清的声明形式,例如: int* a, b;其中只有a是指针而b不是.如果想要将这两个变量都是指针,则需要将他们分开书写.而在Go则可以轻松的将他们都声明为指针类型:

```golang
var a, b *int
```

变量的命名规则遵循骆驼命名法,即首个单词小写,每个新单词的首字母大写,例如: userFiles和systemInfo.

> 变量定义的标准格式为

```golang
var 变量名 变量类型
```

变量声明以关键字 var 开头,后置变量类型,行尾无需分号

使用关键字 var 和括号,可以将一组变量定义放在一起

```go
var (
    addr string
    phone string
)
```

var 形式的声明语句往往是用于需要显示指定变量类型的地方,或者因为变量稍后会被重新赋值而初始值无关紧要的地方.

当一个变量被声明之后,如果没有显示的给它赋值,系统会自动赋予它该类型的零值:

- 整型和浮点型变量的默认值为0和0.0
- 字符串变量的默认值为空字符串
- 布尔型变量的默认为false
- 切片,函数,指针变量的默认为null

#### 变量的初始化

> 变量初始化的基本格式

- var 变量名 类型 = 值(表达式)

比如,我想定义go的一些信息,我可以这么表示

```go
var name string = "go"
var age int = 10

fmt.Println("name:%s,age:%d",name,age)
```

这里的name和age就是变量名,name的类型是string,age的类型是int

> 短变量声明并初始化

```go
name1 := "go"
age1 := 10

fmt.Println("name:%s,age:%d",name1.age1)
```

这是go的推导声明写法,编译器会自动根据右边值的类型判断出左边值的对应类型

它可以自动的推导出一些类型,但是使用也是有限制的

- 定义变量,同时显示初始化
- 不能提供数据类型
- 只能在函数内部。不能随便到处定义

短变量声明被广泛用于大部分的局部变量的声明和初始化

注意：由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，会发生编译错误。

#### 理解变量(内存地址)

```go
var num int
num = 100
fmt.Printf("num的值: %d，num的内存地址: %p\n",num,&num)
num = 200
fmt.Printf("num的值: %d，num的内存地址: %p\n",num,&num)
```

#### 变量的交换

```go
package main

import "fmt"

/*
	在编程中最简单的算法就是变量的交换，最常见的方式就是定义中间变量
	var a int=100
	var b int=200
	var t int
	t = a
	a = b
	b = t
	fmt.Println(a,b)
*/

func main() {
    var a int = 100
    var b int = 200
    //在go语言中，可以直接这样来实现换值，不需要中间变量
    b,a = a,b
    fmt.Println(a.b)
}
```

#### 匿名变量

匿名变量的特点是一个下划线"_",下划线本身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给他），但任何赋给这个标识符的值都将被抛弃，因此这些值不能再后续的代码中使用，也不可以使用这个标识符作为变量对其他变量进行赋值或运算。使用匿名变量时，只需要在变量声明的地方使用下划线替换即可。例如：

```go
package main

import "fmt"

/*
定义一个test函数,他会返回两个int类型的值,每次调用会返回100和200两个数值
*/
func test() (int, int) {
	return 100, 200
}
func main() {
	a, _ := test()
	//这里需要获取第一个返回值,所以第二个返回值定义为匿名变量
	_, b := test()
	//这里需要获取第二个返回值,所以第一个返回值定义为匿名变量
	fmt.Println(a, b)
}
```

在编码过程中，可能会遇到没有名称的变量、类型或方法。虽然这不是必须的，但有时候这样做可以极大的提高代码的灵活性，这些变量被统称为匿名变量。

匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。

#### 变量的作用域

一个变量（常量、类型或函数）在程序中都有一定的作用范围，被称为作用域。

了解变量的作用域很重要，因为Go语言会在编译时检查每个变量是否使用过，一旦出现未使用的变量，就会报编译错误

> 局部变量

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，函数的参数和返回值变量都属于局部变量

```go
package main

import "fmt"

func main() {
	var a int = 3
	var b int = 7
	//声明局部变量a,b并赋值
	c := a + b
	//声明局部变量c并计算a+b的值
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}
```

> 全局变量

在函数体外声明的变量称之为全局变量，全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，当然，不包含这个全局变量的源文件需要使用"import"关键字引用全局变量所在的源文件之后才能使用这个全局变量。

全局变量声明必须以"var"关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。

```go
package main

import "fmt"

var c int

// 定义一个全局变量
func main() {
	var a, b int
	//定义局部变量
	a = 3
	b = 7
	//初始化参数
	c = a + b
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}
```

Go语言中全局变量和局部变量名称可以相同，但是函数体内的局部变量会被优先考虑。

#### 常量

**常量是一个简单值的标识符，在程序运行时，不会被修改的量**

常量中的数据类型只可以是布尔型、数字型（整数型，浮点型和复数）和字符串型

```go
const identifier [type] = value
```

可以省略类型说明符[type]因为编译器可以根据变量的值来推断其类型。

- 显式类型定义：const b string = "abc"
- 隐式类型定义：const b = "abc"

多个相同类型的声明可以简写为：

```go
const c_name1 , c_name2 = value1 , value2
```

```go
package main

import "fmt"

func main() {
	const URL string = "www.baidu.com"     //显式类型
	const URL1 = "www.baidu.com"           //隐式类型
	const a, b, c = 3.1415926, true, "yes" //多重赋值
	fmt.Println(URL)
	fmt.Println(URL1)
	fmt.Println(a, b, c)
}
```

#### iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量，iota是go语言的常量计数器。

iota在const关键字出现时将被重置为0（const内部的第一行之前），const中每新增一行常量声明将使iota计数一次（iota可理解为const语句块中的行引索）























