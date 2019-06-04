部署
=====

## 1.环境

* Archlinux
* miniconda

## 2.配置

### 2.1安装 conda
* 下载 `Miniconda`: https://repo.continuum.io/miniconda/

* 安装 `Miniconda` 根据文档[Document](https://docs.anaconda.com/anaconda/install/)

* 设置`conda`使用中国镜像源：

  ```shell
  conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
  conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/main/
  conda config --set show_channel_urls yes
  ```

* Disable changeps1(可选)

  ```shell
  conda config --set changeps1 False
  ```

* 为`homo-core`创建python环境：

  ```shell
  conda create -n homo-core python=3.6.8	# Create an environment named with python3.6.8
  conda activate homo-core	# 进入环境
  conda deactivate #退出
  conda env remove -n homo-core #删除环境
  ```

### 2.2安装依赖库

```shell
conda install -c anaconda twisted
pip install -r requirements.txt
pip install git+https://github.com/mit-nlp/MITIE.git
```

## 3.启动http服务

```shell
python -m rasa_nlu.server \
       -c configs/rasa/config_jieba_mitie_sklearn.yml \
       --path models
```

* `-c`：使用的配置文件，和训练Rasa NLU模型时使用同样的配置文件
* `--path`：使用的模型所在的目录

