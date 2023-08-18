## Recursive file search
Written in Golang. It uses recursive multithreading to search for given file in provided path

Usage
```
go run main.go [ROOT_PATH] [FILENAME]
```

### Example
```
$ go run . ./ README.md       

Searchig for README.md in ./
Searchig for README.md in .git
Searchig for README.md in .git/logs
Searchig for README.md in .git/info
Searchig for README.md in .git/logs/refs
Searchig for README.md in .git/logs/refs/remotes
Searchig for README.md in .git/refs
Searchig for README.md in .git/logs/refs/heads
Searchig for README.md in .git/logs/refs/remotes/origin
Searchig for README.md in .git/refs/tags
Searchig for README.md in .git/refs/heads
Searchig for README.md in .git/refs/remotes
Searchig for README.md in .git/objects
Searchig for README.md in .git/hooks
Searchig for README.md in .git/refs/remotes/origin
Searchig for README.md in .git/objects/pack
Searchig for README.md in .git/objects/e6
Searchig for README.md in .git/objects/8c
Searchig for README.md in .git/objects/info
Searchig for README.md in .git/objects/b7
Searchig for README.md in .git/objects/f5
Searchig for README.md in .git/objects/e7
✅ Matched README.md
ℹ️ Total matched files: 1
```