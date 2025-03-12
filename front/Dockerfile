# Utiliser l'image officielle de Node.js
FROM node:16-alpine AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier le package.json et package-lock.json (ou yarn.lock)
COPY package*.json ./

# Installer les dépendances
RUN npm install

# Copier tous les fichiers du projet
COPY . .

# Construire le projet Vue.js
RUN npm run build

# Utiliser une image plus légère pour servir le projet
FROM nginx:alpine

# Copier les fichiers de build dans le répertoire de Nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Exposer le port 80 pour Nginx
EXPOSE 80

# Lancer Nginx pour servir l'application
CMD ["nginx", "-g", "daemon off;"]
