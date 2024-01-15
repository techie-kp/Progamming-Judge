#!/bin/bash

function cleanup() {
    if [-f "$1-code-output.txt"]; then
        rm $1-code-output.txt
    fi
    if [-f "$1-diff-messages.txt"]; then
        rm $1-diff-messages.txt
    fi
}

# Input format
# $1 -> id of submission
# $2 -> extension of code file
# $3 -> bind mounted directory
# (relative to user home)
#
# The corresponding source, input and output
# should be placed in the "bind_mnt_dir"
# directory with the following naming convention:
#
# source file = {id} + "-main." + {extension}
#
# input file = {id} + "-input.txt"
#
# output file = {id} + "-output.txt"

a=1
flag=0
while [ -e "$3/$1-input/input-$a.txt" ]
do
    touch $1-code-output.txt

    # Execute and trap output
    timeout $4 pypy3 $3/$1-main.$2 < $3/$1-input/input-$a.txt &> $1-code-output.txt 

    res=$?

    if [ $res -eq 124 ]; then 
        echo "time limit exceeded on test $a"
        
        cleanup $1
        flag=1
        exit
    elif [ $res -eq 137 ]; then
        echo "memory limit exceeded on test $a"
        cleanup $1
        flag=1
        exit
    elif [ $res != 0 ]; then
        echo "run failed on test $a", $res
        cleanup $1
        flag=1
        exit
    fi

    # Check if output matches
    diff --strip-trailing-cr $1-code-output.txt $3/$1-output/output-$a.txt > $1-diff-messages.txt
    if [ $? != 0 ]; then
        echo "wrong output on test case $a"
        cleanup $1
        flag=1
        exit
    fi

    cleanup $1

    a=$((a+1))
done

if [ $flag -eq 0 ]; then
    echo "successfully executed"
fi