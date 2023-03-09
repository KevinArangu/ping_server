docker compose up -d
docker build -t stats .
docker run -d -p 3000:3000 -p 6379 --env PORT_HTTP=3000 stats