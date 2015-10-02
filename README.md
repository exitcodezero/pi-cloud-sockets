# pi-cloud-sockets

Database Migrations
====================

Database migrations are handled by [Flyway](http://flywaydb.org/) and files are stored in the `/sql` directory. Migrations are automatically applied when running tests with Nose. You can run migrations manually in the development environment using `docker-compose` too. The included script `local_migrate.sh` uses an environment variable created by [Docker Compose](https://docs.docker.com/compose/env/) to find the IP address assigned to the database and execute the Flyway command to run migrations.

```
docker-compose run flyway ./local_migrate.sh
```
