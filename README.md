# replace

字串取代

## Features

- [可指定多個副檔名](https://github.com/CarsonSlovoka/replace/blob/40066b77040f0f694a7558fcf2b561f903cd3f14/src/.replace.json#L2)(不指定則視為全部對象)
- [可以指派資料夾](https://github.com/CarsonSlovoka/replace/blob/40066b77040f0f694a7558fcf2b561f903cd3f14/src/.replace.json#L12-L13): 相對路徑, 絕對路徑都支持
- [能使用正規式進行取代](https://github.com/CarsonSlovoka/replace/blob/40066b77040f0f694a7558fcf2b561f903cd3f14/src/.replace.json#L4-L5)

## Install

```
cd src
go install -ldflags "-s -w"
```

## USAGE

在您的工作目錄新增檔案(例如:my-replace)，內容可以參考[.replace.json](src/.replace.json)

```
reaplceAll.exe -f="my-replace.json"
```
