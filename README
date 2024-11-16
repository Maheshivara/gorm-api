# API Básica
## O que é
> Uma API simples construída em golang usando [GORM](https://gorm.io), [gin-gonic](https://gin-gonic.com) e [swaggo](https://github.com/swaggo/swag).

Desenvolvida para fins avaliativos na disciplina de DevOps do Instituto Federal de Alagoas durante o curso de Bacharelado em Sistemas de Informação (2024.2) ministrada pelo professor mestre Italo Carlos Lopes Silva.

O conjunto pode ser dividido (como pode ser observado no [compose](./compose.yml)) entre a API e o banco de dados (PostgreSQL)
## API
Usando o gin-gonic para dividir os endpoints e o GORM para gerenciar o acesso ao banco de dados a API recebe as variáveis de ambiente que estão no [exemplo](./.env.example) e monta a string de conexão com o banco de dados no arquivo [driver.go](./src/driver/driver.go), sendo dependente do banco de dados mas ainda iniciando antes que o ultimo estivesse propriamente configurado, foi necessário criar um [script](./docker/server/start.sh) para esperar alguns segundos antes de iniciar a API.
### Modelos
A API é bastante simples possuindo apenas um modelo de `Food` o qual é mapeado via GORM para o banco de dados.
### Endpoints
A API possui apenas 3 endpoints (4 se o redirect for considerado), sendo eles:
#### **GET** /api/foods
- Retorna a lista de comidas cadastradas no banco de dados, com paginação.
- Possui dois parâmetros de Query:
  - `perPage`: Define a quantidade de itens a ser retornada em uma página.
  - `page`: Define a página a ser retornada (com base na quantidade de items do `perPage`).
#### **POST** /api/foods
- Adiciona uma nova comida ao banco de dados e a retorna em caso de sucesso.
- Recebe no corpo da requisição:
  - `name`: O nome da nova comida.
  - `price` O preço da nova comida.
#### **GET** /api/docs/index.html
- Possui a documentação da API através do SwaggerUI, construído pelo middleware do swaggo.
#### **GET** /api/docs
- Redireciona para `/api/docs/index.html`
## Banco de Dados
Usando a imagem do postgres como base tive que criar um script para criar o usuário pelo qual a API poderia acessar, evitando privilégios de administrador, para tal o [script de belpro-ci](https://gist.github.com/beldpro-ci/bc8d1a48f6a012a1b494460aac84796a#file-01-filladb-sh) foi de grande ajuda, embora tenha sido necessário algumas pequenas alterações para o funcionamento correto nessa aplicação.

Nesse script é criado um novo usuário (conforme definido nas variáveis de ambiente) e garantido a ele privilegio de operar em um banco de dados especifico (junto ao esquema public desse mesmo banco de dados).
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