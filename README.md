# git-open [![Actions Status](https://github.com/sminamot/git-open/workflows/Go/badge.svg)](https://github.com/sminamot/git-open/actions)

![git-open](https://user-images.githubusercontent.com/26164869/73953393-ae0dd500-4943-11ea-91e2-fd7b4278c71b.gif)

## Installation
```
$ go get github.com/sminamot/git-open
```
or download binary from [release page](https://github.com/sminamot/git-open/releases)

## Usage
```
# origin / current branch
$ git open

# specified remote / current branch
$ git open upstream

# specified remote / specified branch
$ git open upstream develop
```

## Completion
### zsh
download completion script
```
$ wget https://raw.githubusercontent.com/sminamot/git-open/master/completion/_git-open -P ~/.zsh/completion/
```
add this in `~/.zshrc`
```
fpath=(~/.zsh/completion $fpath)
autoload -Uz compinit && compinit -i
```

![git-open_competion](https://user-images.githubusercontent.com/26164869/74239397-1d584000-4d1b-11ea-925e-8fc7cd633e88.gif)
