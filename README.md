# addchan
Bot-provided command to add a new text channel

## Setup
* `./docker-compose.yaml`
    ```yaml
    version: "3"
    services:
      app:
        image: ghcr.io/teamecri/addchan:latest
        restart: always
        volumes:
          - type: bind
            source: ./.env
            target: /app/.env
    ```
* `./.env`
    ```
    TOKEN=YOUR_TOKEN_GOES_HERE
    ```
* Then, run `docker-compose up -d`.

## License
MIT license.