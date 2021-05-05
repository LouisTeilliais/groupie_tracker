package handlers

import (
	"text/template"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	data "./data"
)

type PageDataEvents struct {
	Groups []data.Group
	Relations []data.Relations
	ThisGroup data.OneGroup
}
type PageDataEvents2 struct {
	Groups []data.Group
	GroupsFound []data.Group
	ValueSearch string
	ThisGroup data.OneGroup
	Relations []data.Relations
}
func Events(w http.ResponseWriter, req *http.Request){
	groups := data.GetGroups()
	sort.SliceStable(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})
	t := template.Must(template.ParseFiles("./template/events.html", "./template/layout/header.html"))
	var tabGroupsChecked []data.Group
	var valueSearch string
	if req.Method == "POST" {
		fmt.Print("Requete OK", "\n")
		search := req.FormValue("search")

		for _, v := range groups {
			checkName, _ := regexp.MatchString("(?i)"+search, v.Name)
			if checkName {
				tabGroupsChecked = append(tabGroupsChecked, v)
				continue
			} else {
				for _, member := range v.Members{
					checkMember, _ := regexp.MatchString("(?i)"+search, member)
					if checkMember {
						tabGroupsChecked = append(tabGroupsChecked, v)
						break
					}
				}
			}
		}
		valueSearch = search
	} else {
		tabGroupsChecked = nil
		valueSearch = ""
	}

	var thisGroup data.OneGroup
	var groupRelations []data.Relations
	if req.FormValue("oneGroup") == "" {
		groupRelations = nil

	} else {
		thisGroup = data.GetOneGroup(req.FormValue("oneGroup"))
		groupRelations = data.GetEvents(req.FormValue("oneGroup"))
	}
		
	pageEvents := PageDataEvents2{
		Groups: groups,
		GroupsFound: tabGroupsChecked,
		ValueSearch: valueSearch,
		ThisGroup: thisGroup,
		Relations: groupRelations,
	}
	fmt.Print("Events - ✅\n")
	t.Execute(w, pageEvents)

}