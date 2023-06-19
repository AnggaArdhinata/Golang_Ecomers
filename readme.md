# Simple Golang Ecommers

Hello, welcome to my repository.

## 🔥 Showcase
- [Live-Project](link)
- [Postman Docs](https://documenter.getpostman.com/view/20559835/2s93sjToKu)


## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/AnggaArdhinata/Golang_Ecomers.git
```

2. Install dependencies

```bash
go get -u ./...
```

3. Database Setup
Unfortunately this App does not support migration, you must create database manually, in this case, i using Postgresql, and then you can run Sql statement on sql file, that called INIT_DATABASE.

4. Environment Variables
First of all you can rename .env_example with .env, and then you can replace the configuration with your own configuration.

5. Seeding
This app does not support database seeding, and again, you can run the sql file that called DUMMY_DATA and replace email with real email to running email notif feature.
Or you can create data manually with Postman or etc.

6. Run the app

```bash
go run server.go
```

🌟 You are all set!

## Contributor
- [Muhammad Angga Ardhinata](https://github.com/AnggaArdhinata): as Developer

