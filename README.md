# 🎲 Random Number Generator API

Este projeto é uma API REST escrita em Go que gera números aleatórios utilizando o algoritmo **Mersenne Twister**. Os números são persistidos em um banco de dados SQLite e protegidos por autenticação via API Key.

---

## 📦 Funcionalidades

- Geração de número aleatório com base em uma **seed (timestamp)**.
- Persistência dos dados (número, seed e data de criação) em banco de dados SQLite.
- **Validação de unicidade**: impede duplicações no banco.
- **Autenticação via API Key**.
- Testes unitários cobrindo geração, reprodutibilidade e duplicidade.

---

## 🧠 Justificativa da Escolha: Go

**Go foi escolhido por sua eficiência, simplicidade e desempenho** em aplicações back-end que exigem **processamento rápido**, **concorrência segura** e **entrega compacta**. Essas características o tornam ideal para a proposta deste projeto de **geração e verificação de números aleatórios**.


## ⚙️ Instalação e Execução

### ✅ Pré-requisitos

- Go 1.21+ instalado
- Git instalado
- Variável `CGO_ENABLED=1` (necessária para o SQLite)

### 📥 Clonar o repositório

```bash
git clone https://github.com/Kendal-Katherine/random-number.git
cd random-number
```

### 📁 Configurar variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto (ou em outro caminho, mas ajuste o `LoadEnv()` para isso), com:

```env
API_KEY=minha-api-secreta-123
```

### 🧱 Instalar dependências

```bash
go mod tidy
```

### 🏃‍♀️ Executar o projeto

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

---

## 💾 Banco de Dados

- O projeto usa **SQLite**, sem necessidade de instalação externa.
- O arquivo `database.db` será criado automaticamente na primeira execução.
- O GORM é usado como ORM para abstração das queries.

---

## 🔐 Autenticação

Todas as chamadas para os endpoints devem incluir a seguinte header:

```
X-API-Key: minha-api-secreta-123
```

Substitua pelo valor definido no seu `.env`.

---

## 📡 Exemplo de chamada (via cURL)

```bash
curl --request POST   --url http://localhost:8080/api/v1/random-number   --header 'Content-Type: application/json'   --header 'X-API-Key: minha-api-secreta-123'   --data '{
    "seed": "2025-06-02T12:00:00.123Z"
}'
```

### 📥 Resposta esperada:

```json
{
  "id": "f0cbff8d-60a2-4344-b9bb-644f4009baba",
  "number": 81,
  "seed": "2025-06-02T12:00:00.123Z",
  "createdAt": "2025-06-11T01:28:37-03:00",
  "message": "Random number generated successfully"
}
```

---

## 🔢 Mersenne Twister

O algoritmo **Mersenne Twister (MT19937)** é um dos mais populares geradores de números pseudoaleatórios, conhecido pela sua:

- **Alta qualidade estatística**
- **Extenso período (2¹⁹⁹³⁷ − 1)**
- Boa performance

### 🧠 Como é usado neste projeto?

- A **seed** enviada pelo usuário (ou gerada automaticamente) é convertida para um valor inteiro (`UnixNano()`).
- Esse valor inicializa o algoritmo Mersenne Twister.
- O número gerado (`ExtractNumber() % 100`) garante um valor entre 0 e 99.
- O algoritmo foi implementado no pacote `internal/mt19937`.

---

## 🧪 Testes

Para rodar os testes unitários:

```bash
set CGO_ENABLED=1
go test -v ./...
```

Os testes cobrem:

- Geração de número com seed específica
- Reprodutibilidade com mesma seed
- Rejeição de duplicatas no banco

---

## 🧑‍💻 Estrutura de Pastas

```
random-number/
├── cmd/
│ ├── .db/                  # Arquivo de banco de dados SQLite (gerado)
│ └── main.go               # Entrada principal da aplicação
├── internal/
│ ├── handler/              # Handlers da API (lógica de negócios)
│ ├── random/               # Implementação do algoritmo Mersenne Twister
│ └── router/               # Definição das rotas e middlewares
├── infrastructure/
│ ├── config/               # Configurações, logger e carregamento de .env
│ └── database/             # Modelos e conexão com o banco via GORM
├── .env                    # Variáveis de ambiente (API key, etc.)
└── go.mod / go.sum         # Gerenciamento de dependências Go
```

---


## 🤖 Uso de Inteligência Artificial Generativa

Este projeto utilizou ferramentas de IA generativa, em especial o **ChatGPT**, para apoiar nas seguintes etapas do desenvolvimento:

### Como a IA foi usada:
- **Geração de código**: apoio na escrita de funções como geração de número aleatório com Mersenne Twister, configuração de autenticação por API key e criação de testes unitários.
- **Depuração**: auxílio na identificação e resolução de erros relacionados ao uso do GORM, `CGO_ENABLED`, carregamento do `.env` e estruturação correta do middleware.
- **Documentação**: geração do conteúdo deste `README.md` e estrutura de diretórios com comentários claros.

### Onde a IA foi aplicada:
- `internal/random`: implementação do algoritmo Mersenne Twister.
- `internal/handler`: lógica de geração de número, validação de seed, e testes unitários.
- `infrastructure/config`: configuração do `.env` e autenticação por chave.
- `README.md`: documentação geral, estrutura de projeto e exemplos de uso.

A IA foi usada como uma **ferramenta de apoio**, mas todas as decisões técnicas e adaptações finais foram feitas com entendimento próprio do projeto.



## 👩‍💻 Desenvolvido por

**Kendal Katherine Correia**  
Jr Developer | Golang, TypeScript, PostgreSQL  
🚀 Projeto acadêmico com fins de aprendizado

---

## 📄 Licença

MIT License