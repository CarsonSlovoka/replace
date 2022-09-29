package main

import (
	"fmt"
	"regexp"
)

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
