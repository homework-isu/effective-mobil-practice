
-- без ошибки
begin;
INSERT INTO departments (title) VALUES ('Testing');
INSERT INTO projects (title) VALUES ('Project D');
commit;


-- с ошибкой
begin;
INSERT INTO departments (title) VALUES ('Testing');
INSERT INTO projects (title) VALUES (hello);
rollback;