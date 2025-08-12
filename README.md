# Go Project Generator

Uma CLI para gerar projetos Go com estrutura padronizada usando Chi router.

## Instalação

```bash
go install github.com/seu-usuario/go-project-generator@latest
```

Ou clone o repositório e compile:

```bash
git clone https://github.com/seu-usuario/go-project-generator.git
cd go-project-generator
go build -o gpg main.go
```

## Uso

### Criar novo projeto

```bash
gpg create meu-projeto
```

Isso criará a seguinte estrutura:

```
meu-projeto/
├── go.mod
├── cmd/
│   └── meu-projeto/
│       └── main.go
└── internal/
    └── api/
        ├── api.go
        └── router.go
```

### Flags opcionais

```bash
gpg create meu-projeto --port 8080 --module github.com/usuario/meu-projeto
```

- `--port`: Define a porta padrão do servidor (default: 3000)
- `--module`: Define o nome do módulo Go (default: nome do projeto)

## Estrutura gerada

### `cmd/meu-projeto/main.go`
- Ponto de entrada da aplicação
- Configuração básica do servidor
- Inicialização do router

### `internal/api/api.go`
- Struct do servidor
- Configurações do Chi
- Métodos de inicialização

### `internal/api/router.go`
- Definição das rotas
- Middleware básico
- Handlers de exemplo

## O que é incluído

- ✅ Chi router configurado
- ✅ Estrutura de pastas padronizada
- ✅ Middleware básico (logger, recoverer)
- ✅ Handlers de exemplo
- ✅ `go.mod` inicializado
- ✅ Dependências instaladas automaticamente

## Próximos passos após criar o projeto

1. Entre no diretório do projeto:
   ```bash
   cd meu-projeto
   ```

2. Execute o servidor:
   ```bash
   go run cmd/meu-projeto/main.go
   ```

3. Acesse `http://localhost:3000` para ver o servidor rodando

## Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## Licença

MIT License - veja o arquivo [LICENSE](LICENSE) para detalhes.

## Roadmap

- [ ] Adicionar templates para diferentes tipos de handlers
- [ ] Suporte para geração de middlewares customizados
- [ ] Templates para testes unitários
- [ ] Integração com bancos de dados (PostgreSQL, MySQL)
- [ ] Geração de Dockerfile
- [ ] Templates para APIs GraphQL
