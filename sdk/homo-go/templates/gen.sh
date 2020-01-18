#!/bin/sh

if [[ $1 == *node* ]]
then
    sed "s/{{.MODULE}}/$1/g" ./templates/Makefile-node > ../../homo-$1/Makefile
elif [[ $1 == *python* ]]
then
    sed "s/{{.MODULE}}/$1/g" ./templates/Makefile-python > ../../homo-$1/Makefile
else
    sed "s/{{.MODULE}}/$1/g" ./templates/Makefile-go > ../../homo-$1/Makefile
    sed "s/{{.MODULE}}/$1/g" ./templates/Dockerfile-go > ../../homo-$1/Dockerfile
    sed "s/{{.MODULE}}/$1/g" ./templates/package-go.yml > ../../homo-$1/package.yml
fi

cat ./templates/Makefile >> ../../homo-$1/Makefile