package main

const (
	CREATING_TABLE  = "Creating %v table in %v"
	CREATING_BUCKET = "Creating %v bucket in %v"

	ERR_CREATING_TABLE  = "Error creating %v table in %v"
	ERR_CREATING_BUCKET = "Error creating %v bucket in %v"
)

var ctToExt = map[string]string{
	"application/octet-stream":				  ".mov",
	"video/quicktime":				  ".mov",
	"text/css":        ".css",
	"image/gif":                      ".gif",
	"image/jpeg":                     ".jpeg",
	"text/javascript": ".js",
	"application/json":               ".json",
	"application/pdf":                ".pdf",
	"image/png":                      ".png",
	"image/webp":                     ".webp",
}
