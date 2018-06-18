FROM node:alpine AS build

WORKDIR /usr/track-expenses

ADD *.json ./
ADD *.lock ./
RUN yarn install

ADD . .
RUN yarn run lint
RUN yarn run build --configuration=production

FROM nginx:alpine

WORKDIR /usr/share/nginx/html
COPY --from=build /usr/track-expenses/dist /usr/share/nginx/html/

EXPOSE 80
