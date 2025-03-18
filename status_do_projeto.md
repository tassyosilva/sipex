# Status do Projeto: Sistema de Gestão de Requisições Periciais - IML/RR


## Visão Geral


Este sistema visa digitalizar o fluxo de trabalho do Instituto Médico Legal da Polícia Civil de Roraima, substituindo o atual processo baseado em papel por uma solução digital integrada.


Sistema em português brasileiro, sem opções em inglês.


## Fluxo de Trabalho Atual (a ser digitalizado)


1. Requisição da autoridade (delegado, MP ou Juiz, etc.) chega ao IML em 3 vias.
2. É feito um protocolo e gerado um número de protocolo dessa entrada da requisição e encaminhado para o Diretor do IML.
3. O Diretor recebe essa requisição e envia para o perito disponível no dia.
4. O Perito disponível recebe a requisição, elabora o laudo pericial (exame de integridade física, exame cadavérico, etc.) e assina o laudo.
5. Após confecção e assinatura do laudo ele envia para o setor de envios do IML.
6. Esse setor de envios digitaliza o laudo e retorna para a autoridade que fez a requisição.


## Cronograma de Desenvolvimento


### Fase 1: Estrutura Básica e Autenticação


- [x] Criação da estrutura de pastas do projeto
- [x] Configuração inicial do backend (Go)
- [x] Configuração inicial do frontend (React/TypeScript)
- [x] Implementação da conexão com banco de dados PostgreSQL
- [x] Implementação do sistema de autenticação (JWT)
- [x] Criação da tela de login
- [x] Criação do dashboard básico


### Fase 2: Gestão de Requisições


- [ ] Modelagem da entidade Requisição
- [ ] Implementação do CRUD de Requisições no backend
- [ ] Criação das telas de listagem, cadastro e detalhes de Requisições
- [ ] Implementação do fluxo de recebimento e protocolo de Requisições
- [ ] Implementação do fluxo de encaminhamento para o Diretor


### Fase 3: Gestão de Laudos Periciais


- [ ] Modelagem da entidade Laudo
- [ ] Implementação do CRUD de Laudos no backend
- [ ] Criação das telas de elaboração e visualização de Laudos
- [ ] Implementação do fluxo de elaboração de Laudos pelos peritos
- [ ] Implementação do fluxo de envio de Laudos para o setor de envios


### Fase 4: Finalização e Refinamentos


- [ ] Implementação de relatórios e estatísticas
- [ ] Refinamento da interface do usuário
- [ ] Testes integrados
- [ ] Documentação do sistema
- [ ] Implantação e treinamento


## Status Atual


**Data:** 2023-11-27


**Fase:** Fase 1 - Estrutura Básica e Autenticação


**Progresso:** Fase 1 concluída. Sistema de autenticação implementado, tela de login e dashboard básico criados. Interface de login melhorada com as cores padrão da Polícia Civil (branco, cinza, preto e dourado) e adicionada mensagem de erro para credenciais inválidas.


## Arquitetura do Sistema


### Backend (Go)


#### Autenticação JWT
- Implementado no pacote `/internal/auth/jwt.go`
- Geração e validação de tokens JWT
- Armazenamento de informações do usuário no token (ID, nome, papel)


#### Conexão com o Banco de Dados
- Implementado em `/internal/database/config.go`
- Utiliza o driver PostgreSQL (`github.com/lib/pq`)
- Configuração via variáveis de ambiente com valores padrão para desenvolvimento
- Função `ConnectDB()` para estabelecer conexão com o banco


#### Handlers (Controladores)
- Implementados em `/internal/handlers/`
- Padrão de injeção de dependência (repositories são injetados nos handlers)
- Estrutura típica:
1. Definição do tipo handler (struct)
2. Função construtora "New" que recebe repositories necessários
3. Métodos para manipular diferentes endpoints HTTP


#### Middleware
- Implementado em `/internal/middleware/auth_middleware.go`
- Funções principais:
- `AutenticarMiddleware`: Valida token JWT em requisições
- `VerificarPapel`: Verifica se o usuário tem o papel necessário


#### Modelos
- Implementados em `/internal/models/`
- Definem estruturas de dados usadas na aplicação
- Contêm tags para serialização/desserialização JSON


#### Repositories
- Implementados em `/internal/repository/`
- Encapsulam todas as operações com o banco de dados
- Estrutura típica:
1. Definição do tipo repository (struct com campo `db *sql.DB`)
2. Função construtora "New" que recebe conexão com banco
3. Métodos para operações CRUD e consultas específicas


### Padrões de Fluxo de Dados no Backend
1. Requisição HTTP recebida
2. Middlewares aplicados (autenticação, verificação de papel)
3. Handler correspondente processa a requisição
4. Repository executa operações no banco de dados
5. Handler formata e envia a resposta HTTP


### Estrutura do Backend
```
backend/
├── cmd
│ └── api
│ └── main.go
├── go.mod
├── go.sum
└── internal
├── auth
│ └── jwt.go
├── database
│ └── config.go
├── handlers
│ ├── auth_handler.go
│ └── usuario_handler.go
├── middleware
│ └── auth_middleware.go
├── models
│ └── user.go
└── repository
└── usuario_repository.go
```


### Frontend (React/TypeScript)


#### Componentes Principais
- Layout compartilhado com menu lateral (`Layout.tsx`)
- Rotas protegidas com autenticação (`PrivateRoute.tsx`)
- Páginas principais: Login, Dashboard, Acesso Negado


#### Configuração da API
- Implementado em `/web/frontend/src/config/api.ts`
- Define a URL base da API com detecção de ambiente (desenvolvimento/produção)


#### Roteamento
- Implementado em `App.tsx` usando React Router
- Rotas protegidas com componente `PrivateRoute`
- Layout compartilhado para páginas internas


#### Autenticação no Frontend
- Serviço de autenticação em `/web/frontend/src/services/authService.ts`
- Tokens JWT armazenados no `sessionStorage`
- Enviados no cabeçalho `Authorization` em requisições à API
- Redirecionamento para login quando token inválido


#### Tema e Estilização
- Tema claro implementado com Material UI
- Configuração em `/web/frontend/src/theme/lightTheme.ts`
- Cores padrão da Polícia Civil (branco, cinza, preto e dourado) aplicadas ao tema


#### Formulários
- Uso de estado React para gerenciar dados de formulário
- Validação de entrada
- Feedback visual para erros e sucesso


### Padrões de Comunicação Frontend-Backend


#### Endpoints da API
- Prefixo padrão: `/api`
- Formato de resposta: JSON
- Autenticação: Bearer Token (JWT)


#### Chamadas à API
- Uso de `fetch` API para requisições HTTP
- Headers padrão:
- 'Content-Type': 'application/json'
- 'Authorization': `Bearer ${sessionStorage.getItem('token')}`


#### Tratamento de Erros
- Verificação de `response.ok` para detectar erros HTTP
- Extração de mensagens de erro do corpo da resposta
- Feedback visual para o usuário


### Estrutura do Frontend
```
web/frontend/src/
├── App.css
├── App.tsx
├── assets
│ └── react.svg
├── components
│ ├── Layout.tsx
│ └── PrivateRoute.tsx
├── config
│ └── api.ts
├── hooks
├── index.css
├── main.tsx
├── pages
│ ├── AcessoNegado.tsx
│ ├── Dashboard.tsx
│ └── Login.tsx
├── services
│ └── authService.ts
├── styles
├── theme
│ └── lightTheme.ts
├── types
│ └── usuario.ts
└── vite-env.d.ts
```


## Próximos Passos


1. Iniciar a Fase 2 com a modelagem da entidade Requisição
2. Implementar o CRUD de Requisições no backend
3. Criar as telas de listagem, cadastro e detalhes de Requisições no frontend
4. Implementar o fluxo de recebimento e protocolo de Requisições


## Tecnologias Utilizadas


- **Backend:** Go (Golang)
- **Frontend:** React com TypeScript
- **Banco de Dados:** PostgreSQL
- **Autenticação:** JWT (JSON Web Token)
- **UI Components:** Material UI


## Atualizações


- Criada a estrutura básica de pastas para o backend e frontend
- Criado o arquivo go.mod com as dependências iniciais do backend
- Implementada a conexão com o banco de dados PostgreSQL
- Criado o modelo de usuário
- Implementada a autenticação JWT
- Criado o arquivo main.go para iniciar o servidor
- Implementado o repositório de usuários
- Criado o handler de autenticação
- Implementado o middleware de autenticação
- Criado o handler de usuários
- Configuradas as rotas da API
- Configurado o frontend com React e TypeScript
- Implementado o tema claro com Material UI
- Criada a tela de login
- Implementado o serviço de autenticação no frontend
- Criado o componente de rotas protegidas
- Implementado o layout com menu lateral
- Criado o dashboard básico
- Adicionado usuário administrador padrão para primeiro acesso
- Melhorada a interface de login com as cores padrão da Polícia Civil (branco, cinza, preto e dourado)
- Adicionada mensagem de erro para credenciais inválidas na tela de login
