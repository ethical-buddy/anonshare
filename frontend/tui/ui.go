package tui

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

func StartTUI() {
    app := tview.NewApplication()

    textView := tview.NewTextView().
        SetText("AnonShare\n\nPress Q or ESC to quit.\n").
        SetTextAlign(tview.AlignCenter)

    textView.SetDoneFunc(func(key tcell.Key) {
        if key == tcell.KeyEscape {
            app.Stop()
        }
    })

    textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        case tcell.KeyEscape:
            app.Stop()
            return nil
        default:
            switch event.Rune() {
            case 'q', 'Q':
                app.Stop()
                return nil
            }
        }
        return event
    })

    app.SetRoot(textView, true)

    if err := app.Run(); err != nil {
        panic(err)
    }
}

