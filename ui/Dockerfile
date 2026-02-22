# Etapa 1: Construcción (Node.js)
FROM node:22-alpine AS build
WORKDIR /app

# Instalar dependencias primero para aprovechar la caché de Docker
COPY package*.json ./
RUN npm install --legacy-peer-deps

# Copiar el resto del código y compilar
COPY . .
# RUN npm run build -- --configuration production
RUN npx ng build --configuration production

# Etapa 2: Servidor Web (Nginx)
FROM nginx:stable-alpine
WORKDIR /usr/share/nginx/html

# Limpiar archivos por defecto de nginx
RUN rm -rf ./*

# Copiar los archivos compilados desde la etapa de construcción
# Nota: Ajusta 'dist' si tu carpeta de salida tiene otro nombre
COPY --from=build /app/dist/*/browser/ /usr/share/nginx/html/

# Copiar configuración personalizada de Nginx para Angular (ver paso 2)
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]