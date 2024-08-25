find ../ -type d -exec sh -c 'cd "$1" && go fmt || true' _ {} \;
