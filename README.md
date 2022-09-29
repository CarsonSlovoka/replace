# replace

字串取代

## Features

- 可指定多個副檔名(不指定則視為全部對象)
- 可以指派資料夾
- 能使用正規式進行取代

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
