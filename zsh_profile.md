对于mac的终端配置，增加zsh_profile：
```
alias ll="ls -l -a"

export GOROOT="/usr/local/go"
export GOPATH="/Users/llwei/thirdparty/go"
export GOBIN="$GOPATH/bin"
export PATH="$PATH:$GOBIN"

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export LS_OPTIONS='--color=auto'
export CLICOLOR='Yes'
export LSCOLORS='ExfxcxdxbxegedabagGxGx'
export PS1='%n: %1~ $ '
```
