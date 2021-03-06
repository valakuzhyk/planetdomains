package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/valakuzhyk/planetdomains/card"
)

// PickCard gives the user the ability to choose a card from a list, and
// optionally go back.
func PickCard(question string, cards []card.Card, isRequired bool) int {
	cardsToSelectFrom := cards
	offset := 0
	if !isRequired {
		cardsToSelectFrom = append([]card.Card{nil}, cards...)
		offset = 1
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U00002192 {{if .}} {{ .String  }} ({{ .GetCost}}) {{else}} Go Back {{end}}",
		Inactive: "  {{if .}} {{ .String  }} ({{ .GetCost}}) {{else}} Go Back {{end}}",
		Selected: "{{ .String }}",
		Details: `
--------- Card ----------
{{ "Name:" | faint }}	{{with .}}{{.GetName }}{{end}}
{{ "Cost:" | faint }}	{{with .}}{{.GetCost }}{{end}}
{{ "Play Effects:" | faint }}	{{ range .GetPlayEffects }}{{.String}}, {{else}} {{ end }}
{{ "Ally Effects:" | faint }}	{{ range .GetAllyEffects }}{{.String}}, {{else}} {{ end }}
{{ "Scrap Effects:" | faint }}	{{ range .GetScrapEffects }}{{.String}}, {{else}} {{ end }}`,
	}
	prompt := promptui.Select{
		Label:     question,
		Items:     cardsToSelectFrom,
		Templates: templates,
		Size:      7,
	}

	i := -1
	var err error
	if isRequired {
		for i < 0 || err != nil {
			i, _, err = prompt.Run()
		}
	} else {
		i, _, err = prompt.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}
	return i - offset // To correct for the insertion of "Go Back"
}

func pause() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
