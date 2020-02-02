#! /bin/bash

HelpApp()
{
	echo " 附加命令:"
	echo " -p/--pull      获取/更新 镜像，默认从docker hub获取，使用 '-p a' 或 '-p ali' 从阿里云获取"
	echo " -a/--aliyun    运行阿里云Docker镜像(默认运行docker hub镜像)"
	echo " -d/--debug     用于调试，使用Docker容器内的bash，默认调试docker hub镜像，使用 '-d a' 或 '-d ali' 调试阿里云镜像"
	echo " -h/--help      显示此帮助信息"
}

PullDocker()
{
    if [ "$1" = "a" -o "$1" = "ali" ]; then
        docker pull registry.cn-hangzhou.aliyuncs.com/codist/aiicy:latest
    else
        docker pull countstarlight/aiicy:latest
    fi
}

RunDocker()
{
    DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
    docker run --name=aiicy --rm $2 \
           -v "$DIR"/conf:/home/aiicy/aiicy/conf \
           -v "$DIR"/sphinx/en-us:/home/aiicy/aiicy/sphinx/en-us \
           -v "$DIR"/sphinx/cmusphinx-zh-cn-5.2:/home/aiicy/aiicy/sphinx/cmusphinx-zh-cn-5.2 \
           -v "$DIR"/nlu/models:/home/aiicy/aiicy/nlu/models \
           -v "$DIR"/nlu/data/rasa:/home/aiicy/aiicy/nlu/data/rasa \
           -v /tmp/.X11-unix:/tmp/.X11-unix \
           --device /dev/snd \
           --device /dev/dri \
           --group-add $(getent group audio | cut -d: -f3) \
           -e DISPLAY=unix"$DISPLAY" $1 $3
}

RunDockerHub()
{
    if [[ "$(docker images -q countstarlight/aiicy:latest 2> /dev/null)" == "" ]]; then
    echo -e "\033[33m提示: 没有在本地找到镜像 countstarlight/aiicy:latest 开始从docker hub 获取(使用 '-a' 从阿里云拉取并运行镜像)\033[0m"
    PullDocker
    fi
    xhost +SI:localuser:$(id -un)
    RunDocker countstarlight/aiicy:latest
}

RunDockerAli()
{
    if [[ "$(docker images -q registry.cn-hangzhou.aliyuncs.com/codist/aiicy:latest 2> /dev/null)" == "" ]]; then
        echo -e "\033[33m提示: 没有在本地找到镜像 registry.cn-hangzhou.aliyuncs.com/codist/aiicy:latest 开始从阿里云获取\033[0m"
        PullDocker "a"
    fi
    xhost +SI:localuser:$(id -un)
    RunDocker registry.cn-hangzhou.aliyuncs.com/codist/aiicy:latest "" ""
}

DebugDocker()
{
    if [ "$1" = "a" -o "$1" = "ali" ]; then
        RunDocker registry.cn-hangzhou.aliyuncs.com/codist/aiicy:latest "-itd" "/bin/bash"
    else
        RunDocker countstarlight/aiicy:latest "-itd" "/bin/bash"
    fi
    docker exec -it aiicy /bin/bash
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
