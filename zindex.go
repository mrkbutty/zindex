/*
zindex scans a directory of zip files and creates a Windows shortcut to the files within resolving duplicates


*/

package main

import "fmt"
import "os"
import "log"
import "flag"
import "time"


var flagVerbose bool
var flagQuiet bool
var flagTestOnly bool




func main() {
	flag.BoolVar(&flagVerbose, "v", false, "Prints detailed operations")
	flag.BoolVar(&flagQuiet, "q", false, "No output apart from errors")
	flag.BoolVar(&flagTestOnly, "t", false, "Test only do not change")
	flag.Parse()

	//items := []string{"."}  // default arguments to use if omitted

	if flag.NArg() < 2 {
		flag.PrintDefaults()
		return
	}

	start := time.Now()

	fmt.Printf("source dir %s\n" , cflag.Args[0])
	fmt.Printf("target dir %s\n" , cflag.Args[1])


	if !flagQuiet {
		elapsed := time.Since(start)
		fmt.Printf("Processed %d total items in %v\n" , total, elapsed)
	}
}


