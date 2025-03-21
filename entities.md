# Conceptualisation des entités principales

Les entités principales de votre projet sont les suivantes :

1. **Match**
2. **Round**
3. **Action**
4. **Rapport**
5. **Objectif**
6. **Utilisateur**

Ces entités vont interagir les unes avec les autres, et voici comment elles peuvent être reliées.

---

## 1. Entité : Match

**Propriétés importantes** :
- `id` : Identifiant unique du match (UUID).
- `player1_id` : Identifiant du joueur 1.
- `player2_id` : Identifiant du joueur 2.
- `date` : Date et heure du match.
- `replay_url` : URL du replay associé (si disponible).
- `notes` : Notes personnalisées sur le match (texte libre).
- `rounds` : Liste des rounds associés à ce match.
- `reports` : Liste des rapports associés à ce match.
- `objectives` : Liste des objectifs liés au match.

**Relation** :
- Un match peut avoir plusieurs **rounds**.
- Un match peut avoir plusieurs **rapports** associés.
- Un match peut être lié à plusieurs **objectifs**.

---

## 2. Entité : Round

**Propriétés importantes** :
- `id` : Identifiant unique du round.
- `match_id` : Identifiant du match auquel ce round appartient.
- `number` : Numéro du round (1, 2, ou 3).
- `actions` : Liste des actions qui ont eu lieu pendant ce round.

**Relation** :
- Un **round** appartient à un **match**.
- Un **round** peut avoir plusieurs **actions**.

---

## 3. Entité : Action

**Propriétés importantes** :
- `id` : Identifiant unique de l'action.
- `type` : Type de l'action (par exemple : "Jump In", "Dash In", "Poke", etc.).
- `category` : Catégorie de l'action (par exemple : "Neutral", "Offense", "Defense", "System").
- `hitContext` : Contexte de l'impact de l'action ("On Hit", "On Block", "Whiff", etc.).
- `player` : Joueur ayant effectué l'action (par exemple : "P1", "P2").
- `timestamp` : Timestamp de l'action (date/heure à laquelle l'action s'est produite).
- `round_id` : Identifiant du round auquel l'action appartient.

**Relation** :
- Une **action** appartient à un **round**.
- Une **action** est liée à un **match** indirectement par le round.

---

## 4. Entité : Rapport

**Propriétés importantes** :
- `id` : Identifiant unique du rapport.
- `match_id` : Identifiant du match auquel ce rapport appartient.
- `type` : Type de rapport (par exemple : "Analyse", "Feedback", etc.).
- `content` : Contenu du rapport (texte).
- `created_at` : Date et heure de création du rapport.
- `author_id` : Identifiant de l'utilisateur qui a rédigé le rapport (par exemple : un coach ou un analyste).

**Relation** :
- Un **rapport** est lié à un **match**.
- Un **rapport** peut être rédigé par un **utilisateur**.

---

## 5. Entité : Objectif

**Propriétés importantes** :
- `id` : Identifiant unique de l'objectif.
- `match_id` : Identifiant du match auquel cet objectif est lié.
- `description` : Description de l'objectif.
- `status` : Statut de l'objectif (par exemple : "En cours", "Atteint", "Non atteint").
- `created_at` : Date de création de l'objectif.

**Relation** :
- Un **objectif** est lié à un **match**.
- Un **objectif** peut être suivi tout au long du match.

---

## 6. Entité : Utilisateur (Coach / Joueur)

**Propriétés importantes** :
- `id` : Identifiant unique de l'utilisateur.
- `username` : Nom d'utilisateur.
- `email` : Adresse email de l'utilisateur.
- `role` : Rôle de l'utilisateur (par exemple : "Coach", "Joueur").
- `password_hash` : Mot de passe hashé pour la sécurité.
- `created_at` : Date de création du compte.

**Relation** :
- Un **utilisateur** peut être lié à un ou plusieurs **rapports** (si l'utilisateur est un coach ou analyste).
- Un **utilisateur** peut être lié à plusieurs **matchs** si nécessaire (en tant que joueur ou analyste).

---

## Modèle de données relationnel (Schéma SQL)

### Tables :

1. **Matches**
   - `id` (UUID)
   - `player1_id` (UUID, référence à `Utilisateur`)
   - `player2_id` (UUID, référence à `Utilisateur`)
   - `date` (Timestamp)
   - `replay_url` (String, optionnel)
   - `notes` (Text, optionnel)

2. **Rounds**
   - `id` (UUID)
   - `match_id` (UUID, référence à `Matches`)
   - `number` (Integer)

3. **Actions**
   - `id` (UUID)
   - `type` (String)
   - `category` (String)
   - `hitContext` (String)
   - `player` (String, 'P1' ou 'P2')
   - `timestamp` (Timestamp)
   - `round_id` (UUID, référence à `Rounds`)

4. **Reports**
   - `id` (UUID)
   - `match_id` (UUID, référence à `Matches`)
   - `type` (String)
   - `content` (Text)
   - `created_at` (Timestamp)
   - `author_id` (UUID, référence à `Utilisateur`)

5. **Objectives**
   - `id` (UUID)
   - `match_id` (UUID, référence à `Matches`)
   - `description` (Text)
   - `status` (String)
   - `created_at` (Timestamp)

6. **Utilisateurs**
   - `id` (UUID)
   - `username` (String)
   - `email` (String)
   - `role` (String, par exemple : 'coach' ou 'joueur')
   - `password_hash` (String)
   - `created_at` (Timestamp)

### Relations clés (Foreign Keys) :
- Un **match** a plusieurs **rounds**.
- Un **round** contient plusieurs **actions**.
- Un **match** peut avoir plusieurs **rapports** associés.
- Un **match** peut être lié à plusieurs **objectifs**.
- Un **rapport** est écrit par un **utilisateur**.

---

## Recommandation sur l'implémentation

1. **Base de données** : Utilisez **PostgreSQL** pour gérer ces entités relationnelles, avec des types comme **UUID** pour les identifiants uniques, **Text** pour les descriptions et notes, et **Timestamp** pour gérer les dates.
2. **Backend** : Vous pouvez créer des API RESTful (ou GraphQL si vous préférez) pour exposer ces données. **Echo** (avec Go) ou **Gin** pourrait être un bon choix pour la gestion des routes et des API.
3. **Tests unitaires** : Utilisez des outils comme **Go Test** ou **Testify** pour tester vos API et vous assurer que les interactions entre les entités sont correctes.
