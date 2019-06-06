#! /bin/bash

#Get and kill running pid:
kill -9 $(lsof -i tcp:5055 -t)

#starts the action server:
#https://rasa.com/docs/core/customactions/#customactions
python -m rasa_core_sdk.endpoint --actions actions &

python -m rasa_core.train interactive \
    -o models/dialogue \
    -d configs/dialog/domain.yml \
    -c configs/dialog/policy.yml \
    -s data/train_dialog/ \
    --nlu models/rasa/ini \
    --endpoints configs/dialog/endpoints.yml
