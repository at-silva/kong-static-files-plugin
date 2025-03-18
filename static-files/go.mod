module github.com/at-silva/kong-plugin-static-files

go 1.21.3
toolchain go1.23.7

require (
	github.com/Kong/go-pdk v0.10.0
	github.com/onsi/gomega v1.29.0
)

replace github.com/Kong/go-pdk v0.10.0 => github.com/at-silva/go-pdk v0.10.1

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/net v0.36.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
