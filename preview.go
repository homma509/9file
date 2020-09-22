package main

import (
	"io/ioutil"

	"github.com/rivo/tview"
)

// PreviewPanel プレビュー画面の構造体
type PreviewPanel struct {
	*tview.TextView
}

// NewPreviewPanel プレビューのインスタンスの生成
func NewPreviewPanel() *PreviewPanel {
	p := &PreviewPanel{
		TextView: tview.NewTextView(),
	}

	p.SetBorder(true).
		SetTitle("preview").
		SetTitleAlign(tview.AlignLeft)

	return p
}

// UpdateView プレビュー画面の更新
func (p *PreviewPanel) UpdateView(name string) {
	var content string
	b, err := ioutil.ReadFile(name)
	if err != nil {
		content = err.Error()
	} else {
		content = string(b)
	}

	p.Clear().SetText(content)
}
