package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/valakuzhyk/planetdomains/card"
)

func pickCard(cards []card.Card) int {
	cardsToSelectFrom := append([]card.Card{nil}, cards...)

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
		Label:     "Which card would you like to purchase:",
		Items:     cardsToSelectFrom,
		Templates: templates,
		Size:      7,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}
	return i - 1 // To correct for the insertion of "Go Back"
}

func pause() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
