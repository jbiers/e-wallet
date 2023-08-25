# e-wallet

- created database schema on https://dbdiagram.io/
![Alt text](image.png)

- run postgresql with docker
    ```
        docker pull postgres:15.4-alpine
        docker run --name database -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=root -p 5432:5432 -d postgres:15.4-alpine
        docker exec -it database psql -U root
    ```
