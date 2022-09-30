<p align="center">
  <a href="asset/img/site/favicon.svg">
    <img alt="replace" src="asset/img/site/favicon.svg" width="128"/>
  </a><br>
  <a href="http://golang.org">
      <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Made with Go">
  </a>

  <img src="https://img.shields.io/github/go-mod/go-version/CarsonSlovoka/replace?filename=src%2Fgo.mod" alt="Go Version">

  <a href="https://GitHub.com/CarsonSlovoka/replace/releases/">
      <img src="https://img.shields.io/github/release/CarsonSlovoka/replace" alt="Latest release">
  </a>
  <a href="https://github.com/CarsonSlovoka/replace/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/CarsonSlovoka/replace.svg" alt="License">
  </a>
</p>

# replace

字串取代

## Features

- [可指定要匹配的檔案**名稱**](https://github.com/CarsonSlovoka/replace/blob/1230a78f5e29ab84177b362fff48e27264c97aba/src/.replace.json#L2-L3)(不指定則視為全部對象)
- [可以指派資料夾](https://github.com/CarsonSlovoka/replace/blob/1230a78f5e29ab84177b362fff48e27264c97aba/src/.replace.json#L13-L14): 相對路徑, 絕對路徑都支持
- [能使用正規式進行取代](https://github.com/CarsonSlovoka/replace/blob/1230a78f5e29ab84177b362fff48e27264c97aba/src/.replace.json#L4-L5)

## Install

```
git clone https://github.com/CarsonSlovoka/replace.git
cd replace/src
go install -ldflags "-s -w"
```

## USAGE

在您的工作目錄新增檔案(例如:my-replace)，內容可以參考[.replace.json](src/.replace.json)

```
reaplceAll.exe -f="my-replace.json"
```
