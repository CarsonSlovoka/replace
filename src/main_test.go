package main

import (
	"fmt"
	"os"
	"regexp"
)

func chdir(dir string) (back2orgDirFunc func()) {
	orgDir, _ := os.Getwd()
	_ = os.Chdir(dir) // 注意test會跑所有Example, Test的案例，所以當有chdir用的是相對路徑，那麼就要確保相對路徑都是正確的
	return func() {
		err := os.Chdir(orgDir)
		if err != nil {
			panic(err)
		}
	}
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
