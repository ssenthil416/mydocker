## Sample : Two Containers sharing using volume

## Communication between server and client Containers
## Every min server will append a string to a file.
## Every two min client will read whole file and print

## Docker Image using Makefile and Dockerfile
>  make
/Library/Developer/CommandLineTools/usr/bin/make -C server
➭ Cleaning artifacts
➭ Building server statically
 ..done
➭ Building server Version 0-0
sha256:6982046734e99362fa514f50e3b250af417c114218bff67b1dda136b896c79f7
server:0-0
/Library/Developer/CommandLineTools/usr/bin/make -C client
➭ Cleaning artifacts
➭ Building client statically
 ..done
➭ Building client Version 0-0
sha256:90ce5fc56639f45bac5b0ad05223884ff375165b1813658678d8e8c79ae67f84
client:0-0


##Create docker-composer.yml file to launch docker image

//Launch docker image for server and client
>  docker-compose -f docker-compose.yml up
Recreating sharing_client_1 ... 
Recreating sharing_server_1 ... 
Recreating sharing_client_1
Recreating sharing_client_1 ... done
Attaching to sharing_server_1, sharing_client_1
server_1  | In Server = SUNDAR--2017-08-29 20:17:23.646962149 +0000 UTC
client_1  | In Client =  SUNDAR--2017-08-29 13:34:45.315389321 -0400 EDT
client_1  | In Client =  SUNDAR--2017-08-29 13:35:45.315567382 -0400 EDT
client_1  | In Client =  SUNDAR--2017-08-29 13:36:45.316250079 -0400 EDT
client_1  | In Client =  SUNDAR--2017-08-29 13:37:45.316709413 -0400 EDT
client_1  | In Client =  SUNDAR--2017-08-29 13:38:45.317263028 -0400 EDT
client_1  | In Client =  SUNDAR--2017-08-29 13:39:45.317924206 -0400 EDT
client_1  | In Client =  SUNDAR--2017-08-29 20:17:23.646962149 +0000 UTC
server_1  | In Server = SUNDAR--2017-08-29 20:18:23.648651087 +0000 UTC
server_1  | In Server = SUNDAR--2017-08-29 20:19:23.652799689 +0000 UTC
client_1  | In Client =  SUNDAR--2017-08-29 13:34:45.315389321 -0400 EDT
client_1  | In Client =  SUNDAR--2017-08-29 13:35:45.315567382 -0400 EDT


## Stop docker 
>  docker-compose -f docker-compose.yml down

## Remove docker image
> docker image rm <image name>(eg)<server:0-0>

