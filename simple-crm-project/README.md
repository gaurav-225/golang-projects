In this project, I have utilized **Fiber** and **GORM** to abstract and streamline our development process.
**Fiber** handles routing, allowing us to define endpoints. 
**GORM** provides an interface for performing database operations, eliminating the need to write explicit SQL queries. Instead, I used GORM’s built-in functions such as:

- `Create` – for inserting records into the database.
- `Delete` – for removing records.
- `Find(&entity, id)` – for retrieving a single record by its ID.
- `Find(&entities)` – for retrieving multiple records.


# Install local sqlite in server

```bash
# ubuntu server
sudo apt-get install sqlite3 libsqlite3-dev gcc
```

# Run
```bash
go run main.go
```

> I followed the YouTube tutorial: [Go Project For Absolute Beginners Using Go-Fiber - Simple CRM](https://www.youtube.com/watch?v=3JtZqqrJFmM)

