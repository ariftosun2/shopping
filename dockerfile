FROM golang:1.16-alpine
WORKDIR /home/arif/dev/shopping/shopping-servis

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["./shopping-servis"]

#container calistirmak icin: docker run -p 8080:8081 -it my-shopping-servis
#container olusturmak icin:docker build -t my-shopping-servis . 