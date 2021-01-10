# Installing Go

## MacOs

1 - Install Golang using homebrew
```
brew update&& brew install golang
````

2 - Create your workspace
```
mkdir -p $HOME/go/{bin,src,pkg}
```

3 - Set up GOPATH and GOROOT
```
export GOPATH="/Users/username/go"
exort GOROOT="/usr/local/go"
```
Add them to your `.bashrc` or `.zshrc`

4 - You may start creating your projects in `~/go/src` as per [go_workspaces](../01_housekeeping/0102_go_workspace/0102_go_workspaces.md)

