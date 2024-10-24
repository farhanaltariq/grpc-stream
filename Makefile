run-api:
	cd api && go run .
run-view:
	cd view && cargo run
generate:
	cd proto && make generate