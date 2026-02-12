
- goose -dir "db/migrations" create create_user_table sql

- change the file 
```
-- +goose Up
CREATE TABLE post (
    id int NOT NULL,
    title text,
    body text,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE post;
```

- goose -dir "db/migrations" postgres "host=127.0.0.1 user=minhaz_hossain password=12345 dbname=auth_dev port=5432 sslmode=disable TimeZone=UTC" up


