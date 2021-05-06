# Groupie Tracker

## Présentation 

*** 

Ce projet consiste a réaliser un site web qui répertorie une cinquantaine d'artistes musicaux et leurs concerts ainsi que des informations sur eux tels que leur premier album ou encore les membres des groupes. Nous avions le détail de ces informations sur cette [API](http://groupietrackers.herokuapp.com/api) au format JSON. Sur ce site se trouve plusieurs fonctionnalités, que nous détaillerons plus tard. Ce projet est réalisé principalement en GO, JS et HTML/CSS. Nous devions tout d'abord avoir plusieurs pages. Nous les avons découpé ainsi : `Accueil` , `Artistes`, `Maps` et `Concerts`.

Dans ce projet nous sommes trois, Olivia MOREAU, Antoine BOUFFARD et Louis TEILLIAIS. 

## Fonctionnement
***

Pour lancer notre site web, il suffit de se rendre dans notre dossier `groupie-tracker` sur un éditeur de code comme **Visual Studio Code** par exemple, puis de lancer dans la console la commande **go run server.go** puis de se rendre sur un navigateur et l'aller à l'url **localhost:8000**. On tombe sur la page d'Accueil du site, on peut ansi changer de page en cliquant sur les boutons de notre header.

Sur la page `Artistes`, on retrouve tout les artistes affichés sous forme de carte avec leurs informations, artistes triés par ordre alphabétique. On y trouve un système de filtres, par **Date de création**, **Date du premier album** et **Nombre de membres**.

Sur la page `Maps`, on retrouve tout simplement une carte avec la localisation de l'utilisateur dans un premier temps. On peut ensuite sélecter un artiste et le marqueur se déplacera sur une des villes où l'artiste se rpoduit en concert.

Sur la page `Concerts`, on retrouve une barre de recherche où l'on peut chercher un groupe en particulier pour ensuite afficher les lieux et dates de concerts.

## Organisation du code
***

Nous avons organisé notre code de façon a mettre tout les fichiers de mêmes langages dans les mêmes dossier :

> Un dossier **template** où nous avons tous nos fichier HTML, entre autres `artistes.html, events.html, home.html, locations.html`, pour chacune de nos pages. Dans ce dossier, nous avons aussi un dossier "fixe" qui se nomme **layout** où se trouve notre fichier `header.html` qui se trouve sur toutes nos pages.

> Un dossier **static** où se trouve plusieurs dossiers **css**, **images** et **scripts**. Le fichier **css** contient tous nos fichiers pour le style des pages, ex : `artistes.css, events.css, header.css, etc...`. Le dossier **images**, comme son nom l'indique contient juste l'image que l'on a besoin. Le dossier **scripts**  contient lui les fichier JS que l'on utilise pour nos pages : `events.js, artistes.js, maps.js.`

> Pour finir, un dossier **src**, qui lui contient les fichiers qui vont nous permettre de gérer les fonctionnalités de nos différentes pages, en GO : `events.go, artistes.go, home.go, etc...` ce dossier contient aussi un dossier **data** qui contient les fichiers permettant la récupération des données de l'API : `city.go, groups.go, relations.go, etc...`.

Nous avons aussi a la racine notre serveur en GO qui nous permet de lancer notre site, il se nomme **server.go**.

## Organisation du travail 
***
Pour l'organisation du travail de groupe, nous avons dès le début créé un [GitHub](https://github.com/LouisTeilliais/groupie_tracker) pour pouvoir centraliser nos codes. Nous avons une branche `main`, et une `develop` pour pouvoir ajouter les fonctionnalités d'abord sur la branche `develop` puis pour le résultat final tout ajouter à la branche `main`. À partir de la branche `develop`, nous avons créé nos branches personnelles qui nous permettaient d'avoir une branche pour chaque fonctionnalités que l'on travaillait.

Pour se situer dans l'avancement du projet, nous avons créé ce [Trello](https://trello.com/b/mybA3dMd/groupie-tracker) pour pouvoir organiser nos tâches et définir ce que l'on avait à faire. Nous avons décidé des tâches qui nous intéressaient le plus pour pouvoir travailler chacun sur notre partie et être le plus efficace possible dans notre production. Lorsque nous rencontrions des difficultés nous nous sommes penchés à plusieurs sur le problème et/ou délégué une fonctionnalité à une autre personne.

Antoine a travaillé en partie sur la partie front avec le HTML et le CSS, ainsi que sur la gestion d'erreur en GO. Louis a travaillé sur la partie de la carte principalement ainsi que sur la création du serveur ou la récupération de certaines données. Olivia a travaillée sur la recupération des données, les filtres et la pagination.

## Idées pour améliorer

***

Ce projet propose des améliorations : 
- La page artiste peut contenir un tri spécifique par ordre croissant/décroissant sur les dates de création, le premier album ou encore le nombre de membres.
- La page Maps pourrait faire apparaître plusieurs marqueurs correspondant à toutes les villes où l'artiste se produit.
- La page d'accueil mérite d'être un peu plus fournie et travaillée pour exposer correctement notre travail.
- Une des erreurs n'a pas été traitée : l'erreur 400.

