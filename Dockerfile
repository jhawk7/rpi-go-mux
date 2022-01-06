# inherit from the Go official Image
FROM arm32v7/golang:1.17.5-alpine3.15

# set a workdir inside docker
WORKDIR ~/go/src/github.com/jhawk7/rpi-go-mux

# copy . (all in the current directory) to . (WORKDIR)
COPY . .

# run a command - this will run when building the image
RUN go build -o rpi-go-mux

# the port we wish to expose
EXPOSE 8888

# run a command when running the container
CMD ./rpi-go-mux
