# [WIP] Larking

| “Now Bert, none of your larking about.” - Mary Poppins

Reflective gRPC gateway proxy.

- [Transcoding protobuf descriptors REST/HTTP to gRPC](https://cloud.google.com/endpoints/docs/grpc/transcoding)
- [Follows Google API Design principles](https://cloud.google.com/apis/design)
- [Dynamically load descriptors via gRPC server reflection](https://github.com/grpc/grpc/blob/master/doc/server-reflection.md)

### Debugging

Checkout [protobuf](https://github.com/golang/protobuf) at the latest v2 relase.
Go install each protoc generation bin.

Regenerate protoc buffers:

```
larking$ protoc -I=. --go_out=:. --go-grpc_out=:. grpc/reflection/v1alpha/*.proto

protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. testpb/test.proto

bazel run //:gazelle -- update-repos -from_file=go.mod
```

### Protoc

Just link API to `/usr/local/include/google` so protoc can find it.
```
ln -s ~/src/github.com/googleapis/googleapis/google/api/ /usr/local/include/google/
```
