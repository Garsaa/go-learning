# Exemplos

Esta pasta guarda programas pequenos criados para isolar um conceito.

Cada exemplo executável deve ficar em sua própria pasta para que possa possuir
um `package main` independente:

```text
examples/
├── fundamentos/
│   ├── variaveis/
│   │   └── main.go
│   └── controle/
│       └── main.go
└── colecoes/
    ├── slices/
    │   └── main.go
    └── maps/
        └── main.go
```

Exemplo de execução:

**PowerShell e Bash:**

```shell
go run ./examples/fundamentos/variaveis
```

Exemplos devem ser pequenos. Código que cresce, persiste dados ou combina
vários conceitos deve ser promovido para `projects/`.
