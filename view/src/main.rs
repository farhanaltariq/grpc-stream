pub mod protomessagergateway {
    tonic::include_proto!("protomessagergateway"); // Match this with the package name in your .proto file
}
use protomessagergateway::map_service_client::MapServiceClient;
use protomessagergateway::{MapRequest, MapResponse};
// use tonic::transport::Channel;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Connect to the gRPC server
    let mut client = MapServiceClient::connect("http://localhost:50051").await?;

    // Create a request
    let request = tonic::Request::new(MapRequest {
        alt: "foo".to_string(),
        lat: "42".to_string(),
        lon: "42".to_string(),
    });

    // Call the gRPC method
    let response: MapResponse = client.map(request).await?.into_inner();

    // Print the response
    println!("Response: {:?}", response);

    Ok(())
}