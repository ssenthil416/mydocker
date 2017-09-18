## Sample to create Hello World  Docker image. 

## Docker Image using Makefile and Dockerfile
> make
➭ Cleaning artifacts
➭ Building myhelloworld statically
 ..done
➭ Building myhelloworld Version 0-0
sha256:2edc4f6946f44b9dc5d6d56ce44e4cd2798f5d9533a5d02e51036f5cae3b94f5
myhelloworld:0-0


##Create docker-composer.yml file to launch docker image

//Launch docker image 
>  docker-compose -f docker-compose.yml up
Creating helloworld_hello_1 ... 
Creating helloworld_hello_1 ... done
Attaching to helloworld_hello_1
hello_1  | My Hello World!!!
helloworld_hello_1 exited with code 0

## Stop docker 
>  docker-compose -f docker-compose.yml down

## Remove docker image
> docker image rm <image name>(eg)<myhelloworld:0-0>

