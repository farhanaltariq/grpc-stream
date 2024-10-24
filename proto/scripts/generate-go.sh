#!/run/current-system/sw/bin/bash
dir="$(dirname $(realpath ${BASH_SOURCE[0]}))" && cd $dir
dir=$dir/..

cd $dir

export GOBIN=$dir/bin
export GO111MODULE=on

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go get -d github.com/favadi/protoc-go-inject-tag@v1.3.0
export PATH=$GOBIN:$GOPATH/bin:$HOME/go/bin:$PATH

if ! [ -x "$(command -v protoc)" ]; then
    echo "protoc is required"
    exit
fi

if ! [ -x "$(command -v protoc-go-inject-tag)" ]; then
    echo "protoc-go-inject-tag is required"
    exit
fi

cd $dir

goprotocout=$dir/go

rm -rf $goprotocout && mkdir -p $goprotocout

echo "Using protoc version" $(protoc --version), "protoc-gen-go version", $(protoc-gen-go --version)

cd $dir
for entry in $(find ./ -iname "*.proto"); do
    protoc --go_out=$goprotocout \
        --go_opt=paths=import \
        --go_opt=module=github.com/farhanaltariq/proto \
        --go-grpc_out=$goprotocout \
        --go-grpc_opt=paths=import \
        --go-grpc_opt=module=github.com/farhanaltariq/proto \
        $entry
done

for entry in $(find go/ -iname "*.go"); do
    protoc-go-inject-tag -input=$entry
    sed -i '/\/\/ versions:/d' $entry && sed -i '/\/\/ 	protoc/d' $entry
done

# Define the parent directory containing the subdirectories
PARENT_DIR="./go"  # Adjust this to your parent directory

# Loop through each directory in the parent directory
for dir in "$PARENT_DIR"/*/; do
    if [ -d "$dir" ]; then  # Check if it's a directory
        echo "Entering directory: $dir"
        cd "$dir" || { echo "Failed to enter directory: $dir"; continue; }
        
        # Run mockery --all in the current directory
        mockery --all
        
        # Return to the parent directory
        cd - > /dev/null
    fi
done

cd $dir
go mod tidy