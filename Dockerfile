FROM golang:1.21 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o /go/bin/datum

FROM gcr.io/distroless/static

# Copy the binary that goreleaser built
COPY --from=build /go/bin/datum /datum

# Run the web service on container startup.
ENTRYPOINT ["/datum"]
CMD ["serve","--debug","--pretty","--dev"]