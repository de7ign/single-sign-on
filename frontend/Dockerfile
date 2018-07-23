FROM node:8.11-alpine
ENV NODE_ENV production
WORKDIR /usr/src/app

RUN npm install -g serve
CMD serve -s build
EXPOSE 3000

COPY ["package.json", "package-lock.json*", "npm-shrinkwrap.json*", "./"]
RUN npm install --production --silent
COPY . .
RUN npm run build --production
RUN mv node_modules ../