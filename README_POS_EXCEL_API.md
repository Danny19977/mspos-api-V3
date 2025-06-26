# API de Rapport Excel pour les Points de Vente (POS)

## Description
Cette API permet de générer des rapports Excel détaillés des données des Points de Vente (POS) avec toutes les fonctionnalités de filtrage avancé.

## Endpoint
```
GET /api/pos/export/excel
```

## Fonctionnalités

### 📊 Contenu du Rapport
- **Résumé Exécutif** : Statistiques globales incluant le nombre total de POS, POS actifs/inactifs, nombre de provinces, aires, sous-aires
- **Données Détaillées** : Liste complète des POS avec toutes leurs informations
- **Formatage Professionnel** : Styles Excel avec couleurs, bordures et formatage automatique
- **Auto-ajustement** : Colonnes automatiquement ajustées pour une lisibilité optimale

### 🔍 Filtres Supportés
L'API supporte tous les mêmes filtres que l'API de pagination :

#### Recherche Générale
- `search` : Recherche dans nom, shop, type POS, gérant, quartier, référence

#### Filtres Géographiques
- `country` : Filtrer par pays
- `province` : Filtrer par province
- `area` : Filtrer par aire
- `subarea` : Filtrer par sous-aire
- `commune` : Filtrer par commune

#### Filtres POS Spécifiques
- `posType` : Filtrer par type de POS
- `status` : Filtrer par statut (active/inactive)
- `gerant` : Filtrer par nom du gérant (recherche partielle)
- `quartier` : Filtrer par quartier (recherche partielle)

#### Filtres Utilisateur
- `userFullname` : Filtrer par nom complet d'utilisateur exact
- `userSearch` : Recherche partielle dans les noms d'utilisateurs

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
- `startDate` : Date de début (format YYYY-MM-DD) - Filtre les POS créés à partir de cette date
- `endDate` : Date de fin (format YYYY-MM-DD) - Filtre les POS créés jusqu'à cette date

**⚠️ Note Importante :** Les filtres `startDate` et `endDate` sont **uniquement disponibles pour l'API de rapport Excel** (`/api/pos/export/excel`). Ils ne sont pas supportés par les autres endpoints de pagination. Ces filtres peuvent être utilisés ensemble pour définir une plage de dates, ou individuellement.

## Exemples d'Utilisation

### 1. Rapport Complet (Tous les POS)
```
GET /api/pos/export/excel
```

### 2. Rapport des POS Actifs d'une Province
```
GET /api/pos/export/excel?province=Kinshasa&status=active
```

### 3. Rapport des POS par Gérant
```
GET /api/pos/export/excel?gerant=Jean
```

### 4. Rapport des POS Créés cette Semaine
```
GET /api/pos/export/excel?quickDate=last7days
```

### 5. Rapport Combiné (Province + Status + Date)
```
GET /api/pos/export/excel?province=Kinshasa&status=active&quickDate=last30days
```

### 6. Rapport par Recherche Générale
```
GET /api/pos/export/excel?search=Pharmacie
```

### 7. Rapport par Plage de Dates Personnalisée
```
GET /api/pos/export/excel?startDate=2024-01-01&endDate=2024-12-31
```

### 8. Rapport à partir d'une Date
```
GET /api/pos/export/excel?startDate=2024-06-01
```

### 9. Rapport jusqu'à une Date
```
GET /api/pos/export/excel?endDate=2024-06-30
```

### 10. Rapport Combiné (Province + Plage de Dates + Status)
```
GET /api/pos/export/excel?province=Kinshasa&startDate=2024-01-01&endDate=2024-06-30&status=active
```

## Structure du Fichier Excel

### Feuille 1: "Rapport POS"

#### Section 1: En-tête du Rapport
- Titre du rapport
- Nom de l'entreprise
- Date de génération
- Auteur (système automatique)

#### Section 2: Résumé Exécutif
- Total des POS
- POS Actifs
- POS Inactifs
- Total Provinces
- Total Aires
- Total Sous-Aires
- Date de génération

#### Section 3: Données Détaillées
Colonnes incluses :
1. **UUID** : Identifiant unique du POS
2. **Nom POS** : Nom du point de vente
3. **Shop** : Nom du magasin
4. **Type POS** : Type de point de vente
5. **Gérant** : Nom du gérant
6. **Avenue** : Adresse avenue
7. **Quartier** : Quartier
8. **Référence** : Référence de localisation
9. **Téléphone** : Numéro de téléphone
10. **Pays** : Nom du pays
11. **Province** : Nom de la province
12. **Aire** : Nom de l'aire
13. **Sous-Aire** : Nom de la sous-aire
14. **Commune** : Nom de la commune
15. **Utilisateur** : Nom complet de l'utilisateur assigné
16. **ASM** : Area Sales Manager
17. **Superviseur** : Superviseur
18. **DR** : District Representative
19. **Cyclo** : Cyclo
20. **Statut** : Actif/Inactif (avec couleur)
21. **Date Création** : Date de création
22. **Date Modification** : Date de dernière modification

## Limitations de Sécurité

- **Limite de 10 000 enregistrements** : Pour éviter les problèmes de mémoire
- Si plus de 10 000 enregistrements correspondent aux critères, l'API retournera une erreur demandant d'utiliser des filtres plus spécifiques

## Différences avec les APIs de Pagination

Cette API de rapport Excel supporte **tous les filtres des APIs de pagination standard** plus des **filtres de plage de dates exclusifs** :

### Filtres Communs (Disponibles partout)
- Tous les filtres géographiques, POS spécifiques, hiérarchie commerciale, et filtres rapides par date

### Filtres Exclusifs à l'API Excel
- `startDate` et `endDate` : Filtres de plage de dates personnalisée
- Ces filtres permettent une sélection précise des données par période pour les rapports Excel

## Réponses API

### Succès (200)
- Type de contenu : `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- En-tête de disposition : `attachment; filename=rapport_pos_YYYY-MM-DD_HH-MM-SS.xlsx`
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
rapport_pos_YYYY-MM-DD_HH-MM-SS.xlsx
```

Exemple : `rapport_pos_2024-12-26_14-30-25.xlsx`

## Notes Techniques

- Utilise la bibliothèque `excelize/v2` pour la génération Excel
- Formatage professionnel avec styles prédéfinis
- Gestion automatique des colonnes larges (au-delà de la colonne Z)
- Tri par date de modification décroissante
- Gestion sécurisée des données manquantes
- Preloading optimisé des relations de base de données

## Sécurité et Performance

- Validation des entrées utilisateur
- Limite stricte sur le nombre d'enregistrements
- Gestion d'erreurs complète
- Optimisation des requêtes de base de données
- Formatage sécurisé des données dans Excel

---

*Généré automatiquement par le système MSPOS API v3*
