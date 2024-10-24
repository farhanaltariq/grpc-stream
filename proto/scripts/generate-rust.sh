#!/run/current-system/sw/bin/bash

# Get the current directory of the script
dir="$(dirname $(realpath ${BASH_SOURCE[0]}))" && cd $dir
dir=$dir/..

# Move to the root directory
cd $dir

# Check if `protoc` is installed
if ! [ -x "$(command -v protoc)" ]; then
    echo "protoc is required"
    exit 1
fi

# Check if `protoc-gen-prost` is installed
if ! [ -x "$(command -v protoc-gen-prost)" ]; then
    echo "protoc-gen-prost is required, installing..."
    cargo install protoc-gen-prost
fi

# Directory to output generated Rust files (external folder)
rust_protoc_out="$dir/rust"

# Ensure the output directory exists
rm -rf "$rust_protoc_out" && mkdir -p "$rust_protoc_out"

echo "Using protoc version" $(protoc --version)

# Iterate through all .proto files and generate Rust code using `prost`
for entry in $(find ./ -iname "*.proto"); do
    # Extract the base name of the proto file (without the .proto extension)
    base_name=$(basename "$entry" .proto)

    # Create a separate directory for this proto file's output
    output_dir="$rust_protoc_out/$base_name"
    mkdir -p "$output_dir"

    # Generate Rust code in the specific output directory
    protoc --prost_out="$output_dir" "$entry"
done

echo "Protobuf generation complete! Files generated in $rust_protoc_out"

# Optional: If you want to build a Rust project in the current directory
# Ensure the Rust module uses the correct dependencies
if [ -f "Cargo.toml" ]; then
    cargo build
    cargo fmt

    # Optional cleanup
    cargo clean
else
    echo "No Cargo.toml found. Skipping cargo build."
fi

# Clean up Rust modules (format the code)
