 构建NLU训练语料和模型
 ======

Homo自然交互系统的自然语言理解核心

<!-- TOC -->

- [1.获取和处理语料](#1获取和处理语料)
    - [1.1 下载原始数据](#11-下载原始数据)
    - [1.2 抽取正文](#12-抽取正文)
    - [1.3 繁体转简体](#13-繁体转简体)
    - [1.4 符号处理](#14-符号处理)
- [2.对语料文件分词](#2对语料文件分词)
- [3.训练MITIE模型](#3训练mitie模型)
- [4.添加示例数据](#4添加示例数据)
    - [4.1 数据格式](#41-数据格式)
    - [4.2 生成json数据](#42-生成json数据)
- [5.训练Rasa NLU模型](#5训练rasa-nlu模型)
    - [5.1 算法解读：训练流程](#51-算法解读训练流程)
- [6.进行意图和实体判断](#6进行意图和实体判断)
    - [6.1 启动HTTP服务](#61-启动http服务)
    - [6.2 发送测试请求](#62-发送测试请求)

<!-- /TOC -->

# 1.获取和处理语料

从[awesome-chinese-nlp](https://github.com/crownpku/awesome-chinese-nlp)获取收集好的语料，这里使用的是[wikipedia dump](https://dumps.wikimedia.org/zhwiki/)

## 1.1 下载原始数据

从上述wikipedia dump下载`.xml.bz2`文件并解压：

```shell
bzip2 -d filename.bz2
```

## 1.2 抽取正文 

这里使用[Wikipedia Extractor](https://github.com/attardi/wikiextractor)抽取维基百科的正文：

```shell
git clone https://github.com/attardi/wikiextractor.git wikiextractor
cd wikiextractor
python setup.py install
./WikiExtractor.py -b 1024M -o extracted zhwiki-latest-pages-articles.xml.bz2
```

其中：

* `-b 1024M`：以1024M为单位切分文件，默认是1M，把参数设置的大一些可以保证最后的抽取结果全部存在一个文件里。这里我们设为1024M，可以分成一个1G的大文件和一个36M的小文件，后续的步骤可以先在小文件上实验，再应用到大文件上。

得到两个文件：wiki_00, wiki_01

## 1.3 繁体转简体 

根据系统安装`opencc`，再进行转换：

```shell
opencc -i wiki_00 -o zh_wiki_00 -c zht2zhs.ini
opencc -i wiki_01 -o zh_wiki_01 -c zht2zhs.ini
```

## 1.4 符号处理

Wikipedia Extractor抽取正文时，会将有特殊标记的外文直接剔除。将「」『』这些符号替换成引号，顺便删除空括号，相关工具脚本在`scripts/convert.py`：

```shell
python scripts/convert.py zh_wiki_00
python scripts/convert.py zh_wiki_01
```

# 2.对语料文件分词

用jieba对语料文件分词：

```shell
python -m jieba -d " " ./zh_wiki_00 > ./zh_wiki_cut00
python -m jieba -d " " ./zh_wiki_01 > ./zh_wiki_cut01
```

# 3.训练MITIE模型

* 1.获取MITIE

  ```shell
  git clone https://github.com/mit-nlp/MITIE.git
  ```

* 2.编译wordrep

  ```shell
  cd MITIE/tools/wordrep
  mkdir build
  cd build
  cmake ..
  cmake --build . --config Release
  ```

* 3.训练

  ```shell
  ./wordrep -e /path/to/your/folder_of_cutted_text_files
  ```

  这里的目录是上一步分词好的文件存放的目录

  这大概需要两到三天时间，需要保证硬盘有100G以上剩余空间

# 4.添加示例数据

## 4.1 数据格式

`nlu/data/rasa/raw_data.txt`为示例数据，格式如下：
```
text,intent
现在几点了|inform_time
几点了|inform_time
你叫什么名字|ask_name
你叫什么|ask_name

text,intent,food
我想吃火锅啊|restaurant_search|火锅
找个吃拉面的店|restaurant_search|拉面
```

每段第一行为接下来的数据格式，段之间用换行符隔开

`text` 为 `文本`，`intent` 为 `文本` 对应的意图

`food` 意为实体的分类，对应 `文本` 中的实体名称

如 "我想吃火锅啊" 的文本 `text` 为 `我想吃火锅啊` ；对应的意图 `intent` 为 `restaurant_search` ；实体 `火锅` 所属分类为 `food`

## 4.2 生成json数据

运行脚本`scripts/trainsfer_raw_to_rasa.py`将原文本转换为rasa需要的json格式，生成的文件在`nlu/data/train_nlu/train_file.json`：

```shell
python -m scripts.trainsfer_raw_to_rasa
```

# 5.训练Rasa NLU模型

```shell
python -m rasa_nlu.train \
       -c configs/rasa/config_jieba_mitie_sklearn.yml \
       --data data/rasa/train_file_new.json \
       --project rasa \
       --fixed_model_name ini \
       --path models
```

* `-c`：配置文件，包括使用的MITIE模型和生成的模型文件
* `--data`：用来做意图识别和实体识别的训练数据的示例数据
* `--project`：生成的模型上层目录名
* `--fixed_model_name`：生成的模型文件名
* `--path`：生成的模型文件存放的路径

## 5.1 算法解读：训练流程

* 1.初始化MITIE
* 2.用jieba分词
* 3.mite+synonyms进行实体识别
* 4.为意图识别做特征提取
* 5.用sklearn对意图识别分类

# 6.进行意图和实体判断

## 6.1 启动HTTP服务

运行命令：

```bash
# 进入创建好的python环境
source env3.6/bin/activate

# 启动HTTP服务
python -m rasa_nlu.server \
       -c configs/config_jieba_mitie_sklearn.yml \
       --path models
```

* `-c`：加载配置文件 `config_jieba_mitie_sklearn.yml`

* `--path`：加载 `models` 中的模型文件

或直接运行脚本：

```bash
./nlu_server.sh
```

## 6.2 发送测试请求
使用命令行工具`curl`进行测试：

```shell
curl --request POST \
  --url http://localhost:5000/parse \
  --data '{"q": "你好","project": "rasa","model": "ini"}'
```

得到的响应：

```json
{
  "intent": {
    "name": "greet",
    "confidence": 0.7295111471652841
  },
  "entities": [],
  "intent_ranking": [
    {
      "name": "greet",
      "confidence": 0.7295111471652841
    },
    {
      "name": "affirm",
      "confidence": 0.05498027264874326
    },
    {
      "name": "goodbye",
      "confidence": 0.0386203516283928
    },
    {
      "name": "request_search",
      "confidence": 0.0351016353164412
    },
    {
      "name": "thanks",
      "confidence": 0.03148191953326129
    },
    {
      "name": "inform_other_phone",
      "confidence": 0.026396659309389947
    },
    {
      "name": "deny",
      "confidence": 0.025463685159462516
    },
    {
      "name": "ask_name",
      "confidence": 0.016934272738612297
    },
    {
      "name": "inform_package",
      "confidence": 0.008298538094301765
    },
    {
      "name": "restaurant_search",
      "confidence": 0.007565581251946566
    }
  ],
  "text": "你好",
  "project": "rasa",
  "model": "ini"
}
```