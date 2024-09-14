# mysqldatabase

## [Docker](https://docs.docker.com/reference/cli/docker/)

[Start container](https://docs.docker.com/reference/cli/docker/compose/up/)

```sh
docker compose up -d db
```

[Stop container](https://docs.docker.com/reference/cli/docker/compose/down/)

```sh
docker compose down db
```

[Restart container](https://docs.docker.com/reference/cli/docker/compose/restart/)

```sh
docker compose restart db
```

## `db` container

[Execute a command in the container](https://docs.docker.com/reference/cli/docker/container/exec/), `sh` shell

```sh
docker exec -it db sh
```

[Invoke `mysql`](https://dev.mysql.com/doc/refman/8.4/en/mysql.html)

```sh
mysql --user=root --password=password
```

## [MySQL](https://dev.mysql.com/doc/refman/8.4/en/sql-statements.html)

[List databases](https://dev.mysql.com/doc/refman/8.4/en/show-databases.html)

```sql
SHOW DATABASES;
```

[Set `database` as the default database](https://dev.mysql.com/doc/refman/8.4/en/use.html)

```sql
USE database;
```

[List `database` tables](https://dev.mysql.com/doc/refman/8.4/en/show-tables.html)

```sql
SHOW TABLES;
```

[Describe `users` table structure](https://dev.mysql.com/doc/refman/8.4/en/describe.html)

```sql
DESCRIBE users;
```

[Retrieve rows from `users`](https://dev.mysql.com/doc/refman/8.4/en/select.html)

```sql
SELECT * FROM users;
```
