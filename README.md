# Hacktiv8 Golang Course - Assignment 2
## Package Used

- [mux](https://github.com/gorilla/mux)
```
go get -u github.com/gorilla/mux
```

- [gorm](https://gorm.io)
```
go get -u gorm.io/gorm
```

- [postgresql driver](https://gorm.io/driver/postgres)
```
go get -u gorm.io/driver/postgres
```

## Database
```
// file docker-compose.yml
// diisi dengan credential postgresql sendiri

version: "3.9"
services:
  database:
    container_name: database
    image: postgres:12.8
    restart: always
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=example
      - POSTGRES_DB=orders_by
    ports:
      - "5432:5432"
    volumes:
      - db:/var/lib/postgresql/data

volumes:
  db:
```

## API
API yang diserve menggunakan ```localhost :8080```
