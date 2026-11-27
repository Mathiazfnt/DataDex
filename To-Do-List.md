# DataDex - To-do list

## 00 — Préparation
- [ ] [Commun] Extraire et intégrer l'architecture ZIP
- [ ] [Commun] Retirer les extensions .todo
- [ ] [Commun] Initialiser les déclarations de packages Go
- [ ] [Commun] Vérifier Go et les dépendances
- [ ] [Commun] Configurer le dépôt Git et les accès
- [ ] [Commun] Vérifier le .gitignore
- [ ] [Commun] Mettre à jour le README pour Battle.net
- [ ] [Commun] Confirmer l'autorisation de travailler à deux

## 01 — Conception et exploration
- [ ] [Gautier] Créer le client développeur Battle.net
- [ ] [Gautier] Tester OAuth dans Postman
- [ ] [Gautier] Tester les index et détails des races
- [ ] [Gautier] Tester les index et détails des classes
- [ ] [Gautier] Tester les index et détails des familiers
- [ ] [Gautier] Identifier les propriétés JSON réellement disponibles
- [ ] [Mathias] Concevoir les maquettes des pages
- [ ] [Mathias] Définir les routes locales
- [ ] [Commun] Valider les structures et contrats de fonctions communs

## 02 — Configuration et OAuth
- [ ] [Gautier] Déclarer la structure Config
- [ ] [Gautier] Créer LoadConfig
- [ ] [Gautier] Charger les variables d'environnement
- [ ] [Gautier] Contrôler les paramètres obligaGautierres
- [ ] [Gautier] Déclarer TokenResponse
- [ ] [Gautier] Développer RequestAccessToken
- [ ] [Gautier] Décoder la réponse OAuth
- [ ] [Gautier] Calculer l'expiration du token
- [ ] [Gautier] Développer TokenManager
- [ ] [Gautier] Développer IsTokenValid
- [ ] [Gautier] Développer GetAccessToken
- [ ] [Gautier] Implémenter le renouvellement automatique
- [ ] [Gautier] Protéger les accès concurrents au token
- [ ] [Gautier] Tester l'authentification

## 03 — Serveur HTTP et templates
- [ ] [Mathias] Initialiser le serveur HTTP
- [ ] [Mathias] Configurer les routes
- [ ] [Mathias] Créer le système de rendu HTML
- [ ] [Mathias] Créer le header commun
- [ ] [Mathias] Créer le footer commun
- [ ] [Mathias] Créer la page d'accueil
- [ ] [Mathias] Créer la page des collections
- [ ] [Mathias] Servir les ressources statiques
- [ ] [Mathias] Créer les premières pages d'erreur
- [ ] [Mathias] Tester les routes

## 04 — Client API
- [ ] [Gautier] Configurer le client HTTP et son timeout
- [ ] [Gautier] Créer ApiRequest
- [ ] [Gautier] Construire les URL API
- [ ] [Gautier] Ajouter namespace et locale
- [ ] [Gautier] Ajouter le header Bearer
- [ ] [Gautier] Contrôler les statuts HTTP
- [ ] [Gautier] Gérer les erreurs réseau
- [ ] [Gautier] Gérer les réponses 401 avec un seul nouvel essai
- [ ] [Gautier] Gérer les erreurs 429 et 5xx
- [ ] [Gautier] Tester le client API

## 05 — Collections
- [ ] [Gautier] Déclarer ResourceReference
- [ ] [Gautier] Déclarer Race et RaceIndex
- [ ] [Gautier] Récupérer l'index des races
- [ ] [Gautier] Récupérer les détails des races
- [ ] [Gautier] Construire la collection complète des races
- [ ] [Mathias] Afficher la collection des races
- [ ] [Mathias] Afficher les détails d'une race
- [ ] [Commun] Valider la première intégration API et HTML
- [ ] [Gautier] Déclarer Class et ClassIndex
- [ ] [Gautier] Récupérer les classes et leurs détails
- [ ] [Mathias] Afficher les classes et leurs détails
- [ ] [Gautier] Déclarer Pet et PetIndex
- [ ] [Gautier] Récupérer les familiers et leurs détails nécessaires
- [ ] [Mathias] Afficher les familiers et leurs détails

## 06 — Recherche et filtres
- [ ] [Mathias] Créer la structure des paramètres de recherche
- [ ] [Mathias] Développer NormalizeSearch
- [ ] [Mathias] Développer SearchRaces
- [ ] [Mathias] Rechercher sur deux propriétés
- [ ] [Mathias] Créer la page de résultats
- [ ] [Mathias] Définir trois propriétés filtrables pertinentes
- [ ] [Mathias] Développer FilterRaces
- [ ] [Mathias] Générer les options de filtrage
- [ ] [Mathias] Implémenter les filtres cumulables
- [ ] [Mathias] Permettre la réinitialisation des filtres
- [ ] [Mathias] Gérer les résultats vides
- [ ] [Mathias] Tester les combinaisons recherche et filtres
- [ ] [Mathias] Étendre les traitements aux autres collections selon le périmètre validé

## 07 — Cache
- [ ] [Gautier] Définir la structure du cache
- [ ] [Gautier] Définir le TTL
- [ ] [Gautier] Ajouter la lecture du cache
- [ ] [Gautier] Ajouter l'écriture du cache
- [ ] [Gautier] Gérer l'expiration des données
- [ ] [Gautier] Protéger les accès concurrents
- [ ] [Gautier] Intégrer le cache au service de collections
- [ ] [Gautier] Vérifier les contraintes de conservation Blizzard
- [ ] [Gautier] Tester le cache

## 08 — Pagination
- [ ] [Mathias] Déclarer la structure Pagination
- [ ] [Mathias] Calculer le nombre de pages
- [ ] [Mathias] Calculer les indices de début et de fin
- [ ] [Mathias] Découper les slices
- [ ] [Mathias] Ajouter les boutons précédent et suivant
- [ ] [Mathias] Ajouter le choix du nombre de ressources par page
- [ ] [Mathias] Conserver les filtres et la recherche
- [ ] [Mathias] Gérer les pages inexistantes
- [ ] [Mathias] Tester une collection volumineuse
- [ ] [Mathias] Tester les collections vides

## 09 — Favoris
- [ ] [Gautier] Déclarer Favorite
- [ ] [Gautier] Créer LoadFavorites
- [ ] [Gautier] Créer SaveFavorites
- [ ] [Gautier] Créer AddFavorite
- [ ] [Gautier] Créer RemoveFavorite
- [ ] [Gautier] Créer IsFavorite
- [ ] [Gautier] Gérer les doublons
- [ ] [Gautier] Gérer les fichiers absents et corrompus
- [ ] [Gautier] Protéger les opérations de lecture et écriture
- [ ] [Gautier] Tester la persistance après redémarrage
- [ ] [Mathias] Créer les routes POST des favoris
- [ ] [Mathias] Créer les handlers d'ajout et de suppression
- [ ] [Mathias] Ajouter les boutons de favoris
- [ ] [Mathias] Créer la page de gestion des favoris
- [ ] [Mathias] Afficher les messages de confirmation
- [ ] [Commun] Valider les favoris depuis deux navigateurs

## 10 — Médias, sécurité et design
- [ ] [Gautier] Récupérer les médias des ressources retenues
- [ ] [Gautier] Contrôler les URL externes
- [ ] [Gautier] Vérifier l'absence de secrets dans Git
- [ ] [Gautier] Sécuriser les entrées et les identifiants
- [ ] [Gautier] Contrôler les accès concurrents
- [ ] [Gautier] Vérifier les erreurs et les logs
- [ ] [Commun] Sécuriser les formulaires POST
- [ ] [Mathias] Définir la palette graphique
- [ ] [Mathias] Créer les composants CSS
- [ ] [Mathias] Intégrer les médias
- [ ] [Mathias] Harmoniser les cartes des collections
- [ ] [Mathias] Finaliser le responsive
- [ ] [Mathias] Finaliser les pages d'erreur
- [ ] [Mathias] Vérifier la navigation et la lisibilité

## 11 — Tests et livraison
- [ ] [Gautier] Tester OAuth et son renouvellement automatique
- [ ] [Gautier] Tester les interruptions de l'API
- [ ] [Gautier] Tester les erreurs JSON et HTTP
- [ ] [Gautier] Tester le stockage concurrent
- [ ] [Mathias] Tester toutes les routes
- [ ] [Mathias] Tester la recherche
- [ ] [Mathias] Tester les trois filtres cumulables
- [ ] [Mathias] Tester la pagination combinée aux filtres
- [ ] [Mathias] Tester les favoris depuis l'interface
- [ ] [Commun] Réaliser une revue croisée du code
- [ ] [Commun] Documenter les routes
- [ ] [Commun] Documenter les endpoints
- [ ] [Commun] Documenter l'architecture
- [ ] [Commun] Documenter l'installation et les variables d'environnement
- [ ] [Commun] Répondre aux quatre questions de la FAQ
- [ ] [Commun] Finaliser les mentions et la confidentialité
- [ ] [Commun] Tester depuis un clone propre
- [ ] [Commun] Vérifier chaque critère de la grille
- [ ] [Commun] Préparer la présentation au formateur

## 12 — Extensions facultatives
- [ ] [Commun] Confirmer que les exigences obligaGautierres sont terminées
- [ ] [Gautier] Ajouter la collection des montures si pertinent
- [ ] [Mathias] Ajouter le tri alphabétique
- [ ] [Mathias] Ajouter des statistiques
- [ ] [Commun] Démontrer le bonus OAuth au formateur