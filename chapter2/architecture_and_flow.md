# 第二章代码架构设计与运行流程详解

## 一、程序概述

本章实现了一个基于RSS的新闻搜索程序，可以从多个新闻网站获取RSS feed并搜索指定关键词。程序充分展示了Go语言的并发特性，使用goroutines和channels来并行处理多个RSS源，从而提高搜索效率。

## 二、架构设计

### 1. 整体架构

程序采用模块化设计，主要分为以下几个模块：

1. **main包** - 程序入口点
2. **search包** - 核心搜索逻辑和并发控制
3. **matchers包** - 不同类型feed的匹配器实现
4. **data目录** - 配置文件存储

### 2. 核心组件

#### 2.1 main.go
程序入口文件，负责初始化日志输出并启动搜索流程。

#### 2.2 search包
核心搜索模块，包含以下文件：
- [feed.go](./sample/search/feed.go) - Feed数据结构定义和获取
- [match.go](.sample/search/match.go) - 匹配结果结构和接口定义
- [search.go](.sample/search/search.go) - 搜索流程控制和并发管理
- [default.go](.sample/search/default.go) - 默认匹配器实现

#### 2.3 matchers包
匹配器实现模块，目前包含：
- [rss.go](.sample/matchers/rss.go) - RSS类型feed的匹配器实现

#### 2.4 data目录
配置文件目录，包含[data.json](.sample/data/data.json)文件，定义了需要搜索的RSS源列表。

### 3. 设计模式

#### 3.1 插件化设计
通过定义[Matcher](.sample/search/match.go#L16-L18)接口，程序支持不同类型feed的匹配器，可以轻松扩展支持新的feed类型。

#### 3.2 并发处理
使用goroutines和channels实现并发搜索，每个RSS源由一个独立的goroutine处理。

#### 3.3 生产者-消费者模式
搜索结果通过channel传递，实现了生产者-消费者模式。

## 三、运行流程

### 1. 程序启动
1. 执行[main.go](.sample/main.go)中的[main()](.sample/main.go#L18-L21)函数
2. 初始化日志输出到stdout
3. 调用[search.Run("president")](.sample/search/search.go#L20-L66)开始搜索

### 2. 加载配置
1. [Run()](.sample/search/search.go#L20-L66)函数调用[RetrieveFeeds()](.sample/search/feed.go#L17-L35)从[data.json](.sample/data/data.json)加载RSS源列表

### 3. 并发搜索
1. 为每个feed创建一个goroutine
2. 根据feed类型选择合适的匹配器（目前只实现RSS类型）
3. 每个匹配器goroutine执行以下操作：
   - 通过HTTP获取RSS feed内容
   - 解析XML格式的RSS数据
   - 在标题和描述中搜索关键词
   - 将匹配结果发送到结果channel

### 4. 结果处理
1. 使用[WaitGroup](.sample/search/search.go#L27-L27)等待所有搜索goroutine完成
2. 所有搜索完成后关闭结果channel
3. [Display()](.sample/search/match.go#L35-L42)函数从channel读取并显示结果

### 5. 程序结束
结果显示完毕后程序退出

## 四、并发模型

程序使用了Go语言的典型并发模型：

1. **Goroutines** - 为每个RSS源启动一个goroutine进行并行处理
2. **Channels** - 用于在goroutines之间传递搜索结果
3. **WaitGroup** - 等待所有goroutines完成执行
4. **互斥锁** - 在匹配器注册时保证线程安全

## 五、数据流向

```
main() 
  ↓
search.Run()
  ↓
RetrieveFeeds() ← data.json
  ↓
为每个feed启动goroutine
  ↓
匹配器搜索 → HTTP获取RSS → XML解析 → 关键词匹配
  ↓
结果发送到channel
  ↓
Display()函数从channel读取并显示结果
```

## 六、扩展性设计

1. **Matcher接口** - 可以轻松添加新的feed类型匹配器
2. **插件化注册** - 通过import _方式自动注册匹配器
3. **配置文件** - 通过JSON文件管理RSS源列表，便于修改和扩展