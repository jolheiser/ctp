package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	// init models, we can reset them at any time anyway
	models = []tea.Model{NewInitialModel(), NewSpinnerParent()}
	m := models[initialView]
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		panic(err)
	}
}

// I put all the globals here :shrug:
var (
	models []tea.Model
	// current will be used to track the current model being returned from the
	// list of models
	current  int
	EnterVal string
	Cloned   bool // Planning to use this to determine when to exit the spinner when the repo is cloned.
)

const (
	initialView = iota
	spinnerView
)

type errMsg error
