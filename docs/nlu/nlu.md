自然语言理解
======

Homo自然交互系统的核心

采用的技术：

* [MITIE](https://github.com/mit-nlp/MITIE)：信息提取模型，用于构建Rasa NLU进行实体识别和意图识别的模型。
* [jieba](https://github.com/fxsjy/jieba)：“结巴”中文分词，用于分词。
* [scikit-learn](https://github.com/scikit-learn/scikit-learn)：机器学习框架，使用它提供的分类方法，对意图识别分类。

## 训练流程

* 1.初始化MITIE
* 2.用jieba分词
* 3.mite+synonyms进行实体识别
* 4.为意图识别做特征提取
* 5.用sklearn对意图识别分类

## 部署和启动`homo-core`

* 1.安装依赖和配置环境参见：[deploy.md](deploy.md)

* 2.获取语料和构建MITIE模型参见：[dataset.md](dataset.md)

* 3.启动http服务参见：[deploy.md](deploy.md)

# 对话或只进行意图和实体判断

## 1.对话

* 创建和标记对话用语料并启动对话参见：[dialog.md](dialog.md)

### 1.1启动流程

* 1.训练得到MITIE模型参见：[dataset.md](dataset.md)(耗时最长，完整大概需要2~3天)

* 2.生成示例数据，参见：[dataset.md](dataset.md)：

  ```shell
  python -m scripts.trainsfer_raw_to_rasa
  ```

* 3.训练得到Rasa NLU模型，参见：[dataset.md](dataset.md)：(短对话也最少需要半个小时)

  ```shell
  ./train_nlu.sh
  ```

* 4.训练对话模型，参见：[dialog.md](dialog.md)：

  ```shell
  ./train_dialog.sh
  ```

* 5.启动对话，参见：[dialog.md](dialog.md)：

  ```shell
  ./dialog.sh
  ```

### 1.2完善对话(标记意图和应该的回答)

参见：[dialog.md](dialog.md)：

```shell
./interaction.sh
```

完善后需要重新训练对话模型

### 1.3完善行为

在`configs/domain.yml`里添加新的意图和行为后，需要添加响应的示例数据在`data/rasa/raw_data.txt`，

* 1.重新生成示例json数据:

  ```shell
  ./gen_dialog.sh
  ```

* 2.重新训练nlu:

  ```shell
  ./train_nlu.sh
  ```

* 3.参照上一节，完善对话:

  ```shell
  ./interaction.sh
  ```

* 4.重新训练对话:

  ```shell
  ./train_dialog.sh
  ```


## 2.只进行意图和实体判断

启动nlu 服务：

```shell
./nlu_server.sh
```

运行在5000端口：

```shell
curl --request POST \
  --url http://localhost:5000/parse \
  --data '{"q": "你好","project": "rasa","model": "ini"}'
```

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



