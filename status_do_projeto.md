# Status do Projeto: Sistema de Gestão de Requisições Periciais - IML/RR

## Visão Geral
Este sistema visa digitalizar o fluxo de trabalho do Instituto Médico Legal da Polícia Civil de Roraima, substituindo o atual processo baseado em papel por uma solução digital integrada.
Sistem em português brasileiro. sem opções em inglês.

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
- [ ] Configuração inicial do backend (Go)
- [ ] Configuração inicial do frontend (React/TypeScript)
- [ ] Implementação da conexão com banco de dados PostgreSQL
- [ ] Implementação do sistema de autenticação (JWT)
- [ ] Criação da tela de login
- [ ] Criação do dashboard básico

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
**Data:** 2023-11-20  
**Fase:** Fase 1 - Estrutura Básica e Autenticação  
**Progresso:** Iniciando o projeto com a criação da estrutura de pastas.

## Próximos Passos
1. Configurar o ambiente de desenvolvimento backend
2. Implementar a conexão com o banco de dados
3. Desenvolver o sistema de autenticação básico
4. Criar a tela de login no frontend

## Tecnologias Utilizadas
- **Backend:** Go (Golang)
- **Frontend:** React com TypeScript
- **Banco de Dados:** PostgreSQL
- **Autenticação:** JWT (JSON Web Token)
- **UI Components:** Material UI

Vamos atualizar o status do projeto para refletir o progresso:
- Criada a estrutura básica de pastas para o backend e frontend
- Criado o arquivo go.mod com as dependências iniciais do backend

## Próximos Passos
1. Implementar a conexão com o banco de dados PostgreSQL
2. Criar o modelo de usuário
3. Implementar a autenticação JWT" >> status_do_projeto.md

## Atualização: 2023-11-21
- Implementada a conexão com o banco de dados PostgreSQL
- Criado o modelo de usuário
- Implementada a autenticação JWT
- Criado o arquivo main.go para iniciar o servidor

## Próximos Passos
1. Implementar o repositório de usuários
2. Criar o handler de autenticação
3. Implementar o middleware de autenticação
4. Iniciar a configuração do frontend

## Atualização: 2023-11-23
- Implementado o repositório de usuários
- Criado o handler de autenticação
- Implementado o middleware de autenticação
- Criado o handler de usuários
- Configuradas as rotas da API

## Próximos Passos
1. Iniciar a configuração do frontend (com tema claro)
2. Implementar a tela de login
3. Implementar o dashboard básico
4. Implementar a gestão de usuários no frontend
