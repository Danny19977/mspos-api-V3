# API Documentation - User Filter Functions

## Nouvelles fonctionnalités ajoutées

### 1. GetUserByParams - Récupère un utilisateur selon différents paramètres

**Endpoint:** `GET /api/users/filter`

Cette fonction permet de récupérer UN SEUL utilisateur en appliquant différents filtres via les paramètres de requête (query parameters). La fonction retourne le premier utilisateur qui correspond aux critères.

#### Paramètres de filtrage disponibles :

- `pos` : UUID d'un POS spécifique
- `posform` : UUID d'un POSFORM spécifique  
- `role` : Rôle de l'utilisateur (ex: "DR", "ASM", "SUP", etc.)
- `status` : Statut de l'utilisateur ("true" ou "false")
- `country` : UUID du pays
- `province` : UUID de la province
- `area` : UUID de la zone
- `subarea` : UUID de la sous-zone
- `commune` : UUID de la commune
- `manager` : UUID du manager
- `asm` : UUID de l'ASM
- `sup` : UUID du SUP
- `dr` : UUID du DR
- `cyclo` : UUID du CYCLO
- `search` : Recherche textuelle dans fullname, title, email

#### Exemples d'utilisation :

```bash
# Récupérer le premier utilisateur avec le rôle "DR"
GET /api/users/filter?role=DR

# Récupérer l'utilisateur associé à un POS spécifique
GET /api/users/filter?pos=123e4567-e89b-12d3-a456-426614174000

# Récupérer l'utilisateur d'une province spécifique
GET /api/users/filter?province=456e7890-e89b-12d3-a456-426614174001

# Combiner plusieurs filtres
GET /api/users/filter?role=DR&status=true&province=456e7890-e89b-12d3-a456-426614174001

# Filtrer par POSFORM
GET /api/users/filter?posform=789e0123-e89b-12d3-a456-426614174002&search=manager
```

#### Réponse (succès) :

```json
{
  "status": "success",
  "message": "User retrieved successfully",
  "data": {
    "uuid": "user-uuid",
    "fullname": "John Doe",
    "email": "john@example.com",
    "role": "DR",
    "status": true,
    "country": {...},
    "province": {...},
    "area": {...},
    "sub_area": {...},
    "commune": {...}
  },
  "filters_applied": {
    "pos": "",
    "posform": "",
    "role": "DR",
    "status": "true",
    "country": "",
    "province": "456e7890-e89b-12d3-a456-426614174001",
    "search": "",
    ...
  }
}
```

#### Réponse (aucun utilisateur trouvé) :

```json
{
  "status": "error",
  "message": "No user found matching the criteria",
  "data": null
}
```

### 2. GetUserWithRelations - Récupère un utilisateur avec ses relations POS/POSFORM

**Endpoint:** `GET /api/users/with-relations`

Cette fonction récupère UN SEUL utilisateur avec ses relations POS et POSFORM directement incluses dans la réponse.

#### Paramètres disponibles :

- `user` : UUID d'un utilisateur spécifique
- `role` : Rôle de l'utilisateur
- `status` : Statut de l'utilisateur ("true" ou "false")
- `include_pos` : Inclure les relations POS ("true" par défaut)
- `include_posform` : Inclure les relations POSFORM ("true" par défaut)
- `search` : Recherche textuelle

#### Exemples d'utilisation :

```bash
# Récupérer un utilisateur spécifique avec ses POS et POSFORM
GET /api/users/with-relations?user=123e4567-e89b-12d3-a456-426614174000

# Récupérer le premier DR avec seulement ses POS
GET /api/users/with-relations?role=DR&include_posform=false

# Récupérer le premier utilisateur actif avec toutes ses relations
GET /api/users/with-relations?status=true

# Recherche avec relations complètes
GET /api/users/with-relations?search=manager&include_pos=true&include_posform=true
```

#### Réponse (succès) :

```json
{
  "status": "success",
  "message": "User with relations retrieved successfully",
  "data": {
    "uuid": "user-uuid",
    "fullname": "John Doe",
    "email": "john@example.com",
    "role": "DR",
    "status": true,
    "country": {...},
    "province": {...},
    "pos": [
      {
        "uuid": "pos-uuid",
        "name": "Shop ABC",
        "shop": "ABC Store",
        "postype": "Retail",
        "country": {...},
        "province": {...}
      }
    ],
    "pos_forms": [
      {
        "uuid": "posform-uuid",
        "price": 1500,
        "comment": "Visit completed",
        "pos": {
          "uuid": "pos-uuid",
          "name": "Shop ABC"
        },
        "country": {...}
      }
    ]
  },
  "metadata": {
    "include_pos": true,
    "include_posform": true,
    "filters": {
      "user": "",
      "role": "",
      "status": "true",
      "search": "manager"
    }
  }
}
```

#### Réponse (aucun utilisateur trouvé) :

```json
{
  "status": "error",
  "message": "No user found matching the criteria",
  "data": null
}
```

## Cas d'usage recommandés

### 1. Récupération d'un utilisateur spécifique
- Utiliser `GetUserByParams` pour trouver un utilisateur selon des critères précis
- Exemple: Trouver l'utilisateur responsable d'un POS spécifique

### 2. Détails complets d'un utilisateur
- Utiliser `GetUserWithRelations` pour obtenir un utilisateur avec ses POS et POSFORM associés
- Exemple: Afficher le profil complet d'un DR avec toutes ses activités

### 3. Recherche par critères
- Utiliser `GetUserByParams` avec des filtres géographiques ou hiérarchiques
- Exemple: Trouver le manager d'une zone spécifique

### 4. API pour applications mobiles
- Ces endpoints retournent un seul utilisateur, parfait pour les applications mobiles
- Exemple: Login et récupération du profil utilisateur avec ses données associées

## Notes importantes

1. **Retour unique** : Les deux fonctions retournent UN SEUL utilisateur (le premier trouvé selon les critères)
2. **Gestion des erreurs** : Retourne une erreur 404 si aucun utilisateur ne correspond aux critères
3. **Performance** : Plus rapide que les versions avec pagination car ne retourne qu'un seul résultat
4. **Relations** : Les relations sont préchargées avec leurs propres relations (Country, Province, etc.)
5. **Ordre** : Les résultats sont triés par `updated_at DESC`, donc le plus récent en premier
