# png2ascii API

## Overview

An HTTP service that accepts an image, and creates an ASCII version of that image
that can be retrieved later. 

### Features: 

- [x] `GET /images` returns a list of all available images 
- [x] `GET /images/<image_id>` returns the ASCII representation of an image
- [x] `POST /images` accepts a new PNG image to be converted to ASCII
- [x] `POST /images` returns a newly minted ID for the uploaded image
- [x] ASCII-fied Data Persistence as files on Disk 

## Implementation

The `png2ascii-api` uses the `Fiber` web framework to capture PNG 
image files and ASCII-fy them (with `image2ascii`), responding with a unique 
identifier that can be used to retrieve the ASCII-fied image. The ASCII images 
are persisted, so if the service is stopped and restarted, the data should 
still be available.

## Usage 

* `make build` to build a Docker container 
* `make service` to start the service on port 8080 -- Note: host & port is hard-coded
* `make test` to run API feature tests -- includes testing ASCII-fying and directory existence

## Improvements

- Testing could be improved (i.e start the server and then perform tests)
- Logging needs work 
- Concurrent requests could be implemented 
- Storing temporary image files is not super efficient and could be improved
- Cache management could be implemented for obtaining all available images / requested image

## Resources

* [Go image library](https://golang.org/pkg/image/)
* [image2ascii library](https://github.com/qeesung/image2ascii)
* [Fiber library](https://github.com/gofiber/fiber)
* [uuid library](https://github.com/google/uuid)
