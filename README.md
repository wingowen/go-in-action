# 中文RSS搜索项目

## 项目概述

本项目基于Go语言实现了一个中文RSS搜索案例，并提供了React前端界面。项目主要功能是从国内新闻RSS源中搜索指定关键词的新闻内容，并以友好的方式展示结果。

## 项目建立思考

1. **需求分析**：将英文搜索案例修改为中文搜索案例，替换为国内RSS源，并创建前端界面展示搜索结果
2. **技术选型**：
   - 后端：Go语言（保持原项目语言）
   - 前端：React + Vite（现代前端技术栈，快速构建）
3. **实施策略**：先完成后端本地化修改，再构建前端界面

## 环境配置

### Windows 系统环境变量

- GOROOT：Go 安装路径（默认 C:\Go，安装程序自动设置）
- GOPATH：工作区路径（建议自定义，如 D:\go_workspace）
- PATH：添加 %GOROOT%\bin 和 %GOPATH%\bin

### 配置代理

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go env | findstr "GOPROXY"
```

## 项目建立过程

### 1. 后端修改

#### 1.1 替换RSS源
将原项目中的英文RSS源替换为国内中文新闻源：
替换的国内RSS源包括：腾讯新闻、网易新闻、新浪新闻、澎湃新闻、界面新闻、凤凰新闻和36氪。

#### 1.2 修改搜索关键词
将main.go中的搜索关键词从英文"president"修改为中文"中国"，并翻译相关注释。

#### 1.3 本地化日志和错误信息
修改以下文件中的英文日志和错误信息为中文：
- rss.go：搜索日志和HTTP错误信息
- search.go：匹配器注册日志
- match.go：结果显示格式

### 2. 前端创建

#### 2.1 创建Vite+React项目
```bash
# 在项目根目录下执行
npm create vite@latest frontend -- --template react
```

#### 2.2 安装依赖
```bash
# 设置npm国内源
npm config set registry https://registry.npmmirror.com

# 安装依赖
cd frontend
npm install
```

#### 2.3 实现搜索界面
修改App.jsx文件，实现包含搜索框、搜索按钮和结果展示区域的中文RSS搜索界面。

#### 2.4 美化样式
修改App.css文件，为搜索界面提供现代、简洁的样式。

### 3. 项目运行

#### 3.1 后端运行
```bash
# 进入sample目录
cd chapter2/sample

# 运行后端程序
go run main.go
```

#### 3.2 前端运行
```bash
# 进入frontend目录
cd frontend

# 运行前端开发服务器
npm run dev
```

## 遇到的问题及解决方案

1. **npm源问题**：
   - 问题：淘宝npm镜像证书过期
   - 解决方案：切换到阿里云npm镜像 `https://registry.npmmirror.com`

2. **vite命令找不到问题**：
   - 问题：全局vite命令不可用
   - 解决方案：使用项目内安装的vite `node_modules/.bin/vite`

## 项目结构

```
go-in-action/
├── chapter2/
│   └── sample/
│       ├── data/
│       │   └── data.json  # RSS源配置
│       ├── main.go        # 程序入口
│       ├── matchers/
│       │   └── rss.go     # RSS匹配器实现
│       └── search/
│           ├── match.go   # 匹配器接口和结果显示
│           └── search.go  # 搜索逻辑实现
├── frontend/
│   ├── index.html         # 前端入口HTML
│   ├── package.json       # 前端依赖配置
│   ├── src/
│   │   ├── App.css        # 搜索界面样式
│   │   ├── App.jsx        # 搜索界面前端实现
│   │   └── main.jsx       # 前端入口文件
│   └── vite.config.js     # Vite配置
└── README.md              # 项目说明文档
```

## 使用说明

1. 确保已安装Go和Node.js
2. 运行后端程序：`cd chapter2/sample && go run main.go`
3. 运行前端服务器：`cd frontend && npm run dev`
4. 在浏览器中打开前端页面（默认地址：http://localhost:5173）
5. 在搜索框中输入关键词（如"中国"），点击搜索按钮查看结果

## 后续改进方向

1. 实现真实后端API接口，替换模拟数据
2. 添加搜索结果过滤和排序功能
3. 优化前端界面响应式设计
4. 添加用户认证和个性化设置
5. 实现搜索历史记录功能