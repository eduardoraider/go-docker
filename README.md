# Using Distroless Docker Image for Go Applications

How to reduce the size of your Docker image to a few megabytes by using the Distroless base image for running your Go application.

## Dockerfile Example

Check the Dockerfile in this repository for an example of how to reduce the size of your Go application Docker image using the Distroless base image.
## Steps

1. **Build the Image**:
   
   ```bash
   docker build . -t go-server-distroless
   ```

2. **Check Image Size**:

   ```bash
   docker images
   ```

3. **Run the Container**:

   ```bash
   docker run -p 8080:8080 go-server-distroless
   ```

4. **Check the Endpoint**:

   ```bash
   curl http://localhost:8080
   ```

## Explanation

[Distroless](https://github.com/GoogleContainerTools/distroless) is a set of Docker images that are designed to be as minimal as possible. These images contain only your application and its runtime dependencies, without any unnecessary packages or operating system libraries. By using Distroless, you can significantly reduce the size of your Docker images, making them smaller and more secure.

In this example, we use the Distroless base image to build a Docker image for a Go server application. By following the steps above, you can build and run the Docker image, and then test the endpoint to verify that the application is working correctly.

---

#### by Eduardo Raider
🛠 🥋 **Software Engineer**