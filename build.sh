#!/usr/bin/env bash

GO_PATH_DEFAULT="/Users/DaisyZhu/go"

go_path=$GO_PATH_DEFAULT
go_yacc="${go_path}/bin/goyacc"


function info {
    echo
    echo "./build [-h | --help]"
    echo "./build [-p | --go-path path]"
    echo
    echo "Named argumetns:"
    echo "  -p | --go-path   path to golang SDK root"
    echo
}

function parse_args {
    # named args
    while [[ "$1" != "" ]]; do
        case "$1" in
            -p | --go-path )    go_path="$2";             shift;;
            -o | --out )        out_path="$2";            shift;;
            -p | --package )    package_name="$2";        shift;;
            -a | --antlr )      jcp_path="$2";            shift;;
            -h | --help )       info;                     exit;; # quit and show usage
            * )                 echo "unrecognized: $1"
        esac
        shift # move to next kv pair
    done

    if [[ ! -f ${go_yacc} || ! -d ${go_path} ]]; then 
        echo "Could not find goyacc: \"${go_yacc}\" or gopath: \"${go_path}\" "
        echo "  Make sure it is installed or GOPATH is specified correctly"
        echo "  You can install go path with 'go get -u golang.org/x/tools/cmd/goyacc'"
    else
        ${go_yacc} -l -o lure_y.go -p Lure gramma/lure.y
    fi
}

parse_args "$@"

