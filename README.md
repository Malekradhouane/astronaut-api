#Test Astronaut

## Developpement

We use a `docker-compose.yaml` file to launch app dependency dockers like :
```
docker-compose up --build
```

You will have to create an `.env` file from the` .env.example` template and put your parameters there for the dev.

Run the following command to launch dependencies for local dev:

```
make deploy
```

This one you tearn down the dependencies:

```
make teardown
```


## Endpoints

### Health check

- `GET /status`

  Example response:

  ```json
  {
    "status": "ok"
  }
  ```
- ``POST /astronauts ``

To insert astronaut to DB

  ```
  
      {
          "firstname": string,
          "lastname": string,
          "email": string
      }
  
  ```

- ``Get /astonauts ``
  To get all astronauts from DB

- ``Get /astronauts/:id ``
  To get astronaut by id from DB

- ``PATCH /astronauts/:id ``

To update astronaut

  ```
  
      {
          "firstname": string,
          "lastname": string,
          "email": string
      }
  
  ```


- ``Delete /astronauts/:id ``
  Delete astronaut by id

## Tests

Unit tests:

```sh
go test ./...
```