# Evaluator

# Setup:
    go mod download
    cd images/python3
    docker build --tag python3-eval .
    cd ../pypy3
    docker build --tag pypy3-eval .
    cd images/java
    docker build --tag java-eval .
    cd images/c
    docker build --tag c-eval .
    cd images/cpp
    docker build --tag cpp-eval .

# Run:
    cd ../../src
    go run .

# Test it with some http requests

    localhost:7070/submit/eval?id=multi&lang=python3&timelimit=2s&memorylimit=64mb
    localhost:7070/submit/eval?id=multi&lang=pypy3&timelimit=2s&memorylimit=64mb
    localhost:7070/submit/eval?id=multi&lang=java&timelimit=2s&memorylimit=64mb
    localhost:7070/submit/eval?id=multi&lang=c&timelimit=2s&memorylimit=64mb
    localhost:7070/submit/eval?id=multi&lang=cpp14&timelimit=2s&memorylimit=64mb

    **Default time limit = 1s**
    **Default memory limit = 64MB**

# Invalid requests:
    localhost:7070/submit/eval?id=korakora&lang=python
    localhost:7070/submit/eval?id=korakora&lang=python2
    localhost:7070/submit/eval?id=korakora&lang=python3&timelimit=0.5s
    localhost:7070/submit/eval
    localhost:7070/submit/eval?id=korakora
    