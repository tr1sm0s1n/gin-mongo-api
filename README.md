# Gin-Mongo-API

Gin API for CRUD operations in MongoDB.

## 🛠 Built With

<div align="left">
<a href="https://go.dev/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/go.svg" width="36" height="36" alt="Go" /></a>
<a href="https://gin-gonic.com/docs/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/gin.svg" width="36" height="36" alt="Gin" /></a>
<a href="https://www.mongodb.com/" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/DEMYSTIF/DEMYSTIF/main/assets/icons/mongodb.svg" width="36" height="36" alt="MongoDB" /></a>
</div>

## ⚙️ Run Locally

Clone the project

```bash
git clone https://github.com/DEMYSTIF/gin-mongo-api.git
cd gin-mongo-api
```

Start the database

```bash
docker compose up -d
```

View the database (optional)

```bash
docker exec -it gin-mongo mongosh -u root -p rootpw
```

Start the application

```bash
go run .
```

For live reload, install Air (optional)

```bash
go install github.com/cosmtrek/air@latest
```

Run the application with Air

```bash
air
```
