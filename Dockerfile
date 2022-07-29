FROM golang:alpine

WORKDIR /app

# copy Go modules and dependencies to image
COPY ./ /app

# download Go modules and dependencies
RUN go mod download

# compile application
RUN go build -o /go-webapi
 
# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start a container
CMD [ "/go-webapi" ]