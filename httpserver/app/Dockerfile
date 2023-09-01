# Установка базового образа
FROM nginx:latest

# Копирование контента в приложение
COPY . /usr/share/nginx/html

# Установка порта
EXPOSE 80

# Основная команда
CMD ["nginx", "-g", "daemon off;"]

FROM base-image

COPY app/app.conf /path/to/app.conf
