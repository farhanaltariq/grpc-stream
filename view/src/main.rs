// src/main.rs

use proto_files::map::A; // Import MapRequest from the generated module
use proto_files::protomessagergateway::MapResponse; // Import MapResponse from the messager module

fn main() {
    // Now you can use the `map` and `messager` modules here
    let map_data = A::default(); // Use the correct method to instantiate
    let messager_data = MapResponse::default(); // Instantiate the message

    println!("{:?}", map_data);
    println!("{:?}", messager_data);
}
