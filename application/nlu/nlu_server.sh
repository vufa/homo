#! /bin/bash

python -m rasa_nlu.server \
       -c configs/config_jieba_mitie_sklearn.yml \
       --path models
