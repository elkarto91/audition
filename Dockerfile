FROM golang:1.13.4-alpine3.10
ADD assets src/audition/assets
ADD main/main src/audition/main
WORKDIR src/audition
RUN ls
EXPOSE 8080