FROM node:9-alpine
RUN apk update && apk add --no-cache libstdc++ libgcc tini
ADD ./package.json ./package-lock.json ./
RUN npm install
ENV PORT 7878
EXPOSE ${PORT}
ENTRYPOINT ["/sbin/tini", "--"]

CMD ["/node_modules/http-echo-server/bin/http-echo-server"]
