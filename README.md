Homo
======== 

English | [简体中文](README_CN.md)

An open source natural interaction system based on offline wake-up, natural language understanding and sentiment analysis

<p align="center">
  <a href="https://travis-ci.org/countstarlight/homo">
    <img src="https://travis-ci.org/countstarlight/homo.svg?branch=master" alt="Build Status">
  </a>
  <a href="https://hub.docker.com/r/countstarlight/homo">
    <img src="https://img.shields.io/microbadger/layers/countstarlight/homo.svg" alt="Docker Layers">
  </a>
  <a href="https://hub.docker.com/r/countstarlight/homo">
    <img src="https://img.shields.io/microbadger/image-size/countstarlight/homo.svg" alt="Docker Image Size">
  </a>
  <a href="https://hub.docker.com/r/countstarlight/homo">
    <img src="https://img.shields.io/docker/pulls/countstarlight/homo.svg" alt="Docker Pulls">
  </a>
  <a href="https://goreportcard.com/report/github.com/countstarlight/homo">
    <img src="https://goreportcard.com/badge/github.com/countstarlight/homo" alt="Go Report">
  </a>
  <a href="https://github.com/countstarlight/homo/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg?style=flat" alt="MIT License">
  </a>
</p>

Demo Video(Chinese): [BiliBili](https://www.bilibili.com/video/av54654613)

**Notice:** A version under reconstruction is located at [dev branch](https://github.com/countstarlight/homo/tree/dev) and named [Aiicy](https://aiicy.org/). Aiicy is designed for IoT and User Terminal, and will support IoT devices on different platforms, allowing users to interact with it through browsers. Aiicy and documents are in active development stages, so stay tuned.

**Features**

* Offline Keyword Research
  * Based on open source lightweight speech recognition engine [PocketSphinx](https://github.com/cmusphinx/pocketsphinx)
  * Offline language model training using the open source toolset [CMUCLMTK](http://www.speech.cs.cmu.edu/SLM/toolkit_documentation.html)
* Online speech recognition
  * Using Baidu Online Speech Recognition API
* Online Text-to-Speech
  * Using Baidu Online Text-to-Speech API
* Natural Language Understanding
  * Based on open source natural language understanding framework [Rasa NLU](https://github.com/RasaHQ/rasa)
  * Using open source information extraction toolset [MITIE](https://github.com/mit-nlp/MITIE) to build models for entity recognition and intent recognition in Rasa NLU
  * Using open source machine learning framework [scikit-learn](https://github.com/scikit-learn/scikit-learn) to do Intent recognition classification
  * Using open source word segmentation component [jieba](https://github.com/fxsjy/jieba) to do Chinese word segmentation
* Text Sentiment Analysis
  * Sentiment Analysis Using Support Vector Machine(SVM)
  * Using open source topic modelling tool [Gensim](https://github.com/RaRe-Technologies/gensim) to build word2vec model
  * (Optional)Sentiment Analysi based on Logistic Regression Classification

**Contents**

<!-- TOC -->

- [Quick start(Linux)](#quick-startlinux)
- [Road map](#road-map)
- [Contributing](#contributing)
- [License](#license)

<!-- /TOC -->

# Quick start(Linux)

Get source code with git:

```bash
git clone https://github.com/countstarlight/homo.git
```

Download the dataset for Homo refer document(Chinese): [https://homo.codist.me/docs/dataset/](https://homo.codist.me/docs/dataset/)

Make sure Docker is installed then run(`run.sh` needs root privileges if the current user is not in the `docker` group):

```bash
cd homo
cp conf/example_app.ini conf/app.ini
./run.sh
```

This will download and launch the image from the docker hub by default, or use the image built by Alibaba Cloud:

```bash
./run.sh -a
```

`run.sh` supported commands:

```bash
$ ./run.sh -h
 Usage:
 -p/--pull      Get/Update image from docker hub by default, using '-p a' or '-p ali' to get from Alibaba Cloud
 -a/--aliyun    Launch Alibaba Cloud Docker image (launch docker hub image by default)
 -d/--debug     For debugging, use bash in Docker container, debug docker hub image by default, '-d a' or '-d ali' for debugging Alibaba Cloud image
 -h/--help      show help
```

# Road map

- [ ] Plug-in system
    - [ ] Custom actions

- [ ] Improve documentation
    - [x] Custom wake word
    - [x] Expanding Natural Language Understanding
    - [ ] Custom actions

- [ ] Support for English
    - [x] Offline Keyword Research
    - [ ] Speech Recognition
    - [x] Text-to-Speech
    - [ ] Documentation

# Contributing

Questions and suggestions are welcomed through [issues](https://github.com/countstarlight/homo/issues), or changes submitted to the project through [Pull Requests](https://github.com/countstarlight/homo/pulls)

# License

[MIT](https://github.com/countstarlight/homo/blob/master/LICENSE)

Copyright (c) 2019-present Codist
