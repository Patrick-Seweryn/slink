package main
import (
    "fmt"
    "strings"
    "os"
)



func caseHelp() {
  /*
    fmt.Println(`
Help (verbose):
    Create symlink/s on existing item:
    'slink /path/to/item /path/for/symlink /optional/path/for/other/symlink ...'
	NOTE: Replacing /path/to/item with a path to existing symlink will just
	      point to the original item


    List all an item/s symlinks
    'slink list /path/to/item /optional/path/to/other/item ...'
	NOTE: This only lists those created using slink.


    Move an existing item or symlink to ensuring nothing is broken
    'slink move /path/to/item /path/to/new/directory'


    Destroy an item or symlink 
    'slink delete /path/to/existing/item/or/symlink'
	NOTE: If deleting an *item*, all its symlinks will also be deleted (with a
	      'Y/n' clarification on whether to execute)
	NOTE: If said item is a directory, the command will carry onto any symlinks
	      or symlinked items inside (recursively).


    Recursively search a directory for broken symlinks:
    'slink searchBroken /optional/path/to/directory /optional/path/to/other/directory ...'
	NOTE: If you don't specify a directory, it will search the current one.


When managing symlinks or symlinked items, PLEASE use 'slink' instead of
typical file management commands like 'mv' or 'rm'. Broken symlinks will
clutter your file system (very annoying).

I recommend only using 'slink' on files you control. Don't create symlinks
to stuff handled by package managers.
    `)
    */

    fmt.Println(`
Usage:
  slink [optional subcommand] [file/directory ...]

NOTE: Here, 'item' refers to a file or directory. and 'symlink' refers to
symbolic link.
Optional Subcommands:
  list		List all symlinks to items, that were created using 'slink'.
		e.g. slink list /path/to/item /optional/path/to/other/item ..

  move          Move existing item / symlink.
		e.g. slink move /path/to/item/or/symlink /new/path

  delete 	Destroy existing items / symlinks, applies recursively on
		directories. If deleting symlink, it does not delete the
		original item.
		e.g. slink delete /path/to/item/or/symlink /optional/path/to/
		       other/item/or/symlink ...

  searchBroken  Recursively search directories for broken symlinks.
		e.g. slink searchBroken /path/to/search/directory /optional/
		       path/to/other/search/directory

  help          Output this help text into the console.

Create a symbolic link with:
  slink /path/to/existing/file/or/directory /path/of/the/symolic/link
    /optional/path/of/another/symbolic/link ...

Only use 'slink' on items you control. Don't create symlinks to stuff
handled by package managers.

    `)
}




func caseList(arguments []string) {
    joinedString := strings.Join(arguments, "  ")
    fmt.Println("Listing symlinks to items: " + joinedString)
}




func caseMove(arguments []string) {
    fmt.Println("Moving item " + arguments[0] + " to " arguments[1])
}




func caseDelete(arguments []string) {
    joinedString := strings.Join(arguments, "  ")
    fmt.Println("Deleting item " + joinedString)
}




func caseSearchBroken(arguments []string) {
    joinedString := strings.Join(arguments, "  ")
    fmt.Println("Searching directories " + joinedString + " for broken symlinks")
}




func caseCreateSymlink(arguments []string) {
    symlinkPaths := arguments[1:]
    joinedString := strings.Join(symlinkPaths, "  ")
    fmt.Println("Creating symlinks to " + arguments[0] + " at " + joinedString)
}




func parseArguments() {
    const firstArgumentIdx = 1
    arguments := os.Args[firstArgumentIdx:]
    argumentCount = len(arguments)

    switch (arguments[firstArgumentIdx]) {
    case "list":
	caseList(arguments)

    case "move":
	caseMove(arguments)

    case "delete":
	caseDelete(arguments)

    case "searchBroken":
	caseSearchBroken(arguments)

    case "help":
	caseHelp()

    default:
	caseCreateSymlink()
    }
}












