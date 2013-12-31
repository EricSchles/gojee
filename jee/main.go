package main

import (
    "github.com/nytlabs/gojee"
    "fmt"
    "encoding/json"
    "os"
    "io/ioutil"
)


var info = `jee 0.1.0`

func main() {
    var umsg jee.BMsg

    if len(os.Args) != 2 {
        fmt.Println(info)
        os.Exit(1)
    }

    l, err := jee.Lexer(string(os.Args[1]))
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    tree, err := jee.Parser(l)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    j, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    err = json.Unmarshal(j, &umsg)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    result, err := jee.Eval(tree, umsg)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    r, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(r))
}