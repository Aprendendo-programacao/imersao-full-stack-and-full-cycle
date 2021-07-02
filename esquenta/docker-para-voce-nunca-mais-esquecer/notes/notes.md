# Docker para você nunca mais esquecer

### Container vs Máquina virtual

* **Máquina virtual**

  * Máquina com sistema operacional completo

  * Boot é demorado

  * Gasto de muito recurso (memória, cpu, armazenamento)

* **Container**

  * Um **único processo** rodando na máquina local e gerenciado por com _Container Runtime_ (Docker)

  * Cada processo é **isolado** (recurso do Linux)

  * Muito rápido e leve

### Comandos Docker

> `$ docker` é uma CLI (Commando Line Interface) que chama o Docker Engine (responsável pela comunicação com o _Docker daemon_)

* **Criar um container**

  * Sintaxe: `$ docker run <imagem>`

  * Exemplo: `$ docker run nginx`

    > **nginx** é um servidor web de proxy reverso, ou seja, ele repassa o tráfego de rede recebido para um conjunto de servidores

* **Listagem dos containers em execução** 

  * Sintaxe: `$ docker ps` ou `$ docker container ls`

* **Listagem de todos os container**: 

  * Sintaxe: `$ docker ps -a` ou `$ docker container ls -a`

* **Binding de portas (máquina local -> container)**

  * Sintaxe: `$ docker run -p <porta maquina local>:<porta container>`

  * Exemplo: `$ docker run -p 8080:80 nginx`

* **Remover um container pelo ID/nome**

  * Sintaxe: `$ docker rm <ID/nome>`

  * Exemplo: `$ docker rm 069961c31294`

* **Remover, de forma forçada, um container pelo ID/nome**

  * Sintaxe: `$ docker rm <ID/nome> -f`

  * Exemplo: `$ docker rm 069961c31294 -f`

* **Executar um comando em um container existente**

  * Sintaxe: `$ docker exec -it <container> <comando>`

    > **-it** = modo interativo

  * Exemplo: `$ docker exec -it 069961c31294 bash`

* **Rodar um container em segundo plano**

  * Sintaxe: `$ docker run -d <imagem>`

    > **-d** = _detach_, ou seja, separar o comando de execução de um novo container do terminal em uso

  * Exemplo: `$ docker run -d -p 8080:80 nginx`

* **Remover automaticamente o container quando ele sair (`status = EXITED`)**

  * Sintaxe: `$ docker run --rm <imagem>`

  * Exemplo: `$ docker run --rm -d -p 8080:80 nginx`

* **Criar uma imagem a partir de um Dockerfile**

  * Sintaxe: `$ docker build -t <nome> <contexto>`

    * **Padrão para nomear imagens**: `<usuário>/<nome>:<tag>`

      * **usuário** = nome do usuário no Docker Hub

      * **nome da imagem** = nome da imagem

      * **tag** = versão da imagem (v1, v2, latest)

  * Exemplo: `$ docker build -t imgabreuw/esquenta3-nginx:v1 .`

* **Criar os containers, em segundo plano, com Docker Compose**

  * Sintaxe: `$ docker-compose up -d`

* **Recriar as imagens do Docker Compose**

  * Sintaxe: `$ docker-compose up -d --build`

* **Listagem dos containers criados a partir do docker-compose**

  * Sintaxe: `$ docker-compose ps`

* **Remover todos os containers presentes no Docker Compose**

  * Sintaxe: `$ docker-compose down`

### Dockerfile

* **Definição**: é uma "receita de bolo" para gerar imagens próprias

* **Propriedades**

  * `FROM` = definir a imagem base

    > IMPORTANTE: é um boa prática indicar a tag (versão) da imagem base, pois assim é possível garantir a idempotência

  * `COPY` = copiar arquivo(s) da maquina local para o container

* **Exemplo de um Dockerfile**

  ```dockerfile
  FROM nginx:latest 
  COPY index.html /usr/share/nginx/html/index.html
  ```

### Volume

* **Definição**: são diretórios, da máquina local, compartilhados com os do container, ou seja, **alterações locais refletem no container** 

### Docker Compose

* **Definição**: permitir que vários container subam ao mesmo tempo

* **Propriedades**

  * `version` = versão do Docker Compose (versão atual: 3)

  * `services` = containers que serão criados

  * `build` = caminho até o Dockerfile que será utilizado para a criação das imagens

  * `ports` = portas liberadas

    ```yml
    ports:
      - <porta maquina local>:<porta container>
    ```

  * `volumes` = diretórios que serão compartilhados com o container

    ```yml
    volumes:
      - <diretório local>:<diretório container>
    ```

* **Exemplo de um docker-compose**

  ```yml
  version: "3"

  services:
    web: 
      build: nginx/
      ports:
        - 8080:80
      volumes: 
        - ./nginx/:/usr/share/nginx/html/

    nodeapp:
      build: nodeapp/
      ports:
        - 3000:3000
      volumes:
        - ./nodeapp/:/usr/src/app
  ```