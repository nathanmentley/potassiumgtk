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
type ButtonComponent struct {
    button *gtk.Button

    potassium.Component
    gtkComponent
}
func NewButtonComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    button, err := gtk.ButtonNew()
    if err != nil {
        log.Fatal("Unable to create button:", err)
    }

    return &ButtonComponent{button, potassium.NewComponent(parent), newGtkComponent(&button.Widget)}
}
//iGtkComponent
func (b *ButtonComponent) getGtkWidget() gtk.IWidget {
    return b.button
}
//component callback methods
func (b *ButtonComponent) onClick(processor potassium.IComponentProcessor) {
    if onClick, ok := processor.GetProps()["onClick"].(func ()); ok && onClick != nil {
        onClick()
    }
}
//IComponent
func (b *ButtonComponent) ComponentWillUpdate(processor potassium.IComponentProcessor) {
    b.Component.ComponentWillUpdate(processor)
    b.gtkComponent.componentWillUpdate(processor)
}
func (b *ButtonComponent) ComponentDidMount(processor potassium.IComponentProcessor) {
    b.Component.ComponentDidMount(processor)
    b.gtkComponent.componentDidMount(processor)
    
    b.button.Connect("clicked", func() { 
        b.onClick(processor)
    })
}
func (b *ButtonComponent) ComponentDidUpdate(processor potassium.IComponentProcessor) {
    b.Component.ComponentDidUpdate(processor)
    b.gtkComponent.componentDidUpdate(processor)
}
func (b *ButtonComponent) ComponentWillUnmount(processor potassium.IComponentProcessor) {
    b.Component.ComponentWillUnmount(processor)
    b.gtkComponent.componentWillUnmount(processor)
}
func (b *ButtonComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if title, ok := processor.GetProps()["title"].(string); ok {
        b.button.SetLabel(title)
    }

    return &potassium.RenderResult{}
}
