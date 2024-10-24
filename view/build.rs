fn main() {
    tonic_build::compile_protos("../proto/source/messagergateway/messager.proto")
        .unwrap();
}