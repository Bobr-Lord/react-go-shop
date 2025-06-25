# React Go Shop
<p align="center">
  <img src="./img/img.png" alt="preview" />
</p>

<p align="center">
  <img alt="Go" src="https://img.shields.io/badge/Backend-Go-00ADD8?logo=go&logoColor=white" />
  <img alt="React" src="https://img.shields.io/badge/Frontend-React-61DAFB?logo=react&logoColor=white" />
  <img alt="PostgreSQL" src="https://img.shields.io/badge/Database-PostgreSQL-336791?logo=postgresql&logoColor=white" />
  <img alt="Docker" src="https://img.shields.io/badge/Container-Docker-2496ED?logo=docker&logoColor=white" />
  <img alt="JWT" src="https://img.shields.io/badge/Auth-JWT-orange?logo=jsonwebtokens&logoColor=white" />
  <img alt="REST API" src="https://img.shields.io/badge/API-REST-blue?logo=api&logoColor=white" />
</p>


Интернет-магазин с фронтендом на React и бэкендом на Go. Проект демонстрирует архитектуру с двумя независимыми Go-сервисами (например, `auth` и `shop`), и фронтендом, взаимодействующим с ними через REST API.

---

---

![Превью сайта](./img/img.png)

---

## 🚀 Возможности

- Регистрация и авторизация (JWT + refresh)
- Каталог товаров и корзина
- Проверка email, защита через middleware
- Взаимодействие через REST API между фронтом и двумя Go-сервисами
- Docker-окружение для полноценного запуска

---

## 🐳 Быстрый старт с Docker

1. Убедитесь, что установлены `Docker` и `Docker Compose`

2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/Bobr-Lord/react-go-shop.git
   cd react-go-shop
   ```

3. Соберите и запустите проект:
   ```bash
   docker compose up --build
   ```

4. Приложение будет доступно:
   - Фронтенд: [http://localhost:80](http://localhost:80)

---

## 📁 Структура проекта

```
react-go-shop/
├── backend/
│   ├── auth/            # Auth-сервис
│   ├── shop/            # Магазин-сервис
│   └── docker-compose.yml
├── frontend/            # React-приложение
│   └── src/
├── nginx/               # Конфигурация nginx для reverse proxy
├── keys/                # RSA-ключи (в .gitignore)
└── README.md
```

---

## ⚙️ Конфигурация

Перед запуском убедитесь, что у вас есть:

- `keys/private.pem`, `keys/public.pem` (используются для подписи JWT)


---

## 🧪 Возможности для улучшения

- Полноценная админка
- CI/CD пайплайн
- Хранение изображений и загрузка товаров
- OAuth авторизация (Google, GitHub)
- GraphQL-версия API
- Юнит- и интеграционные тесты

---

## 📝 Лицензия

MIT © 2025 @Bobr-Lord