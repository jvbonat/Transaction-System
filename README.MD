# Desafio Transações

Este é o repositório do **Desafio Transações**, um sistema para gerenciar transações financeiras com GORM e MySQL.

## Tecnologias Utilizadas

- **Go**: Linguagem principal do projeto.
- **Gin**: Framework web usado para gerenciar as rotas e requisições HTTP. 
- **GORM**: ORM para interagir com o banco de dados MySQL.
- **MySQL**: Banco de dados utilizado para armazenar os dados do sistema em tabelas.
- **Dotenv**: Para carregar variáveis de ambiente (se necessário).

## Pré-requisitos

- **Go** (versão 1.18 ou superior).
- **MySQL** instalado e em execução.

## Configuração do Ambiente

1. **Clone o repositório**:
   Clone o repositório do projeto para o seu ambiente local.
   ```bash
   mkdir transactions-challenge
   cd transactions-challenge
   git clone https://github.com/jvbonat/Transaction-System.git
Configure o banco de dados MySQL: Certifique-se de que o MySQL está instalado e em execução na sua máquina.

2. **Configure o Banco de Dados MySQL**
Certifique-se de que o MySQL está instalado e em execução na sua máquina.
Crie um banco de dados chamado transactions:
CREATE DATABASE transactions;
Garanta que você tenha as credenciais de acesso corretas.

3. **Crie o Arquivo .env**
Crie um arquivo .env na raiz do projeto. O conteúdo do arquivo pode ser o seguinte:
- Em DB_PASSWORD ponha sua senha do MySQL Workbench ou outro software que esteja usando como base relacional de dados
DB_USER=root
DB_PASSWORD=
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=transactions

4. **Instale as Dependências**
Execute o comando abaixo para instalar as dependências do Go:

go mod tidy

**Como Rodar**
1. Execute a Aplicação
Após configurar o ambiente, basta rodar o seguinte comando na raiz do projeto para executá-lo:

go run main.go

### 📌 Inicialização do Banco de Dados

O projeto utiliza GORM para a modelagem do banco de dados. No primeiro uso, o sistema executa `db.AutoMigrate()` para criar as tabelas automaticamente.  

Após a criação inicial do banco, **recomenda-se comentar ou remover essa linha** depois de rodar o comando go run main.go para evitar tentativas desnecessárias de recriação das tabelas. Após parar a aplicação e comentar ou remover a linha, pode rodá-la novamente sem problemas.

Caso precise atualizar a estrutura do banco, descomente a linha.

📡 Endpoints da API
Aqui estão os principais endpoints da API:

Método	Endpoint	Descrição
POST	/accounts	Criar uma nova conta
GET	/accounts/:accountId	Buscar uma conta pelo ID
POST	/transactions	Criar uma transação financeira

🔹 Exemplo de Uso
Criar uma conta:

### Exemplo de Requisição para Criar uma Conta

Para criar uma nova conta, envie uma requisição POST para o endpoint `/accounts` com o número do documento no corpo da requisição:

curl --location "http://localhost:8080/api/accounts" \
     --header "Content-Type: application/json" \
     --data '{"document_number": "12345678900"}'


Criar um tipo de operação:

curl --location 'http://localhost:8080/api/operations' \
--header 'Content-Type: application/json' \
--data '{"description": "PAGAMENTO"}'


Criar uma transação:

- Essencial que o Account esteja criado previamente e o Operation_type_id também

curl --location "http://localhost:8080/api/transactions" \
     --header "Content-Type: application/json" \
     --data '{
           "account_id": 1,
           "operation_type_id": 4,
           "amount": 150.75
         }'

✅ Rodando os Testes

- Crie um .env dentro da pasta de testes

O projeto contém testes unitários de inserção para garantir o funcionamento correto da API. Para executar os testes, use:

go test -v -run TestCreateAccount | go test -v -run TestCreateTransaction

