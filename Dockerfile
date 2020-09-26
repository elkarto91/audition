FROM golang:1.14
ADD assets src/audition/assets
ADD main/main src/audition/main
WORKDIR src/audition
RUN ls
CMD ["./main"]