对话
=====

`configs/dialog/domain.yml`：定义对话所有的意图、槽、实体和系统可能采取的action

`data/stories.md`：对话训练语料， 开始可以用于预训练对话模型，后期加入在线学习过程中生成的语料，更新完善。

一个story的例子：

```markdown
## Generated Story -3503963698851273627
* greet
    - utter_greet
* inform_time{"item": "?"}
    - slot{"item": "?"}
    - action_get_time

## Generated Story -7639399201402861128
* greet
    - utter_greet
* inform_time
    - action_get_time
```

具体story的格式要求参考[官方文档](https://core.rasa.ai/stories.html) ，下面通过初始对话语料，测试强化学习（online learning），并生成对话语料。

## 训练对话模型

```shell
python -m rasa_core.train \
       -d configs/dialog/domain.yml \
       -s data/dialog/stories.md \
       -o models/dialogue \
       -c configs/dialog/policy.yml
```

## 完善对话(标记意图和应该的回答)

```shell
python -m rasa_core_sdk.endpoint --actions actions&
python -m rasa_core.train interactive \
       -o models/dialogue \
       -d configs/dialog/domain.yml \
       -c configs/dialog/policy.yml \
       -s data/dialog/stories.md \
       --nlu models/rasa/ini \
       --endpoints configs/dialog/endpoints.yml
```

调用了rasa core 的 Python接口，具体代码可以在项目代码中看到。从输出中可以看到其使用的kerse 网络结构(单层的LSTM)：

```
_________________________________________________________________
Layer (type)                 Output Shape              Param #   
=================================================================
masking_1 (Masking)          (None, 2, 31)             0         
_________________________________________________________________
lstm_1 (LSTM)                (None, 32)                8192      
_________________________________________________________________
dense_1 (Dense)              (None, 11)                363       
_________________________________________________________________
activation_1 (Activation)    (None, 11)                0         
=================================================================
Total params: 8,555
Trainable params: 8,555
Non-trainable params: 0
```

教对话的过程中，按照提示0，将对话输出并保存，替换原有对话语料文件，再次训练对话模型

## 启动对话

```shell
python -m rasa_core_sdk.endpoint --actions actions&
python -m rasa_core.run \
       -d models/dialogue \
       -u models/rasa/ini \
       --endpoints configs/dialog/endpoints.yml \
       -o logs/rasa_core.log
```

