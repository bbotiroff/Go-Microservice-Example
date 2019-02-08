
initdb: 
	docker cp config/db_init.cql containerName:/
	docker exec -it containerName cqlsh -f db_init.cql

rundb:
	# TODO: add if the containerName exists docker rm containerName
	docker run -p :9042:9042 --name containerName cassandra