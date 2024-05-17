# Génération de PDF à partir d'Excel

## Structure du projet

```
microentreprise/
|-- backend/
| |-- calculator.go
| |-- main.go
|-- build/
| |-- build.go
|-- data/
| |-- examples/
| |-- templates/
| | |-- template.xlsx
|-- dist/
|-- frontend/
| |-- assets/
| | |-- images/
| | |-- scripts/
| | | |-- main.js
| | | |-- translations.js
| | |-- styles/
| | | |-- _base.scss
| | | |-- _components.scss
| | | |-- _layout.scss
| | | |-- _responsive.scss
| | | |-- _variables.scss
| | | |-- main.scss
| | |-- translations/
| | | |-- en.yaml
| | | |-- fr.yaml
| |-- .htaccess
| |-- 404.html
| |-- index.html
| |-- robots.txt
| |-- sitemap.xml
|-- .gitignore
|-- go.mod
|-- go.sum
|-- LICENSE.md
|-- README.md
```

## Démarrage

### Prérequis

- [Go](https://golang.org/doc/install) doit être installé sur votre machine.
- [Node.js](https://nodejs.org/) doit être installé pour gérer les dépendances frontend et les outils de build.

### Instructions

1. **Télécharger le code :**

   Clonez le repository depuis GitHub :

   ```bash
   git clone https://github.com/votre-utilisateur/microentreprise.git
   cd microentreprise
   ```

2. **Installer les dépendances :**

   Installez les dépendances nécessaires :

   * Pour le backend : ``go mod tidy``
   * Pour le frontend : ``npm install -g sass``

3. **Compiler les fichiers Sass :**

    Compilez les fichiers Sass en CSS :
    ``sass frontend/assets/styles/main.scss frontend/assets/styles/main.css``
    Ou écoutez les modifications pour une compilation automatique :
    ``sass --watch main.scss:main.css``

4. **Construire le projet :**

    Utilisez le script de build pour préparer les fichiers de distribution :
    ``go run build/build.go``

5. **Lancer le serveur :**

    Lancez le serveur Go :
    ``go run backend/main.go``

6. **Ouvrir dans le navigateur :**

    Ouvrez [localhost](http://localhost:8080) dans votre navigateur.

7. **Utilisation :**

* Téléchargez le template Excel depuis la section dédiée.
* Remplissez le template avec vos données.
* Déposez le fichier rempli sur le site web.
* Un PDF avec vos déclarations sera généré et téléchargé automatiquement.
* Consultez les indices utiles sur votre entreprise directement sur la page web.

## Notes

* Ce projet est un exemple et peut être modifié selon vos besoins.
* Personnalisez les fichiers avec vos données.
* Testez bien avant de publier.

## Contribution

Les contributions sont les bienvenues ! Veuillez ouvrir une issue ou soumettre une pull request pour toute suggestion ou amélioration.

## License

Ce projet est sous licence MIT.