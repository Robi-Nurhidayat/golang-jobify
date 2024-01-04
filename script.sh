# jalanin container dari mysql
docker run --name mysql82 -p 3306:3307 -v /my/custom:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.2

# masuk ke mysql
docker run -it  mysql82 mysql root -p

# Membuat migrations
migrate create -ext sql -dir db/migrations -seq initial_schema

# menjalankan migrations
migrate -database "mysql://root:secret@tcp(localhost:3307)/jobify?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up