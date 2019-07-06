#! /bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

HelpApp()
{
	echo " 附加命令:"
	echo " -p/--pull      从阿里云获取镜像(默认从docker hub获取)"
	echo " -a/--aliyun    运行阿里云Docker镜像(默认运行docker hub镜像)"
	echo " -h/--help      显示此帮助信息"
}

PullDocker()
{
    docker pull countstarlight/homo:latest
}

PullDockerAli()
{
    docker pull registry.cn-hangzhou.aliyuncs.com/codist/homo:latest
}

RunDocker()
{
    if [[ "$(docker images -q countstarlight/homo:latest 2> /dev/null)" == "" ]]; then
    echo -e "\033[33m提示: 没有在本地找到镜像 countstarlight/homo:latest 开始从docker hub 获取(使用 '-a' 从阿里云拉取并运行镜像)\033[0m"
    PullDocker
    fi
    xhost +SI:localuser:$(id -un)
    docker run --name=homo --rm \
           -v "$DIR"/conf:/home/homo/homo/conf \
           -v "$DIR"/sphinx/en-us:/home/homo/homo/sphinx/en-us \
           -v "$DIR"/sphinx/cmusphinx-zh-cn-5.2:/home/homo/homo/sphinx/cmusphinx-zh-cn-5.2 \
           -v "$DIR"/nlu/models:/home/homo/homo/nlu/models \
           -v "$DIR"/nlu/data/rasa:/home/homo/homo/nlu/data/rasa \
           -v /tmp/.X11-unix:/tmp/.X11-unix \
           --device /dev/snd \
           --device /dev/dri \
           --group-add $(getent group audio | cut -d: -f3) \
           -e DISPLAY=unix"$DISPLAY" \
           countstarlight/homo:latest
}

RunDockerAli()
{
    if [[ "$(docker images -q registry.cn-hangzhou.aliyuncs.com/codist/homo:latest 2> /dev/null)" == "" ]]; then
    echo -e "\033[33m提示: 没有在本地找到镜像 registry.cn-hangzhou.aliyuncs.com/codist/homo:latest 开始从阿里云获取\033[0m"
    PullDockerAli
    fi
    xhost +SI:localuser:$(id -un)
    docker run --name=homo --rm \
           -v "$DIR"/conf:/home/homo/homo/conf \
           -v "$DIR"/sphinx/en-us:/home/homo/homo/sphinx/en-us \
           -v "$DIR"/sphinx/cmusphinx-zh-cn-5.2:/home/homo/homo/sphinx/cmusphinx-zh-cn-5.2 \
           -v "$DIR"/nlu/models:/home/homo/homo/nlu/models \
           -v "$DIR"/nlu/data/rasa:/home/homo/homo/nlu/data/rasa \
           -v /tmp/.X11-unix:/tmp/.X11-unix \
           --device /dev/snd \
           --device /dev/dri \
           --group-add $(getent group audio | cut -d: -f3) \
           -e DISPLAY=unix"$DISPLAY" \
           registry.cn-hangzhou.aliyuncs.com/codist/homo:latest
}

if [ -z "$1" ]; then
	RunDocker
	exit 0
fi

case $1 in
	"-p" | "--pull")
		PullDockerAli
		;;
	"-a" | "--aliyun")
		RunDockerAli
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
