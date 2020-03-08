FROM node:13

WORKDIR /usr/src/app
COPY package.json ./
RUN yarn

VOLUME /usr/src/app/src
VOLUME /usr/src/app/public

EXPOSE 3000
ENTRYPOINT yarn start
