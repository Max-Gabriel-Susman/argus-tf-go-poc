module github.com/Max-Gabriel-Susman/argus-tf-go-poc

go 1.23.4

require (
	github.com/wamuir/graft v0.3.0
	gocv.io/x/gocv v0.39.0
	golang.org/x/image v0.23.0
)

replace github.com/wamuir/graft v0.9.0 => github.com/wamuir/graft v0.3.0

require google.golang.org/protobuf v1.34.2 // indirect
