app := "greenbutterfly-go"
version_pkg := "github.com/shsiddhant/greenbutterfly-go/cmd"

version := `git describe --tags --always --dirty 2>/dev/null || echo 0.0.0`

# Build binary
build:
    go build -ldflags="-X {{version_pkg}}.Version={{version}}" -o bin/{{app}} .

# Clean generated binaries
clean:
    rm -rf bin/
