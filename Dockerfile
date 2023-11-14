FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-torq"]
COPY baton-torq /