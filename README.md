# api-golang
Esta documentação tem como objetivo descrever a estrutura da API em Golang e como ela foi desenvolvida. A aplicação utiliza diversas bibliotecas, incluindo o GORM, Chi, Bcrypt e JWT, para criar uma estrutura robusta e segura.

### GORM
O GORM é uma biblioteca ORM (Object-Relational Mapping) para Go que permite interagir com bancos de dados de forma mais simples e intuitiva. Ele é utilizado na API para realizar operações no banco de dados PostgreSQL.

### Chi
O Chi é um roteador HTTP em Go que oferece diversas funcionalidades, incluindo middleware e roteamento dinâmico. Ele é utilizado na API para gerenciar as rotas e as requisições HTTP.

### Bcrypt
O Bcrypt é uma biblioteca de hashing de senhas que oferece um alto nível de segurança. Ele é utilizado na API para proteger as senhas dos usuários armazenadas no banco de dados.

### JWT
O JWT (JSON Web Token) é um padrão aberto que define um formato compacto e autocontido para transmitir informações entre as partes como um objeto JSON. Ele é utilizado na API para autenticar as requisições HTTP.
SOLID

A API foi desenvolvida seguindo os princípios do SOLID, que é um conjunto de princípios de programação orientada a objetos que visam tornar o código mais legível, sustentável e fácil de manter.
#
# main.go
O arquivo main.go é o ponto de entrada da aplicação. Ele é responsável por carregar as variáveis de ambiente, definir as rotas da API, conectar ao banco de dados e iniciar o servidor para ouvir as solicitações HTTP.

A função main() começa carregando as variáveis de ambiente por meio da função Load() no arquivo de configuração (config.go). Em seguida, é exibida uma mensagem de log informando que a API está iniciando.

Depois, o arquivo main.go cria um novo roteador utilizando o pacote chi e compila todas as rotas da aplicação por meio do módulo "router".

Em seguida, o arquivo main.go se conecta ao banco de dados por meio da função Connect() do módulo "database". Se a conexão for estabelecida com sucesso, é exibida uma mensagem de log informando que a aplicação está conectada ao banco de dados. Se houver algum erro na conexão, a função main() exibe uma mensagem de log informando o erro e a aplicação é encerrada.

Por fim, é exibida outra mensagem de log informando que a API está ouvindo e servindo na porta definida pelas variáveis de ambiente. A função ListenAndServe() do pacote "http" é usada para iniciar o servidor HTTP e começar a ouvir as solicitações na porta definida. Certifique-se de revisar e atualizar essa documentação de acordo com as necessidades específicas do seu aplicativo.

```go
func main() {
	config.Load()
	utils.Log("Golang API Starting...")

	r := chi.NewRouter()
	r.Route("/", router.Routes)

	_, erro := database.Connect()
	if erro != nil {
		utils.Error(erro)
		return
	}
	utils.Log("Golang API Conected to Database.")

	utils.Log("Golang API Listening and Serving!")
	http.ListenAndServe(config.ApiPort, r)
}
```

# Auth
O diretório "auth" contém a lógica de autenticação do aplicativo. Ele é composto por dois diretórios principais: "controllers" e "routes".

O diretório "controllers" contém os controladores para lidar com as funcionalidades de autenticação do aplicativo, como login, refresh token e alteração de senha. Esses controladores são responsáveis por receber as requisições HTTP correspondentes e processá-las de acordo com a lógica de negócios do aplicativo.

O diretório "routes", por sua vez, contém a configuração das rotas para esses controladores. Isso significa que esse diretório é responsável por mapear as URLs das requisições HTTP para os respectivos controladores.

Ao separar as responsabilidades de configuração de rotas e lógica de negócios, o módulo "auth" torna o código mais organizado e fácil de manter. Além disso, essa separação permite que os controladores sejam reutilizados em diferentes rotas, caso necessário.


# Config
O diretório "config" possui o arquivo "config.go" é responsável por carregar as variáveis de ambiente do aplicativo. As variáveis de ambiente são usadas para armazenar informações confidenciais, como senhas de banco de dados ou chaves de API, que devem ser mantidas fora do código-fonte.

O arquivo "config.go" define valores padrão para as variáveis de ambiente caso elas não sejam definidas no ambiente em que o aplicativo está sendo executado. Isso garante que o aplicativo possa ser executado em diferentes ambientes sem precisar configurar manualmente as variáveis de ambiente.

## .env exemple
```env
# Database
DB_PORT=
DB_HOST=
DB_USER=
DB_NAME=
DB_PASS=

# Auth
SECRET_KEY_ACCESS=
SECRET_KEY_REFRESH=

# Api
API_PORT=
```


# Database
O diretório "database" contém o arquivo "database.go", que é responsável por configurar o ORM GORM para se conectar e interagir com um banco de dados PostgreSQL.

Para usar o GORM em nosso aplicativo, precisamos importar o pacote "gorm" e o driver de banco de dados "postgres". Além disso, precisamos definir as informações de conexão do banco de dados, como nome do host, porta, nome do banco de dados, nome de usuário e senha. Essas informações são definidas como variáveis de ambiente para que possam ser facilmente configuradas em diferentes ambientes.

Depois de definir as informações de conexão do banco de dados, podemos usá-las para abrir uma conexão com o banco de dados usando o GORM. A partir daí, podemos criar e interagir com tabelas no banco de dados usando os modelos definidos em nossa aplicação.


# Middlewares
O diretório "middlewares" contém os middlewares do aplicativo, que são funções que podem ser usadas para processar as requisições HTTP antes que elas sejam enviadas aos controladores.

Os middlewares são úteis para adicionar funcionalidades adicionais a uma rota, como autenticação, validação de entrada ou registro de log. Eles são executados antes dos controladores e podem ser usados para modificar a requisição ou a resposta.

O arquivo "ensureAuthenticated.go" é um middleware que garante que o usuário esteja autenticado antes de permitir que a requisição seja processada. Ele faz isso validando o token JWT enviado pelo cliente e, se o token for válido, retorna o ID do usuário na requisição.

A validação do token é importante porque garante que apenas usuários autenticados possam acessar determinados recursos no aplicativo. Isso ajuda a manter a segurança e a privacidade dos dados do usuário.


# Models
O diretório "models" contém os modelos das entidades do aplicativo e suas funções para tratamento. Os modelos são estruturas de dados que representam as entidades do aplicativo, como usuários, produtos, pedidos e assim por diante.

Os modelos são usados para interagir com o banco de dados e realizar operações de leitura, gravação, atualização e exclusão. Eles também podem conter funções personalizadas que são específicas para as entidades que representam. Essas funções podem ser usadas para realizar tarefas adicionais, como validação de dados, formatação de valores ou cálculos complexos.

Os modelos devem ser definidos de acordo com as necessidades do aplicativo e sua estrutura de banco de dados. Eles devem incluir todos os campos necessários para representar a entidade e seus relacionamentos com outras entidades.


# Modules
O diretório "modules" contém um diretório para cada entidade do aplicativo. Dentro de cada diretório de entidade, há um diretório "controllers", um diretório "routes" e um diretório "repository".

## Controllers
O diretório "controllers" contém os controladores da entidade, que são responsáveis por lidar com as solicitações HTTP e chamar as funções apropriadas para manipular os dados da entidade.

## Routes
O diretório "routes" contém as rotas HTTP da entidade, que são usadas para mapear as solicitações HTTP para as funções apropriadas nos controladores da entidade.

## Repository
O diretório "repository" contém os repositórios da entidade, que são responsáveis por manipular os dados da entidade no banco de dados. Isso pode incluir funções para criar, ler, atualizar e excluir registros da entidade.


# Router
O diretório "router" contém o arquivo "router.go", que é responsável por compilar as rotas dos demais subdiretórios de rotas e prepará-las para serem carregadas no arquivo "main.go".

O arquivo "router.go" cria e configura o roteador HTTP global do aplicativo. Ele também carrega todas as rotas definidas nos subdiretórios de rotas de cada entidade e as registra no roteador HTTP global. Isso permite que o servidor HTTP gerencie todas as rotas do aplicativo de forma centralizada.


# Utils
O diretório "utils" contém funções úteis que podem ser usadas em várias partes da API. Essas funções incluem manipulação de tokens JWT, criptografia e descriptografia de dados, e validação de entradas do usuário, entre outras.

Diferente dos outros diretórios, o "utils" não possui um propósito específico em relação a uma entidade ou funcionalidade do aplicativo. Em vez disso, ele contém funções genéricas que podem ser usadas em vários lugares.
