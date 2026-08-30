# Go Learning

Repositório de estudo da linguagem Go por meio de anotações, exemplos,
exercícios e pequenos projetos.

## Objetivo

Aprender não apenas a sintaxe, mas também o modelo de tipos, a biblioteca
padrão, testes, concorrência e a construção de aplicações backend idiomáticas.

## Convenção dos exemplos de terminal

Os blocos marcados como **PowerShell e Bash** usam a mesma sintaxe nos dois
shells. O identificador `shell` é usado no bloco Markdown porque uma cerca de
código aceita apenas um identificador de linguagem. Quando os comandos diferem
entre Windows e Linux, cada shell recebe seu próprio exemplo.

## Trilha de aprendizagem

A ordem completa dos estudos, os exercícios e os critérios para avançar estão
em [Trilha de aprendizagem](docs/01-trilha-de-aprendizado.md).

Progresso geral:

- [x] Toolchain, módulos e comandos essenciais
- [ ] Fundamentos da linguagem
- [ ] Tipos compostos e modelo de dados
- [ ] Métodos, interfaces e tratamento de erros
- [ ] Biblioteca padrão e persistência local
- [ ] Testes avançados
- [ ] Concorrência
- [ ] Runtime e ferramentas de diagnóstico
- [ ] Backend real

## Documentação

- [Toolchain e comandos úteis](docs/00-toolchain.md)
- [Trilha de aprendizagem](docs/01-trilha-de-aprendizado.md)

## Organização do repositório

```text
go-learning/
├── cmd/          # Pontos de entrada dos executáveis (package main)
├── docs/         # Anotações e trilha de estudos
├── examples/     # Exemplos pequenos, cada conceito em sua própria pasta
├── projects/     # Projetos incrementais e pacotes reutilizáveis
├── assets/       # Imagens usadas pela documentação
├── go.mod
└── README.md
```

O repositório usa um único módulo. Cada diretório com código representa um
pacote; novos programas executáveis devem ficar em uma subpasta de `cmd/`.

## Projeto atual: calculadora

A lógica está no pacote `projects/calculadora`, e o executável está em
`cmd/calculadora`.

**PowerShell e Bash:**

```shell
go run ./cmd/calculadora
go test ./...
go test -race ./...
go vet ./...
```

## Rotina de estudo

Para cada tópico:

1. Escrever uma explicação curta com as próprias palavras.
2. Criar um exemplo mínimo em `examples/`.
3. Fazer pelo menos um exercício sem copiar a solução.
4. Adicionar testes para o comportamento aprendido.
5. Aplicar o conceito em um projeto de `projects/`.
6. Executar formatação, análise estática e testes.

**PowerShell e Bash:**

```shell
go fmt ./...
go vet ./...
go test ./...
```

## Referências principais

- [Documentação oficial](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/)
- [Biblioteca padrão e pacotes](https://pkg.go.dev/std)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
