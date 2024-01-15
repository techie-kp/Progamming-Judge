#!/bin/sh

function cleanup(){
    if [ -f "$1-code-output.txt" ]; then
        rm $1-code-output.txt
    fi
    if [ -f "$1-diff-messages.txt" ]; then
        rm $1-diff-messages.txt
    fi
    if [ -f "$1-main.out" ]; then
        rm $1-main.out
    fi
}
 
# $1 - id
# $2 - extension
# $3 - bind_mnt_dir
# $4 - timelimit

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

# touch $1-code-output.txt

# compile the c code 
gcc $3/$1-main.$2 -o $1-main.out

if [ $? != 0 ]; then
    echo "compile failed"
    cleanup $1
    exit
fi

a=1
flag=0
while [ -e "$3/$1-input/input-$a.txt" ]
do
    touch $1-code-output.txt

    # Execute and trap output
    timeout $4 ./$1-main.out < $3/$1-input/input-$a.txt &> $1-code-output.txt 

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
    diff $1-code-output.txt $3/$1-output/output-$a.txt > $1-diff-messages.txt

    if [ $? != 0 ]; then
        echo "wrong output on test case $a"
        #cat $1-diff-messages.txt
        cleanup $1
        flag=1
        exit
    fi

    rm $1-code-output.txt
    rm $1-diff-messages.txt

    a=$((a+1))
done

cleanup $1

if [ $flag -eq 0 ]; then
    echo "successfully executed"
fi