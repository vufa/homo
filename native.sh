#! /bin/bash

CheckProcess()
{
    if [ "$1" = "" ];
    then
        return 1
    fi

    PROCESS_NUM=`ps -ef | grep "$1" | grep -v "grep" | wc -l`
    if [ $PROCESS_NUM -eq 0 ];
    then
        return 1
    else
        return 0
    fi
}

KillProcess()
{
    ps -ef | grep "$1" | grep -v grep | awk '{print $2}' | xargs kill
}

StartMaster()
{
    CheckProcess "homo-master"
    Check_RET=$?
    if [ $Check_RET -eq 0 ]; then
        echo -e "\033[33native: Program 'homo-master' already running\033[0m"
        exit 0
    fi
    ./homo-master -c conf/example_master.yml start &
}

StopMaster()
{
    CheckProcess "homo-master"
    Check_RET=$?
    if [ $Check_RET -eq 0 ]; then
        KillProcess "homo-master"
    fi
    echo -e "[\033[32m ok \033[0m] kill program \"homo-master\""
}

StartHub()
{
    CheckProcess "homo-hub"
    Check_RET=$?
    if [ $Check_RET -eq 0 ]; then
        echo -e "\033[33native: Program 'homo-hub' already running\033[0m"
        exit 0
    fi
    ./homo-hub -c conf/example_hub.yml &
}

StopHub()
{
    CheckProcess "homo-hub"
    Check_RET=$?
    if [ $Check_RET -eq 0 ]; then
        KillProcess "homo-hub"
    fi
    echo -e "[\033[32m ok \033[0m] kill program \"homo-hub\""
}

StartFunction()
{
    SetEnv
    CheckProcess "homo-function"
    Check_RET=$?
    if [ $Check_RET -eq 0 ]; then
        echo -e "\033[33native: Program 'homo-function' already running\033[0m"
        exit 0
    fi
    ./homo-function -c conf/example_function.yml &
}

StopFunction()
{
    CheckProcess "homo-function"
    Check_RET=$?
    if [ $Check_RET -eq 0 ]; then
        KillProcess "homo-function"
    fi
    echo -e "[\033[32m ok \033[0m] kill program \"homo-function\""
}

StartAll()
{
    StartMaster
    StartHub
    StartFunction
}

StopAll()
{
    StopMaster
    StopHub
    StopFunction
}

SetEnv()
{
    export HOMO_SERVICE_MODE="native"
    export HOMO_MASTER_API_ADDRESS="unix://var/run/homo.sock"
    export HOMO_API_ADDRESS="unix://var/run/homo/api.sock"
    export HOMO_MASTER_API_VERSION="v1"
}

HelpApp()
{
    echo " Extra Commands:"
    echo " -s/--start          Start all homo module"
    echo " -m/--master         Start homo-master"
    echo " -b/--hub            Start homo-hub"
    echo " -f/--function       Start homo-function"
    echo " -k/--kill           Stop all module of homo"
    echo " -e/--env            Set env"
    echo " -h/--help           Show program help info"
}

if [ -z $1 ]; then
    StartAll
    exit 0
fi
case $1 in
    "-s" | "--start")
        StartAll
    ;;
    "-m" | "--master")
        StartMaster
    ;;
    "-b" | "--hub")
        StartHub
    ;;
    "-f" | "--function")
        StartFunction
    ;;
    "-k" | "--kill")
        StopAll
    ;;
    "-e" | "--env")
        SetEnv
    ;;
    "-h" | "--help")
        HelpApp
    ;;
    *)
        echo -e "\033[31native: unrecognized option '$1' \033[0m"
        echo "Use -h|--help to get help"
        exit 1
    ;;
esac
exit 0