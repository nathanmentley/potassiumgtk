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
type LabelComponent struct {
    label *gtk.Label

    potassium.Component
    gtkComponent
}
func NewLabelComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    label, err := gtk.LabelNew("")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }

    return &LabelComponent{label, potassium.NewComponent(parent), newGtkComponent(&label.Widget)}
}
func (l *LabelComponent) ComponentWillUpdate(processor potassium.IComponentProcessor) {
    l.Component.ComponentWillUpdate(processor)
    l.gtkComponent.componentWillUpdate(processor)
}
func (l *LabelComponent) ComponentDidMount(processor potassium.IComponentProcessor) {
    l.Component.ComponentDidMount(processor)
    l.gtkComponent.componentDidMount(processor)
}
func (l *LabelComponent) ComponentDidUpdate(processor potassium.IComponentProcessor) {
    l.Component.ComponentDidUpdate(processor)
    l.gtkComponent.componentDidUpdate(processor)
}
func (l *LabelComponent) ComponentWillUnmount(processor potassium.IComponentProcessor) {
    l.Component.ComponentWillUnmount(processor)
    l.gtkComponent.componentWillUnmount(processor)
}
//iGtkComponent
func (l *LabelComponent) getGtkWidget() gtk.IWidget {
    return l.label
}
//IComponent
func (l *LabelComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if text, ok := processor.GetProps()["text"].(string); ok {
        if l.label != nil && l.label.GetLabel() != text {
            l.label.SetLabel(text)
        }
    }

    return &potassium.RenderResult{}
}
