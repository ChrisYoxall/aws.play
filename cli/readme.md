
Don't want to install CLI on my machine unless needed.


There is an image on Docker Hub at https://hub.docker.com/r/amazon/aws-cli with instructions.


To shell into container do: docker run -it --rm --entrypoint /bin/bash amazon/aws-cli:2.13.11
 

To experiment, I installed Go in the CLI container by doing: yum update && yum install golang.
