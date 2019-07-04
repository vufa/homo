Homo
======== 

一个基于离线唤醒，自然语言理解和情感分析的开源自然交互系统

<p align="center">
  <a href="https://travis-ci.org/countstarlight/homo">
    <img src="https://travis-ci.org/countstarlight/homo.svg?branch=master" alt="Build Status">
  </a>
  <a href="https://hub.docker.com/r/countstarlight/homo">
    <img src="https://img.shields.io/microbadger/layers/countstarlight/homo.svg" alt="MicroBadger Layers">
  </a>
  <a href="https://goreportcard.com/report/github.com/countstarlight/homo">
    <img src="https://goreportcard.com/badge/github.com/countstarlight/homo" alt="Go Report">
  </a>
  <a href="https://github.com/countstarlight/homo/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat" alt="MIT License">
  </a>
</p>

**演示视频(BiliBili)：**
https://www.bilibili.com/video/av54654613

[![https://www.bilibili.com/video/av54654613](screenshot/screenshot.jpg)](https://www.bilibili.com/video/av54654613)

**功能**

* 离线唤醒
  * 基于开源轻量级语音识别引擎[PocketSphinx](https://github.com/cmusphinx/pocketsphinx)实现
  * 使用开源工具集[CMUCLMTK](http://www.speech.cs.cmu.edu/SLM/toolkit_documentation.html)进行离线语言模型训练
* 在线语音识别
  * 调用百度在线语音识别API
* 语音合成：
  * 调用百度在线语音合成API
* 自然语言理解
  * 基于开源自然语言理解框架[Rasa NLU](https://github.com/RasaHQ/rasa)实现
  * 采用开源信息提取工具集[MITIE](https://github.com/mit-nlp/MITIE)构建用于Rasa NLU进行实体识别和意图识别的模型
  * 意图识别分类采用开源机器学习框架[scikit-learn](https://github.com/scikit-learn/scikit-learn)
  * 中文分词采用开源分词组件[jieba](https://github.com/fxsjy/jieba)
* 文本情感分析
  * 基于支持向量机(SVM)算法进行情感极性分析
  * word2vec模型构建采用开源主题建模工具[Gensim](https://github.com/RaRe-Technologies/gensim)
  * (可选)基于逻辑回归(Logistic Regression)算法的情感极性分类器实现

**目录**

<!-- TOC -->

- [安装和配置](#安装和配置)
- [运行](#运行)
    - [1.运行自然语言理解引擎](#1运行自然语言理解引擎)
    - [2.运行主程序](#2运行主程序)
- [使用指南](#使用指南)
    - [1. 意图理解范围](#1-意图理解范围)
- [自定义](#自定义)
    - [1. 自定义唤醒词](#1-自定义唤醒词)
- [文件结构](#文件结构)
- [发展路线](#发展路线)
- [贡献](#贡献)
- [捐赠](#捐赠)
- [协议](#协议)

<!-- /TOC -->

# 安装和配置

安装和配置请阅读文档: [https://homo.codist.me/docs/install/](https://homo.codist.me/docs/install/)

# 运行

## 1.运行自然语言理解引擎

进入`nlu`的文件夹，`source`对应的python虚拟环境并启动http服务器：

```bash
cd nlu
source env3.6/bin/activate

python -m rasa_nlu.server \
       -c configs/rasa/config_jieba_mitie_sklearn.yml \
       --path models
```

或者直接运行脚本`nlu_server.sh`：

```bash
cd nlu
./nlu_server.sh
```

**注意：加载word2vec模型需要花费5~7分钟时间**

## 2.运行主程序

```bash
./homo-webview
```

了解详细启动参数，请阅读文档：[https://homo.codist.me/docs/run/](https://homo.codist.me/docs/run/)

# 使用指南

## 1. 意图理解范围

了解Homo自带的意图理解的范围，请阅读文档：[https://homo.codist.me/docs/intent/](https://homo.codist.me/docs/intent/)

# 自定义

## 1. 自定义唤醒词

自定义唤醒Homo时的唤醒词，请阅读文档：[https://homo.codist.me/docs/wake-up/](https://homo.codist.me/docs/wake-up/)
# 文件结构

* `cmd`：用户交互部分，Golang实现
  * `interact`：控制台UI实现(已停止维护)
  * `webview`：webview UI实现
* `module`：主体架构各模块，Golang实现
  * `audio`：底层音频硬件交互
  * `baidu`：baidu在线语音识别&合成API交互
  * `nlu`：自然语言理解引擎交互
  * `sphinx`：语音识别引擎`sphinx`交互
  * `com`：通用模块
* `sentiment`：文本情感分析引擎，用到的数据集，模型构建和模型加载，用Python实现
* `nlu`：自然语言理解引擎，用到的数据集，模型构建和模型加载，用Python实现
* `sphinx`：离线唤醒模块，包括数据集及模块构建

# 发展路线

- [ ] 插件系统
    - [ ] 扩展行为

- [ ] 多平台支持
    - [ ] Windows
    - [ ] Macos

- [ ] 完善文档
    - [ ] 在多个平台上的编译配置
    - [x] 自定义唤醒词
    - [ ] 扩展自然语言理解
    - [ ] 扩展行为

- [ ] 完善单元测试
    - [ ] 离线唤醒模块

- [ ] Python部分用Go或Rust或C++重写
    - [ ] 替代用到的机器学习库
    - [ ] 文本情感分析部分：SVM...
    - [ ] 自然语言理解部分：MITIE...

- [ ] 添加对英文的支持
    - [ ] 离线唤醒
    - [ ] 语音识别
    - [ ] 语音合成
    - [ ] 文档

# 贡献

欢迎通过 [issues](https://github.com/countstarlight/homo/issues) 提出问题和建议，或通过 [Pull Requests](https://github.com/countstarlight/homo/pulls) 向本项目提交修改

# 捐赠

如果你觉得这个项目对你有帮助，可以请作者喝一杯咖啡，支持作者继续开发

![donate.png](screenshot/donate.png)

# 协议

[MIT](https://github.com/countstarlight/homo/blob/master/LICENSE)

Copyright (c) 2019-present Codist
