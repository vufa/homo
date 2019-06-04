 构建语料和模型
 ======

## 1.获取和处理语料

从[awesome-chinese-nlp](https://github.com/crownpku/awesome-chinese-nlp)获取收集好的语料，这里使用的是[wikipedia dump](https://dumps.wikimedia.org/zhwiki/)

### 1.1下载原始数据

从上述wikipedia dump下载`.xml.bz2`文件并解压：

```shell
bzip2 -d filename.bz2
```

### 1.2抽取正文 

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

### 1.3繁体转简体 

根据系统安装`opencc`，再进行转换：

```shell
opencc -i wiki_00 -o zh_wiki_00 -c zht2zhs.ini
opencc -i wiki_01 -o zh_wiki_01 -c zht2zhs.ini
```

### 1.4符号处理

Wikipedia Extractor抽取正文时，会将有特殊标记的外文直接剔除。将「」『』这些符号替换成引号，顺便删除空括号，相关工具脚本在`scripts/convert.py`：

```shell
python scripts/convert.py zh_wiki_00
python scripts/convert.py zh_wiki_01
```

## 2.对语料文件分词

用jieba对语料文件分词：

```shell
python -m jieba -d " " ./zh_wiki_00 > ./zh_wiki_cut00
python -m jieba -d " " ./zh_wiki_01 > ./zh_wiki_cut01
```

## 3.训练MITIE模型

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

## 4.生成示例数据

`data/rasa/raw_data.txt`为原始示例数据，运行脚本`scripts/trainsfer_raw_to_rasa.py`将原始数据转换为rasa需要的json格式，生成的文件在`data/rasa/train_file_new.json`

```shell
python -m scripts.trainsfer_raw_to_rasa
```

## 5.训练Rasa NLU模型

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