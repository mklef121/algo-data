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
3. `exclude` directive: An exclude directive prevents a module version from being loaded by the go command. Example:
    ```
    exclude golang.org/x/net v1.2.3

    exclude (
        golang.org/x/crypto v1.4.5
        golang.org/x/text v1.6.7
    )```




