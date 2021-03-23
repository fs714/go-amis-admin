package version

var BaseVersion = "0.0.1-dev"
var GitVersion string
var GoVersion string
var BuildTime string

var Version = BaseVersion + " build on " + BuildTime + "\nGit Commit on " + GitVersion + "\n" + GoVersion
