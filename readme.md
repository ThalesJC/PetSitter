# 🐾 PetSitter

> Cuide da saúde do seu pet com organização e tranquilidade.  
> O **Petsitter** é o app que te ajuda a não esquecer nenhum cuidado importante com seus animais de estimação.  
> Registre vacinas, vermífugos, antiparasitários e tratamentos — o Petsitter cria automaticamente a agenda de medicamentos com lembretes nos horários certos.  
> Tudo salvo, sincronizado e sempre à mão. Seu pet saudável, você mais tranquilo.

---

## ✨ Funcionalidades (Versão 1)

- Cadastro e autenticação de usuários
- Cadastro de múltiplos pets por usuário
- Edição, visualização e remoção de pets
- Middleware de autenticação com JWT
- Agenda automática de medicamentos
- Lembretes com suporte a notificações no mobile (KMP)

---

## 🛠 Tecnologias Utilizadas

- [**Go (Golang)**](https://go.dev/) — Back-end
- [**Fiber**](https://gofiber.io/) — Web framework
- [**GORM**](https://gorm.io/) — ORM para Go
- [**PostgreSQL**](https://www.postgresql.org/) — Banco de dados relacional
- [**JWT**](https://en.wikipedia.org/wiki/JSON_Web_Token) — Autenticação segura
- [**Kotlin Multiplatform**](https://www.jetbrains.com/kotlin-multiplatform/) — Front-end mobile (Android)

---

## 🔐 Rotas da API

### Usuários & Autenticação

| Método | Rota                        | Ação                                      |
|--------|-----------------------------|-------------------------------------------|
| POST   | `/api/v1/user`              | Criar novo usuário                        |
| POST   | `/api/v1/auth/login`        | Fazer login                               |
| POST   | `/api/v1/auth/refresh`      | Renovar token de acesso                   |
| GET    | `/api/v1/user/me`           | Obter dados do usuário autenticado        |
| PUT    | `/api/v1/user/me`           | Atualizar dados pessoais do usuário       |
| DELETE | `/api/v1/user/me`           | Deletar conta do usuário autenticado      |

### Pets

| Método | Rota                   | Ação                                 |
|--------|------------------------|--------------------------------------|
| POST   | `/api/v1/pets`         | Cadastrar novo pet                   |
| GET    | `/api/v1/pets`         | Listar todos os pets do usuário      |
| GET    | `/api/v1/pets/:id`     | Obter dados de um pet específico     |
| PUT    | `/api/v1/pets/:id`     | Atualizar informações de um pet      |
| DELETE | `/api/v1/pets/:id`     | Remover um pet                       |

---
