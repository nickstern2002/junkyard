### Welcome to the Junkyard
This is a working web server ready to be built as a container image
and deployed into k8s cluster

### Build the Image
`podman build -t junkyard -f Dockerfile .`

### Confirm Successful Build
`podman images`

### Save as TarBall
`podman save --format docker-archive -o <tarball-name>.tar <image-path>:latest`

### Load into Kind Cluster
`kind load image-archive --name <kind-cluster-name> <tarball-name>.tar   `


### Testing the App in Cluster
From here, the image can be accessed

To test the service, deploy a Junkyard pod and run 

`k port-forward svc/junkyard-service 3333:3333`

From here the service can be accessed at `http://localhost:3333/`
