# API de Rapport Excel pour les Formulaires POS (PosForm)

## Description
Cette API permet de générer des rapports Excel détaillés des données des Formulaires POS avec toutes les fonctionnalités de filtrage avancé.

## Endpoint
```
GET /api/posforms/export/excel
```

## Fonctionnalités

### 📊 Contenu du Rapport
- **Résumé Exécutif** : Statistiques globales incluant le nombre total de formulaires, formulaires complets/incomplets, nombre de provinces, aires, sous-aires, prix total
- **Données Détaillées** : Liste complète des formulaires POS avec toutes leurs informations
- **Formatage Professionnel** : Styles Excel avec couleurs, bordures et formatage automatique
- **Auto-ajustement** : Colonnes automatiquement ajustées pour une lisibilité optimale

### 🔍 Filtres Supportés
L'API supporte tous les mêmes filtres que l'API de pagination plus des filtres exclusifs :

#### Recherche Générale
- `search` : Recherche dans les commentaires

#### Filtres Géographiques
- `country` : Filtrer par pays
- `province` : Filtrer par province
- `area` : Filtrer par aire
- `subarea` : Filtrer par sous-aire
- `commune` : Filtrer par commune

#### Filtres Commerciaux
- `price` : Filtrer par prix exact
- `status` : Filtrer par statut (complete/incomplete)
- `brandCount` : Filtrer par nombre de marques
  - `0` : Aucune marque
  - `5` : Exactement 5 marques
  - `5-10` : Entre 5 et 10 marques
  - `11+` : 11 marques ou plus
- `posType` : Filtrer par type de POS
- `posSearch` : Recherche dans nom/shop du POS

#### Filtres Hiérarchie Commerciale
- `asm` / `asmSearch` : Filtrer par ASM (exact/recherche)
- `supervisor` / `supervisorSearch` : Filtrer par superviseur (exact/recherche)
- `dr` / `drSearch` : Filtrer par DR (exact/recherche)
- `cyclo` / `cycloSearch` : Filtrer par cyclo (exact/recherche)

#### Filtres Temporels
- `quickDate` : Filtres rapides par date
  - `today` : Créés aujourd'hui
  - `yesterday` : Créés hier
  - `last7days` : Créés dans les 7 derniers jours
  - `last30days` : Créés dans les 30 derniers jours

#### Filtres de Plage de Dates (Exclusifs à l'API Excel)
- `startDate` : Date de début (format YYYY-MM-DD) - Filtre les formulaires créés à partir de cette date
- `endDate` : Date de fin (format YYYY-MM-DD) - Filtre les formulaires créés jusqu'à cette date

**⚠️ Note Importante :** Les filtres `startDate` et `endDate` sont **uniquement disponibles pour l'API de rapport Excel** (`/api/posforms/export/excel`). Ils ne sont pas supportés par les autres endpoints de pagination.

## Exemples d'Utilisation

### 1. Rapport Complet (Tous les Formulaires)
```
GET /api/posforms/export/excel
```

### 2. Rapport des Formulaires Complets d'une Province
```
GET /api/posforms/export/excel?province=Kinshasa&status=complete
```

### 3. Rapport des Formulaires par Prix
```
GET /api/posforms/export/excel?price=1000
```

### 4. Rapport des Formulaires Créés cette Semaine
```
GET /api/posforms/export/excel?quickDate=last7days
```

### 5. Rapport par Plage de Dates
```
GET /api/posforms/export/excel?startDate=2024-01-01&endDate=2024-12-31
```

### 6. Rapport par Nombre de Marques
```
GET /api/posforms/export/excel?brandCount=5-10
```

### 7. Rapport Combiné
```
GET /api/posforms/export/excel?province=Kinshasa&status=complete&startDate=2024-01-01&brandCount=5+
```

## Structure du Fichier Excel

### Feuille 1: "Rapport PosForm"

#### Section 1: En-tête du Rapport
- Titre du rapport
- Nom de l'entreprise
- Date de génération
- Auteur (système automatique)

#### Section 2: Résumé Exécutif
- Total des Formulaires
- Formulaires Complets
- Formulaires Incomplets
- Total Provinces
- Total Aires
- Total Sous-Aires
- Prix Total
- Date de génération

#### Section 3: Données Détaillées
Colonnes incluses :
1. **UUID** : Identifiant unique du formulaire
2. **Prix** : Prix du formulaire
3. **Commentaire** : Commentaire du formulaire
4. **Latitude** : Coordonnée latitude
5. **Longitude** : Coordonnée longitude
6. **Signature** : Signature du formulaire
7. **Pays** : Nom du pays
8. **Province** : Nom de la province
9. **Aire** : Nom de l'aire
10. **Sous-Aire** : Nom de la sous-aire
11. **Commune** : Nom de la commune
12. **Utilisateur** : Nom complet de l'utilisateur
13. **POS Nom** : Nom du point de vente
14. **POS Shop** : Nom du magasin
15. **ASM** : Area Sales Manager
16. **Superviseur** : Superviseur
17. **DR** : District Representative
18. **Cyclo** : Cyclo
19. **Nombre d'Articles** : Nombre d'articles dans le formulaire
20. **Statut** : Complet/Incomplet (avec couleur)
21. **Date Création** : Date de création
22. **Date Modification** : Date de dernière modification

## Limitations de Sécurité

- **Limite de 10 000 enregistrements** : Pour éviter les problèmes de mémoire
- Si plus de 10 000 enregistrements correspondent aux critères, l'API retournera une erreur demandant d'utiliser des filtres plus spécifiques

## Différences avec les APIs de Pagination

Cette API de rapport Excel supporte **tous les filtres des APIs de pagination standard** plus des **filtres de plage de dates exclusifs** :

### Filtres Communs (Disponibles partout)
- Tous les filtres géographiques, commerciaux, hiérarchie commerciale, et filtres rapides par date

### Filtres Exclusifs à l'API Excel
- `startDate` et `endDate` : Filtres de plage de dates personnalisée
- Ces filtres permettent une sélection précise des données par période pour les rapports Excel

## Réponses API

### Succès (200)
- Type de contenu : `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- En-tête de disposition : `attachment; filename=rapport_posform_YYYY-MM-DD_HH-MM-SS.xlsx`
- Corps : Fichier Excel binaire

### Erreur (400) - Trop de Données
```json
{
  "status": "error",
  "message": "Trop de données pour le rapport Excel. Maximum 10000 enregistrements autorisés, XXXX trouvés. Veuillez utiliser des filtres plus spécifiques.",
  "data": null
}
```

### Erreur (500) - Erreur Serveur
```json
{
  "status": "error",
  "message": "Description de l'erreur",
  "error": "Détails techniques de l'erreur"
}
```

## Format du Nom de Fichier
Les fichiers générés suivent le format :
```
rapport_posform_YYYY-MM-DD_HH-MM-SS.xlsx
```

Exemple : `rapport_posform_2024-12-26_14-30-25.xlsx`

## Statistiques Spécialisées

### Métriques Calculées
- **Formulaires Complets** : Formulaires avec un POS assigné
- **Formulaires Incomplets** : Formulaires sans POS assigné
- **Prix Total** : Somme de tous les prix des formulaires
- **Nombre d'Articles** : Comptage des PosFormItems associés

### Formatage Spécialisé
- **Colonnes Prix** : Formatage numérique avec séparateurs de milliers
- **Colonnes Statut** : Couleur verte pour "Complet", orange pour "Incomplet"
- **Colonnes Dates** : Formatage date standard
- **Colonnes Nombres** : Formatage numérique pour les compteurs

## Notes Techniques

- Utilise la bibliothèque `excelize/v2` pour la génération Excel
- Formatage professionnel avec styles prédéfinis
- Gestion automatique des colonnes larges (au-delà de la colonne Z)
- Tri par date de modification décroissante
- Gestion sécurisée des données manquantes
- Preloading optimisé des relations de base de données
- Calculs statistiques en temps réel

## Sécurité et Performance

- Validation des entrées utilisateur
- Limite stricte sur le nombre d'enregistrements
- Gestion d'erreurs complète
- Optimisation des requêtes de base de données
- Formatage sécurisé des données dans Excel

---

*Généré automatiquement par le système MSPOS API v3*
