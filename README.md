# Microservices example
Simple authentication example app built using microservices architecture.
<br />
It has few services and client app which consumes these services.
<br />
Services are built using: **Go, Kafka, Redis, PostgreSQL**.
<br />
Client is simple **ReactJS** application.


### Quick start
To start the app quickly with default configurations follow the next steps:

1. Create **email-service.env** file in [./services/email](./services/email)
```
# SMTP #
SMTP_USER=user_email_address
SMTP_PASS=user_email_password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

2. Run ``docker-compose up`` to start all services and client app.

3. Go to client app and try all services on address: ``http://localhost:3000``


### Start with custom configuration <!-- omit in toc -->
To start the app with custom configurations follow the next steps:

1. Create **email-service.env** file in [./services/email](./services/email)
```
# SMTP #
SMTP_USER=user_email_address
SMTP_PASS=user_email_password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587

# KAFKA #
KAFKA_TOPIC=verify-user
KAFKA_CLIENT_ID=kafka-client-id
KAFKA_BROKER_HOST=kafka
KAFKA_BROKER_PORT=9092
```
2. Create **redis-storage-service.env** file in [./services/redis-storage](./services/redis-storage)
```
# KAFKA #
KAFKA_REG_TOPIC=register-user
KAFKA_VER_TOPIC=verify-user
KAFKA_CLIENT_ID=kafka-client-id
KAFKA_BROKER_HOST=kafka
KAFKA_BROKER_PORT=9092

# REDIS #
REDIS_HOST=redis
REDIS_PORT=6379
```
3. Create **user-service.env** file in [./services/user](./services/user)
```
# USER SERVICE #
PORT=8080

# KAFKA #
KAFKA_REG_TOPIC=register-user
KAFKA_VER_TOPIC=verified-user
KAFKA_CLIENT_ID=kafka-client-id
KAFKA_BROKER_HOST=kafka
KAFKA_BROKER_PORT=9092

# REDIS #
REDIS_HOST=redis
REDIS_PORT=6379
```
4. Create **user-storage-service.env** file in [./services/user-storage](./services/user-storage)
```
# USER STORAGE SERVICE #
DB_DRIVER=postgres
DB_HOST=database
DB_PORT=5432
DB_NAME=kafka-example
DB_USER=admin
DB_PASS=admin

# KAFKA #
KAFKA_VER_TOPIC=verified-user
KAFKA_CLIENT_ID=kafka-client-id
KAFKA_BROKER_HOST=kafka
KAFKA_BROKER_PORT=9092
```
5. Create **auth-service.env** file in [./services/auth](./services/auth-service)
```
# AUTH SERVICE #
PORT=8081

# REDIS #
REDIS_HOST=redis
REDIS_PORT=6379

JWT_SECRET=supersecret

# DB #
DB_DRIVER=postgres
DB_HOST=database
DB_PORT=5432
DB_NAME=kafka-example
DB_USER=admin
DB_PASS=admin
```

6. Run ``docker-compose up`` to start all services and client app.

7. Go to client app and try all services on address: ``http://localhost:3000``

**Note:**
If you try to use your gmail account to send email, most probably it will fail with error "Username and Password not accepted". You need to enable less secure apps. Check out this link: https://hotter.io/docs/email-accounts/secure-app-gmail/
