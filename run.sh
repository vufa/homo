#! /bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

docker run -d --name=homo \
           -v "$DIR"/conf:/home/homo/homo/conf \
           -v "$DIR"/sphinx/en-us:/home/homo/homo/sphinx/en-us \
           -v "$DIR"/sphinx/cmusphinx-zh-cn-5.2:/home/homo/homo/sphinx/cmusphinx-zh-cn-5.2 \
           -v "$DIR"/nlu/models:/home/homo/homo/nlu/models \
           -v /tmp/.X11-unix:/tmp/.X11-unix \
           --device /dev/snd \
           -e DISPLAY=unix"$DISPLAY" \
           countstarlight/homo:latest
