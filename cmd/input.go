package cmd

import (
	"log"
	"nto_cli/utils"
	"os"

	"github.com/rivo/tview"
)

func SelectionInput() ([]string, string) {
	if len(os.Args) == 1 {
		log.Fatalf("Please provide path to models.go")
	}

	path := os.Args[1]

	structNames := utils.GetStructList(path)

	var result []string

	app := tview.NewApplication()

	form := tview.NewForm()
	var checkboxes []*tview.Checkbox

	for _, name := range structNames {
		cb := tview.NewCheckbox().SetLabel(name)
		checkboxes = append(checkboxes, cb)
		form.AddFormItem(cb)
	}

	form.AddButton("Confirm", func() {
		for i, cb := range checkboxes {
			if cb.IsChecked() {
				result = append(result, structNames[i])
			}
		}
		app.Stop()
	})

	if err := app.SetRoot(form, true).Run(); err != nil {
		panic(err)
	}

	return result, path
}
