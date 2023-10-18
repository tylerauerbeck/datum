FROM gcr.io/distroless/static:nonroot

# `nonroot` coming from distroless
USER 65532:65532

# Copy the binary that goreleaser built
COPY  datum /datum

# Run the web service on container startup.
ENTRYPOINT ["/datum"]
CMD ["serve"]