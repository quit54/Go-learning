# GO语言module精粹（参照zerotolearngo目录）
*参考文档：(https://golang.google.cn/doc/tutorial/call-module-code,"官方技术文档")*
### go语言的起点，什么叫做module？
- module是包的集合，包是一堆离散而有用的函数的集合，可以单独创建分别负责不同模块功能的包，之后将包进行集合为module的utils工具，之后任何调用可以通过调用这个module的函数，来直接使用。
### module内涵内容
module指定运行代码所需的依赖项，包括Go版本和它所需的其他模块集，在调用的时候发挥作用。
### module指令集合
1. '''go mod init + (module名称)'''初始化命令,会生成一个go.mod文件，关联依赖项，一开始文件只包含module的模型和GO语言版本，后面可增加依赖项。
2. '''$ go mod edit -replace example.com/greetings=../greetings'''go语言编辑指令
3. '''$ go mod tidy   go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000'''go语言tidy指令，同步example.com/hello模块的依赖项，增加本地指令
### （卡点）module应用解析
- 为了保重我的每一个module模型正常运行，需要反复使用go mod init,为代码创建依赖项跟踪。
- 当你想在一个main包中调用其他module的代码时，需要做的是：import{"fmt"标准库,"example.com/greetings"其他module的名称}
- 这样你才可以调用其他module的函数，编辑example.com/hello模块以使用本地的example.com/greetings模块。其中greetings包是存放各种工具的，hello是调用工具的。
**这里设立greetings的module模型之后，还需要对go.mod文件进行重新编辑，将Go工具从模块路径（模块不存在的地方）重定向到本地目录（模块所在的地方），编辑module为了之后它可以在本地文件系统被调用**
'''
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings #依赖项 
'''
这里达到设立依赖项的目的。
**之后，$ go mod tidy指令，向模型提供一个需要指令，才可以运行。**
'''
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
'''
- 在hello.go调用greetings的时候，会创建依赖项的。模块路径后面的数字是一个伪版本号——一个生成的数字，用来代替语义版本号（模块还没有）。
- 最终，使用一个带有版本号标签的require指令。
# 总结module与依赖项的关系
**1. Module (模块)：这个文件就像是你这个项目的身份证和食谱。**
**2. 依赖项 (Dependencies)：你的项目为了运行，所需要的外部代码库**
**关系是： 一个 Module 会声明它所需要的所有依赖项。 这些声明就写在 go.mod 文件里。**
- require 指令：我需要什么，用来声明项目正常编译和运行所必须依赖的另一个模块及其版本。
- replace 指令：替换“我需要的”为“我指定的”，replace 指令用于将某个依赖的模块路径，替换为另一个本地路径或代码库路径。它就像一个“重定向”或“覆盖”规则。这是一个高级且临时的操作，通常在特定开发场景下使用。
## 关于replace指令的深度思考：
### 1.为什么import指令引用的是URL，而不是本地库文件？
1. 全球唯一的命名空间：Go 设计这个规则是为了让每个包都有一个全球唯一的名字。使用代码托管的URL（如 GitHub, GitLab 等）是最简单、最自然的方式，可以确保不同作者、不同项目的包不会重名。
2. 可重定位性：这个路径告诉 Go 工具链 “去哪里找到这个包”。当你运行 go mod tidy 或 go build 时，Go 命令会读取这些导入路径，然后通过与模块代理（proxy.golang.org）或版本控制系统（如 git）交互，将这些路径解析并下载为实际的代码到你的本机缓存中（$GOPATH/pkg/mod）。
3. 永久的、稳定的标识符：你的代码应该永远引用 github.com/gin-gonic/gin，无论这个库的代码实际存储在哪里（也许未来 Gin 团队换了平台，但会有重定向机制保证这个路径依然有效）。这保证了你的代码的长期稳定性。
- 它的核心目的不是“下载”，而是“唯一标识”。所以，Import Path 是一个永恒的、唯一的名字，而不是一个临时的下载指令。
### 2.为什么不能直接修改代码里的 Import Path？
1. 破坏协作和构建：你的项目 collaborators（协作者）clone 你的代码后，他们的电脑上根本没有 ../my-local-gin 这个目录，会导致编译立即失败。
2. 污染代码库：你将一个本地化的、临时的配置决策，永久地写入了源代码中。这是非常糟糕的实践。
3. 忘记改回来：你很容易忘记把这个 import 路径改回正确的官方路径，从而把错误代码提交到仓库，或者甚至发布出去。
**replace 指令的强大之处就在于它完美地解耦了“代码标识（What）”和“代码来源（Where）”。**
*思考是对的。Go强制要求Import Path是全局唯一的URL，是为了保证依赖管理的一致性和可靠性。而 replace 正是为了在不破坏这个核心原则的前提下，为开发者开的一个“后门”或“绿色通道”，用于处理本地开发、调试、临时fork等特殊场景。*