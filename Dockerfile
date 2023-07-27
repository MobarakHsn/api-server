FROM golang:alpine AS builder
RUN mkdir /api-server
COPY . /api-server
WORKDIR /api-server
RUN go build .

FROM alpine
WORKDIR /api-server
COPY --from=builder /api-server/ /api-server/

# Expose port 8080 to the outside world
#EXPOSE 8081

# Command to run the executable
ENTRYPOINT ["./api-server"]
CMD ["serve"]