### DB Container URL

environment|url
--|--
host|localhost
docker run|ip
docker-compose|contaier name


### up db container

```
$ docker-compose up db
```

### up go container

```
$ docker-compose build api
$ docker-compose run api dep init
```

### run api container program

```
$ docker-compose run
```

### run host program

```
$ ./node_modules/.bin/ts-node index.ts
```