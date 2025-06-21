// src/main.go

package main

import (
    "fmt"
    "os"
    "strings"
    "encoding/json"
)

type Env struct {
    REQUEST     map[string]string
    FILES       map[string]string
    HTTP        map[string]string
    SERVER      map[string]string
    RESPONSE    map[string]string
}

type EnvRef struct {
    envContext  *map[string]string
    field       string
}

func main() {
    fmt.Println("Content-Type: text/plain\n")
    envVars := Env{
        REQUEST:    map[string]string{},
        FILES:      map[string]string{},
        HTTP:       map[string]string{},
        SERVER:     map[string]string{},
        RESPONSE:   map[string]string{},
    }

    env_map                             := map[string]EnvRef{}
    env_map["REQUEST_METHOD"]           = EnvRef{envContext: &envVars.REQUEST, field: "METHOD"}
    env_map["QUERY_STRING"]             = EnvRef{envContext: &envVars.REQUEST, field: "QUERY_STRING"}
    env_map["REQUEST_SCHEME"]           = EnvRef{envContext: &envVars.REQUEST, field: "SCHEME"}

    env_map["HTTP_ACCEPT"]              = EnvRef{envContext: &envVars.HTTP, field: "ACCEPT"}
    env_map["HTTP_CACHE_CONTROL"]       = EnvRef{envContext: &envVars.HTTP, field: "CACHE_CONTROL"}
    env_map["HTTP_COOKIE"]              = EnvRef{envContext: &envVars.HTTP, field: "COOKIE"}
    env_map["HTTP_CONNECTION"]          = EnvRef{envContext: &envVars.HTTP, field: "CONNECTION"}
    env_map["HTTP_ACCEPT_ENCODING"]     = EnvRef{envContext: &envVars.HTTP, field: "COMPRESS"}
    env_map["HTTP_HOST"]                = EnvRef{envContext: &envVars.HTTP, field: "HOST"}

    env_map["SERVER_ADMIN"]             = EnvRef{envContext: &envVars.SERVER, field: "ADMIN"}
    env_map["SERVER_NAME"]              = EnvRef{envContext: &envVars.SERVER, field: "NAME"}
    env_map["SERVER_PORT"]              = EnvRef{envContext: &envVars.SERVER, field: "PORT"}
    env_map["SERVER_PROTOCOL"]          = EnvRef{envContext: &envVars.SERVER, field: "PROTOCOL"}

    env_map["SERVER_PROTOCOL"]          = EnvRef{envContext: &envVars.RESPONSE, field: "REDIRECT_STATUS"}

    for _, map_item := range env_map {
        (*map_item.envContext)[map_item.field] = ""
    }

    envs := os.Environ()
    for _, env := range envs {
        pair := strings.SplitN(env, "=", 2) // separa em chave e valor
        key := pair[0]
        val := ""
        if len(pair) > 1 {
            val = pair[1]
        }
        fmt.Printf("%s = %s\n", key, val)

        if map_item, ok := env_map[key]; ok {
            (*map_item.envContext)[map_item.field] = val
        }
    }

    fmt.Println(toJson(envVars))
    // fmt.Printf("%+v", envVars)
    // fmt.Printf("%#v", envVars)
}

func toJson(text Env) string {
    out, err := json.MarshalIndent(text, "", "  ")
    if err != nil {
        fmt.Println("Erro ao converter:", err)
        os.Exit(0)
    }

    return string(out)
}
