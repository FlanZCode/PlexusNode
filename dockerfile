# Utilise l'image officielle de Go depuis Docker Hub
FROM golang:1.20.6

# Définit le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copie le code source de l'application dans le conteneur
COPY . .

# Construit l'exécutable de l'application avec un nom de sortie personnalisé
RUN go build -o plexus-node-exec .

# Expose le port sur lequel le démon écoute
EXPOSE 2009

# Commande pour démarrer le démon Plexus Node dans le conteneur
CMD ["./plexus-node-exec"]
