FROM golang:1.13
RUN mkdir src/blogServer
ADD . src/blogServer

WORKDIR src/blogServer
# RUN ls

# RUN go get "github.com/gorilla/mux"
# RUN go get "github.com/jinzhu/gorm"
# RUN go get "github.com/jinzhu/gorm/dialects/sqlite"

RUN go build -o main .

RUN ls

# ENTRYPOINT [ "/go/src/blogServer/main", "host" ]
CMD ["/go/src/blogServer/main", "host"]