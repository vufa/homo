基于PocketSphinx的唤醒系统
========

# 1. 安装

## 1.1 安装sphinxbase

下载源码：

```bash
git clone https://github.com/cmusphinx/sphinxbase.git
```

编译安装：

```bash
cd sphinxbase
./autogen.sh
./configure --prefix=/usr
make
#check
make check
sudo make install
```

## 1.2 安装PocketSphinx

下载源码：

```bash
git clone https://github.com/cmusphinx/pocketsphinx.git
```

编译安装：

```bash
cd pocketsphinx
./autogen.sh
./configure --prefix=/usr
#make clean all
make check
```

## 1.3 测试语音识别

使用golang绑定：

```bash
go get github.com/xlab/pocketsphinx-go/sphinx
```

获取示例：

```bash
go get github.com/xlab/pocketsphinx-go/example/gortana
```

### 1.3.1 英文语音识别

```bash
gortana --hmm "pocketsphinx/model/en-us/en-us" \
        --dict "pocketsphinx/model/en-us/cmudict-en-us.dict" \
        --lm "pocketsphinx/model/en-us/en-us.lm.bin"
```

使用`pocketsphinx_continuous`：

```bash
pocketsphinx_continuous -lm 2916.lm -dict 2916.dic
```

### 1.3.2 中文语音识别

下载中文语音模型

https://sourceforge.net/projects/cmusphinx/files/Acoustic%20and%20Language%20Models/Mandarin/

解压后测试：

```bash
gortana --hmm "cmusphinx-zh-cn-5.2/zh_cn.cd_cont_5000" \
        --dict "cmusphinx-zh-cn-5.2/zh_cn.dic" \
        --lm "cmusphinx-zh-cn-5.2/zh_cn.lm.bin"
```

要播放录制的音频：

```bash
ffplay -f s16le -ar 16000 -ac 1 -i 000000000.raw
```

# 2. 中文唤醒词识别

## 2.1 创建语料库

语料库实际上就是一些文本的集合，包含了你需要识别的语音的文字的一些集合，例如句子啊，词啊等等。

编辑一个自定义的`keyword.txt`文本，里面写入打算唤醒的中文词语，和发音可能混淆的词（如果拼音相同只记录一个就行）。

再添加一些其他的乱七八糟的词，这样匹配的时候就不会一直匹配唤醒词了（唤醒词的重点），以小贝为例，则`keyword.txt`中的内容如下：

```
小贝
小魏
巧倍
啊
呵呵
哈哈
么么哒
```

## 2.2 利用在线工具LMTool建立语言模型

在[http://www.speech.cs.cmu.edu/tools/lmtool-new.html] 上面训练上一步的keyword文本，会生成字典文件`* .dic` 和语言模型文件 `*.lm`，下载这两个文件就可以。用来替代语言模型和拼音字典。

```
如：
1234.lm
1234.dic
```

编辑下载的`随机数.dic`文件，对照着`zh_broadcastnews_utf8.dic`的拼音字典，更改成与其同样格式的内容。原字典中不一定会有相同的词语，有的话，就按照原先的写，没有的话，就按照单个发音的写上就可以：

```
小贝 x i ao b ei
小魏 x i ao w ei
巧倍 q i ao b ei
啊 a as
.
.
.
```

在代码中，替换掉对应的lm和dic路径。

```python
import os
from pocketsphinx import LiveSpeech, get_model_path

model_path = get_model_path()

speech = LiveSpeech(
    verbose=False,
    sampling_rate=16000,
    buffer_size=2048,
    no_search=False,
    full_utt=False,
    hmm=os.path.join(model_path, 'zh/zh_broadcastnews_16k_ptm256_8000'),
    lm=os.path.join(model_path, 'zh/1234.lm'),  # 这个目录位置自己设置
    dic=os.path.join(model_path, 'zh/1234.dic')  # 同上
)
for phrase in speech:
    print("phrase:", phrase)
    print(phrase.segments(detailed=True))
    # 只要命中上述关键词的内容，都算对
    if str(phrase) in ["小贝", "小魏", "巧倍"]:
    print("正确识别唤醒词")
```

## 2.3 本地语言模型生成

如果语言模型较小（例如小的语音指令集或者任务），而且是英文的，那就可以直接上CMU提供的网络服务器上面训练，如果较大的话，一般使用`CMUCLMTK`语言模型工具来训练。

上面用在线工具生成语言模型，这里尝试本地生成，需要安装语言模型训练工具`CMUCLMTK`，语言模型训练工具的说明见：

http://www.speech.cs.cmu.edu/SLM/toolkit_documentation.html

### 2.3.1 安装CMUCLMTK

获取源码：

```bash
svn co https://svn.code.sf.net/p/cmusphinx/code/trunk/cmuclmtk/
```

编译安装：

```bash
cd cmuclmtk/
./autogen.sh
make
sudo make install
```

### 2.3.2 准备语料库

创建文件`homo.txt`，输入如下内容，记住结尾不可留“\n”(实验证明了这一点)。每个utterances由 `<s>` 和 `</s>`来分隔：

```bash
<s> 天气 </s>

<s> 有雨 </s>

<s> 晴朗 </s>

<s> 多云 </s>

<s> 雷电 </s>
```

### 2.3.3 生成词汇表(vocabulary)文件

```bash
text2wfreq < homo.txt | wfreq2vocab > homo.vocab
```

* `text2wfreq`：统计文本文件中每个词出现的次数，例如：

```bash
text2wfreq < homo.txt
```

输出：

```
text2wfreq : Reading text from standard input...
雷电 1
<s> 5
天气 1
多云 1
晴朗 1
有雨 1
</s> 5
text2wfreq : Done.
```

表示词`雷电`、`<s>`和`</s>`在训练文本中出现的次数依次为1、5、5。

* `wfreq2vocab`：统计文本文件中含有多少个词，即有哪些词。如数字识别中包含10个数字和两个静音，故共有12个词，按照拼音顺序依次是：`</s>`、`<s>`、`八`、`二`、`九`、`零`、`六`、`七`、`三`、`四`、`五`、`一`。

生成文件`homo.vocab`

### 2.3.4 生成 arpa格式的语言模型

```bash
text2idngram -vocab homo.vocab -idngram homo.idngram < homo.txt
idngram2lm -vocab_type 0 -idngram homo.idngram -vocab homo.vocab -arpa homo.arpa
```

* `text2idngram`：列举文本中出现的每一个n元语法。产生一个二进制文件，含有一个n元数组的数值排序列表，对应于与词有关的的N-Gram。超出文本范围的词汇映射值为0

  生成文件`homo.idngram`

* `idngram2lm`：输入文件包括一个`.idngram`文件，一个`.vocab`文件和一个ccs文件，输出是一个后缀为`.arpa`的语言模型文件，其中ccs文件指句首和句尾的静音`<s>`和`</s>`

  生成文件`homo.arpa`

### 2.3.5 转换为 CMU的二进制格式 (DMP)

如果你的语言模型比较大的话，最好就转换为CMU的二进制格式 (DMP)，这样可以加快加载语言模型的速度，减少解码器初始化的时间。但对于小模型来说，就没有这个必要，因为sphinx3能处理这两种后缀名的语言模型文件。

```bash
sphinx_lm_convert -i homo.arpa -o homo.lm.bin
```

生成了语言模型文件`homo.lm.bin`，此文件为解码器端所需要的文件格式。

### 2.3.6 创建对应的字典文件

和2.2一样，创建字典文件`homo.dic`：

```
多云 d uo1 vv vn2
天气 t ian1 q i4
晴朗 q ing2 l ang3
有雨 ii iu3 vv v3
雷电 l ei2 d ian4
```

测试语音识别：

```bash
gortana --hmm "cmusphinx-zh-cn-5.2/zh_cn.cd_cont_5000" \
        --dict "homo.dic" \
        --lm "homo.lm.bin"
```

## 2.4 语音唤醒优化

假设给机器人起名字：“小贝”

确定一个唤醒词，最好用四字词，如 小贝同学

然后用手机语音识别软件测试，如语音识别输入法，对着手机说：“小贝”和“小贝同学”，看手机上识别出的结果，可能的结果有：“小为”、“交杯”、“小类”、“小贝”“小白” 和 “小贝同学”、“小飞同学”、“小薇同学”、“这位同学”等

记录这些词到`keyword.txt`文件中，重新生成语言模型

重新运行`gortana`，在嘈杂的环境中测试，统计识别出最多的结果，绝大部分都是二字词，而该系统只有识别到含有名字 “小贝” 或者 “小贝**” 的结果才能唤醒

假设嘈杂环境识别到最多的是“小为”，把“小为”写入`keyword.txt`，对应调整`dic`文件

