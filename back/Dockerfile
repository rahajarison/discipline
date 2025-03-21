# Étape 1 : Utiliser l'image officielle Go comme image de base
FROM golang:1.24 AS builder

# Définir le répertoire de travail à l’intérieur du container
WORKDIR /back

# Copier les fichiers de votre projet dans le container
COPY . .

# Télécharger les dépendances et compiler l’application
RUN go mod tidy && go build -o main .

# Étape 2 : Créer une image plus légère pour l'exécution
FROM alpine:latest

# Installer les certificats SSL (nécessaire pour certaines connexions réseau)
RUN apk --no-cache add ca-certificates

# Définir le répertoire de travail
WORKDIR /root/

# Copier l'exécutable compilé depuis le builder
COPY --from=builder /back/main .

# Exposer le port que l'application va écouter (ex. 8080)
EXPOSE 8080

# Lancer l'application
CMD ["./main"]
