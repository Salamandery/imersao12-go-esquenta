# 🚀 Imersão 12 Go Esquenta - API de Produtos

[![Go Reference](https://pkg.go.dev/badge/github.com/salamandery/imersao12-go-esquenta.svg)](https://pkg.go.dev/github.com/salamandery/imersao12-go-esquenta)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/docker-ready-blue)](https://www.docker.com/)

> **API REST em Go para cadastro e listagem de produtos, com integração a Kafka e MySQL.**
> 
> 🏆 Projeto desenvolvido durante o evento **Imersão 12 Go Esquenta**.

---

## 📑 Sumário
- [Sobre o Projeto](#sobre-o-projeto)
- [Principais Tecnologias](#principais-tecnologias)
- [Padrões de Projeto e Arquitetura](#padrões-de-projeto-e-arquitetura)
- [Setup e Execução](#setup-e-execução)
- [Endpoints](#endpoints)
- [Autor](#-autor)

---

## 💡 Sobre o Projeto

API para cadastro e listagem de produtos, utilizando arquitetura limpa, integração com banco de dados MySQL e mensageria Kafka.

---

## 🛠️ Principais Tecnologias

- **Go 1.21+**
- **[Chi v5](https://github.com/go-chi/chi)** · Router HTTP
- **[MySQL](https://www.mysql.com/)** · Banco de dados relacional
- **[confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)** · Cliente Kafka
- **[uuid](https://github.com/google/uuid)** · Geração de UUIDs
- **[Docker](https://www.docker.com/)** e **[Docker Compose](https://docs.docker.com/compose/)**

---

## 🏗️ Padrões de Projeto e Arquitetura

- **Clean Architecture:** Separação clara entre camadas de entidade, usecase, infraestrutura e web
- **Repository Pattern:** Abstração do acesso a dados via interface `ProductRepository`
- **DTOs:** Para entrada e saída dos casos de uso
- **Handlers:** Camada web desacoplada dos casos de uso
- **Injeção de Dependências:** Construção dos casos de uso e repositórios via construtores

```
internal/
  entity/        # Entidades e interfaces de domínio
  usecase/       # Casos de uso (aplicação)
  infra/         # Infraestrutura (DB, Kafka, Web)
    repository/  # Implementação dos repositórios
    akafka/      # Integração Kafka
    web/         # Handlers HTTP
cmd/app/         # Ponto de entrada da aplicação
```

---

## ⚙️ Setup e Execução

### Pré-requisitos
- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

### Subindo o ambiente com Docker Compose

```sh
docker-compose up --build
```
A aplicação estará disponível em [`http://localhost:8000`](http://localhost:8000).

### Rodando localmente (sem Docker)
1. Suba o MySQL e o Kafka (pode usar o `docker-compose.yaml` para isso)
2. Instale as dependências Go:
   ```sh
   go mod tidy
   ```
3. Execute a aplicação:
   ```sh
   go run ./cmd/app/main.go
   ```

---

## 📬 Endpoints

- `POST /products` — Cadastra um novo produto
  - Body: `{ "name": "Produto", "price": 100 }`
- `GET /products` — Lista todos os produtos

---


## 👤 Autor
by **Rodolfo M. F. Abreu**

<div align="center">
  <sub>Projeto desenvolvido durante a <a href="https://imersao.fullcycle.com.br/">Imersão 12 Go Esquenta</a> 🚀</sub>
</div>