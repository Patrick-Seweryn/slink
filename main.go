package main
import (
    "fmt"
    "os"
)




type FilePath = string




func listItemSymlinks(item FilePath) {
    
}




func parseArguments(arg string) {

    switch arg {
    case "list":
	var item: FilePath = os.Args[2]
	return listItemSymlinks(item)

    case "move":
	var item FilePath = os.Args[2]
	var destination FilePath = os.Args[3]
	return moveItemToNewDir(item, destination)

    case "delete":
	var item FilePath = os.Args[2]
	return deleteItem(item)

    case "search":
	var directory = os.Args[2]
	return searchDirForBrokenSymlinks(directory)

    default:
	return false
    }
}




func parseArguments() {
}




func main() {
    argCount := len(os.Args)
    const firstArgIdx = 1

    
    for i := firstArgIdx; i < argCount; i++ {
	fmt.Println(i)
    }
    
}
