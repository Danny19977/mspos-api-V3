# 🚀 API Filtres Avancés PosForm - Guide Complet

## 🎯 Nouveaux Endpoints Disponibles

### 1. **Pagination avec Filtres Avancés**
```
GET /api/posforms/advanced?[paramètres]
```

### 2. **Valeurs Uniques pour Filtres**
```
GET /api/posforms/filter-values
```

### 3. **Hiérarchie Commerciale Filtrée**
```
GET /api/posforms/filtered-hierarchy?[recherche]
```

### 4. **Statistiques des Filtres**
```
GET /api/posforms/filter-statistics?[filtres]
```

### 5. **Test des Filtres**
```
GET /api/posforms/test-filters?[tous_filtres]
```

## 🔍 Filtres Disponibles

### **Filtres Géographiques**
- `country` - Nom exact du pays
- `province` - Nom exact de la province  
- `area` - Nom exact de l'area
- `subarea` - Nom exact de la subarea
- `commune` - Nom exact de la commune

### **Filtres Commerciaux**
- `search` - Recherche dans les commentaires (ILIKE)
- `price` - Prix exact de la visite
- `status` - `complete` | `incomplete`
- `posType` - Type de point de vente
- `posSearch` - Recherche dans nom/shop du POS (ILIKE)

### **Hiérarchie Commerciale avec Recherche Intégrée**
- `asm` - ASM exact
- `asmSearch` - Recherche dans ASM (ILIKE)
- `supervisor` - Supervisor exact  
- `supervisorSearch` - Recherche dans Supervisor (ILIKE)
- `dr` - DR exact
- `drSearch` - Recherche dans DR (ILIKE)
- `cyclo` - Cyclo exact
- `cycloSearch` - Recherche dans Cyclo (ILIKE)

### **Filtres Spéciaux**
- `brandCount` - `0` | `1` | `2-5` | `6+`
- `quickDate` - `today` | `yesterday` | `last7days` | `last30days`

### **Pagination Standard**
- `page` - Numéro de page (défaut: 1)
- `limit` - Taille de page (défaut: 15)
- `start_date` - Date de début (ISO format)
- `end_date` - Date de fin (ISO format)

## 🧪 Exemples d'Utilisation

### **Exemple 1: Recherche Basique**
```
GET /api/posforms/advanced?search=bon accueil&page=1&limit=20
```

### **Exemple 2: Filtres Géographiques**
```
GET /api/posforms/advanced?country=RDC&province=Kinshasa&area=Gombe
```

### **Exemple 3: Hiérarchie Commerciale avec Recherche**
```
GET /api/posforms/advanced?asmSearch=Jean&supervisorSearch=Pierre
```

### **Exemple 4: Filtres de Statut et Marques**
```
GET /api/posforms/advanced?status=complete&brandCount=2-5&price=1000
```

### **Exemple 5: Filtres Temporels**
```
GET /api/posforms/advanced?quickDate=last7days&posType=Supermarché
```

### **Exemple 6: Recherche Multi-Critères Complexe**
```
GET /api/posforms/advanced?search=formation&country=RDC&province=Kinshasa&asmSearch=Jean&status=complete&brandCount=6+&quickDate=last30days&page=1&limit=30
```

## 📊 Réponses des Endpoints

### **Structure de Réponse Paginée**
```json
{
  "status": "success",
  "message": "PosForm avec filtres avancés récupérés avec succès",
  "data": [
    {
      "uuid": "...",
      "price": 1000,
      "comment": "Bon accueil, formation nécessaire",
      "asm": "Jean Dupont",
      "sup": "Pierre Martin",
      "dr": "Marie Claire",
      "cyclo": "Paul Nyembo",
      "Country": { "name": "RDC" },
      "Province": { "name": "Kinshasa" },
      "Area": { "name": "Gombe" },
      "SubArea": { "name": "Centre" },
      "Commune": { "name": "Gombe" },
      "Pos": { "name": "Super Market", "type": "Supermarché" },
      "PosFormItems": [
        {
          "Brand": { "name": "Coca-Cola" },
          "number_farde": 5,
          "sold": 3
        }
      ]
    }
  ],
  "pagination": {
    "total_records": 1250,
    "total_pages": 42,
    "current_page": 1,
    "page_size": 30,
    "complete_reports": 980,
    "incomplete_reports": 270,
    "has_filters": true,
    "active_filters": 6
  }
}
```

### **Valeurs Uniques pour Filtres**
```json
{
  "status": "success",
  "data": {
    "countries": ["RDC", "Congo", "Cameroun"],
    "provinces": ["Kinshasa", "Lubumbashi", "Matadi"],
    "areas": ["Gombe", "Bandalungwa", "Kasa-Vubu"],
    "sub_areas": ["Centre", "Nord", "Sud"],
    "communes": ["Gombe", "Limete", "Ngaliema"],
    "prices": [500, 1000, 1500, 2000],
    "pos_types": ["Supermarché", "Boutique", "Dépôt"],
    "asms": ["Jean Dupont", "Marie Claire", "Paul Nyembo"],
    "supervisors": ["Pierre Martin", "Alice Mbala", "Joseph Kala"],
    "drs": ["Marie Claire", "Jean-Baptiste", "Sarah Lungu"],
    "cyclos": ["Paul Nyembo", "Michel Kaba", "David Tshombe"]
  }
}
```

### **Hiérarchie Filtrée**
```json
{
  "status": "success",
  "data": {
    "filtered_asms": ["Jean Dupont", "Jean-Baptiste"],
    "filtered_supervisors": ["Pierre Martin", "Pierre Kabila"],
    "filtered_drs": ["Marie Claire", "Marie Antoinette"],
    "filtered_cyclos": ["Paul Nyembo", "Paul Mbuyi"]
  }
}
```

## 🎯 Logique des Filtres

### **Correspondance Exacte vs Recherche**
- **Filtres de sélection** (`country`, `asm`, `supervisor`) : Correspondance exacte
- **Filtres de recherche** (`search`, `asmSearch`, `posSearch`) : Recherche partielle (ILIKE)

### **Filtres Combinés**
- Tous les filtres sont combinés avec **AND**
- Les filtres hiérarchie + recherche sont complémentaires :
  - `asm=Jean Dupont` : Sélection exacte d'un ASM
  - `asmSearch=Jean` : Recherche tous les ASMs contenant "Jean"

### **Gestion des Dates**
- `start_date` et `end_date` : Filtres sur `created_at`
- `quickDate` : Raccourcis prédéfinis qui override les dates personnalisées

### **Statuts des Rapports**
- `complete` : `pos_uuid` n'est ni NULL ni vide
- `incomplete` : `pos_uuid` est NULL ou vide

### **Comptage des Marques**
- Utilise des sous-requêtes sur `pos_form_items`
- Supporte les plages (`2-5`, `6+`)

## 🔧 Intégration Frontend

### **TypeScript Interface**
```typescript
interface PosFormFilters {
  search?: string;
  country?: string;
  province?: string;
  area?: string;
  subarea?: string;
  commune?: string;
  price?: string;
  status?: 'complete' | 'incomplete' | '';
  brandCount?: '0' | '1' | '2-5' | '6+' | '';
  posType?: string;
  posSearch?: string;
  asm?: string;
  asmSearch?: string;
  supervisor?: string;
  supervisorSearch?: string;
  dr?: string;
  drSearch?: string;
  cyclo?: string;
  cycloSearch?: string;
  quickDate?: 'today' | 'yesterday' | 'last7days' | 'last30days' | '';
  page?: number;
  limit?: number;
  start_date?: string;
  end_date?: string;
}
```

### **Service Angular**
```typescript
getPosFormsWithFilters(filters: PosFormFilters): Observable<any> {
  const params = new HttpParams({ fromObject: filters as any });
  return this.http.get(`${this.apiUrl}/posforms/advanced`, { params });
}

getFilterValues(): Observable<any> {
  return this.http.get(`${this.apiUrl}/posforms/filter-values`);
}

getFilteredHierarchy(searchParams: any): Observable<any> {
  const params = new HttpParams({ fromObject: searchParams });
  return this.http.get(`${this.apiUrl}/posforms/filtered-hierarchy`, { params });
}
```

## 🚨 Notes Importantes

1. **Performance** : Les filtres utilisent des index sur les colonnes fréquemment filtrées
2. **Sécurité** : Tous les paramètres sont validés et échappés
3. **Pagination** : Limite maximale recommandée de 100 enregistrements par page
4. **Cache** : Les valeurs uniques peuvent être mises en cache côté frontend
5. **Recherche** : Insensible à la casse (ILIKE)
6. **Dates** : Format ISO 8601 recommandé

## 🎉 Fonctionnalités Impressionnantes

✅ **Recherche en temps réel** dans la hiérarchie commerciale  
✅ **Filtres géographiques** en cascade  
✅ **Statistiques dynamiques** des filtres  
✅ **Comptage intelligent** des marques  
✅ **Raccourcis temporels** intuitifs  
✅ **Combinaison flexible** de tous les critères  
✅ **Performance optimisée** avec JOINs efficaces  
✅ **Compatibilité totale** avec votre frontend existant  

Cette implémentation transforme votre système de filtrage en une **machine de recherche ultra-puissante** ! 🔥
