package cmd

import (
	"fmt"
	"nto_cli/utils"
	"os"

	"github.com/rivo/tview"
)

func Input() (string, string) {
	fmt.Print("struct name, path to file (including struct): ")
	var structName, path string
	fmt.Scan(&structName, &path)
	return structName, path
}

func SelectionInput() ([]string, string) {
	path := os.Args[1]

	structNames := utils.GetStructList(path)

	result := []string{}

	app := tview.NewApplication()

	form := tview.NewForm()
	checkboxes := []*tview.Checkbox{}

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
