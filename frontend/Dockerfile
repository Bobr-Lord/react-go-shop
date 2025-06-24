FROM node:20 AS build
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM alpine:latest
RUN apk add --no-cache bash coreutils
WORKDIR /app
VOLUME /usr/share/nginx/html
COPY --from=build /app/build /tmp/build

CMD ["sh", "-c", "cp -r /tmp/build/* /usr/share/nginx/html && tail -f /dev/null"]
