
build:
	docker build -t png2ascii_api:latest .

service: 
	docker run -d -t -p 127.0.0.1:8080:8080 png2ascii_api:latest /server

test:
	docker run -d -t png2ascii_api:latest go test
