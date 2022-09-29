package main

import (
	"encoding/json"
	"fmt"
	. "github.com/CarsonSlovoka/go-pkg/v2/fmt"
	"github.com/CarsonSlovoka/go-pkg/v2/slices"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

type Config struct {
	SearchExtensions []string `json:"ext"` // 要處理的附檔名
	Dirs             []string // 要處理的資料夾路徑
	Regexp           string   // 正規式
	Substitution     string   // 取代內容
	MaxLoading       int      // 每一個routine可以處理的檔案上限
	Verbose          bool
}

func NewConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	cfg := new(Config)
	err = json.NewDecoder(f).Decode(cfg)
	return cfg, err
}

var (
	pWarn *ColorPrinter
	pErr  *ColorPrinter
)

func init() {
	pWarn = NewColorPrinter(255, 255, 255, 0, 0, 255)
	pErr = NewColorPrinter(255, 255, 255, 255, 0, 0)
}

func main() {
	cfg, err := NewConfig(".replace.json")
	if err != nil {
		log.Fatal(err)
	}

	allFileList := make([]string, 0) // 需要被處理的檔案路徑
	for _, dirPath := range cfg.Dirs {
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			pWarn.Printf("path not found: %s\n", dirPath)
			continue
		}
		_ = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error { // 會遞迴的找所有內容
			if info.IsDir() {
				return nil
			}

			if len(cfg.SearchExtensions) == 0 {
				allFileList = append(allFileList, path)
				return nil
			}

			if len(cfg.SearchExtensions) > 0 && slices.Any(cfg.SearchExtensions, filepath.Ext(path)) { // .md, .rst, ...
				allFileList = append(allFileList, path)
				return nil
			}
			return nil
		})
	}
	wg := sync.WaitGroup{}
	sTime := time.Now()
	subFileList := slices.ChunkBy(allFileList, cfg.MaxLoading)
	re := regexp.MustCompile(cfg.Regexp)
	for _, subFiles := range subFileList {
		wg.Add(1)
		go func(files []string) {
			defer wg.Done()
			for _, fPath := range files {
				f, err := os.OpenFile(fPath, os.O_RDWR, 0666)
				if err != nil {
					log.Println(pErr.Sprintln(err))
					continue
				}
				defer func() {
					_ = f.Close()
				}()

				bs, err := io.ReadAll(f)
				if err != nil {
					log.Println(pErr.Sprintln(err))
				}

				if !re.Match(bs) {
					return
				}

				newStr := re.ReplaceAllString(string(bs), cfg.Substitution)
				_ = f.Truncate(0)     // 我們使用O_RDWR所以可以再寫入，把所有內容截斷(清除內文)
				_, err = f.Seek(0, 0) // 指標回到0,0的位置再開始重新寫入
				if err != nil {
					log.Println(err)
					continue
				}
				if _, err = f.Write([]byte(newStr)); err != nil {
					_ = f.Close()
					log.Println(err)
					continue
				}
				_ = f.Close()
				if cfg.Verbose {
					log.Printf("File:%s changed done.\n", fPath)
				}
			}
		}(subFiles)
	}
	wg.Wait()
	fmt.Printf("%.0f seconds in total\n", time.Now().Sub(sTime).Seconds())
}
