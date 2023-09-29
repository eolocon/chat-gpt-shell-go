# chat-gpt-shell-go

## Running locally
- [install go](https://go.dev/doc/install)
- open a command prompt
- execute the following command from the project root directory:  
    `OPENAI_AUTH_TOKEN=<your-token> go run ./app`

## Running with Docker
- [install Docker Desktop](https://docs.docker.com/get-docker/) or [install the Docker Engine](https://docs.docker.com/engine/install/)
- open a command prompt
- execute the following commands from the project root directory:  
    ```
    docker build -t <name-you-want-to-give-to-the-image>
    docker run --rm -it -e OPENAI_AUTH_TOKEN=<your-token> <name-you-want-to-give-to-the-image>
  ```
