/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassiumgtk

import (
    "log"
    
    "github.com/gotk3/gotk3/gtk"
    "github.com/nathanmentley/potassium"
)

//Component
type WindowComponent struct {
    window *gtk.Window

    potassium.Component
    gtkComponent
}
func NewWindowComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unable to create window:", err)
    }

    return &WindowComponent{win, potassium.NewComponent(parent), newGtkComponent()}
}
//iGtkComponent
func (w *WindowComponent) getGtkWidget() gtk.IWidget {
    return w.window
}
//component callback methods
func (w *WindowComponent) onClose() {
    gtk.MainQuit()
}
//IComponent
func (w *WindowComponent) ComponentDidMount(processor potassium.IComponentProcessor) {
    w.Component.ComponentDidMount(processor)
    
    if title, ok := processor.GetProps()["title"].(string); ok {
        w.window.SetTitle(title)
        w.window.Connect("destroy", w.onClose)
    }
}
func (w *WindowComponent) ComponentDidUpdate(processor potassium.IComponentProcessor) {
    // Recursively show all widgets contained in this window.
    w.window.ShowAll()
}
func (w *WindowComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if title, ok := processor.GetProps()["title"].(string); ok {
        w.window.SetTitle(title)
    }

    return &potassium.RenderResult{processor.GetChildren()}
}
