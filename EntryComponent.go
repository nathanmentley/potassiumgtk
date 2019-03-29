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
type EntryComponent struct {
    entry *gtk.Entry

    potassium.Component
    gtkComponent
}
func NewEntryComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    entry, err := gtk.EntryNew()
    if err != nil {
        log.Fatal("Unable to create entry:", err)
    }

    return &EntryComponent{entry, potassium.NewComponent(parent), newGtkComponent(&entry.Widget)}
}
//component callback methods
func (e *EntryComponent) onChange(processor potassium.IComponentProcessor) {
    if onChange, ok := processor.GetProps()["onChange"].(func(string)); ok {
        if value, ok := processor.GetProps()["value"].(string); ok {
            text, err := e.entry.GetText()
            if err == nil && text != value {
                onChange(text)
            }
        }
    }
}
func (e *EntryComponent) ComponentWillUpdate(processor potassium.IComponentProcessor) {
    e.Component.ComponentWillUpdate(processor)
    e.gtkComponent.componentWillUpdate(processor)
}
func (e *EntryComponent) ComponentDidMount(processor potassium.IComponentProcessor) {
    e.Component.ComponentDidMount(processor)
    e.gtkComponent.componentDidMount(processor)
    
    e.entry.Connect("changed", func() { 
        e.onChange(processor)
    })
}
func (e *EntryComponent) ComponentDidUpdate(processor potassium.IComponentProcessor) {
    e.Component.ComponentDidUpdate(processor)
    e.gtkComponent.componentDidUpdate(processor)
}
func (e *EntryComponent) ComponentWillUnmount(processor potassium.IComponentProcessor) {
    e.Component.ComponentWillUnmount(processor)
    e.gtkComponent.componentWillUnmount(processor)
}
//iGtkComponent
func (e *EntryComponent) getGtkWidget() gtk.IWidget {
    return e.entry
}
func (e *EntryComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if value, ok := processor.GetProps()["value"].(string); ok {
        e.entry.SetText(value)
    }

    return &potassium.RenderResult{}
}
