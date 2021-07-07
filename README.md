# Gin-api-frame

gin-api-frame提供一种基于[Gin](https://github.com/gin-gonic/gin)快速开发API服务的框架。其中使用到的日志库为[logrus](https://github.com/sirupsen/logrus)，orm库为[gorm](https://github.com/go-gorm/gorm)，命令行框架为[cobra](https://github.com/spf13/cobra)。

# Framework

整体代码目录结构如下：

```bash
├── build
├── CHANGELOG
├── cmd
├── docs
├── go.mod
├── go.sum
├── hack
├── LICENSE
├── Makefile
├── _output
├── pkg
└── README.md
```

- build: Dockerfile目录，用来构建Docker镜像。
- cmd: main函数入口，包含参数解析和配置文件。
- docs: 文档目录。
- hack: 编译构建脚本，部署文件及脚本。
- pkg: 核心代码逻辑，主要包括handler、controller、model等。
- _output：构建产物存储路径。

核心代码逻辑为`pkg`包，具体目录功能如下：

```bash
pkg
├── apis
├── constant
├── controller
├── errors
├── handler
├── model
├── server
├── types
├── util
├── validation
└── version
```

- constant: 常量包。
- controller：实际的业务控制器逻辑。
- errors: 定义error常量。
- handler: gin框架的handler逻辑。
- model: 数据库增删改查操作逻辑。
- server: gin框架的路由定义逻辑。
- types: 定义项目的结构体类型对象。
- util: 通用的工具包。
- validation: 请求参数的合法性校验。
- version: version构建包。

# Usage

### step1. 下载项目代码，替换项目名称



```
git clone https://github.com/huweihuang/gin-api-frame.git

grep -rl gin-api-frame . | xargs sed -i 's/gin-api-frame/{you-project}/g' 
```

### step2. 定义配置参数，项目结构体参数。

- cmd/server/app/config/config.go: 默认有日志配置、数据库配置、Etcd配置。可自定义增加配置项。
配置文件为toml格式，具体参考[hack/conf/conf.toml]。

- pkg/types: 定义项目所需结构体对象。

### step3. 补充业务逻辑代码及路由。

- model: 定义数据库增删改查逻辑。
- controller：定义业务控制逻辑。
- handler: 定义handler处理逻辑。
- server/router: 定义路由逻辑。
- validation: 定义入参校验逻辑。

### step4. 编译构建镜像及部署。

```bash
# 构建二进制、构建镜像、推送镜像(修改构建镜像仓库)
make

# 只构建二进制
make build

# 修改配置文件
修改hack/deploy/deploy.yaml中的配置文件和镜像地址

# 使用k8s集群部署
kubectl create -f hack/deploy/deploy.yaml
```

# Have fun ^_^

Please <a class="github-button" href="https://github.com/huweihuang/gin-api-frame" data-icon="octicon-star" aria-label="Star huweihuang/hexo-theme-huweihuang on GitHub">Star</a> this Project if you like it! <a class="github-button" href="https://github.com/huweihuang" aria-label="Follow @huweihuang on GitHub">Follow</a> would also be appreciated!
