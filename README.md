#** API Golang **

# Objetivos:

Desenvolvimento de uma API bancária em linguagem Golang com banco de dados Postgres em container Docker.

# Pré requisitos:

    Linguagem Golang: https://go.dev/doc/install
    Postgresql: https://www.postgresql.org/download/
    Docker:https://docs.docker.com/get-docker
    Postman

# Setup

    Clonar o repositório git clone:https://github.com/Crisscoelho/API.git

    Instalar a imagem do docker com o seguinte comando: docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres
    Rodas as migrates com o seguinte comando: migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up
    Executar o comando: docker start postgres
    Executar o comando: go run main.go

# Banco de dados postgres

Banco de dados postgre(https://dbdiagram.io/d/6541641c7d8bbd64653a4ba2


#Busca todos os bancos
<imagens src=/assets/imagens.png/Get banks.png">

#Busca todas as contas
<imagens src=/assets/imagens.png/Get accounts.png">

#Busca todas os clientes
<imagens src=/assets/imagens.png/Get customers.png">

