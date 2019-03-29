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
type boxComponent struct {
    box *gtk.Box

    potassium.Component
    gtkComponent
}
func newBoxComponent(parent potassium.IComponentProcessor) potassium.IComponent {
    return &boxComponent{nil, potassium.NewComponent(parent), newGtkComponent(nil)}
}
//iGtkComponent
func (b *boxComponent) getGtkWidget() gtk.IWidget {
    return b.box
}


//IComponent
func (b *boxComponent) ComponentWillUpdate(processor potassium.IComponentProcessor) {
    b.Component.ComponentWillUpdate(processor)
    b.gtkComponent.componentWillUpdate(processor)
}
func (b *boxComponent) ComponentDidMount(processor potassium.IComponentProcessor) {
    b.Component.ComponentDidMount(processor)
    b.gtkComponent.componentDidMount(processor)

    if orientation, ok := processor.GetProps()["orientation"].(gtk.Orientation); ok {
        box, err := gtk.BoxNew(orientation, 1)
        b.gtkComponent.widget = &box.Widget

        if err != nil {
            log.Fatal("Unable to create box:", err)
        } else {
            b.box = box
        }
    }
}
func (b *boxComponent) ComponentDidUpdate(processor potassium.IComponentProcessor) {
    b.Component.ComponentDidUpdate(processor)
    b.gtkComponent.componentDidUpdate(processor)
}
func (b *boxComponent) ComponentWillUnmount(processor potassium.IComponentProcessor) {
    b.Component.ComponentWillUnmount(processor)
    b.gtkComponent.componentWillUnmount(processor)
}
func (b *boxComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    return &potassium.RenderResult{processor.GetChildren()}
}
