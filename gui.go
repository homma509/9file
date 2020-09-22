package main

import (
	"os"

	"github.com/rivo/tview"
)

// GUI 画面全体を管理するGUI構造体
type GUI struct {
	App          *tview.Application
	Pages        *tview.Pages
	FilePanel    *FilePanel
	PreviewPanel *PreviewPanel
}

// NewGUI GUIインスタンスの生成
func NewGUI() *GUI {
	return &GUI{
		App:          tview.NewApplication(),
		Pages:        tview.NewPages(),
		FilePanel:    NewFilePanel(),
		PreviewPanel: NewPreviewPanel(),
	}
}

// Run TUIの実行
func (g *GUI) Run() error {
	cur, err := os.Getwd()
	if err != nil {
		return err
	}
	files, err := Files(cur)
	if err != nil {
		return err
	}

	g.FilePanel.SetFiles(files)
	g.FilePanel.UpdateView()

	file := g.FilePanel.SelectedFile()
	if file != nil {
		g.PreviewPanel.UpdateView(file.Name())
	}

	g.SetKeyBinding()

	grid := tview.NewGrid().SetColumns(0, 0).
		AddItem(g.FilePanel, 0, 0, 1, 1, 0, 0, true).
		AddItem(g.PreviewPanel, 0, 1, 1, 1, 0, 0, true)

	g.Pages.AddAndSwitchToPage("main", grid, true)

	return g.App.SetRoot(g.Pages, true).Run()
}

// SetKeyBinding 管理画面のキーバインド設定
func (g *GUI) SetKeyBinding() {
	g.FilePanel.Keybinding(g)
}
