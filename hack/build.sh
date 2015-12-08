## Builds the scripts for the given GOOS and GOARCH

TARGET_DIR=$WSO2SCRIPTS_BUILD_DIR/$GOOS/$GOARCH

SCRIPTS=(*.go)

for script in "${SCRIPTS[@]}"
do
    echo "Building $script...."
    go build -o $TARGET_DIR/"${script%.*}" -v $script
done

echo "All scripts built"
