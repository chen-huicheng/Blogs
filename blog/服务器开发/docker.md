# Docker简介与基本操作

## 容器技术的起源

我们要开发一个APP，程序员从头到尾搭建一套环境并编写代码，代码写完了交给测试人员，测试人员**从头到尾搭建这套环境，**然后对程序进行测试，测试完成后，终于要上线了，这时运维同学又要重新从**头到尾搭建这套环境，**上线后发现系统崩溃了？

这时候程序员就会有一个大大的疑惑，明明在我电脑上可以正常运行。

从这个例子可以看到一下三点：

1.  应用程序的运行依赖特定的系统环境。
2.  开发、测试、上线是在不同的机器或实例上运行。
3.  保存程序在各个机器/实例上运行一致，需要保证应用程序运行环境的一致性。

这个时候大家可能会想到虚拟机，先搭好一套虚拟机环境，然后给测试和运维clone出来不就可以了吗？

在没有容器技术之前，这确实是一个好办法，只不过这个办法还没有那么好。

主要包含一下内容

-   虚拟机怎么解决环境隔离问题，以及优缺点
-   容器怎么解决环境隔离问题
-   容器技术的一些基础概念和整体结构
-   容器技术实现原理
-   容器技术的应用场景
-   使用容器技术实现应用容器化的简单步骤

### 虚拟机

虚拟机通过软件模拟了一个操作系统，而操作系统是一个很复杂、很笨重的应用程序，同时这个虚拟的操作系统本身也会占用大量的资源。

假设有一台机器，16G内存，需要部署三个应用，那么使用虚拟机技术可以这样划分：

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=N2JiMzRkYzU2ZTkyMDcxYmY3ODI1MmFhZjZlYzUyMzhfa1RIUDJIczc0MVJHVUw0VVRwRzE2NmE1eThtdmtMdVRfVG9rZW46Ym94Y25GdkIzVWVwa3p2SmtGMkFkWmNRNFliXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

在这台机器上开启三个虚拟机，每个虚拟机上部署一个应用，其中应用分别占用1G、2G和4G内存，同时虚拟机本身却占用了4G、3G和2G内存。

这里有一个明显的问题：虚拟机技术的特性导致了虚拟环境本身会占用大量的系统资源。

同时还有另一个问题，那就是启动时间，虚拟机技术是虚拟了一个操作系统，启动一个虚拟机相当于启动了一个操作系统，而操作系统启动的时间是非常慢的。

有没有技术可以让避免把内存浪费在“无用”的操作系统上，同时提升应用的启动时间那就太好了？

答案就是容器技术：

### 容器

容器，就是一种轻量级的虚拟化技术，目的和虚拟机一样，都是为了创造“隔离环境”。但是它不像VM采用操作系统级的资源隔离，容器采用的是进程级的系统隔离。

容器本质上就是一个 Linux 进程。

Docker容器具有以下三大特点：

-   轻量化：一台主机上运行的多个Docker容器可以共享主机操作系统内核；启动迅速，只需占用很少的计算和内存资源。
-   标准开放：Docker容器基于开放式标准，能够在所有主流Linux版本、Microsoft Windows以及包括VM、裸机服务器和云在内的任何基础设施上运行。
-   安全可靠：Docker赋予应用的隔离性不仅限于彼此隔离，还独立于底层的基础设施。Docker默认提供最强的隔离，因此应用出现问题，也只是单个容器的问题，而不会波及到整台主机。

 作为一种新兴的虚拟化方式，`Docker` 跟传统的虚拟化方式相比具有众多的优势。

### 容器VS虚拟机

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=MTVhMjZmM2M0NDljZDFhZmQ3YmY0NzE1OWFhMDcyYTVfdU1oSHZMQ2NycVJnajVPeFo4SWJLdlRqUDBuRmdtNWFfVG9rZW46Ym94Y253cHZnTmIyaUY0ZVFPR3BqVVd2a2hlXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

| **特性**   | **容器**           | **虚拟机**  |
| ---------- | ------------------ | ----------- |
| 启动       | 秒级               | 分钟级      |
| 硬盘使用   | 一般为 `MB`        | 一般为 `GB` |
| 性能       | 接近原生           | 弱于        |
| 系统支持量 | 单机支持上千个容器 | 一般几十个  |

这里我们大概知道容器比虚拟机要好，启动快，占用资源少。容器为什么好，它是通过什么技术或方法实现的呢。

## 容器的基本原理

### Docker 基础概念

**镜像**：docker镜像是一个特殊的文件系统，除了提供容器运行时所需的程序、库、资源、配置等文件外，还包含了一些为运行时准备的一些配置参数（如匿名卷、环境变量、用户等）。镜像 **不包含** 任何动态数据，其内容在构建之后也不会被改变。

**容器**：**镜像运行时的实体。容器的实质是进程，但与直接在宿主执行的进程不同，容器进程运行于属于自己的独立的****命名空间****。**因此容器可以拥有自己的 `root` 文件系统、网络配置、进程空间，甚至自己的用户 ID 空间。容器内的进程是运行在一个隔离的环境里，使用起来，就好像是在一个独立于宿主的系统下操作一样。这种特性使得容器封装的应用比直接在宿主运行更加安全。

**镜像仓库**：存放**docker**镜像的仓库。仓库会包含同一个软件不同版本的镜像，而标签就常用于对应该软件的各个版本。我们可以通过 `<仓库名>:<标签>` 的格式来指定具体是这个软件哪个版本的镜像。

**镜像（**`Image`**）和容器（**`Container`**）的关系，就像是面向对象程序设计中的** **`类`** **和** **`实例`** **一样，镜像是静态的定义，容器是镜像运行时的实体。**

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=MDNkMmMzNjA2M2Y0ZTYxZDYzYmZkMzkxNjc3N2MyZWJfTzRiQll5U2Q0dnM5cTgzbjBOU3FRQk9vTjBCem8wOUpfVG9rZW46Ym94Y25WT251N3lmWnlxQm93enRBcXI2clFnXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

### Docker引擎

运行一个容器的基本过程，docker client docker server 的交互。

Docker容器的运行逻辑如下图所示，Docker使用客户端/服务器 (C/S) 架构模式，Docker守护进程（Docker daemon）作为Server端接收Docker客户端的请求，并负责创建、运行和分发Docker容器。Docker守护进程一般在Docker主机后台运行，用户使用Docker客户端直接跟Docker守护进程进行信息交互。

三者的关系大致如下图：

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=OWM5ZjA3ODhiMzM5ZjZhNjYzMTM5MzU1NzNjYTRiZTNfY0FXbVFMeTkyM2JzWnlIbnQ4TTFTRjRpZDVSWks0RDNfVG9rZW46Ym94Y240Yjllc2tEbXptTG1ZSkloWnI4NmJlXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

客户端和服务端既可以运行在一个机器上，也可通过 `socket` 或者 `RESTful API` 来进行通信。

1.  构建镜像

-   如橙色流程所示，执行Docker构建指令会根据Docker文件构建一个镜像存放于本地Docker主机。

1.  拉取镜像

-   如蓝色流程所示，执行Docker拉取指令会从云端镜像仓库拉取镜像至本地Docker主机。

1.  运行镜像

-   如橙色流程所示，执行Docker启动指令会将镜像安装至容器并启动容器。

## Docker镜像

OCI（Open Container Initiative）规范是事实上的容器标准，已经被大部分容器实现以及容器编排系统所采用，包括 Docker 和 Kubernetes。它的出现是一段关于开源商业化的有趣历史：它由 Dokcer 公司作为领头者在 2015 年推出，但如今 Docker 公司在容器行业中已经成了打工仔。

规范要求镜像内容必须包括以下 3 部分：

-   **Image Manifest**：提供了镜像的配置和文件系统层定位信息，可以看作是镜像的目录，文件格式为 `json` 。
-   **Image Layer Filesystem Changeset**：序列化之后的文件系统和文件系统变更，它们可按序一层层应用为一个容器的 rootfs，因此通常也被称为一个 `layer`（与下文提到的镜像层同义），文件格式可以是 `tar` ，`gzip` 等存档或压缩格式。
-   **Image Configuration**：包含了镜像在运行时所使用的执行参数以及有序的 rootfs 变更信息，文件类型为 `json`。

镜像是一个多层结构。如果指令向镜像中增添、修改文件或程序，则会新建一个镜像层，添加配置信息不会新建镜像层。

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=ZjJkMTUxMDVjNzg3YzFlNGMxZmY4YzRiZTVjMWU1YmFfcjVadEsxQ2NsQXRrSkdQblF0cXczV003RTdNUG82bXhfVG9rZW46Ym94Y25ONUFjVG5PWDZnQ3lZdGF5NlJvMWNjXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

容器引擎还启动**容器运行时：**

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=MjJlYmE4MDg1OWRlMjU1ZGNiN2YyZjA5NWYzMDBiZDRfcnlXQ0dDZ2sxQVJxelg2ajRoeUVQekI0bE9OMWZMMVRfVG9rZW46Ym94Y242bURMSkVBZFhDQUR5V0dNN1U5NFVkXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

**docker** **容器镜像包含以下两部分内容**：

-   **Rootfs（容器根文件系统）：**系统上的一个目录，看起来像操作系统的标准根 (`/`)。例如，带有`/usr`, `/var`,`/home`等的目录。（镜像由多个松耦合镜像层组成）
-   **JSON文件（容器配置）：**指定如何运行rootfs；例如：容器启动时在 rootfs 中运行什么**命令**或**入口点；** 为容器设置的**环境变量；**容器的 **工作目录**；和其他一些设置。

runC 是一个轻量级、可移植的容器运行时。它包括 Docker 用来与容器相关的系统功能进行交互的所有管道代码。

## Docker容器

容器本质上就是一个 Linux 进程。

特殊之处在于容器是在一个特殊的隔离环境中运行的进程。

隔离环境是由Linux Namespace、 Linux Cgroups、rootfs三种技术构建出来的。

Linux Namespace 实现视图的隔离，Linux Cgroups 实现资源的隔离与限制，rootfs实现文件系统的隔离。

### 内核命名空间 Linux Namespace

内核命名空间为容器进程在其中运行提供了一个虚拟化世界，使得容器化进程与运行在同一 Linux 主机上的其他进程隔离。对内核资源进行分区，以便一组进程看到一组资源，而另一组进程看到另一组资源。

例如，“PID”命名空间使容器化进程只能看到该容器内的其他进程，但看不到共享主机上其他容器的进程。

容器内

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=YmZhMmI1NGE4YjkyNjBlZTNlM2RjNTA1ZTc4Y2M0MzBfRjR3d1BkOXdUOUttVHR3NUFYa3NuVVo3Y2xCOGRwQkdfVG9rZW46Ym94Y25NTWZQSm5qUmlSbmFBcUtTYmV3Q0pjXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

宿主机

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=ODljMDBmYjdiNjVkMGZkYjQxMzRkYjQ1MDI0ZjUwMzdfNG1RdWxWQnBSSzBnS3VseFBwdXJWMTVrQ0FsUmdEQWhfVG9rZW46Ym94Y24zRHFCZ2NKNTdReUlVS2wybGx5UkhkXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=NWUzMjIxMDVkOGM4NzFjMzUxN2U3NTRlN2FkMzk2ODJfWlhZYUxPVHdWQ2ZZS1NhVVY0VmRkU1Q0WUlDOE1LSDdfVG9rZW46Ym94Y243Wmh1Zm1teTk5a0tIcVRkTFRhcEx4XzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

https://tinylab.org/pid-namespace/

核命名空间为主要是一下数据进行隔离，保证部分数据只对这些进程可见。

8种不同类型的命名空间https://man7.org/linux/man-pages/man7/namespaces.7.html

1.  cgroups — 隔离根目录
2.  IPC — 隔离进程间通信
3.  Network — 隔离网络堆栈
4.  Mount — 隔离挂载点
5.  PID — 隔离进程 ID
6.  Time — 隔离进程间时间
7.  User — 隔离用户和组 ID
8.  UTS — 隔离主机名和域名

运行一个容器时，将创建一些新的 namespace，`init` 进程将被加入到这些 namespace；在一个容器中运行一个新进程时，新进程将加入创建容器时所创建的 namespace。

### 控制组 Linux Cgroups（Control Groups）

Linux Cgroups实现资源的限制，每个容器进程消耗的资源（内存、cpu、I/O 等）被限制在指定的范围内。

比如：生成一个组(生成后此组为空，里面没有进程)，设置其CPU使用率为10%，并把一个进程丢进这个组中，那么这个进程最多只能使用CPU的10%，如果将多个进程丢进这个组，这个组的所有进程平分这个10%。

当一个容器被创建时，将为每种类型的资源创建一个新的 cgroup，在容器中运行的所有进程都将加入到这些 cgroup 中。通过控制容器中运行的所有进程，cgroups 实现了对容器的资源限制。

这种隔离容器化进程并限制它们消耗的资源的能力使多个应用程序容器能够在共享的 Linux 主机上更安全地运行。隔离和资源限制的结合使 Linux 进程成为 Linux 容器。换句话说，容器就是 Linux。

如果您将容器定义为具有资源约束、Linux 安全约束和命名空间的进程，根据定义，Linux 系统上的每个进程都在一个容器中。这就是为什么我们常说[Linux就是容器，容器就是Linux](https://www.redhat.com/en/blog/containers-are-linux)。**容器运行时**是修改这些资源约束、安全性和命名空间并启动容器的工具。

**容器就是Linux，Linux就是容器**

### 联合文件系统

从前面描述可知容器镜像由多个镜像层组成，如何将多个镜像层应用成一个文件系统呢？

联合文件系统（Union File System）也叫 UnionFS，主要的功能是将多个不同位置的目录联合挂载（union mount）到同一个目录下。

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=N2I4MzJiMTQ4NzVlNDBhMGVkNmZjNmY4Mjk3MzA0MjRfOFRhR2NVbHdsRHVPRkszcmF5N29OV0t6N0FueWFlWXFfVG9rZW46Ym94Y241RTZ0TEphdVF5M1NFQkp3bHR4WUtDXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

提供文件系统的隔离

## 容器的应用场景

### 快速部署

同一个镜像部署在多个平台

**快速部署交付** 

持续集成和持续部署 (CI/CD)，持续集成 (CI) 和持续部署 (CD)是现代运维的基础，开发人员和运维人员持续沟通，CI/CD 创造了一种实时反馈机制，持续地传输小型迭代更改，加速更改，提高质量。

**代码流水线**（Code Pipeline）管理

代码从开发者的机器到最终在生产环境上的部署，需要经过很多的中间环境。而每一个中间环境都有自己微小的差别，Docker给应用提供了一个从开发到上线均一致的环境，让代码的流水线变得简单不少。

### 环境隔离

**简化配置**

这是Docker公司宣传的Docker的主要使用场景。虚拟机的最大好处是能在你的硬件设施上运行各种配置不一样的平台（软件、系统），Docker在降低额外开销的情况下提供了同样的功能。它能让你将运行环境和配置放在代码中然后部署，同一个Docker的配置可以在不同的环境中使用，这样就降低了硬件要求和应用环境之间耦合度。

**提高开发效率**

不同的开发环境中，想把一件事做好。一是想让开发环境尽量贴近生产环境，二是快速搭建开发环境。

### 多应用

**隔离应用**

有很多种原因会让你选择在一个机器上运行不同的应用。

**整合服务器**

正如通过虚拟机来整合多个应用，Docker隔离应用的能力使得Docker可以整合多个服务器以降低成本。由于没有多个操作系统的内存占用，以及能在多个实例之间共享没有使用的内存，Docker可以比虚拟机提供更好的服务器整合解决方案。

**多租户环境**

另外一个Docker有意思的使用场景是在多租户的应用中，它可以避免关键应用的重写。我们一个特别的关于这个场景的例子是为物联网的应用开发一个快速、易用的多租户环境。这种多租户的基本代码非常复杂，很难处理，重新规划这样一个应用不但消耗时间，也浪费金钱。

## 应用的容器化

将应用整合到容器中并运行起来的过程，称为“容器化”，也可叫做“Docker化”。

容器化的基本过程如下：

1.  编写应用代码。
2.  创建一个 Dockerfile，其中包括当前应用的描述、依赖以及如何运行这个应用。
3.  对该 Dockerfile 执行 `docker image buile` 命令。
4.  推送镜像到镜像仓库。
5.  拉取镜像并运行。

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=NWY4M2NiNzNkYTU3YjQzM2MzMGJhNjYyOTFmMmFkZDhfb2FRNzljSmpPZFVRS3d0R1ZHclJMbHJ6Uk1Ccmo1UGZfVG9rZW46Ym94Y25peE1KbVU0R2R3Nm1GVlhaRjJMekxiXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

Step1：编写应用代码 main.go

```Go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)
func myHandler(w http.ResponseWriter, r *http.Request) {
    url := r.URL.String()
    log.Printf("host:%s\n url:%s", r.Host, url)
    fmt.Fprintln(w, url)
}
func initLog(file string) {
    fp, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        panic("open log file err")
    }
    log.SetOutput(fp)
}
func main() {
    initLog("http.log")
    http.HandleFunc("/", myHandler)
    fmt.Println("server begin")
    err := http.ListenAndServe("0.0.0.0:80", nil)
    if err != nil {
        fmt.Println(err)
    }
}
```

Step2：创建一个 Dockerfile

```Dockerfile
FROM golang:1.18.3-alpine as builder   # FROM 指定基础镜像
# FROM debian:10 as builder
COPY . /home   # COPY 将从构建上下文目录中的文件/目录  复制 到的镜像内的位置
WORKDIR /home  # WORKDIR 指定工作目录
RUN go build -o http main.go  # RUN 执行命令

FROM golang:1.18.3-alpine
# FROM debian:10
WORKDIR /home
COPY --from=builder /home/http /home/http
EXPOSE 80
CMD ["./http"]
```

Step3：构建镜像，执行 `docker image buile` 命令

```PowerShell
docker build -t http:1.0.0 .
```

`-t` 参数是给镜像命名并打标签。

Step4：推送镜像到镜像仓库。

```SQL
docker login //输入用户名密码登录到  dockerhub
docker tag http:1.0.0 chenhuicheng/http:1.0.0
docker push chenhuicheng/http:1.0.0
```

推送到dockerhub后的镜像：https://hub.docker.com/r/chenhuicheng/http/tags

Step5：运行镜像

```PowerShell
docker run -d -p 2345:80 chenhuicheng/http:1.0.0
```

-d：守护态运行。

-p：指定端口映射，支持的格式有 `ip:hostPort:containerPort | ip::containerPort | hostPort:containerPort`。将容器的端口映射到主机端口。

访问容器内服务：`ip:hostPort`

![img](https://bytedance.feishu.cn/space/api/box/stream/download/asynccode/?code=ZGYzYmM4OTcyNTdhMDQ4MDViMzk5NDYwYWIxZTVmN2RfY2lwYVN2NTZEdUlCa1Y3T2xkOVN4T0hGejFMTFljZ1VfVG9rZW46Ym94Y25MTW1CQVlHdFpnbjJGZkU5WVU0RHhlXzE2NzUwNjQ4NjU6MTY3NTA2ODQ2NV9WNA)

## 参考文献：

[Docker 入门到实践](https://yeasy.gitbook.io/docker_practice/)

[如何构建并发布自己的 docker 镜像](https://jerrymei.cn/docker-build-or-commit-images/)

[什么是Docker容器？](https://info.support.huawei.com/info-finder/encyclopedia/zh/Docker容器.html)

https://www.redhat.com/en/blog/containers-are-linux 容器是 Linux

https://opensource.com/article/18/8/sysadmins-guide-containers  系统管理员的容器指南

https://medium.com/@Mark.io/beginners-guide-to-runc-1b29cf281752 runc

https://jessicagreben.medium.com/what-is-the-difference-between-a-process-a-container-and-a-vm-f36ba0f8a8f7  process container vm

https://waynerv.com/posts/container-fundamentals-learn-container-with-oci-spec/ 容器技术原理(一)：从根本上认识容器镜像

https://www.zhihu.com/question/22969309 docker 应用场景

