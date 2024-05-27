## InternTest01



### Prerequisites

List the prerequisites required to run your project. Include any software or tools that need to be installed.

- Go Language | [Document](https://golang.org/)
  ```bash
  brew install go
  ```
- Air | Go Package | [Document](https://github.com/cosmtrek/air#installation)
  ```bash
  go install github.com/cosmtrek/air@latest
  ```
- Docker | [Document](https://www.docker.com/)
- Docker Compose | [Document](https://docs.docker.com/compose/)
  ```bash
  for mac
  brew install docker
  brew install docker-compose
  ```

### Installation

How to installing and setting up your service.

1. Clone repository

   ```
   git clone https://github.com/KhingNoi/internTest01.git
   ```

2. Install dependency package

   ```bash
   go get ./cmd
   ```

3. For Localhost | Use `air` package to build and run the project

   ```bash
   air
   ```

### Migration

1. Install go migate

   ```bash
   go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.0
   ```

   if there's a stable and newer version you better install that instead.

2. Run migreate to check versioning should comes out a version number

   ```bash
   migrate -version
   ```

3. Create migration file by using command

   ```bash
   make create-migrations <file name>
   ```

   Up and down file will created empty you should add on migration code yourself (gpt able to do this)
   
5. Run migrations using command

   ```
   make migrations-up
   make migrations-down
   ```

   \*If you run migration up and “no change” even if in your db not have given table or some error occorred
   run migration down to clear all cache before you run migration up


### Running the Containers

1. start the container(s) using docker or docker Compose. Include any necessary commands or configuration files.

   ```bash
   docker-compose up -d
   ```

2. If using docker directly
   [Document](https://hub.docker.com/r/mysql/mysql-server/)
   Pull MySQL image (if not already pulled)

   ```bash
   docker pull mysql/mysql-server:tag
   ```

   Start MySQL container

   ```bash
   docker run --name=mysql1 -d mysql/mysql-server:tag
   ```

### FAQ

1. When we system not found `air` command, we should set PATH environment variable in `~/.bashrc`, `~/.bash_profile`, `~/.zshrc`
   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```
   
## API Specification

<img width="3632" alt="Untitled" src="https://github.com/KhingNoi/internTest01/assets/157562653/6e8b031a-f8b3-4319-a429-2593f11d86fd">


- Orange Background => commit
- White Background => uncommit

## EER Diagram
![image](https://github.com/KhingNoi/internTest01/assets/157562653/d3993fb0-d94a-4373-bd73-8019320dc395)

- table
![image](https://github.com/KhingNoi/internTest01/assets/157562653/79b7f70b-f846-4860-91e6-9fe8c8eca1e3)


### [figma click here](https://www.figma.com/board/OErXYD0odN0JQzE068A77f/Untitled?node-id=0-1&t=yvQuoEdeL58iXdOz-0)



