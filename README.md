Homo
======== 

<!-- TOC -->

- [安装](#安装)
    - [1.安装依赖](#1安装依赖)
    - [2.编译](#2编译)
    - [3.运行](#3运行)
        - [运行`webview`](#运行webview)
- [项目结构](#项目结构)
- [开发文档](#开发文档)
    - [1.离线唤醒实现](#1离线唤醒实现)
    - [2.文本情感分析实现](#2文本情感分析实现)

<!-- /TOC -->

# 安装

## 1.安装依赖

录制声音依赖`portaudio`

archlinux需要安装`pulseaudio-alsa`使得和`pulseaudio`一起工作时不会造成崩溃

## 2.编译

编译`homo-webview`：

```bash
make webview
```

生成的文件为：`homo-webview`

## 3.运行

### 运行`webview`

调试模式：

```bash
./homo-webview -d
```

调试模式下日志会打印具体的代码文件和函数信息，webview界面也能调用控制台查看html和js/css

# 项目结构

* `cmd`：用户交互部分，Golang实现
  * `interact`：控制台UI实现(已停止维护)
  * `webview`：webview UI实现
* `module`：主体架构各模块，Golang实现
  * `audio`：底层音频硬件交互
  * `baidu`：baidu在线API交互
  * `nlu`：自然语言理解引擎交互
  * `sphinx`：语音识别引擎`sphinx`交互
  * `com`：通用模块
* `sentiment`：文本情感分析模块，包括数据集及模型构建，Python实现
* `sphinx`：离线唤醒模块，包括数据集及模块构建，Shell脚本
* `docs`：项目详细设计文档

# 开发文档

## 1.离线唤醒实现

参见文档：[wake_up.md](docs/sphinx/wake_up.md)

## 2.文本情感分析实现

用语音识别得到的文本进行情感分析

参见文档：[sentiment.md](docs/sentiment/sentiment.md)