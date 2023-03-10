docker compose up -d
docker build -t stats_image .
docker run -d -p 3000:3000 -p 6379 --name stats_container stats_image