# Vue 3 + Vite

This template should help get you started developing with Vue 3 in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

Learn more about IDE Support for Vue in the [Vue Docs Scaling up Guide](https://vuejs.org/guide/scaling-up/tooling.html#ide-support).


# Technological stack

Langage : Go  
Framework Web : Gin ou Echo  
Base de données : PostgreSQL  
ORM : GORM  
Hébergement : GCP Cloud Run / Kubernetes avec Docker  
Tâches en arrière-plan : Goroutines  
Tests Unitaires : Testify  
Monitoring : Prometheus + Grafana  

Conclusion : Pourquoi choisir PostgreSQL ?
Gestion des transactions complexes : PostgreSQL excelle dans la gestion des transactions complexes et des relations entre plusieurs entités de manière robuste et fiable.
Support avancé des types de données : Si vous avez des besoins avancés comme la gestion de données géospatiales, des documents JSON ou des types personnalisés, PostgreSQL est un excellent choix.
Performances pour des charges lourdes : Il est particulièrement adapté pour des applications avec des transactions simultanées et à forte concurrence.
Écosystème d’extensions et outils : PostgreSQL possède un écosystème riche d'extensions qui peuvent facilement étendre ses capacités.

---- 
Utilisation de secrets : Pour un environnement de production, il est préférable d'utiliser un gestionnaire de secrets pour éviter d’exposer les mots de passe dans les fichiers docker-compose.yml. Cela peut être fait avec des secrets Docker ou des outils comme Vault de HashiCorp.

Persisting data : Le volume postgres_data permet de persister les données entre les redémarrages. Assurez-vous que les données critiques soient stockées dans des volumes persistants pour éviter la perte de données.

Logs et monitoring : Vous pouvez également configurer des outils comme Prometheus ou Grafana pour surveiller vos services Docker. Pour des applications critiques, vous pourriez envisager l'intégration d'outils de monitoring de performance.

## Backend
### Structure
- `/cmd` : Points d'entrée de l'application
- `/internal` : Logique métier
- `/pkg` : Bibliothèques réutilisables
- `/api` : Configuration API et handlers
- `/config` : Fichiers de configuration
- `/migrations` : Scripts de migration DB
- `/test` : Tests unitaires

### Commands

Lancer le développement avec Docker Compose
Construire l'image et démarrer le conteneur : Ouvrez votre terminal dans le répertoire contenant le fichier docker-compose.yml et exécutez les commandes suivantes :

`docker-compose up --build`
Cela va :

Construire l'image à partir du Dockerfile.
Démarrer à la fois votre back-end Go et PostgreSQL dans des conteneurs séparés.
Note : Si vous ne voulez pas reconstruire l'image à chaque fois, vous pouvez utiliser `docker-compose up` sans le --build.

Accéder à l'application : Une fois le conteneur lancé, vous pouvez accéder à votre application Go en visitant http://localhost:8080.

docker-compose exec backend go run main.go
Stopper et redémarrer le conteneur :
Vous pouvez arrêter et redémarrer les services sans perdre vos modifications locales avec :

`docker-compose down`
`docker-compose up`

### Differences dockerfiles

##### Go mod / Go sum
`COPY go.mod go.sum ./` : Cette commande copie uniquement les fichiers go.mod et go.sum dans le conteneur. Cela permet à Docker de mettre en cache ces fichiers lors de la construction de l'image. En d'autres termes, seules les dépendances Go sont téléchargées dans l'image, et ce processus est mis en cache.
`RUN go mod download` : Cette commande télécharge toutes les dépendances Go définies dans le go.mod et go.sum sans toucher au code source du projet. Cela permet de ne télécharger les dépendances que lorsque les fichiers go.mod ou go.sum changent.

Avantages :
Cela améliore les performances en mettant en cache le téléchargement des dépendances Go. Cela signifie que si vous ne modifiez pas vos dépendances, Docker n'aura pas besoin de télécharger les dépendances à chaque modification du code source. Cela est particulièrement utile lors du développement local, où vous ne modifiez généralement que votre code sans toucher aux dépendances.

##### Go mod tidy
`COPY . .` : Cette commande copie tout le répertoire de travail (le code source et les fichiers de configuration) dans le conteneur. Cela signifie que vous copiez l'intégralité de votre application, y compris le code source.
`RUN go mod tidy && go build -o main .` : Cette commande fait deux choses :
`go mod tidy` : Cette commande nettoie le fichier go.mod et ajoute les dépendances manquantes et supprime celles inutilisées. C'est utile pour garder votre fichier go.mod propre.
`go build -o main .` : Cette commande compile le projet Go en un exécutable et le place dans un fichier nommé main.

Avantages :
Moins de cache : Chaque fois que vous modifiez un fichier source (comme un fichier Go), Docker doit copier tous les fichiers et exécuter go mod tidy et go build. C'est une approche plus simple mais moins optimisée en termes de performance si vous modifiez seulement le code sans changer les dépendances.
Simple pour les petits projets : Cela peut convenir pour des projets simples ou lorsque vous êtes en train de développer de manière itérative et que vous voulez que tout soit compilé à chaque fois.

Inconvénients de la version 2 :
Cette approche peut être moins efficace en termes de cache. Par exemple, Docker devra télécharger à chaque fois les dépendances, même si go.mod et go.sum n'ont pas changé.