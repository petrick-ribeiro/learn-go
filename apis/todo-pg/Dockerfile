FROM golang:alpine AS build

WORKDIR /app
ADD ./src/ /app/
RUN go build -o todo

FROM alpine
WORKDIR /todo_app
COPY --from=build /app /todo_app/
ENTRYPOINT [ "./todo" ]
