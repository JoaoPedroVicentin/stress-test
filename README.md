# Stress Test HTTP

Este projeto é uma ferramenta de stress test HTTP desenvolvida em Go, que permite realizar um grande volume de requisições simultâneas a um endpoint HTTP, medindo desempenho, distribuição de status HTTP e eventuais falhas.

## Funcionalidades

- Envio de múltiplas requisições HTTP concorrentes para uma URL informada
- Controle do número total de requisições e do grau de concorrência
- Relatório detalhado ao final da execução, incluindo:
    - Parâmetros utilizados
	- Tempo total de execução
	- Quantidade total de requests
	- Distribuição dos códigos de status HTTP

## Como clonar o repositório

```
git clone https://github.com/JoaoPedroVicentin/stress-test.git
cd stress-test
```

## Como rodar localmente

```
go run main.go --url=<URL> --requests=<TOTAL> --concurrency=<CONCORRENCIA>
```

Exemplo:

```
go run main.go --url=https://www.google.com --requests=100 --concurrency=10
```

## Como rodar com Docker

1. Construa a imagem Docker:

```
docker build -t stress-test-go .
```

2. Execute o container:

```
docker run --rm stress-test-go --url=<URL> --requests=<TOTAL> --concurrency=<CONCORRENCIA>
```

Exemplo:

```
docker run --rm stress-test-go --url=https://www.google.com --requests=100 --concurrency=10
```

<div align="center">
<h3>👨‍💻</h3>
    <h3> Criado por João Pedro Vicentin!</h3>
    <div>
        <h3>
            <a href="https://www.linkedin.com/in/joaopedrovicentin/" target="_blank">Linkedin</a>
            <a href='https://github.com/JoaoPedroVicentin' target='_blank'>Github</a>
            <a href="https://contate.me/joao-pedro-lopes-vicentin" target="_blank">Whatsapp</a>
        </h3>
    </div>
</div>