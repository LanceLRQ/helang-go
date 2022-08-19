# HeLang-Go
Next-Generation Cyber Programming Language from Li Tang, with golang support.

## 介绍

次世代赛博编程语言，诞生于E-SMOKER之乡：赛博理塘。

本项目是由自己会打字的5G键盘，在上班摸鱼期间，搜索并发现了 `https://github.com/kifuan/helang` 项目后，模仿着该项目的源代码，在AirDesk上配合AirPods编写的。

5G键盘发现，原项目运行程序的体验并不太好。毕竟，写好了一段代码，发给你亲爱的朋友们，他们却要抱怨需要安装Python环境。于是自作主张，增加了把它编译成可执行文件的功能！

实在是太酷了，很符合我对未来生活的想象，科技并带着趣味。

## 为什么用He语言

众所周知，He语言非常的酷！

相比起来，C语言的语法存在许多问题。比如对某些人来说，"或" 和 "位或" 运算符都是不会区分的。

学不会C语言不要紧，让我们一起来学习一下he语言吧，简单易懂，一学就会！

## 使用方法

0. 安装`golang`环境，建议`1.16及以上` 
- [国内](https://studygolang.com/dl) https://studygolang.com/dl
- [官网(需要梯子)](https://go.dev/dl/) https://go.dev/dl/

1. 拉取源代码并安装一些依赖（测5G用的）

```shell
> git clone git@github.com:LanceLRQ/helang-go.git
> cd helang-go
> go get
```

2. 编译程序

```shell
go run compiler.go great.he
```

3. 执行编译

```shell
./great
```

windows下为：（我没测过啊，正经赛博人谁用windows啊，不都是用Mac吗？）

```shell
./great.exe
```

## 语法

语法参见原作 [kifuan/helang](https://github.com/kifuan/helang)

下面是开始改写代码时，参考的语法版本

---

## 基本语法

**Saint He** 曾说，一切类型均为`u8`，是什么意思呢？这个词倒过来就是`8u`，看来圣人也喜欢玩贴吧。

除此之外，`u8` 与 `v8` 形似，所以能不能关注[永雏塔菲](https://space.bilibili.com/1265680561)喵，关注永雏塔菲谢谢喵。

如你所见，我们用**bitwise or**，即`|`代替了传统数组的符号。**都什么年代了还在写传统数组**？

```c
u8 a = 1 | 2 | 3;
```

**Saint He** 曾说：`whichKey - 1` ，所以我们数组的下标需要从 `1` 开始。

```c
u8 a = 1 | 2 | 3;
print a[1];
// 1
```

但是，当你设置一个 `u8` 的元素时，你可以用 `0` 作为下标：这意味着所有元素都将被赋值。

```c
u8 a = 1 | 2 | 3;
a[0] = 10;
print a;
// 10 | 10 | 10
```

为了符合最新的技术，我们同样支持多下标操作，所以你再也不用写 `for` 循环了。

```c
u8 a = 1 | 2 | 3;
a[1 | 2] = 0;
print a;
// 0 | 0 | 3
```

同样，我们还提供了一种根据数组长度的初始化方式，可惜这还是传统写法。比如下面的代码，可以初始化一个长度为5的数组。

```c
u8 a = [5];
print a;
// 0 | 0 | 0 | 0 | 0
```

最后，我们结合一下，可以写出下列代码。

```c
u8 forceCon = [68];

forceCon[1 | 2 | 6 | 7 | 11 | 52 | 57 | 58 | 65] = 10;

print forceCon;
```

如此精妙的代码，在地球的人类是无法理解的。我们作为**赛博智能生命体**，也只能给你演示一下日常操作了。

这实在是太酷了，后面我忘了，我也不想翻到文章开头去看。

## Hello, world!

有人认为何语言无法打印出`Hello, world!`意味着它太垃圾了，实则不然。

**Saint He** 专注于单片机应用的开发，哪里有时间顾及字符串？

好在，经过协商，他同意了这个请求，使得我们能在这门语言中打印出`Hello, world!`

```c
sprint 72 | 101 | 108 | 108 | 111 | 44 | 32 | 119 | 111 | 114 | 108 | 100 | 33;
// Hello, world!
```

由于对效率的极端苛刻要求，我们使用字符在 UTF-8 中对应的数字来表示这个字符。

通过降低可读性，换来了指数级的性能提升，不愧赛博世界的唯一真神。

## 自增运算

我们注意到，**Saint He** 的代码中还出现了自增运算：`++`

所以，我们也支持这种运算。

```c
u8 a = 1 | 2 | 3;
a++;
print a;
// 2 | 3 | 4
```

## 变量声明与赋值

在早期版本中，我们的仅仅支持变量的定义，不支持修改和声明，因为 **Saint He** 喜欢 `immutable`。

经过意见征求，现在已经可以做到这三个方面了！

```c
// 现在支持变量先声明后定义，可以写出这种代码：
u8 a;
a = 1 | 2;

// 早期版本仅支持下面这种写法，当然现在也支持：
u8 b = 3 | 4;
```

注意！如果你没有声明或定义一个变量，尝试直接给它赋值，你会收获`CyberNameException`。

```c
c = 1 | 2 | 3;
// helang.exceptions.CyberNameException: c is not defined.
```

我们作为高科技语言，当然是需要严谨的。

## 查看你是否在 Cyber Spaces

为了见到 **Saint He**，你需要身处 **Cyber Spaces**。输入下方命令查看你是否身处其中：

```c
cyberspaces;
// Getting your location...
// Your location is UNITED STATES.
// Congratulations! You are in the Cyber Spaces!
```

## 5G测速

理论上这块也属于语法，但我就是要把它单独摘出来。

很简单，只需要另起一行输入：

```c
test5g;
```

即可从本行开始5G测速。

---

## 单元测试

最新一期视频还在制作中...

## 开源协议

这点破代码，爱干嘛干嘛，随便玩

## 未来更新？

- 单元测试
- 加点别的啥语法？（原版好像有新的内容！有空再搬过来）