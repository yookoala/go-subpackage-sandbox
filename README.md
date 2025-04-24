# Go Subpackage Sandbox

This is a repository of code purely to test the command `go get` ability to
correctly fetch subpackage from a GitHub repository.

The top level parent package (github.com/yookoala/go-subpackage-sandbox) is supposed
to be used by an optionally installed [subpackage](subpackage)
(github.com/yookoala/go-subpackage-sandbox), with its own
[go.mod](subpackage/go.mod) and [go.sum](subpackage/go.sum) file so that:

1. The parent package depends on nothing.
2. The subpackage depends on the parent package and some external library.


## Test Usages

In the [example application 1](examples/app1), it only depends on the parent
package. Running `go mod tidy` should only result in the installation of the
parent package, not the subpackage (nor its dependencies).

In the [example application 2](examples/app2), it will use both the parent
package and subpackage. Running `go mod tidy` should result in no issue with
both in their operations.
