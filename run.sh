#! /bin/bash

HelpApp()
{
	echo " Usage:"
	echo " -p/--pull      Get/Update image from docker hub by default, using '-p a' or '-p ali' to get from Alibaba Cloud"
	echo " -a/--aliyun    Launch Alibaba Cloud Docker image (launch docker hub image by default)"
	echo " -d/--debug     For debugging, use bash in Docker container, debug docker hub image by default, '-d a' or '-d ali' for debugging Alibaba Cloud image"
	echo " -h/--help      show help"
}

PullDocker()
{
    if [ "$1" = "a" -o "$1" = "ali" ]; then
        docker pull registry.cn-hangzhou.aliyuncs.com/codist/homo:latest
    else
        docker pull countstarlight/homo:latest
    fi
}

RunDocker()
{
    DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
    docker run --name=homo --rm $2 \
           -v "$DIR"/conf:/home/homo/homo/conf \
           -v "$DIR"/sphinx/en-us:/home/homo/homo/sphinx/en-us \
           -v "$DIR"/sphinx/cmusphinx-zh-cn-5.2:/home/homo/homo/sphinx/cmusphinx-zh-cn-5.2 \
           -v "$DIR"/nlu/models:/home/homo/homo/nlu/models \
           -v "$DIR"/nlu/data/rasa:/home/homo/homo/nlu/data/rasa \
           -v /tmp/.X11-unix:/tmp/.X11-unix \
           --device /dev/snd \
           --device /dev/dri \
           --group-add $(getent group audio | cut -d: -f3) \
           -e DISPLAY=unix"$DISPLAY" $1 $3
}

RunDockerHub()
{
    if [[ "$(docker images -q countstarlight/homo:latest 2> /dev/null)" == "" ]]; then
    echo -e "\033[33m提示: 没有在本地找到镜像 countstarlight/homo:latest 开始从docker hub 获取(使用 '-a' 从阿里云拉取并运行镜像)\033[0m"
    PullDocker
    fi
    xhost +SI:localuser:$(id -un)
    RunDocker countstarlight/homo:latest
}

RunDockerAli()
{
    if [[ "$(docker images -q registry.cn-hangzhou.aliyuncs.com/codist/homo:latest 2> /dev/null)" == "" ]]; then
        echo -e "\033[33m提示: 没有在本地找到镜像 registry.cn-hangzhou.aliyuncs.com/codist/homo:latest 开始从阿里云获取\033[0m"
        PullDocker "a"
    fi
    xhost +SI:localuser:$(id -un)
    RunDocker registry.cn-hangzhou.aliyuncs.com/codist/homo:latest "" ""
}

DebugDocker()
{
    if [ "$1" = "a" -o "$1" = "ali" ]; then
        RunDocker registry.cn-hangzhou.aliyuncs.com/codist/homo:latest "-itd" "/bin/bash"
    else
        RunDocker countstarlight/homo:latest "-itd" "/bin/bash"
    fi
    docker exec -it homo /bin/bash
}

if [ -z "$1" ]; then
	RunDockerHub
	exit 0
fi

case $1 in
	"-p" | "--pull")
		PullDocker $2
		;;
	"-a" | "--aliyun")
		RunDockerAli
		;;
    "-d" | "--debug")
		DebugDocker $2
		;;
	"-h" | "--help")
		HelpApp
		;;
	*)
		echo "无效的选项: $1"
		echo "使用 -h|--help 来获取帮助"
		exit 1
		;;
esac
exit 0
