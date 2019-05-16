Homo
======== 

# 安装

## 安装依赖

录制声音依赖`portaudio`

archlinux需要安装`pulseaudio-alsa`使得和`pulseaudio`一起工作时不会造成崩溃

## 编译

编译`homo-webview`：

```bash
make webview
```

生成的文件为：`homo-webview`

# 运行

## 运行`webview`

调试模式：

```bash
./homo-webview -d
```

调试模式下日志会打印具体的代码文件和函数信息，webview界面也能调用控制台查看html和js/css

# 具体实现

## 离线唤醒实现

参见文档：[wake_up.md](docs/wake_up.md)

## 情感分析实现

### 1.文本情感分析

用语音识别得到的文本进行情感分析：[sentiment.md](docs/sentiment/sentiment.md)