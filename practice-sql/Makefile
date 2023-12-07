run:
	@docker-compose up -d

db-connect:
	@docker exec -it practice-psql psql -U supervisor -d Company

db-backup:
	@docker exec -it practice-psql pg_dump -U supervisor -d Company -c -f tmp/dump.sql
	
copy-backup:
	@docker cp practice-psql:/tmp/dump.sql ./backup/dump.sql

db-full-backup: db-backup copy-backup

db-copy-backup-to-container:
	@docker cp ./backup/dump.sql practice-psql:/tmp/dump.sql

connect-container:
	@echo use to repair database
	@echo psql -U supervisor -d Company < /tmp/dump.sql
	@docker exec -it practice-psql bash
