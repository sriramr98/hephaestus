Hephaestus

### Setup Nuances

1. If there's any change in `server/db/init.sql`, remove `postgres_data` folder from root, otherwise `docker-compose up` will not re-initialize the new changes
