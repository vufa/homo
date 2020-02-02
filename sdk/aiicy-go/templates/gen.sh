#!/bin/sh

if [[ $1 == *node* ]]
then
    sed "s/{{.MODULE}}/$1/g" ./templates/Makefile-node > ../../aiicy-$1/Makefile
elif [[ $1 == *python* ]]
then
    sed "s/{{.MODULE}}/$1/g" ./templates/Makefile-python > ../../aiicy-$1/Makefile
else
    sed "s/{{.MODULE}}/$1/g" ./templates/Makefile-go > ../../aiicy-$1/Makefile
    sed "s/{{.MODULE}}/$1/g" ./templates/Dockerfile-go > ../../aiicy-$1/Dockerfile
    sed "s/{{.MODULE}}/$1/g" ./templates/package-go.yml > ../../aiicy-$1/package.yml
fi

cat ./templates/Makefile >> ../../aiicy-$1/Makefile