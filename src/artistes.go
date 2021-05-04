package handlers

import (
	"text/template"
	"fmt"
	"net/http"
	data "./data"
	"strconv"
	"regexp"
	"sort"
	)


type NumberMembers struct {
	Checked string
	Number int
}

type RangeDate struct {
	Start string
	End string
}

type PageDataGroups struct {
	Groups []data.Group
	CheckboxMember []NumberMembers
	RangeCreationDate RangeDate
	RangeFirstAlbum RangeDate
	NumberPages []int
}

var nbMembersMax = 8

func Artistes(w http.ResponseWriter, req *http.Request){
	var tabMembers []NumberMembers
	
	for i := 1; i<=nbMembersMax; i++ {
		var member = NumberMembers{"unchecked", i}
		tabMembers = append(tabMembers, member)
	}

	var creationDate = RangeDate{"1958", "2015"}
	var firstAlbum = RangeDate{"1963", "2018"}

	groups := data.GetGroups()
	sort.SliceStable(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})

	t := template.Must(template.ParseFiles("./template/artistes.html", "./template/layout/header.html"))
	var tabGroups []data.Group

	if req.Method == "POST" {

		fmt.Printf("Récupération form POST Artistes ✅\n")
		creationDate = RangeDate{req.FormValue("startCD"), req.FormValue("endCD")}
		firstAlbum = RangeDate{req.FormValue("startFA"), req.FormValue("endFA")}

		fmt.Printf("Années : %s\n",creationDate)
		fmt.Printf("Album : %s\n",firstAlbum)

		// Boucle sur les membres, on voit lesquels sont checked ou pas
		var counterUnchecked = 0
		for j := 1; j<=nbMembersMax; j++{
			n := strconv.Itoa(j)
			if req.FormValue("checkbox-"+n) == n {
				tabMembers[j-1] = NumberMembers{"checked", j}
			} else {
				counterUnchecked += 1
			}
		}
		fmt.Printf("Membres : %v\n",tabMembers)




		// TRI PAR MEMBRES 
		if counterUnchecked != 8 {
			// création du tableau qui contiendra les groupes checkés
			// Aller chercher le nombre de membres par groupe
			for k := range groups {
				lenGroupMembers := len(groups[k].Members)
				for l := range tabMembers{
					if lenGroupMembers == tabMembers[l].Number && tabMembers[l].Checked == "checked"{
						tabGroups = append(tabGroups, groups[k])
					}
				}
			}
		} else {
			tabGroups = groups
		}
		// TRI PAR MEMBRES FINI

		// TRI PAR CREATION DATE
		yearStart, e := strconv.Atoi(creationDate.Start)
		if e != nil{
			fmt.Print("Erreur ATOI")
		}
		yearEnd, e := strconv.Atoi(creationDate.End)
		if e != nil{
			fmt.Print("Erreur ATOI")
		}
		var tabGroups2 []data.Group

		for _,v := range tabGroups {

			if v.CreationDate >= yearStart && v.CreationDate <= yearEnd {
				// ajout des groupes qui correspondent à l'intervalle
				tabGroups2 = append(tabGroups2, v)
			}
		}
		tabGroups = tabGroups2
		//Creation DATE FINI

		// Tri par First Album 
		yearStartFA, e := strconv.Atoi(firstAlbum.Start)
		if e != nil{
			fmt.Print("Erreur ATOI")
		}
		yearEndFA, e := strconv.Atoi(firstAlbum.End)
		if e != nil{
			fmt.Print("Erreur ATOI")
		}
		var tabGroups3 []data.Group

		for _,v := range tabGroups {
			var re, _ = regexp.Compile(`....$`)

			FirstAlbumYear := re.FindStringSubmatch(v.FirstAlbum)[0]
			FirstAlbumYearInt, e := strconv.Atoi(FirstAlbumYear)
			if e != nil{
				fmt.Print("Erreur ATOI")
			}

			if FirstAlbumYearInt >= yearStartFA && FirstAlbumYearInt <= yearEndFA {
				// ajout des groupes qui correspondent à l'intervalle
				tabGroups3 = append(tabGroups3, v)
			}
		}
		tabGroups = tabGroups3
		// First album END
		
	} else{
		tabGroups = groups
	}
	_, tabNumberPages := Paging(tabGroups)

	
	var startElem int
	var endElem int

	if req.Method == "GET"{
		page,_ := strconv.Atoi(req.FormValue("page"))
		startElem = (page-1)*9 + 1
		endElem = startElem + 9
		if endElem>len(tabGroups){
			endElem = startElem + (len(tabGroups)-startElem)
		}
		fmt.Printf("Numéro de page : %v\n", page)
		if page == 0 {
			startElem = 1
			endElem = 10
		}
	}

	
	fmt.Printf("premier index : %v\n", startElem)
	fmt.Printf("dernier index : %v\n", endElem)
	tabGroups = tabGroups[startElem:endElem]


	pageGroups := PageDataGroups{
	Groups: tabGroups,
	CheckboxMember:tabMembers,
	RangeCreationDate:creationDate,
	RangeFirstAlbum: firstAlbum,
	NumberPages: tabNumberPages}
	

	t.Execute(w, pageGroups)
	fmt.Print("Artistes - ✅\n")
	// w.Header().Add("Content-Type", "application/json")
}

func Paging(tabGroups []data.Group) (int, []int){
	numberOfPages := len(tabGroups)/9
	if len(tabGroups)%9!=0{
		numberOfPages++
	}
	var tabNumbers []int
	for i := 1; i<=numberOfPages; i++{
		tabNumbers = append(tabNumbers, i)
	}

	return numberOfPages, tabNumbers
}