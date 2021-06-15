FROM alpine:3.13

COPY ./server /server
COPY ./manage /manage

ENV GIN_MODE release

RUN cd /server/cmd/heretics && go build
RUN go clean -cache

RUN cd /manage && npm install && npm run generate
RUN npm cache clean --force

ENTRYPOINT /app/cmd/heretics/heretics

EXPOSE 8080
