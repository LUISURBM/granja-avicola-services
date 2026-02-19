curl -X GET http://localhost:8080/api/v1/lotes
docker exec -it mariadb mariadb -uroot -ppassword gestion_avicola

docker logs go_backend_avicola
docker compose down
docker compose build --no-cache api-avicola
docker compose up -d

docker compose up -d --build api-avicola