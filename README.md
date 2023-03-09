docker compose up -d
docker build -t stats .
docker run --expose 3000 --env PORT_HTTP=3000 stats