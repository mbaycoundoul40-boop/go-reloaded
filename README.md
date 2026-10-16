# go-reloaded

Outil en Go qui lit un fichier texte, le corrige et écrit le résultat dans un autre fichier.

## Utilisation

```bash
go run . sample.txt result.txt
```

## Modifications effectuées

| Règle | Effet | Exemple |
|---|---|---|
| `(hex)` | Hexadécimal → décimal | `1E (hex)` → `30` |
| `(bin)` | Binaire → décimal | `10 (bin)` → `2` |
| `(up)` | Majuscules | `go (up)` → `GO` |
| `(low)` | Minuscules | `SHOUTING (low)` → `shouting` |
| `(cap)` | Majuscule initiale | `bridge (cap)` → `Bridge` |
| `(up, n)` `(low, n)` `(cap, n)` | Sur les n mots précédents | `so exciting (up, 2)` → `SO EXCITING` |
| Ponctuation `. , ! ? : ;` | Collée au mot précédent | `there ,and` → `there, and` |
| Groupes `...` `!?` | Restent groupés | `thinking ...` → `thinking...` |
| Apostrophes `' '` | Collées aux mots | `' awesome '` → `'awesome'` |
| `a` → `an` | Devant voyelle ou h | `a amazing` → `an amazing` |

## Organisation du code

| Fichier | Rôle |
|---|---|
| `main.go` | Lit les arguments et les fichiers |
| `process.go` | Applique les étapes dans l'ordre |
| `markers.go` | Gère les marqueurs |
| `format.go` | Gère a/an, apostrophes et ponctuation |
| `process_test.go` | Tests unitaires |

## Tests

```bash
go test -v
```

## Packages utilisés

Uniquement la bibliothèque standard de Go : `os`, `fmt`, `strings`, `strconv`, `regexp`, `unicode`, `testing`.

## Auteur

Mbayang Coundoul – Ynov Campus