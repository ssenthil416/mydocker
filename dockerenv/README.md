## Sample to show how to pass env varaible and value to docker image

##Get a Docker image. Using Makefile and Dockerfile 
> make
➭ Cleaning... 
➭ Building dockernenv statically
 ..done
➭ Building dockernenv Version 0-0
sha256:634b2ff3948f649050242c16080a3ef2dd62ab6bc9db7d3fea32b81a4b06609f
dockernenv:0-0

##Create docker-composer.yml and with environment file to launch docker image

##Launch docker 
>  docker-compose -f docker-compose.yml up
Creating network "dockerenv_default" with the default driver
Creating dockerenv_hello_1 ... 
Creating dockerenv_hello_1 ... done
Attaching to dockerenv_hello_1
hello_1  | Hello Sundara Senthil
dockerenv_hello_1 exited with code 0

##Stop docker 
>  docker-compose -f docker-compose.yml down

##Remove docker image
> docker image rm <image name>(eg)<dockernenv:0-0>

