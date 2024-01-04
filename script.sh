# jalanin container dari mysql
docker run --name mysql82 -p 3306:3307 -v /my/custom:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=secret -d mysql:8.2