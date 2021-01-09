# Go Workspaces
Go is very opinionaded as to how code should be organised, for better efficiency and team productivity
this is how go expects a go project dirtree to be:


## Go space organisation
```
~/go
  .
  ├── bin                 # executables
  ├── pkg                 # package objects compiled from src directory 
  └── src                 # go sorce files
      └── github.com      # project repository path
            └── driprado
                └── starting_up_with_go
```

Code arranged this way can be downloaded with the command `go get`

## Environment variables
`GOPATH` should point to your workspace ex: `GOPATH="/Users/username/go"`  
`GOROOT` should point to your go binary location ex: `GOROOT="/usr/local/go"`