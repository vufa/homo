#! /bin/bash

source env3.6/bin/activate &&
    python -m rasa_nlu.train \
        -c configs/config_jieba_mitie_sklearn.yml \
        --data data/train_nlu/ \
        --project rasa \
        --fixed_model_name ini \
        --path models
