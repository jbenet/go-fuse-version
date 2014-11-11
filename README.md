# go-fuse-version

Simple package to get the user's FUSE libraries information.

- Godoc: https://godoc.org/github.com/jbenet/go-fuse-version

**Warning** Currently only supports OSXFUSE. if you want more, add them, it's really trivial now.

## fuseprint

If you dont use Go, you can also install the tiny silly util fuseprint:

```
> go get github.com/jbenet/go-fuse-version/fuseprint
> go install github.com/jbenet/go-fuse-version/fuseprint
> fuseprint
FuseVersion, AgentVersion, Agent
27, 2.7.2, OSXFUSE
```
