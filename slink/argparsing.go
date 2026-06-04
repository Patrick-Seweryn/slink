package main
import (
    "fmt"
    "os"
    "strings"
)




func listItemSymlinks(arguments []string) {
    fmt.Println("Listing item symlinks...")
}




func moveItemOrSymlink(arguments []string) {
    fmt.Println("Moving item or symlink...")
}




func deleteItemAndOrItsSymlinks(arguments []string) {
    fmt.Println("Deleting item and/or its symlinks...")
}




func searchDirForBrokenSymlinks(arguments []string) {
    fmt.Println("Searching dir for broken symlinks...")
}




func outputHelpMessage(arguments []string) {
    fmt.Println(`
Usage:
    slink [behaviour] [file/directory ...]

NOTE: Here, 'item' refers to a file or directory. and 'symlink' refers to
symbolic link.

Behaviour (Max. 1):
    --list          List all symlinks to items, that were created using 'slink'.
                    e.g. slink list /path/to/item /optional/path/to/other/item ..

    --move          Change the path of existing item or symlink.
                    e.g. slink move /path/to/item/or/symlink /new/path

    --delete        Destroy existing items and/or their symlinks, applies
                    recursively on directories. If deleting symlink, it does NOT
                    delete the original item.
                    e.g. slink delete /path/to/item/or/symlink /optional/path/to/
                         other/item/or/symlink ...

    --searchBroken  Recursively search directories for broken symlinks.
                    e.g. slink searchBroken /path/to/search/directory /optional/
                         path/to/other/search/directory

    --help          Output this help text into the console.

Create a symbolic link with:
    slink /path/to/existing/file/or/directory /path/of/the/symolic/link
    /optional/path/of/another/symbolic/link ...

Only use 'slink' on items you control. Don't create symlinks to items handled by
package managers.
    `)
}




func createSymlinks(arguments []string) {
    fmt.Println("Creating symlinks...")
}




func getExecutableArguments() []string {
    const executableNameIdx = 0
    const firstArgumentIdx = executableNameIdx + 1
    lastArgumentIdx := len(os.Args)

    argumentSlice := os.Args[firstArgumentIdx : lastArgumentIdx]
    return argumentSlice
}




type BehaviourFunction func([]string)
func selectBehaviourFromArguments(arguments []string) BehaviourFunction {
    joinedString := strings.Join(arguments, " ")
    fmt.Println(joinedString)
    for _, arg := range arguments {
        switch (arg) {
        case "--list", "--ls", "-l":
            return listItemSymlinks

        case "--move", "--mv", "-m":
            return moveItemOrSymlink

        case "--delete", "--rm", "-d":
            return deleteItemAndOrItsSymlinks

        case "--searchBroken", "-s":
            return searchDirForBrokenSymlinks

        case "--help", "-h":
            return outputHelpMessage
        }
    }
    return createSymlinks
}




func computeExecutable() {
    arguments := getExecutableArguments()
    behaviour := selectBehaviourFromArguments(arguments)
    filePaths := os.Args

    behaviour(filePaths)
}
