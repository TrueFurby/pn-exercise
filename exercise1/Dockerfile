FROM golang

WORKDIR /go/src/getweather
COPY getweather.go .
RUN go get && go install

CMD getweather
