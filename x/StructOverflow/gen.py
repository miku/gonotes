# package main
#
# import "fmt"
#
# func main() {
# 	var a, b struct{}
# 	fmt.Println(&a == &b) // false, but may be true?
# }
#
#
vars = ["s{}".format(i) for i in range(20000)]
assignments = ["fmt.Println({})".format(v) for v in vars]


print("""
package main

import "fmt"
import "time"
import "runtime"

func main() {

    var %s struct{}
    %s

    PrintMemUsage()
    time.Sleep(10 * time.Second)

}
""" % (",".join(vars), "\n".join(assignments)))

print(r"""
// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
""")