## SW-SYSTEM

#### Setup before build image and run container

1. change **sample.env** to **.env**
2. fill all environment variables in .env file
3. make **sql/init.sql** has same db name in **.env** file to create db for api server

then you can choose **Docker run** or **docker-compose up** to create a container.

#### Docker image build:

```shell
docker build -t sw-system-backend .
```

#### Create a docker container :

##### Docker run:

run

```shell
docker run -d --name sw-system-backend -p 3000:3000 sw-system-backend
```

stop and remove container

```shell
docker stop sw-system-backend
docker rm sw-system-backend
```

[comment]: <> (> if you want to connect to local postgres server change environment variable **DB_HOST** to **host.docker.internal** in **.env** file)

##### Docker compose up:

run

```shell
docker-compose up -d
```

stop and remove container

```shell
docker-compose down -d
```
