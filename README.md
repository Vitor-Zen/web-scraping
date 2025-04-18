# 🕷️ Web Scraping (Full-stack App)

Full-stack web scraping application powered by a **Go backend** and a **Next.js frontend** for data presentation and user interaction.

---

## 📁 Estrutura do projeto

```
web-scraping/
├── backend/                 # API e scraping com Go
│   ├── cmd/server/          # Ponto de entrada (main.go)
│   └── internal/            # Auth, tracking, scraping, DB, etc
│
├── frontend/                # Interface do usuário (Next.js)
│   ├── pages/
│   └── components/
│
├── go.mod
└── README.md
```

---

## ✅ Funcionalidades previstas

- Login com conta Google
- Autorização manual de acesso por administrador
- Cadastro de palavras-chave para monitoramento de produtos
- Scraping periódico em sites definidos (Kabum, Pelando, Pichau, etc.)
- Consulta de preços consolidados por palavra-chave
- Exclusão de trackings pelo próprio usuário

---

## 🚀 Como rodar o projeto

### Backend

```
bash
cd backend
go run ./cmd/server
```

### Frontend

```
cd frontend
npm install
npm run dev
```

🔐 Requisitos

- Go v1.20 ou superior
- Node.js v18 ou superior
- npm
- VS Code com extensões para Go e React
- Conta Google configurada no Google Cloud Console para usar o login OAuth2
- SQLite (para testes locais) ou PostgreSQL (em produção)

## 👤 Acesso

> Os usuários se autenticam com o Google, mas precisam ser **autorizados manualmente** para acessar o sistema.

📌 Status

> Projeto em desenvolvimento.

🧑🏻‍💻 Autor
Desenvolvido por Vitor Zen
