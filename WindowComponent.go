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

//Props
type WindowComponentProps struct {
    title string
}
func NewWindowComponentProps(title string) WindowComponentProps {  
    return WindowComponentProps{title}
}

//State
type WindowComponentState struct {
}
func NewWindowComponentState() WindowComponentState {  
    return WindowComponentState{}
}

//Component
type WindowComponent struct {
    window *gtk.Window

    potassium.Component
}
func NewWindowComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Unable to create window:", err)
    }

    return &WindowComponent{win, potassium.NewComponent(parent)}
}

//IComponent
func (w *WindowComponent) SetInitialState(props potassium.IProps) potassium.IState {
    return NewWindowComponentState()
}

func (w *WindowComponent) ComponentDidMount(processor potassium.IComponentProcessor) {
    w.Component.ComponentDidMount(processor)
    
    props, ok := processor.GetProps().(WindowComponentProps)

    if ok {
        w.window.SetTitle(props.title)
        w.window.Connect("destroy", w.onClose)
    }
}

func (w *WindowComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    props, ok := processor.GetProps().(WindowComponentProps)

    if ok {
        w.window.SetTitle(props.title)
    }

    return &potassium.RenderResult{processor.GetChildren()}
}
func (w *WindowComponent) ComponentDidUpdate(processor potassium.IComponentProcessor) {
    // Recursively show all widgets contained in this window.
    w.window.ShowAll()
}

//iGtkComponent
func (w *WindowComponent) getGtkWidget() gtk.IWidget {
    return w.window
}

//component callback methods
func (w *WindowComponent) onClose() {
    gtk.MainQuit()
}

