package trace

import (
	"fmt"
	"runtime"
	"strings"
)

// WhereAmI where am I
func WhereAmI(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	return fmt.Sprintf("File: %s  Function: %s Line: %d", ChopPath(file), runtime.FuncForPC(function).Name(), line)
}

// ChopPath return the source filename after the last slash
func ChopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}

	return original[i+1:]

}
