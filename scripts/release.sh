for GOOS in linux darwin windows; do
    for GOARCH in amd64 arm64; do
        output="videoalchemy-${GOOS}-${GOARCH}"
        [ "$GOOS" = "windows" ] && output="${output}.exe"
        GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-X 'main.version=1.0-RC1' -X 'main.date=$(date -u +%Y-%m-%d)'" -o $output ./cmd/compose
    done
done
