# Trilha de aprendizagem de Go

Esta é a ordem principal de estudo do repositório. Ela combina fundamentos da
linguagem, prática contínua, biblioteca padrão, concorrência e backend.

## Visão geral - Sumário

| Fase | Tópicos | Projeto de consolidação |
|---|---|---|
| 0. Ferramentas | Toolchain, módulos e organização |
| 1. Fundamentos | 1–5 | Conversor de unidades |
| 2. Dados e memória | 6–10 | Analisador de texto |
| 3. Abstração idiomática | 11–14 | Calculadora como biblioteca e CLI |
| 4. Código confiável | 15–16 | CLI de tarefas em JSON |
| 5. Concorrência | 17–19 | Verificador concorrente de URLs |
| 6. Execução e diagnóstico | 20 | Analisar os projetos anteriores |
| 7. Backend | 21 | API de tarefas com banco de dados |

---

## Fase 0 — Toolchain, módulos e organização

- [x] Instalação e versão do Go.
- [x] `go run`, `go build` e `go install`.
- [x] `go fmt`, `gofmt` e `goimports`.
- [x] `go test`, `go vet` e `go doc`.
- [x] `go env`, `GOPATH`, `GOBIN`, `GOCACHE` e `GOMODCACHE`.
- [x] `go.mod`, `go.sum`, módulos e pacotes.
- [x] `go get`, `go mod tidy`, `graph`, `why` e `verify`.
- [x] Padrões `.`, `./...` e `all`.
- [x] Organização básica de um módulo.
Material: [Toolchain e comandos úteis](00-toolchain.md).

---

## Fase 1 — Fundamentos da linguagem

### 1. Estrutura, variáveis, constantes e zero values

Estude:
- [x] Estrutura de um arquivo Go: `package`, `import` e declarações.
- [x] `package main` e `func main()`.
- [x] Declaração com `var` e declaração curta com `:=`.
- [x] Declarações agrupadas.
- [x] Inferência de tipos.
- [x] Constantes tipadas e não tipadas.
- [x] `iota` e quando ele é apropriado.
- [] Escopo de pacote, função e bloco.
- [] Shadowing de variáveis.
- [] Zero value de cada categoria de tipo.

Pratique:

- Declare a mesma informação usando `var`, `:=` e constante.
- Imprima os zero values de vários tipos.
- Crie uma enumeração pequena com `iota` e `String()`.
- Produza um caso de shadowing e explique por que ele pode confundir.

Critério de saída: saber quando usar `var`, `:=` e `const`, e prever o valor de
uma variável que foi declarada mas não inicializada explicitamente.

### 2. Tipos primitivos, strings, bytes e runes

Estude:

- `bool`.
- Inteiros com e sem sinal: `int`, `int8`…`uint64`.
- `uintptr` apenas conceitualmente.
- `float32`, `float64`, `complex64` e `complex128`.
- `string`, literais interpretados e raw strings.
- `byte` como alias de `uint8`.
- `rune` como alias de `int32`.
- UTF-8: bytes não são necessariamente caracteres.
- `len(string)` conta bytes.
- Imutabilidade de strings.
- Zero values e limites numéricos.

Pratique:

- Compare `len("ação")` com a quantidade de runes.
- Percorra uma string por índice e depois com `range`.
- Converta entre `string`, `[]byte` e `[]rune`.
- Use `math.MaxInt` e investigue overflow de inteiros.

Critério de saída: explicar a diferença entre byte, rune e caractere e escolher
um tipo numérico sem depender sempre de `int` ou `float64` por hábito.

### 3. Operadores e conversões

Estude:

- Operadores aritméticos, relacionais e lógicos.
- Operadores bit a bit e deslocamentos.
- Precedência.
- Incremento e decremento como statements.
- Conversões explícitas: `T(valor)`.
- Diferença entre conversão e parsing.
- Comparabilidade de tipos.
- Divisão inteira e ponto flutuante.

Pratique:

- Converta Celsius para Fahrenheit.
- Faça parsing de números com `strconv` e trate o erro.
- Implemente flags simples usando operações bit a bit.
- Compare arrays, structs e slices e observe o que compila.

Critério de saída: não esperar conversões numéricas implícitas e saber que
`int("10")` não substitui `strconv.Atoi("10")`.

### 4. Controle de fluxo: `if`, `switch`, `for` e `range`

Estude:

- `if` com statement inicial.
- `switch` de expressão e switch sem expressão.
- Ausência de fallthrough implícito.
- `for` como único laço da linguagem.
- Formas `for condição`, `for {}` e `for inicialização; condição; passo`.
- `range` sobre inteiros, arrays, slices, maps, strings e channels.
- `break`, `continue`, labels e quando evitá-las.
- Ordem não garantida de iteração de maps.

Pratique:

- FizzBuzz.
- Verificador de número primo.
- Tabuada.
- Contador de runes usando `range`.
- Menu simples usando `switch`.

Critério de saída: resolver exercícios com cada forma de `for` e explicar o
que o índice e o valor representam em cada uso de `range`.

### 5. Funções

Funções foram movidas para antes das coleções porque todos os exercícios
seguintes precisam ser divididos em unidades testáveis.

Estude:

- Parâmetros e retornos.
- Múltiplos retornos.
- Retornos nomeados e uso moderado.
- Funções variádicas.
- Funções como valores.
- Tipos de função.
- Funções anônimas e closures.
- Recursão.
- `defer` em nível introdutório.
- Passagem de argumentos por valor.

Pratique:

- Transforme FizzBuzz em funções pequenas.
- Crie uma função que retorne quociente e resto.
- Implemente `Aplicar(a, b int, op func(int, int) int)`.
- Crie uma closure que mantenha um contador.
- Escreva testes para todas as funções puras.

Critério de saída: decompor um problema em funções pequenas e testar essas
funções sem depender de `fmt.Println` para verificar o resultado.

### Projeto 1 — Conversor de unidades

Crie um comando que converta temperatura, distância e peso.

Requisitos:

- Funções separadas da leitura e impressão.
- Parsing de entrada com tratamento de erro básico.
- Uso de `switch` para escolher a conversão.
- Testes das fórmulas.
- Execução por `go run ./cmd/conversor`.

---

## Fase 2 — Tipos compostos e modelo de dados

### 6. Arrays

Estude:

- Sintaxe `[N]T`.
- Tamanho como parte do tipo.
- Inicialização com literais.
- Comparação de arrays comparáveis.
- Cópia de arrays por valor.
- Inferência de tamanho com `[...]T`.
- Relação entre arrays e slices.

Pratique:

- Implemente operações sobre uma matriz `[3][3]int`.
- Passe um array para uma função e demonstre que ele é copiado.
- Compare `[3]int` com outro `[3]int`.

Critério de saída: explicar por que `[3]int` e `[4]int` são tipos diferentes e
por que arrays aparecem menos que slices em APIs comuns.

### 7. Slices — estudo aprofundado

Este é um dos assuntos mais importantes da linguagem.

Estude:

- Slice como descrição de uma região de um array subjacente.
- Literal, slicing e `make`.
- `len` e `cap`.
- `append` e realocação.
- `copy`.
- Slice `nil` versus slice vazio.
- Compartilhamento do array subjacente.
- Full slice expression: `a[low:high:max]`.
- Remoção e inserção de elementos.
- Armadilhas de retenção de memória.
- Slices multidimensionais.
- Ordenação com `slices`.

Pratique:

- Implemente pilha e fila usando slices.
- Remova um elemento preservando e não preservando a ordem.
- Demonstre quando dois slices compartilham memória.
- Force `append` a realocar e compare os resultados.
- Clone um slice sem compartilhar o array.
- Faça testes envolvendo slices `nil` e vazios.

Critério de saída: olhar uma sequência de `append` e subslices e conseguir
explicar quais valores compartilham memória e quando a capacidade muda.

### 8. Maps

Estude:

- Declaração, literal e `make`.
- Leitura de chave ausente e zero value.
- Forma `valor, ok := mapa[chave]`.
- Inserção, atualização, `delete` e `clear`.
- Map `nil`.
- Tipos permitidos como chave.
- Ordem de iteração não especificada.
- Maps não são seguros para escrita concorrente sem sincronização.
- Funções do pacote `maps`.

Pratique:

- Conte frequência de palavras.
- Agrupe pessoas por cidade.
- Inverta um map quando os valores forem únicos.
- Compare dois maps com `maps.Equal`.

Critério de saída: distinguir “chave ausente” de “chave com zero value” e saber
quando usar a forma de dois resultados.

### 9. Structs

Estude:

- Declaração e literais nomeados.
- Acesso e atualização de campos.
- Struct anônima.
- Comparabilidade.
- Embedding de structs.
- Tags de struct.
- Modelagem por composição.
- Zero value útil.
- Evitar structs gigantes e “objetos de dados” sem comportamento.

Pratique:

- Modele `Produto`, `Item` e `Carrinho`.
- Converta uma struct de e para JSON.
- Use tags `json`.
- Crie construtores somente quando houver invariantes ou padrões úteis.

Critério de saída: modelar um domínio pequeno com structs coesas sem tentar
reproduzir herança de linguagens orientadas a objetos.

### 10. Ponteiros e semântica de valores

Estude:

- Operadores `&` e `*`.
- Ponteiro `nil`.
- Passagem sempre por valor, inclusive quando o valor é um ponteiro.
- Alteração por ponteiro versus retorno de um novo valor.
- Endereço de variáveis locais e escape analysis em nível introdutório.
- Ponteiros para structs.
- Quando evitar `*string`, `*int` e ponteiro para interface.
- Diferença entre copiar uma struct e copiar um slice, map ou channel.

Pratique:

- Escreva versões de uma função com valor e com ponteiro.
- Demonstre a cópia de uma struct.
- Demonstre o compartilhamento de dados de um slice copiado.
- Use `go build -gcflags=-m` apenas para observar escapes simples.

Critério de saída: explicar “Go passa tudo por valor” sem concluir que todos os
valores copiados possuem dados completamente independentes.

### Projeto 2 — Analisador de texto

Leia um texto e produza:

- Quantidade de bytes, runes e palavras.
- Frequência de palavras.
- Palavras mais frequentes.
- Tamanho médio das palavras.
- Resultado ordenado e testado.

Esse projeto deve usar strings, runes, slices, maps, structs e funções puras.

---

## Fase 3 — Métodos, interfaces, erros e generics

### 11. Métodos, receivers, pacotes e APIs

Estude:

- Método versus função.
- Receiver por valor e por ponteiro.
- Method sets.
- Métodos sobre tipos definidos.
- Embedding e promoção de métodos.
- Organização de responsabilidades em pacotes.
- Nomes exportados e não exportados.
- Comentários de documentação.
- `internal/` e limites de importação.
- Evitar pacotes genéricos como `utils` e `common`.

Pratique:

- Adicione comportamento a uma struct de carrinho.
- Crie um tipo `Contador` com receiver por ponteiro.
- Crie um tipo numérico com métodos sem usar uma struct.
- Consulte tudo com `go doc`.

Critério de saída: justificar o tipo de receiver escolhido e criar uma API de
pacote pequena cujo uso seja claro sem ler sua implementação.

### 12. Interfaces e composição — estudo aprofundado

Este é outro assunto central de Go.

Estude:

- Implementação implícita.
- Interface como contrato de comportamento.
- Interfaces pequenas e definidas pelo consumidor.
- `any` e interface vazia.
- Valor dinâmico e tipo dinâmico.
- Interface `nil` versus interface contendo ponteiro `nil`.
- Type assertion.
- Type switch.
- Embedding de interfaces.
- Interfaces importantes: `io.Reader`, `io.Writer`, `fmt.Stringer`, `error`.
- Aceitar interfaces e retornar tipos concretos como orientação, não dogma.
- Mocks/fakes pequenos sem frameworks.

Pratique:

- Implemente `fmt.Stringer` para um tipo seu.
- Faça uma função aceitar `io.Writer` e teste com `bytes.Buffer`.
- Modele um repositório em memória atrás de uma interface definida pelo serviço.
- Produza e explique a armadilha de uma interface não `nil` com ponteiro `nil`.

Critério de saída: criar uma interface a partir da necessidade do consumidor e
substituir uma implementação real por uma fake em teste.

### 13. Errors, `defer`, `panic` e `recover`

Estude:

- Interface `error`.
- `errors.New` e `fmt.Errorf`.
- Encadeamento com `%w`.
- `errors.Is` e `errors.As`.
- Erros sentinela.
- Tipos de erro personalizados.
- Adicionar contexto sem perder a causa.
- Retorno antecipado e fluxo feliz à esquerda.
- `defer` e ordem LIFO.
- Fechamento de recursos.
- Quando `panic` é aceitável.
- Uso restrito de `recover` em limites apropriados.

Pratique:

- Faça divisão retornar erro para divisor zero.
- Encadeie erros em três camadas e encontre a causa com `errors.Is`.
- Crie um tipo de erro que carregue um campo inválido.
- Abra e feche um arquivo corretamente com `defer`.

Critério de saída: retornar erros com contexto, preservar causas e diferenciar
uma falha esperada de uma condição realmente irrecuperável.

### 14. Generics

Estude generics somente depois de interfaces e composição.

Estude:

- Parâmetros de tipo.
- Constraints.
- `comparable` e `any`.
- Aproximação de tipos com `~`.
- Uniões de tipos em constraints.
- Inferência de argumentos de tipo.
- Funções genéricas.
- Tipos genéricos.
- Quando uma interface comum ou função concreta é mais simples.
- Pacotes `slices`, `maps` e `cmp` como exemplos de APIs genéricas.

Pratique:

- Implemente `Contem`, `Filtrar` e `Mapear` genéricos.
- Crie uma pilha genérica.
- Compare uma solução genérica com uma baseada em interface.
- Remova um generic desnecessário e explique por que a versão concreta é melhor.

Critério de saída: usar generics para eliminar duplicação real entre tipos sem
transformar todo código em abstrações antecipadas.

### Projeto 3 — Calculadora como biblioteca e CLI

Evolua o projeto atual:

- Operações representadas por funções ou tipos claros.
- Parsing de argumentos.
- Erros para operação inválida e divisão por zero.
- API documentada.
- Testes table-driven.
- Pelo menos uma abstração por interface somente se houver duas implementações
  ou necessidade real do consumidor.

---

## Fase 4 — Biblioteca padrão e código confiável

### 15. Biblioteca padrão, I/O e persistência local

Não tente memorizar toda a stdlib. Aprenda a ler a documentação e aprofunde os
pacotes mais usados:

- Texto e conversão: `fmt`, `strings`, `strconv`, `bytes`, `unicode/utf8`.
- Algoritmos e coleções: `sort`, `slices`, `maps`, `cmp`.
- I/O: `io`, `bufio`, `os`, `path/filepath`.
- Dados: `encoding/json`, `encoding/csv`.
- Tempo: `time`.
- CLI: `flag` e `os.Args`.
- Logs: `log/slog`.
- Segurança e identificadores: `crypto/rand`.
- HTTP introdutório: `net/http`, ainda sem construir o backend completo.

Conceitos transversais:

- Composição com `io.Reader` e `io.Writer`.
- Streaming versus carregar tudo em memória.
- Permissões e caminhos de arquivos.
- Serialização, validação e compatibilidade de dados.
- Injeção de relógio, leitor, escritor e filesystem quando ajuda os testes.

Pratique:

- Copie dados usando `io.Copy`.
- Leia um arquivo linha a linha com `bufio.Scanner`.
- Persista uma lista de structs em JSON.
- Crie uma CLI com `flag`.
- Escreva logs estruturados com `slog`.

Critério de saída: combinar documentação de três pacotes desconhecidos para
resolver um problema sem procurar uma biblioteca externa imediatamente.

### 16. Testing avançado e qualidade

Testes básicos devem existir desde a fase 1. Aqui o estudo se aprofunda:

- Table-driven tests.
- Subtests com `t.Run`.
- `t.Helper`, `t.Cleanup`, `t.TempDir` e `t.Setenv`.
- Pacote interno versus pacote externo `_test`.
- Exemplos executáveis `ExampleXxx`.
- Cobertura e limites da métrica.
- Benchmarks com `testing.B` e `b.Loop`.
- Alocações com `-benchmem`.
- Fuzzing e corpus de seeds.
- Race detector.
- Testes determinísticos.
- Fakes, stubs e dependency injection simples.
- Golden files e diretório `testdata`.

Comandos:

**PowerShell e Bash:**

```shell
go test -v -count=1 ./...
go test -shuffle=on ./...
go test -race ./...
go test -coverprofile=coverage.out ./...
go test -bench=. -benchmem ./...
go test -fuzz=FuzzNome ./caminho/do/pacote
```

Pratique:

- Converta testes repetidos em table-driven tests.
- Teste erros com `errors.Is` e `errors.As`.
- Teste código de I/O usando `bytes.Buffer` e `t.TempDir`.
- Crie benchmark antes de tentar otimizar.
- Crie um fuzz test para parsing ou round trip de serialização.

Critério de saída: testar comportamento observável, inclusive falhas, sem
acoplar todos os testes aos detalhes internos da implementação.

### Projeto 4 — CLI de tarefas

Comandos sugeridos:

```text
todo add "Estudar slices"
todo list
todo complete 1
todo remove 1
```

Requisitos:

- Struct `Task` com ID, descrição, estado e datas.
- Persistência em JSON.
- Separação entre parsing da CLI, regras e armazenamento.
- Erros encadeados com contexto.
- Testes usando diretório temporário.
- Logs úteis sem misturar log com a saída normal do comando.

---

## Fase 5 — Concorrência

### 17. Goroutines

Estude:

- Palavra-chave `go`.
- Concorrência versus paralelismo.
- Scheduler em nível conceitual.
- Captura de variáveis por closures.
- Tempo de vida de goroutines.
- Goroutine leaks.
- Ausência de garantia de ordem.
- Comunicação e ownership de dados.
- Race detector desde o primeiro exercício concorrente.

Pratique:

- Execute trabalhos independentes sequencialmente e concorrentemente.
- Produza uma goroutine leak e depois corrija.
- Produza uma data race e observe `go test -race`.
- Meça antes e depois; concorrência não garante maior velocidade.

Critério de saída: iniciar goroutines sabendo como elas terminam, como erros são
recolhidos e quem é responsável pelos dados compartilhados.

### 18. Channels e `select`

Estude:

- Channel sem buffer e com buffer.
- Envio, recebimento e bloqueio.
- Channels direcionais.
- `close` e quem deve fechar.
- Recebimento com `valor, ok`.
- `range` sobre channel.
- `select`.
- Timeout e cancelamento.
- Fan-out, fan-in e pipelines.
- Backpressure.
- Channel `nil` e channel fechado.
- Deadlocks.

Pratique:

- Crie produtor e consumidor.
- Implemente fan-out/fan-in.
- Cancele um pipeline sem deixar goroutines bloqueadas.
- Demonstre a diferença entre channel fechado e channel `nil`.

Critério de saída: desenhar quem envia, quem recebe, quem fecha e como cada
goroutine termina antes de implementar um pipeline.

### 19. `sync`, atomics e `context`

Estude:

- `sync.WaitGroup`.
- `sync.Mutex` e `sync.RWMutex`.
- `sync.Once`.
- `sync.Cond` conceitualmente.
- `sync.Map` e por que não é o map padrão concorrente.
- Operações atômicas quando realmente necessárias.
- `context.Context`.
- Cancelamento, deadlines e timeouts.
- Propagação de context como primeiro parâmetro.
- Não guardar context em struct como regra geral.
- Não usar `context.Value` para parâmetros opcionais comuns.
- Concorrência limitada e worker pools.

Pratique:

- Proteja um contador com mutex e compare com atomic.
- Implemente um worker pool com limite configurável.
- Cancele todos os workers no primeiro erro.
- Passe timeout para uma requisição HTTP.

Critério de saída: escolher conscientemente entre channel, mutex ou ownership
exclusivo, e propagar cancelamento até todas as operações bloqueantes.

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

- Versão sequencial primeiro.
- Worker pool com limite de concorrência.
- Timeout global e por requisição.
- Resultados apresentados em ordem estável.
- Erros preservados.
- Testes com `httptest.Server`.
- `go test -race ./...` sem problemas.

---

## Fase 6 — Runtime, compilador e diagnóstico

### 20. Runtime e compiler internals

Estude o suficiente para diagnosticar aplicações; não é necessário ler o
compilador inteiro antes de construir backend.

Runtime:

- Scheduler e modelo G-M-P.
- Stacks que crescem dinamicamente.
- Garbage collector em alto nível.
- Heap versus stack como decisão do compilador.
- Escape analysis.
- Alocação, pressão no GC e pooling.
- Representação conceitual de slices, strings e interfaces.
- Data races e modelo de memória.

Compilação e binários:

- Etapas conceituais: parsing, type checking, SSA, geração de código e link.
- Build cache.
- Build constraints.
- Inlining.
- Cross-compilation com `GOOS` e `GOARCH`.
- Informações de módulo embutidas no binário.

Ferramentas:

**PowerShell e Bash:**

```shell
go build -gcflags=-m ./cmd/calculadora
go test -bench=. -benchmem ./...
go test -cpuprofile=cpu.out ./caminho/do/pacote
go tool pprof cpu.out
go test -trace=trace.out ./caminho/do/pacote
go tool trace trace.out
go build -x ./cmd/calculadora
```

Inspeção do executável compilado:

**PowerShell:**

```powershell
go version -m ./programa.exe
```

**Bash no Linux:**

```bash
go version -m ./programa
```

Deixe para estudo opcional posterior:

- `unsafe`.
- Assembly.
- cgo.
- Código-fonte completo do compilador.
- Implementação interna detalhada do garbage collector.

Critério de saída: usar benchmark, escape analysis e profiler para investigar
uma hipótese, evitando “otimizações” baseadas apenas em intuição.

---

## Fase 7 — Backend real

### 21. Aplicação backend completa

#### HTTP

- `net/http`, `http.Server`, `Handler` e `ServeMux`.
- Métodos, headers, status codes e body.
- JSON de entrada e saída.
- Limites de tamanho e validação de entrada.
- Middleware.
- Context da requisição.
- Timeouts do cliente e servidor.
- `httptest`.
- Graceful shutdown.

Aprenda primeiro `net/http`; frameworks podem ser avaliados depois que o fluxo
de uma requisição HTTP estiver claro.

#### Persistência

- `database/sql`.
- Driver PostgreSQL ou SQLite.
- Queries parametrizadas.
- `QueryContext`, `ExecContext` e `Scan`.
- Transações.
- Connection pool.
- Migrações.
- Constraints e índices no banco.
- Testes de integração.

Comece com SQL explícito. Um ORM fica mais fácil de avaliar depois que você
entende o que ele abstrai.

#### Design e operação

- Separação entre transporte HTTP, regras e persistência.
- Dependências apontando para o domínio, sem camadas artificiais demais.
- Configuração por ambiente.
- `log/slog` e logs estruturados.
- IDs e timestamps.
- Paginação, filtros e ordenação.
- Idempotência quando aplicável.
- Autenticação e autorização.
- Hash de senha com algoritmo apropriado; nunca criptografia caseira.
- CORS, limites de body e rate limiting conforme o contexto.
- Sinais do sistema e shutdown.
- Métricas, tracing e profiling em nível introdutório.
- CI executando format, vet e testes.
- Container e deploy somente depois de a aplicação funcionar localmente.

#### Projeto final — API de tarefas

Endpoints mínimos:

```text
POST   /tasks
GET    /tasks
GET    /tasks/{id}
PATCH  /tasks/{id}
DELETE /tasks/{id}
```

Evolução recomendada:

1. Implementação em memória.
2. Testes HTTP com `httptest`.
3. Persistência SQLite ou PostgreSQL.
4. Migrações e testes de integração.
5. Logs estruturados e tratamento uniforme de erros.
6. Timeouts e graceful shutdown.
7. Autenticação, se fizer sentido para o objetivo.
8. CI.
9. Container e deploy.

Critério de saída: outra pessoa conseguir clonar, configurar, testar, executar
e entender a API usando apenas a documentação do repositório.

---

## Marcos dos projetos

| Após o tópico | Entrega |
|---|---|
| 5 | Conversor de unidades com testes |
| 10 | Analisador de texto |
| 14 | Calculadora como biblioteca e CLI |
| 16 | CLI de tarefas persistida em JSON |
| 19 | Verificador concorrente de URLs |
| 20 | Relatório de benchmark/profile de um projeto |
| 21 | API de tarefas com banco de dados |

## Assuntos para estudar depois

Não são proibidos, mas rendem mais quando os fundamentos já estão sólidos:

- Reflection avançada.
- `unsafe` e Assembly.
- cgo.
- Geração de código.
- WebSockets e gRPC.
- Filas e sistemas distribuídos.
- Cache distribuído.
- Microserviços.
- Kubernetes.
- Escrita de ferramentas de análise estática.
- Contribuição para o compilador ou runtime.

## Checklist de domínio por fase

Antes de avançar, verifique se você consegue:

- [ ] Explicar os conceitos sem depender de termos de outra linguagem.
- [ ] Implementar um exemplo novo sem copiar.
- [ ] Escrever testes de sucesso e falha.
- [ ] Ler a documentação dos pacotes utilizados.
- [ ] Identificar pelo menos uma armadilha comum do assunto.
- [ ] Aplicar o assunto no projeto da fase.
- [ ] Executar `go fmt`, `go vet` e `go test` sem erros.

## Referências

- [A Tour of Go](https://go.dev/tour/)
- [Go User Manual](https://go.dev/doc/)
- [Especificação da linguagem](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Biblioteca padrão](https://pkg.go.dev/std)
- [Go Blog](https://go.dev/blog/)
- [Organização de módulos](https://go.dev/doc/modules/layout)
