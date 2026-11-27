# DataDex — Azeroth Archives : architecture de départ

Cette archive est un squelette pédagogique à intégrer au dépôt Git existant. Aucun code Go, HTML ou CSS n’a été implémenté.

## Installation du squelette

1. Extraire l’archive dans un dossier temporaire.
2. Copier le contenu du dossier DataDex dans le dépôt existant, sans remplacer main.go, go.mod, README.md ni le dossier .git.
3. Le dossier swapi, s’il est toujours vide, peut être supprimé.
4. Vérifier la version de Go déclarée dans le go.mod existant avant de commencer.
5. Les fichiers portant le suffixe  sont des fichiers de travail vides. Retirer le suffixe uniquement au moment de les implémenter. Les fichiers Go vides ne doivent pas avoir l’extension .go, car ils empêchent la compilation.
6. Le fichier de favoris sera créé par l’application au premier usage ; aucun secret ni donnée Blizzard n’est fourni.

## Responsabilités des dossiers

### config/
Chargement et validation des variables d’environnement. Ne pas journaliser les secrets.

- config.go

### battlenet/
Client HTTP centralisé, gestion OAuth 2.0 Client Credentials, récupération des index, des détails et des médias.

- client.go
- auth.go
- races.go
- classes.go
- pets.go
- media.go
- client_test.go
- auth_test.go

### models/
Structures Go représentant les réponses JSON, les références et les données de présentation.

- auth.go
- resource.go
- race.go
- class.go
- pet.go
- media.go
- favorite.go
- pagination.go
- page.go

### cache/
Cache mémoire avec durée de validité, protection des accès concurrents et limitation des appels redondants.

- cache.go
- cache_test.go

### services/
Orchestration des collections, recherche, filtres cumulables, pagination manuelle et favoris.

- catalog.go
- search.go
- filters.go
- pagination.go
- favorites.go
- search_test.go
- filters_test.go
- pagination_test.go
- favorites_test.go

### storage/
Lecture et écriture du fichier JSON unique, gestion des erreurs et accès concurrents.

- favorites.go
- favorites_test.go

### routes/
Association des chemins et méthodes HTTP aux handlers.

- routes.go
- routes_test.go

### handlers/
Validation des entrées HTTP, appel des services, rendu HTML et gestion des erreurs.

- home.go
- collections.go
- details.go
- search.go
- favorites.go
- about.go
- errors.go
- render.go
- handlers_test.go

### templates/
Templates HTML dynamiques : header et footer communs, pages de collection, de détail, recherche, favoris, à-propos, confidentialité et erreurs.

### static/
Feuilles de style CSS et images autorisées. Aucun JavaScript.

### data/
Emplacement du fichier favorites.json, qui sera créé lors de la mission favoris. Ne pas versionner les données générées.

## Première mission : config/config.go

- Définir la structure Config : ClientID, ClientSecret, BaseURL, TokenURL, Namespace, Locale et Port (string), CacheTTL (time.Duration).
- Définir LoadConfig : aucun paramètre ; deux retours, Config et error.
- Lire les variables d’environnement BATTLENET_CLIENT_ID et BATTLENET_CLIENT_SECRET. Vérifier leur présence.
- Définir les valeurs par défaut non sensibles : API Europe, namespace static-eu, locale fr_FR, port 8080, cache 30 minutes.
- Ne pas écrire ni afficher le secret. Ne pas appeler l’API à cette étape.
- Valider avec des identifiants factices, et tester le cas où une variable obligatoire est absente.

## Ordre des missions

Configuration → OAuth initial → client HTTP → renouvellement automatique → serveur et templates → races → recherche → filtres → pagination → favoris → classes et familiers → tests et design → documentation.

Le README existant sera complété progressivement. Les montures restent une extension facultative après validation du socle.
