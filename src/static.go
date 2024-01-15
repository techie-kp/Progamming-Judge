package main

const SECONDS = 's'
const MEGABYTES = "mb"
const DEFAULT_TIME_LIMIT = 1
const DEFAULT_MEMORY_LIMIT = 64

var bind_mnt_dir = "submissions"
var unp_user = "execution_user"
var lang_extension_map = map[string]string{
	"cpp14":   "cpp",
	"cpp17":   "cpp",
	"cpp20":   "cpp",
	"python3": "py",
	"pypy3":   "py",
	"python2": "py",
	"pypy2":   "py",
	"c":       "c",
	"java":    "java",
}
var lang_image_map = map[string]string{
	"c":       "c-eval",
	"cpp14":   "cpp-eval", // TODO
	"python3": "python3-eval",
	"pypy3":   "pypy3-eval",
	"java":    "java-eval",
}
