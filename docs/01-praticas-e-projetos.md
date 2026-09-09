# Práticas e projetos de Go

Este documento reúne os exercícios e projetos de consolidação da
[trilha de aprendizagem](00-trilha-de-aprendizado.md). Faça as práticas depois
de estudar o tópico correspondente e use os projetos para consolidar cada fase.

## Práticas

### Fase 1 — Fundamentos da linguagem

#### 3. Operadores e conversões

- [ ] Converta Celsius para Fahrenheit.
- [ ] Faça parsing de números com `strconv` e trate o erro.
- [ ] Implemente flags simples usando operações bit a bit.
- [ ] Compare arrays, structs e slices e observe o que compila.

#### 4. Controle de fluxo: `if`, `switch`, `for` e `range`

- [ ] FizzBuzz.
- [ ] Verificador de número primo.
- [ ] Tabuada.
- [ ] Contador de runes usando `range`.
- [ ] Menu simples usando `switch`.

#### 5. Funções

- [ ] Transforme FizzBuzz em funções pequenas.
- [ ] Crie uma função que retorne quociente e resto.
- [ ] Implemente `Aplicar(a, b int, op func(int, int) int)`.
- [ ] Crie uma closure que mantenha um contador.
- [ ] Escreva testes para todas as funções puras.

### Fase 2 — Tipos compostos e modelo de dados

#### 6. Arrays

- [ ] Implemente operações sobre uma matriz `[3][3]int`.
- [ ] Passe um array para uma função e demonstre que ele é copiado.
- [ ] Compare `[3]int` com outro `[3]int`.

#### 7. Slices — estudo aprofundado

- [ ] Implemente pilha e fila usando slices.
- [ ] Remova um elemento preservando e não preservando a ordem.
- [ ] Demonstre quando dois slices compartilham memória.
- [ ] Force `append` a realocar e compare os resultados.
- [ ] Clone um slice sem compartilhar o array.
- [ ] Faça testes envolvendo slices `nil` e vazios.

#### 8. Maps

- [ ] Conte frequência de palavras.
- [ ] Agrupe pessoas por cidade.
- [ ] Inverta um map quando os valores forem únicos.
- [ ] Compare dois maps com `maps.Equal`.

#### 9. Structs

- [ ] Modele `Produto`, `Item` e `Carrinho`.
- [ ] Converta uma struct de e para JSON.
- [ ] Use tags `json`.
- [ ] Crie construtores somente quando houver invariantes ou padrões úteis.

#### 10. Ponteiros e semântica de valores

- [ ] Escreva versões de uma função com valor e com ponteiro.
- [ ] Demonstre a cópia de uma struct.
- [ ] Demonstre o compartilhamento de dados de um slice copiado.
- [ ] Use `go build -gcflags=-m` apenas para observar escapes simples.

### Fase 3 — Métodos, interfaces, erros e generics

#### 11. Métodos, receivers, pacotes e APIs

- [ ] Adicione comportamento a uma struct de carrinho.
- [ ] Crie um tipo `Contador` com receiver por ponteiro.
- [ ] Crie um tipo numérico com métodos sem usar uma struct.
- [ ] Consulte tudo com `go doc`.

#### 12. Interfaces e composição — estudo aprofundado

- [ ] Implemente `fmt.Stringer` para um tipo seu.
- [ ] Faça uma função aceitar `io.Writer` e teste com `bytes.Buffer`.
- [ ] Modele um repositório em memória atrás de uma interface definida pelo serviço.
- [ ] Produza e explique a armadilha de uma interface não `nil` com ponteiro `nil`.

#### 13. Errors, `defer`, `panic` e `recover`

- [ ] Faça divisão retornar erro para divisor zero.
- [ ] Encadeie erros em três camadas e encontre a causa com `errors.Is`.
- [ ] Crie um tipo de erro que carregue um campo inválido.
- [ ] Abra e feche um arquivo corretamente com `defer`.

#### 14. Generics

- [ ] Implemente `Contem`, `Filtrar` e `Mapear` genéricos.
- [ ] Crie uma pilha genérica.
- [ ] Compare uma solução genérica com uma baseada em interface.
- [ ] Remova um generic desnecessário e explique por que a versão concreta é melhor.

### Fase 4 — Biblioteca padrão e código confiável

#### 15. Biblioteca padrão, I/O e persistência local

- [ ] Copie dados usando `io.Copy`.
- [ ] Leia um arquivo linha a linha com `bufio.Scanner`.
- [ ] Persista uma lista de structs em JSON.
- [ ] Crie uma CLI com `flag`.
- [ ] Escreva logs estruturados com `slog`.

#### 16. Testing avançado e qualidade

- [ ] Converta testes repetidos em table-driven tests.
- [ ] Teste erros com `errors.Is` e `errors.As`.
- [ ] Teste código de I/O usando `bytes.Buffer` e `t.TempDir`.
- [ ] Crie benchmark antes de tentar otimizar.
- [ ] Crie um fuzz test para parsing ou round trip de serialização.

### Fase 5 — Concorrência

#### 17. Goroutines

- [ ] Execute trabalhos independentes sequencialmente e concorrentemente.
- [ ] Produza uma goroutine leak e depois corrija.
- [ ] Produza uma data race e observe `go test -race`.
- [ ] Meça antes e depois; concorrência não garante maior velocidade.

#### 18. Channels e `select`

- [ ] Crie produtor e consumidor.
- [ ] Implemente fan-out/fan-in.
- [ ] Cancele um pipeline sem deixar goroutines bloqueadas.
- [ ] Demonstre a diferença entre channel fechado e channel `nil`.

#### 19. `sync`, atomics e `context`

- [ ] Proteja um contador com mutex e compare com atomic.
- [ ] Implemente um worker pool com limite configurável.
- [ ] Cancele todos os workers no primeiro erro.
- [ ] Passe timeout para uma requisição HTTP.

## Projetos

### Projeto 1 — Conversor de unidades

Crie um comando que converta temperatura, distância e peso.

Requisitos:

- [ ] Funções separadas da leitura e impressão.
- [ ] Parsing de entrada com tratamento de erro básico.
- [ ] Uso de `switch` para escolher a conversão.
- [ ] Testes das fórmulas.
- [ ] Execução por `go run ./cmd/conversor`.

### Projeto 2 — Analisador de texto

Leia um texto e produza:

- [ ] Quantidade de bytes, runes e palavras.
- [ ] Frequência de palavras.
- [ ] Palavras mais frequentes.
- [ ] Tamanho médio das palavras.
- [ ] Resultado ordenado e testado.

Esse projeto deve usar strings, runes, slices, maps, structs e funções puras.

### Projeto 3 — Calculadora como biblioteca e CLI

Evolua o projeto atual:

- [ ] Operações representadas por funções ou tipos claros.
- [ ] Parsing de argumentos.
- [ ] Erros para operação inválida e divisão por zero.
- [ ] API documentada.
- [ ] Testes table-driven.
- [ ] Pelo menos uma abstração por interface somente se houver duas implementações
  ou necessidade real do consumidor.

### Projeto 4 — CLI de tarefas

Comandos sugeridos:

```text
todo add "Estudar slices"
todo list
todo complete 1
todo remove 1
```

Requisitos:

- [ ] Struct `Task` com ID, descrição, estado e datas.
- [ ] Persistência em JSON.
- [ ] Separação entre parsing da CLI, regras e armazenamento.
- [ ] Erros encadeados com contexto.
- [ ] Testes usando diretório temporário.
- [ ] Logs úteis sem misturar log com a saída normal do comando.

### Projeto 5 — Verificador concorrente de URLs

Entrada:

```text
https://go.dev
https://example.com
```

Saída aproximada:

```text
https://go.dev       200 OK  120ms
https://example.com  200 OK   85ms
```

Requisitos:

- [ ] Versão sequencial primeiro.
- [ ] Worker pool com limite de concorrência.
- [ ] Timeout global e por requisição.
- [ ] Resultados apresentados em ordem estável.
- [ ] Erros preservados.
- [ ] Testes com `httptest.Server`.
- [ ] `go test -race ./...` sem problemas.

### Projeto 6 — API de tarefas

Endpoints mínimos:

```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}
DELETE /tasks/{id}
```

Evolução recomendada:

1. [ ] Implementação em memória.
2. [ ] Testes HTTP com `httptest`.
3. [ ] Persistência SQLite ou PostgreSQL.
4. [ ] Migrações e testes de integração.
5. [ ] Logs estruturados e tratamento uniforme de erros.
6. [ ] Timeouts e graceful shutdown.
7. [ ] Autenticação, se fizer sentido para o objetivo.
8. [ ] CI.
9. [ ] Container e deploy.

Critério de saída: outra pessoa conseguir clonar, configurar, testar, executar
e entender a API usando apenas a documentação do repositório.

### Marcos dos projetos

| Após o tópico | Entrega |
|---|---|
| 5 | Conversor de unidades com testes |
| 10 | Analisador de texto |
| 14 | Calculadora como biblioteca e CLI |
| 16 | CLI de tarefas persistida em JSON |
| 19 | Verificador concorrente de URLs |
| 20 | Relatório de benchmark/profile de um projeto |
| 21 | API de tarefas com banco de dados |
