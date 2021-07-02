# Golang: Performance e simplicidade ao eXtremo

### [Golang](https://golang.org/)

* Criado pelo Google

* Simplicidade, rapidez e alta performance

* Facilidade ao trabalhar com multithreading

* Baixa curva de aprendizagem (Curiosidade: +/- 24 palavras chave)

### Características

* É uma **linguagem tipada**, ou seja, não criação de um variável é necessário definir o seu tipo

* Imports e variáveis não utilizadas = erro (não compila)

* **Não há a propagação de erros** no Golang, ou seja, caso ocorra um, ele é tratado logo após o seu lançamento

  > **OBS**: no Golang, não existe `try-catch`

* É **orientado a objetos**, porém se difere em relação as outras linguagens

### Sintaxe

* **Criação de um função**

  ```go
  func main() {
	  fmt.Println("OLÁ FULL CYCLE")
  }
  ```

* **Criação de um função com parâmetros e retorno**

  ```go
  func soma(a int, b int) int {
	  return a + b
  }
  ```

* **Criação de um função com 2 tipos de retorno**

  ```go
  func soma(a int, b int) (int, error) {
    if a < 0 {
      return 0, errors.New("a < 0")
    }

    return a + b, nil
  }
  ```

* **Declarar uma variável e atribuí-la um valor**

  ```go
  var carro string
  carro = "Fusca"
  ```

  **OU**

  ```go
  carro := "Fusca" // Inferência de tipo
  ```

* **Tratamento de erros**

  ```go
  resultado, err := soma(-10, 10)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resultado)
  ```

* **Criação de uma "classe"**

  ```go
  type Product struct {
	  ID   string
	  Name string
  }
  ```

  ```go
  type Products struct {
	  Product []Product
  }
  ```

* **Instanciação de um "objeto"**

  ```go
	product := entity.Product{}

	product.ID = "abc"
	product.Name = "Fusca"
  ``` 

* **Criação de uma função de um "classe"**

  ```go
  func (p Products) Add(product Product) {
    // A adição de produto ocorre apenas no contexto da função (não altera a lista Product)
    p.Product = append(p.Product, product)
  }

  func main() {
    products := entity.Products{}
    product := entity.Product{}

    product.ID = "abc"
    product.Name = "Fusca"

    products.Add(product)

    fmt.Println(products)
  }
  ```

* **Ponteiro**

  ```go
  func (p *Products) Add(product Product) {
    // A adição é feita na lista Product, a partir da referência em memória
    p.Product = append(p.Product, product)
  }
  ```

  ```go
  func NewProduct() *Product {
    return &Product{
      ID: "",
    }
  }
  ```

* **Acessar atributos de um "objeto"**

  ```go
  fmt.Println(product.Name)
  ``` 

* **Serialização** 

  ```go
  type Product struct {
    ID   string `json:"-"` // campo ignorado ("-")
    Name string `json:"name"` // campo com a chave = name ("name")
  }
  ```

### Projeto **appgo**

* **Criação do Dockerfile**

  ```dockerfile
  FROM golang:1.16.0-stretch

  WORKDIR /go/src/

  ENV PATH="/go/bin:${PATH}"
  ENV CGO_ENABLED=0

  RUN go install github.com/spf13/cobra/cobra@latest

  CMD ["tail", "-f", "/dev/null"]
  ```

* **Criação do Docker Compose**

  ```yml
  version: "3"

  services:
    appgo:
      build: .
      container_name: appgo
      volumes:
      - .:/go/src/
      ports:
      - 8585:8585
  ``` 

* **Criação dos container com docker-compose**: `$ docker-compose up -d`

* **Entrar em modo interativo no container _goapp_**: `$ docker exec -it appgo bash`

* **Executar um arquivo `.go`**: `$ go run <arquivo>`

* **Criação de um executável para qualquer sistema operacional**

  * Executável da máquina em uso: `$ go build <arquivo>`

  * Executável do Windows: `$ GOOS=windows go build <arquivo>`

* **Criar o gerenciador de dependências do Golang** 
  
  * Sintaxe: `$ go mod init <URL do repositório>`

  * Exemplo: `$ go mod init github.com/Aprendendo-programacao/imersao-full-stack-and-full-cycle`

* **Instalar um dependência no Go mod**

  * Sintaxe: `$ go get github.com/<nome do repositório>`

  * Exemplo: `$ go get github.com/satori/go.uuid`

* **Instalar todas as dependências do `go mod`**: `$ go mod tidy`

* **Transformar um aplicação em CLI (cobra)**:

  * Sintaxe:

  * Exemplo: `$ cobra init --pkg-name=github.com/Aprendendo-programacao/imersao-full-stack-and-full-cycle`