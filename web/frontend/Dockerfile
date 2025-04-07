# Build React app
FROM node:20 AS builder

WORKDIR /app

COPY . .

RUN yarn install
RUN yarn build

# Serve with nginx
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
