package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Port sur lequel le démon Plexus Node écoute les requêtes
	port := "2009"

	// Démarrage du serveur TCP sur le port spécifié
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
		os.Exit(1)
	}
	fmt.Println("Le démon Plexus Node est en écoute sur le port", port)

	// Boucle infinie pour accepter les connexions entrantes
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
			continue
		}
		go handleRequest(conn)
	}
}

// Fonction pour gérer les requêtes du panneau de contrôle
func handleRequest(conn net.Conn) {
	defer conn.Close()

	// Lire la requête du panneau de contrôle
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la requête:", err)
		return
	}

	// Convertir les données lues en une chaîne
	request := string(buffer[:n])

	// Traitement de la requête (à compléter en fonction des besoins)
	// Exemple : Afficher la requête reçue
	fmt.Println("Requête reçue:", request)

	// Envoyer une réponse au panneau de contrôle
	response := "Réponse du démon Plexus Node\n"
	conn.Write([]byte(response))
}
