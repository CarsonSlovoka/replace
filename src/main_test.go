package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

func chdir(dir string) (back2orgDirFunc func()) {
	orgDir, _ := os.Getwd()
	_ = os.Chdir(dir) // 注意test會跑所有{Example, Test}的案例，所以當有chdir，那麼就要確保相對路徑都正確！(每一個做完都要還原會去，否則下一個的相對路徑將會基於前一個改變的結果)
	return func() {
		err := os.Chdir(orgDir)
		if err != nil {
			panic(err)
		}
	}
}

// GitHub.action.checkout如果沒有在.gitAttributes設定eol，那麼在windows下，他會自動變成crlf，因此讀到的txt都是crlf結尾，與go中寫的Output的lf結尾不同，就會導致錯誤
func getStdOut(mainFunc func()) string {
	orgStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	mainFunc()

	_ = w.Close()
	bs, _ := io.ReadAll(r)
	_ = r.Close()
	os.Stdout = orgStdout
	return string(bs)
}

func Example_main_caseSensitive() {
	back2orgDirFunc := chdir("test/case_sensitive")
	defer back2orgDirFunc()
	os.Args = []string{os.Args[0], "-f=config.json", "-dry=1"}
	main()
	// Output:
	// H?ll? World
}

func Example_main_caseInsensitive() {
	back2orgDirFunc := chdir("test/case_insensitive")
	defer back2orgDirFunc()
	os.Args = []string{os.Args[0], "-f=config.json", "-dry=1"}
	main()
	// Output:
	// H?ll? W?rld
}

// multiline影響的是匹配開頭的^或者結尾的$符號才會需要用到
func Example_main_multiline() {

	back2orgDirFunc := chdir("test/multiline")
	defer back2orgDirFunc()

	os.Args = []string{os.Args[0], "-f=config.json", "-dry=1"}
	main()
	// Output:
	// Hello World foo
	// Hello World
}

func Example_mainPattern() {
	back2orgDirFunc := chdir("test/pattern")
	defer back2orgDirFunc()

	os.Args = []string{os.Args[0], "-f=config.json", "-dry=1"}
	main()
	// Output:
	// Example0
	// Example00
}

func Example_core() {
	const testContent = `
<glyph name="uni76B8" format="2">
  <advance width="888" height="456"/>
  <unicode hex="76B8"/>
</glyph>
`
	substitution := "<advance width=\"$1\" height=\"1024\"/>"
	re := regexp.MustCompile(`<advance width="(\d*)" height="(\d*)"/>`)
	fmt.Println(re.ReplaceAllString(testContent, substitution))
	// Output:
	// <glyph name="uni76B8" format="2">
	//   <advance width="888" height="1024"/>
	//   <unicode hex="76B8"/>
	// </glyph>
}
