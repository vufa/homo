Homo
======== 

一个高性能，易于扩展且完全开源的自然交互系统

除语音合成和语音识别调用在线API，其他模块均基于完全开源的库实现，所有训练过程和运行过程均采用完全开源的工具集和数据集，不依赖任何在线服务

# 功能

* 离线唤醒
  * 基于开源轻量级语音识别引擎[PocketSphinx](https://github.com/cmusphinx/pocketsphinx)实现
  * 使用开源工具集[CMUCLMTK](http://www.speech.cs.cmu.edu/SLM/toolkit_documentation.html)用于语言模型训练
  * 包含全部的训练数据集，工具集参数配置和预训练模型
* 语音识别(待定)
* (可选)本地离线语音识别：
  * 采用开源语音识别工具集[Kaldi](https://github.com/kaldi-asr/kaldi)
  * 使用清华大学30小时的数据集`thchs30`配置及训练
  * 使用AI SHELL公司开源178小时中文语音语料置及训练
  * 使用cvte预训练模型
* 语音合成：调用百度在线API
* 自然语言理解
  * 基于开源自然语言理解框架[Rasa NLU](https://github.com/RasaHQ/rasa)实现
  * 采用完全开源免费的公共数据集作为语料进行训练
  * 采用开源信息提取工具集[MITIE](https://github.com/mit-nlp/MITIE)构建用于Rasa NLU进行实体识别和意图识别的模型
  * 意图识别分类采用开源机器学习框架[scikit-learn](https://github.com/scikit-learn/scikit-learn)
  * 中文分词采用开源分词组件[jieba](https://github.com/fxsjy/jieba)
* 情感分析
  * 基于支持向量机(SVM)算法进行情感极性分析
  * word2vec模型构建采用开源主题建模工具[Gensim](https://github.com/RaRe-Technologies/gensim)
  * (可选)基于逻辑回归(Logistic Regression)算法的情感极性分类器实现

# 配置运行

## 1. 安装依赖

### 1.1 系统依赖

录制声音依赖`portaudio`

Archlinux需要安装`pulseaudio-alsa`使得和`pulseaudio`一起工作时不会造成崩溃

### 1.2 Python依赖

情感分析部分主要由Python实现，需要安装用到的依赖库，推荐使用`virtualenv`

为了和自然语言理解模块兼容，Python版本推荐使用`3.6.8`

```bash
cd homo/sentiment
# 创建一个python3.6的环境
# 需要已经安装有pyhton3.6
virtualenv --python=python3.6 env3.6
# 进入创建的环境
source env3.6/bin/activate
# 安装依赖
pip install -r requirements.txt
# 使用国内镜像
# pip install -i https://pypi.tuna.tsinghua.edu.cn/simple -r requirements.txt
```

## 1.3 使用pyenv

安装`pyenv`：

```bash
git clone https://github.com/pyenv/pyenv.git ~/.pyenv

# 安装 pyenv-virtualenv
git clone https://github.com/pyenv/pyenv-virtualenv.git $(pyenv root)/plugins/pyenv-virtualenv
```

写入shell配置文件`.zshrc`和`.bashrc`：

```bash
#pyenv
export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/bin:$PATH"
if command -v pyenv 1>/dev/null 2>&1; then
  eval "$(pyenv init -)"
fi
#pyenv-virtualenv
eval "$(pyenv virtualenv-init -)"
```

列出所有可安装的python：

```bash
pyenv install -l
```

列出所有已经安装的python：

```bash
pyenv versions
```

安装python：

```bash
# 安装python 3.6.8
pyenv install 3.6.8
# 卸载
# pyenv uninstall 3.6.8
```

创建环境：

```bash
pyenv virtualenv 3.6.8 env3.6.8
```

## 2. 编译

编译`homo-webview`：

```bash
make webview
```

生成的文件为：`homo-webview`

## 3. 运行

## 3.1 运行自然语言理解引擎

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

## 3.2 运行文本情感分析引擎

进入`sentiment`文件夹，`source`对应的python虚拟环境并启动http服务器：

```bash
cd sentiment
source env3.6/bin/activate
python server.py
```

或直接运行脚本：

```bash
cd sentiment
./server.sh 
```

**注意：加载word2vec模型需要花费5~7分钟时间**

## 3.3 运行主程序

```bash
./homo-webview
```

调试模式下日志会打印具体的代码文件和函数信息，webview界面也能调用控制台查看html和js/css

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
* `sentiment`：文本情感分析子系统，用到的数据集，模型构建和模型加载，用Python实现
* `sphinx`：离线唤醒模块，包括数据集及模块构建
* `docs`：项目详细设计文档

# 开发文档

## 1. 离线唤醒实现

参见文档：[wake_up.md](docs/sphinx/wake_up.md)

## 2. 自然语音理解实现

参见文档：

## 3. 文本情感分析实现

用语音识别得到的文本进行情感分析

参见文档：[sentiment.md](docs/sentiment/sentiment.md)