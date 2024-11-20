# API Básica

## O que é

> Uma API simples construída em golang usando [GORM](https://gorm.io), [gin-gonic](https://gin-gonic.com) e [swaggo](https://github.com/swaggo/swag).

Desenvolvida para fins avaliativos na disciplina de DevOps do Instituto Federal de Alagoas durante o curso de Bacharelado em Sistemas de Informação (2024.2) ministrada pelo professor mestre Italo Carlos Lopes Silva.

O conjunto pode ser dividido (como pode ser observado no [compose](./compose.yml)) entre a API e o banco de dados (PostgreSQL)

## API

Usando o gin-gonic para dividir os endpoints e o GORM para gerenciar o acesso ao banco de dados a API recebe as variáveis de ambiente que estão no [exemplo](./.env.example) e monta a string de conexão com o banco de dados no arquivo [driver.go](./src/driver/driver.go), sendo dependente do banco de dados mas ainda iniciando antes que o ultimo estivesse propriamente configurado, foi necessário criar um _health check_ para o banco de dados, evitando a situação.

### Modelos

A API é bastante simples possuindo apenas um modelo de `Food` o qual é mapeado via GORM para o banco de dados.

### Endpoints

A API possui apenas 3 endpoints (4 se o redirect for considerado), sendo eles:

#### **GET** /api/foods

- Retorna a lista de comidas cadastradas no banco de dados, com paginação.
- Possui dois parâmetros de Query:
  - `perPage`: Define a quantidade de itens a ser retornada em uma página (int).
  - `page`: Define a página a ser retornada (com base na quantidade de items do `perPage`) (int).

#### **POST** /api/foods

- Adiciona uma nova comida ao banco de dados e a retorna em caso de sucesso.
- Recebe no corpo da requisição:
  - `name`: O nome da nova comida.
  - `price` O preço da nova comida.

#### **PUT** /api/foods/:id

- Altera os dados de uma comida no banco de dados e retorna as informações atualizadas em caso de sucesso.
- Recebe no _path_:
   - `id`: ID da comida a ser atualizada (uuid)
- Recebe no corpo da requisição:
  - `name`: O nome da nova comida (string).
  - `price` O preço da nova comida (float).

#### **DELETE** /api/foods/:id

- Remove uma comida do banco de dados.
- Recebe no _path_:
   - `id`: ID da comida a ser removida (uuid)

#### **GET** /api/docs/index.html

- Possui a documentação da API através do SwaggerUI, construído pelo middleware do swaggo.

#### **GET** /api/docs

- Redireciona para `/api/docs/index.html`

## Banco de Dados

Usando a imagem do postgres como base tive que criar um script para criar o usuário pelo qual a API poderia acessar, evitando privilégios de administrador, para tal o [script de belpro-ci](https://gist.github.com/beldpro-ci/bc8d1a48f6a012a1b494460aac84796a#file-01-filladb-sh) foi de grande ajuda, embora tenha sido necessário algumas pequenas alterações para o funcionamento correto nessa aplicação.

Nesse script é criado um novo usuário (conforme definido nas variáveis de ambiente) e garantido a ele privilegio de operar em um banco de dados especifico (junto ao esquema public desse mesmo banco de dados).

## O compose

### Serviços

#### Server

Usando a imagem `golang:alpine3.20` como base para _buildar_ o projeto e a `alpine:3.20` sua [Dockerfile](./docker/server/Dockerfile) pode ser resumida no processo de:

1. Estágio **builder**:
   1. Copiar os arquivos que detalham as dependências.
   2. Realizar o download dessas dependências.
   3. Copiar o código fonte.
   4. Realizar o _build_ do código fonte para o binário `app` (uma das vantagens de usar golang).
2. Estágio Final
   1. Copiar o binário gerado no estágio anterior
   2. Configurar o binário para execução ao iniciar o container.

#### Database

O banco de dados já veio quase que completamente configurado através da imagem `postgres:16-alpine` mas como foi solicitado o uso de um usuário exclusivo e sem privilégios de administrador para a aplicação se fez necessário o uso de um [script shell](./docker/database/001_init.sh), para tal foi criada uma nova [Dockerfile](./docker/database/Dockerfile) na qual simplesmente o script é copiado para o diretório `/docker-entrypoint-initdb.d/` do container, que executa automaticamente durante a inicialização.

O script é dividido em duas funções:

- `check_env_vars_set()` que como o nome diz, verifica se todas as variáveis de ambiente que são usadas existem.

e

- `init_user_and_db()` que usa o comando `psql` para acessar ao banco como administrador e:
  1. Criar o usuário que vai ser usado pela aplicação com uma senha predefinida (pelo `.env`).
  2. Criar o banco de dados no qual a aplicação salvará os dados.
  3. Garantir todos os privilégios do novo banco de dados para o usuário da aplicação.
  4. Acessar o banco de dados criado no passo 2 como administrador e garantir acesso ao `schema public` (para que o gorm possa criar as tabelas necessárias).

Como explicado anteriormente o health check foi necessário para evitar o início prematuro do serviço da API, sendo implementado pela função _build in_ do PostgreSQL `pg_isready`.

### Rede

Para isolar o container do banco de dados (database), foi criada a rede `api-network` que engloba os dois containers, permitindo a comunicação entre eles, mas proibindo acesso de elementos fora dessa rede.

### Volume

O único volume que se fez necessário foi o nomeado `api-data` que mapeia para `/var/lib/postgresql/data` do container database, esse diretório é onde se encontra os dados salvos no banco de dados, fazendo com que eles possam persistir entre execuções dos containers.

## Como Rodar

1. Clone esse repositório em seu diretório de preferência.
2. Dentro da pasta do repositório, copie o arquivo [.env.example](./.env.example) para `.env`
   - Altere as variáveis de ambiente conforme necessário (exceto a `POSTGRES_HOST`, caso vá rodar via docker).
3. Execute o comando:
   ```bash
   docker compose up -d
   ```
   - O `-d` é opcional, apenas indica ao docker para executar os containers em segundo plano.
4. Acesse o SwaggerUI em http://localhost:8080/api/docs/index.html.
