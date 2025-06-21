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

func main() {
    fmt.Println("Content-Type: text/plain\n")
    envVars := Env{
        REQUEST:    map[string]string{},
        FILES:      map[string]string{},
        HTTP:       map[string]string{},
        SERVER:     map[string]string{},
        RESPONSE:   map[string]string{},
    }
    
    envVars.REQUEST["METHOD"]           = ""
    envVars.REQUEST["QUERY_STRING"]     = ""
    envVars.REQUEST["PARAMS"]           = ""
    envVars.REQUEST["SCHEME"]           = ""
    envVars.REQUEST["URL"]              = ""
    envVars.REQUEST["TYPE"]             = ""

    envVars.HTTP["ACCEPT"]              = ""
    envVars.HTTP["CACHE_CONTROL"]       = ""
    envVars.HTTP["COOKIE"]              = ""
    envVars.HTTP["CONNECTION"]          = ""
    envVars.HTTP["COMPRESS"]            = ""
    envVars.HTTP["HOST"]                = ""

    envVars.SERVER["ADMIN"]             = ""
    envVars.SERVER["NAME"]              = ""
    envVars.SERVER["PORT"]              = ""
    envVars.SERVER["PROTOCOL"]          = ""
    envVars.SERVER["OS"]                = ""
    
    envVars.RESPONSE["REDIRECT_STATUS"] = ""

    env_map                             := map[string][2]string{}
    env_map["REQUEST_METHOD"]           = [2]string{"REQUEST", "METHOD"}
    env_map["QUERY_STRING"]             = [2]string{"REQUEST", "QUERY_STRING"}
    env_map["REQUEST_SCHEME"]           = [2]string{"REQUEST", "SCHEME"}

    env_map["HTTP_ACCEPT"]              = [2]string{"HTTP", "ACCEPT"}
    env_map["HTTP_CACHE_CONTROL"]       = [2]string{"HTTP", "CACHE_CONTROL"}
    env_map["HTTP_COOKIE"]              = [2]string{"HTTP", "COOKIE"}
    env_map["HTTP_CONNECTION"]          = [2]string{"HTTP", "CONNECTION"}
    env_map["HTTP_ACCEPT_ENCODING"]     = [2]string{"HTTP", "COMPRESS"}
    env_map["HTTP_HOST"]                = [2]string{"HTTP", "HOST"}

    env_map["SERVER_ADMIN"]             = [2]string{"SERVER", "ADMIN"}
    env_map["SERVER_NAME"]              = [2]string{"SERVER", "NAME"}
    env_map["SERVER_PORT"]              = [2]string{"SERVER", "PORT"}
    env_map["SERVER_PROTOCOL"]          = [2]string{"SERVER", "PROTOCOL"}

    env_map["SERVER_PROTOCOL"]          = [2]string{"RESPONSE", "REDIRECT_STATUS"}

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
            switch map_item[0] {
                case "REQUEST":
                    envVars.REQUEST[map_item[1]]    = val
                case "HTTP":
                    envVars.HTTP[map_item[1]]       = val
                case "SERVER":
                    envVars.SERVER[map_item[1]]     = val
                case "RESPONSE":
                    envVars.RESPONSE[map_item[1]]   = val
            }
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
