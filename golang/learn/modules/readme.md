### Understanding go modules

A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.

> Modules are how Go manages dependencies.


**Each package** within a module is a collection of **source files in the same directory that are compiled together**. A package path is the module path joined with the subdirectory containing the package.
For example, the module "golang.org/x/net" contains a package in the directory "html". That package’s path is "golang.org/x/net/html".


->> **NOTE**

A module is identified by a module path, which is declared in a go.mod file, together with information about the module’s dependencies. *The module root directory is the directory that contains the go.mod file.*

#### Go Mod File

```
module github.com/mklef121/useTest  //The module path
go 1.17

require ( //The module dependencies
	github.com/fsnotify/fsnotify v1.5.1 
)
```


#### Module Paths

A module path is the canonical name for a module, declared with the module directive in the module’s go.mod file. A module’s path is the prefix for package paths within the module. (Recall That a package is a collection of go files in a module directory).

A module path should describe both what the module does and where to find it. This typically includes

- **A repository root path :** The repository root path is the portion of the module path that corresponds to the root directory of the *version control repository* where the module is developed.
- A directory within the repository
- and a major version suffix


#### Versions in modules

A version identifies an immutable snapshot of a module, which may be either a release or a pre-release.
- **Release Version:** A version without a pre-release suffix. For example, v1.2.3
- **Pre-release Version:** A version with a dash followed by a series of dot-separated identifiers immediately following the patch version, for example, v1.2.3-beta4,  v1.2.3-pre. Pre-release versions are considered unstable and are not assumed to be compatible with other versions.


#### go.mod files

A module is defined by a UTF-8 encoded text file named **go.mod** in its root directory. The go.mod file is line-oriented. Each line holds a single directive, made up of a keyword followed by arguments. For example:

```
module example.com/my/thing

go 1.12

require example.com/other/thing v1.0.2
require example.com/new/thing/v2 v2.3.4
exclude example.com/old/thing v1.2.3
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
retract [v1.9.0, v1.9.5]

```

##### Contents of the go.mod file

1. `go` directive: A go directive indicates that a module was written assuming the semantics of a given version of Go. The version must be a valid Go release version: a positive integer followed by a dot and a non-negative integer (for example, 1.9, 1.14). The go directive still affects use of new language features: 
    - For packages within the module, the compiler rejects use of language features introduced after the version specified by the go directive. For example, if a module has the directive go 1.12, its packages may not use numeric literals like 1_000_000, which were introduced in Go 1.13.
    - If an older Go version builds one of the module’s packages and encounters a compile error, the error notes that the module was written for a newer Go version. For example, suppose a module has go 1.13 and a package uses the numeric literal 1_000_000. If that package is built with Go 1.12, the compiler notes that the code is written for Go 1.13.
2. `require` directive: A require directive declares a minimum required version of a given module dependency. For each required module version, the go command loads the go.mod file for that version and incorporates the requirements from that file.
3. `exclude` directive: An exclude directive prevents a module version from being loaded by the go command. Once a module version is excluded, the Minimal version selection (MVS) algorithm used by GO will use the next upper version of the module anywhere it's needed in code. Example:

    ```
    exclude golang.org/x/net v1.2.3

    exclude (
        golang.org/x/crypto v1.4.5
        golang.org/x/text v1.6.7
    )
    ```
    
4. `replace` directive: A replace directive replaces the contents of a specific version of a module, or all versions of a module, with contents found elsewhere. The replacement may be specified with either another module path and version, or a platform-specific file path. If a version is present on the left side of the arrow (=>), only that specific version of the module is replaced; other versions will be accessed normally. If the left version is omitted, all versions of the module are replaced. Example:

    ```
    replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5

    replace (
        golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
        golang.org/x/net => example.com/fork/net v1.4.5
        golang.org/x/net v1.2.3 => ./fork/net
        golang.org/x/net => ./fork/net
    )
    ```


#### Vendoring

When using modules, the go command typically satisfies dependencies by downloading modules from their sources into the module cache, then loading packages from those downloaded copies.

**Vendoring** may be used to ensure that all files used for a build are stored in a single file tree.

To constructs a directory named vendor in the main module’s root directory containing copies of all packages needed to build and test packages in the main module use the `go mod vendor` command.


#### Go Module Commands

- `go get` :  usage: go get [-d] [-t] [-u] [-v] [build flags] [packages] 
    Get resolves its command-line arguments to packages at specific module versions, updates `go.mod` file to require those versions, downloads source code into the module cache, then builds and installs the named packages.
    The flags do the following
    * The `-d` flag tells `go get` not to build or install packages. When -d is used, go get will only manage dependencies in go.mod.
    * The `-u` flag tells go get to upgrade modules providing packages imported directly or indirectly by packages named on the command line. 
    * The `-t` flag tells go get to consider modules needed to build tests of packages named on the command line. When -t and -u are used together, go get will update test dependencies as well.

    ```bash
        #To add a dependency for a package or upgrade it to its latest version:
        go get example.com/pkg

        # To upgrade or downgrade a package to a specific version:
        go get example.com/pkg@v1.2.3 
        
        ## A version query suffix consists of an @ symbol followed by a version query, which may indicate a specific version (v0.3.0), a version prefix (v0.3), a branch or tag name (master), a revision (1234abcd), or one of the special queries latest, upgrade, patch, or none. If no version is given, go get uses the @upgrade query.
        

        ##To remove a dependency on a module and downgrade modules that require it:
        go get example.com/mod@none
    ```

    Once `go get` has resolved its arguments to specific modules and versions, go get will add, change, or remove *require directives* for that particular package in the main module’s go.mod file to ensure the modules remain at the desired versions in the future.

- `go install`: usage: go install [build flags] [packages]
    The **go Install** command compiles and installs the packages named by the import paths. Executables are installed in the directory named by the GOBIN environment variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH environment variable is not set. Non-executable packages are built and cached but not installed. Examples:
    
    ```bash
    # Install the latest version of a program,
    # ignoring go.mod in the current directory (if any).
    go install golang.org/x/tools/gopls@latest

    # Install a specific version of a program.
    go install golang.org/x/tools/gopls@v0.6.4

    # Install a program at the version selected by the module in the current directory.
    go install golang.org/x/tools/gopls

    # Install all programs in a directory.
    go install ./cmd/...
    ```

- `go mod download`: Usage: go mod download [-json] [-x] [modules]
    The `go mod download` command downloads the named modules into the module cache. With no arguments, download applies to all dependencies of the main module. Example:

    ```bash
    go mod download
    go mod download golang.org/x/mod@v0.2.0
    ```

    The go command will automatically download modules as needed during ordinary execution. The go mod download command is useful mainly for pre-filling the module cache or for loading data to be served by a module proxy.

- `go mod init`: Usage: go mod init [module-path]
    The go mod init command initializes and writes a new go.mod file in the current directory, in effect creating a new module rooted at the current directory. The go.mod file must not already exist. `init` accepts one optional argument, the [module path](#module-paths) for the new module. Example:

    ```bash
    go mod init
    go mod init example.com/m
    ```

- `go mod tidy`: Usage: go mod tidy [-e] [-v] [-go=version] [-compat=version]
    go mod tidy ensures that the go.mod file matches the source code in the module.  It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.
    * The -e flag (added in Go 1.16) causes go mod tidy to attempt to proceed despite errors encountered while loading packages.
    * The -v flag causes go mod tidy to print information about removed modules to standard error.
    **go mod tidy** works by loading all of the packages in the main module and all of the packages they import, recursively. This includes packages imported by tests (including tests in other modules). 

- `go mod vendor`: Usage: go mod vendor [-e] [-v] [-o]
    The **go mod vendor** command constructs a directory named vendor in the main module’s root directory that contains copies of all packages needed to support builds and tests of packages in the main module.
    When vendoring is enabled, the go command will load packages from the vendor directory instead of downloading modules from their sources into the module cache and using packages those downloaded copies. 

    go mod vendor also creates the file vendor/modules.txt that contains a list of vendored packages and the module versions they were copied from. When vendoring is enabled, this manifest is used as a source of module version information, as reported by go list -m and go version -m. When the go command reads vendor/modules.txt, it checks that the module versions are consistent with go.mod. If go.mod changed since vendor/modules.txt was generated, go mod vendor should be run again.

    ##### The command options
    The -e flag (added in Go 1.16) causes go mod vendor to attempt to proceed despite errors encountered while loading packages.

    The -v flag causes go mod vendor to print the names of vendored modules and packages to standard error.

    The -o flag (added in Go 1.18) causes go mod vendor to output the vendor tree at the specified directory instead of vendor.

- `go mod verify`: Usage: go mod verify
    go mod verify checks that dependencies of the main module stored in the module cache or vendor directory have not been modified since they were downloaded.  To perform this check, go mod verify hashes each downloaded module .zip file and extracted directory, then compares those hashes with a hash recorded when the module was first downloaded.

- `go clean`: Usage: go clean [-modcache]
    The `-modcache` flag causes go clean to remove the entire module cache, including unpacked source code of versioned dependencies.


#### Module-aware commands

Most go commands may run in Module-aware mode or GOPATH mode. In module-aware mode, the go command uses go.mod files to find versioned dependencies, and it typically loads packages out of the module cache, downloading modules if they are missing. In GOPATH mode, the go command ignores modules; it looks in vendor directories and in GOPATH to find dependencies.

As of Go 1.16, module-aware mode is enabled by default, regardless of whether a go.mod file is present. In lower versions, module-aware mode was enabled when a go.mod file was present in the current directory or any parent directory.

Module-aware mode may be controlled with the GO111MODULE environment variable, which can be set to on, off, or auto.

    If GO111MODULE=off, the go command ignores go.mod files and runs in GOPATH mode.
    If GO111MODULE=on or is unset, the go command runs in module-aware mode, even when no go.mod file is present. Not all commands work without a go.mod file: see Module commands outside a module.
    If GO111MODULE=auto, the go command runs in module-aware mode if a go.mod file is present in the current directory or any parent directory. In Go 1.15 and lower, this was the default behavior. go mod subcommands and go install with a version query run in module-aware mode even if no go.mod file is present.

In module-aware mode, GOPATH no longer defines the meaning of imports during a build, but it still stores downloaded dependencies (in GOPATH/pkg/mod; see Module cache) and installed commands (in GOPATH/bin, unless GOBIN is set).



----------------------------------------------- Phewww, enough of the theory, lets go practical now ----------------------------------

### Using Go Modules


#### Creating a new module

`cd` into a project folder you just started and run 


`go mod init example.com/hello`

This makes the current directory the root of a module.

Packages in subdirectories have import paths consisting of the module path plus the path to the subdirectory. For example, if we created a subdirectory world, we would not need to (nor want to) run go mod init there. The package would automatically be recognized as part of the example.com/hello module, with import path example.com/hello/world.

#### Adding a dependency

use `go get rsc.io/sampler` or `go get rsc.io/sampler@v1.3.1`


#### Removing unused dependencies

use `go list -m all` to list all depenedency in your project

use `go mod tidy` to clean up unused depenedencies















