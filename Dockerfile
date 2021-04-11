# inherit from the Go official Image
FROM hypriot/rpi-golang

# set a workdir inside docker
WORKDIR ~/go/src/github.com/jhawk7/rpi-go-mux

# copy . (all in the current directory) to . (WORKDIR)
COPY . .

# run a command - this will run when building the image
RUN go get -u github.com/gorilla/mux; exit 0
RUN go build -o rpi-go-mux; exit 0

# the port we wish to expose
EXPOSE 8888

# run a command when running the container
CMD ./rpi-go-mux
