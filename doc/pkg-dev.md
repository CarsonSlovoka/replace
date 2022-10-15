# 發佈到pkg.dev的相關事項

go package的類型簡易分成兩類:

1. executable package: 工具包，提供應用程式讓使用者應用. 存在`package main`
2. utility package: 大多數的go專案都是如此，是一種套件或者想像成語法糖，能讓其他專案引用，簡化代碼。

建議根目錄最好存在`go.mod`，不然pkg.dev可能會出不來<sup>v2的放法除外</sup>，而且也有很多應用都吃這套

由於go.mod位於根目錄，又當您的使用的是`package main`，這時候他會用`資料夾的名稱`來當作go install預設產出的執行檔名稱，例如`go install github.com/CarsonSlovoka/app/cool`請查看以下的第二點

簡單的歸類:

1. go.mod與go install的工作目錄同資料夾: 那麼install出來的執行檔名為go.mod所命名的package名稱
    ```yaml
    📂 src/
      - main.go
      - go.mod # module cool

    $ cd src
    $ go install # 產出的執行檔名稱同go.mod所命名的名稱 => cool.exe
    ```

2. go.mod與install不同目錄: 使用`package main`所在的**資料夾**當作預設的執行檔名稱

    ```yaml
    go.mod # github.com/CarsonSlovoka/app
    📂 cool
      - main.go
    $ go install github.com/CarsonSlovoka/app/cool # 它會用最後一個名稱也就是資料夾名稱來當成執行檔的名稱
    ```


