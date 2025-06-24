FROM node:20 AS build
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM alpine:latest
WORKDIR /dist
COPY --from=build /app/build .
