#! /bin/bash

source env3.6/bin/activate && \

python -m rasa_nlu.server \
       -c configs/rasa/config_jieba_mitie_sklearn.yml \
       --path models
