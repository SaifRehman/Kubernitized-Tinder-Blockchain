# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
ENV SRC_DIR=/Users/saifrehman/go/src/github.com
ADD . $SRC_DIR
EXPOSE 1234
RUN cd $SRC_DIR; go build -o app; cp app /app/
ENTRYPOINT ["./app"]