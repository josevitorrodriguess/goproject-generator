# Go Project Generator (GPG)

Uma ferramenta CLI desenvolvida em Go para gerar projetos com estrutura padronizada usando Chi router, CORS configurado e utilitários JSON prontos para uso. Acelere o desenvolvimento de suas APIs REST criando uma base sólida e organizida em segundos.

## ✨ Características

- 🚀 **Geração rápida de projetos** - Crie um projeto completo em segundos
- 📁 **Estrutura padronizada** - Organização seguindo padrões da comunidade Go
- 🌐 **Chi Router v5** - Router HTTP rápido e leve já configurado
- 🔄 **CORS configurado** - Suporte completo para requisições cross-origin
- 🛡️ **Middleware essencial** - Logger, Recovery e RequestID incluídos
- 🔧 **Utilitários JSON** - Funções prontas para encode/decode JSON
- 📦 **Dependências automáticas** - go.mod e todas as dependências instaladas
- 🏥 **Health Check** - Endpoint `/health` configurado por padrão


## 📋 Pré-requisitos

- Go 1.24 ou superior
- Git (para instalação via GitHub)

## 🚀 Instalação

### Via go install (Recomendado)

```bash
go install github.com/josevitorrodriguess/goproject-generator@latest
```

### Via clone e build

```bash
git clone https://github.com/josevitorrodriguess/goproject-generator.git
cd goproject-generator
go build -o gpg main.go
```

## 📖 Uso

### Comando básico

```bash
gpg create meu-projeto
```

### Com opções personalizadas

```bash
gpg create minha-api --port 8080 --module github.com/usuario/minha-api
```

### Flags disponíveis

```bash
gpg create --help
```

## ⚙️ Opções disponíveis

| Flag | Shorthand | Descrição | Valor padrão |
|------|-----------|-----------|--------------|
| `--port` | `-p` | Define a porta do servidor | `3000` |
| `--module` | `-m` | Nome do módulo Go | Nome do projeto |

## 📁 Estrutura gerada

```
meu-projeto/
├── go.mod                           # Módulo Go com dependências
├── cmd/
│   └── meu-projeto/
│       └── main.go                 # Ponto de entrada da aplicação
└── internal/
    ├── api/
    │   ├── api.go                  # Struct do servidor e inicialização
    │   └── router.go               # Rotas, middleware e handlers
    └── jsonutils/
        └── json_utils.go           # Utilitários para JSON
```

### 📄 Descrição detalhada dos arquivos

#### `cmd/meu-projeto/main.go`
- Ponto de entrada da aplicação
- Carregamento de variáveis de ambiente (.env)
- Inicialização do servidor na porta 3050
- Configuração do servidor HTTP

#### `internal/api/api.go`
- Struct `Api` principal
- Método `New()` para criação da instância
- Inicialização do Chi router
- Configuração de middleware e rotas

#### `internal/api/router.go`
- **Middleware configurado:**
  - `Logger` - Log de todas as requisições
  - `Recoverer` - Recovery automático de panics
  - `RequestID` - ID único para cada requisição
  - `CORS` - Configuração completa para cross-origin
- **Rotas padrão:**
  - `GET /health` - Health check endpoint
- **Handler de exemplo** com response padronizada

#### `internal/jsonutils/json_utils.go`
- `EncodeJson[T]()` - Função genérica para encoding JSON
- `DecodeJson[T]()` - Função genérica para decoding JSON
- `SendError()` - Envio padronizado de erros JSON

## 🏃‍♂️ Como executar o projeto gerado

1. **Entre no diretório:**
   ```bash
   cd meu-projeto
   ```

2. **Execute o servidor:**
   ```bash
   go run cmd/meu-projeto/main.go
   ```

3. **Teste o health check:**
   ```bash
   curl http://localhost:3050/health
   ```

   **Response esperado:**
   ```json
   {
     "service": "api",
     "status": "OK"
   }
   ```

## 📦 Dependências incluídas

O projeto é gerado com as seguintes dependências automaticamente instaladas:

- **`github.com/go-chi/chi/v5`** - Router HTTP performático
- **`github.com/go-chi/chi/`** - Middleware adicional
- **`github.com/go-chi/cors`** - Middleware CORS
- **`github.com/joho/godotenv`** - Carregamento de variáveis de ambiente

## 🔧 Recursos prontos para uso

### ✅ Middleware configurado
- **Logger** - Log automático de todas as requisições
- **Recoverer** - Recovery de panics com graceful degradation
- **RequestID** - Tracking de requisições com ID único
- **CORS** - Suporte completo para requisições cross-origin

### ✅ Utilitários JSON
```go
// Encoding com tipos genéricos
response := map[string]string{"status": "ok"}
jsonutils.EncodeJson(w, r, http.StatusOK, response)

// Decoding com tipos genéricos
type UserRequest struct {
    Name string `json:"name"`
}
user, err := jsonutils.DecodeJson[UserRequest](r)

// Envio de erros padronizado
jsonutils.SendError(w, r, http.StatusBadRequest, "Invalid input")
```

### ✅ Health Check
Endpoint `/health` configurado retornando:
```json
{
  "status": "OK",
  "service": "api"
}
```

## 🌟 Exemplo de uso prático

Após executar `gpg create minha-api`, você terá:

```bash
cd minha-api
go run cmd/minha-api/main.go
```

**Console output:**
```
Server is running on http://localhost:3050
```

**Testando:**
```bash
curl http://localhost:3050/health
# {"service":"api","status":"OK"}
```

## 🛠️ Personalizações

### Adicionando novas rotas

No arquivo `internal/api/router.go`:

```go
func (api *Api) setupRoutes() {
    api.Router.Get("/health", api.handleHealth)
    
    // Adicione suas rotas aqui
    api.Router.Get("/users", api.handleGetUsers)
    api.Router.Post("/users", api.handleCreateUser)
}
```


## 🤝 Como contribuir

1. Faça um fork do projeto
2. Crie uma branch para sua feature:
   ```bash
   git checkout -b feature/nova-feature
   ```
3. Commit suas mudanças:
   ```bash
   git commit -am 'feat: adiciona nova feature'
   ```
4. Push para a branch:
   ```bash
   git push origin feature/nova-feature
   ```
5. Abra um Pull Request

### 📝 Padrões de commit

- `feat:` nova funcionalidade
- `fix:` correção de bug
- `docs:` atualização de documentação
- `refactor:` refatoração sem mudança de funcionalidade
- `test:` adição ou correção de testes


## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 🙏 Tecnologias utilizadas

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Chi Router](https://github.com/go-chi/chi) - HTTP router
- [Chi CORS](https://github.com/go-chi/cors) - CORS middleware  
- [GoDotEnv](https://github.com/joho/godotenv) - Environment variables

⭐ **Curtiu o projeto?** Deixe uma estrela no repositório!
