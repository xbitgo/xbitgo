# XbitGo

    为了解决之前碰到的各类痛点而开发的go极简框架    

## 特点

    强大的代码生成工具
    微服务优先 支持grpc和http
    极简 规范 极少依赖 
    领域模型 实践优先
    丰富的功能组件
    高度可定制化
    服务治理集成
    单元测试友好

## 包含三个核心仓库
  
  - [core核心工具](https://github.com/xbitgo/components)
  - [components组件库](https://github.com/xbitgo/components)
  - [cli命令行工具](https://github.com/xbitgo/cli)

## 安装使用(go1.16+)

- 安装 protoc


    根据各系统情况自行安装       
            

- 安装 protoc-gen-go-grpc,
  protoc-gen-gofast,
  protoc-go-inject-tag


    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/gogo/protobuf/protoc-gen-gofast@latest
    go install github.com/favadi/protoc-go-inject-tag@latest


- 安装 xbit

  
    go install github.com/xbitgo/cli@latest


## 创建项目

    执行 xbit init {projectName} 

## 创建应用

    cd {projectPath}
    执行 xbit create {appName}

## 启动应用

    cd {projectPath}
    xbit protoc
    go mod tidy
    cd apps/{appName}
    go run main.go

## 更详细文档完善中...