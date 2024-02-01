
FROM node:20.11.0-alpine as BUILDER_IMAGE

WORKDIR /app/web

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:1.25.3-alpine as SERVER_IMAGE

COPY --from=BUILDER_IMAGE /app/web/dist /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
