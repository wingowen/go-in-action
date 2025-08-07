# 前后端交互实现指南

## 项目结构
```
chapter2/
├── frontend/           # 前端React应用
│   ├── src/
│   │   ├── App.jsx     # 主组件
│   │   └── ...
│   └── ...
├── sample/
│   ├── api/            # 后端API实现
│   │   ├── main.go     # API服务器入口
│   │   └── server.go   # API处理逻辑
│   ├── search/         # 搜索功能包
│   └── ...
└── ...
```

## 实现的功能
1. 后端API服务器，提供搜索接口
2. 前端React应用，调用后端API获取搜索结果
3. 完整的前后端数据交互流程

## 如何运行

### 启动后端API服务器
1. 打开终端，进入到API目录：
   ```
   cd /Users/wingo.wen/Documents/code/go-in-action/chapter2/sample/api
   ```
2. 运行API服务器：
   ```
   go run main.go
   ```
3. 服务器将启动在 http://localhost:8080

### 启动前端开发服务器
1. 打开另一个终端，进入到前端目录：
   ```
   cd /Users/wingo.wen/Documents/code/go-in-action/chapter2/frontend
   ```
2. 运行前端开发服务器：
   ```
   npm run dev
   ```
3. 根据终端提示，在浏览器中打开前端应用

## 测试前后端交互
1. 确保后端API服务器和前端开发服务器都已启动
2. 在前端应用的搜索框中输入关键词（如"中国"）
3. 点击搜索按钮，前端将调用后端API获取搜索结果
4. 查看搜索结果是否正确显示

## 注意事项
1. 确保后端服务器先启动，以便前端能够正确调用API
2. 如果遇到跨域问题，可以在Vite配置中添加代理设置
3. 确保Go环境已正确配置，能够运行Go程序
4. 前端依赖需要先安装：在frontend目录下运行`npm install`

## 可能的改进
1. 添加更多的错误处理
2. 实现分页功能
3. 添加加载状态和动画
4. 优化搜索结果展示样式
5. 添加单元测试和集成测试