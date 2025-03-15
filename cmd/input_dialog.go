package cmd

import (
	"github.com/opbnq-q/nto-cli/model"
	"log"
	"os"

	"github.com/rivo/tview"
)

func SelectionInput(models []model.Model) *[]model.Model {
	unimplementedModels := model.GetNotImplementedModels(models)
	var result []model.Model

	if len(unimplementedModels) == 0 {
		log.Println("No unimplemented models -> nothing to do")
		os.Exit(0)
	}

	app := tview.NewApplication()

	form := tview.NewForm()
	var checkboxes []*tview.Checkbox

	for _, m := range unimplementedModels {
		cb := tview.NewCheckbox().SetLabel(m.Name).SetChecked(true)
		checkboxes = append(checkboxes, cb)
		form.AddFormItem(cb)
	}

	form.AddButton("Generate", func() {
		for i, cb := range checkboxes {
			if cb.IsChecked() {
				result = append(result, unimplementedModels[i])
			}
		}
		app.Stop()
	})

	if err := app.SetRoot(form, true).Run(); err != nil {
		log.Fatalf("Failed to initialize dialog: %s", err)
	}

	return &result
}
