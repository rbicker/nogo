nogo
====

Nogo helps to compile non-go files into go binaries. It might be helpful if you want to get a kickstart while writing your own, small library. If you are looking for a well-written, feature-complete solution, please have a look at: https://github.com/markbates/pkger.

The nogo-method will only work if you are using go modules.

# 1) generate nogo.go
```bash
# install nogogen
go get github.com/rbicker/nogo/cmd/nogogen

# run nogogen to generate a a nogo file within your golang project
nogogen

# be default, nogogen will include a folder called "assets" and all of it's subfolders and -files
# if you want to include other (maybe multiple) directories, use the NOGO_DIRS env variable
NOGO_DIRS="/templates /public" nogogen
# please make sure to use absolute paths, using your project directory as root
```
