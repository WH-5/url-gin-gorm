# postgres
make_postgres:
	docker run --name url-pg \
	-e POSTGRES_USER=ROOT \
	-e POSTGRES_PASSWORD=whwhwhwhwhwh1231 \
	-e POSTGRES_DB=MAN \
	-p 5432:5432 \
	-d postgres

# redis
make_redis:
	docker run --name=url-rd \
	-p 6379:6379 \
	-d redis \
	redis-server --requirepass whwhwhwhwhwh1321



make_db: make_postgres make_redis

start_pg:
	docker start url-pg

start_rd:
	docker start url-rd

start_db: start_pg start_rd
