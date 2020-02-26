package main

import (
        "log"
        "net/http"
        "os"
        "strings"
)

const (
        ENV_VAR_SERVER_PORT = "SERVER_PORT"
        ENV_VAR_SERVER_DIR = "SERVER_DIR"
        DEFAULT_SERVER_PORT = ":4000"
        DEFAULT_SERVER_DIR = "./static"
)

func configuredPort() string {
        return ReadEnvVar(ENV_VAR_SERVER_PORT, DEFAULT_SERVER_PORT)
}

func configuredDir() string {
        return ReadEnvVar(ENV_VAR_SERVER_DIR, DEFAULT_SERVER_DIR)
}

func ReadEnvVar(name string, defaultValue string) string {
        variable := os.Getenv(name)
        if variable == "" {
                variable = defaultValue
        }
        return variable
}

func main() {
        port := configuredPort()
        dir := configuredDir()
        log.Println("static-server serving directory: ", dir, " on port: ", port)
        err := http.ListenAndServe(port, http.FileServer(neuteredFileSystem{http.Dir(dir)}))
        log.Fatal(err)
}

type neuteredFileSystem struct {
        fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
        f, err := nfs.fs.Open(path)
        if err != nil {
                return nil, err
        }

        s, err := f.Stat()
        if s.IsDir() {
                index := strings.TrimSuffix(path, "/") + "/index.html"
                if _, err := nfs.fs.Open(index); err != nil {
                        return nil, err
                }
        }

        return f, nil
}