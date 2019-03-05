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
type EntryComponentProps struct {
    value string
    onChange func(string)
}
func NewEntryComponentProps(value string, onChange func(string)) EntryComponentProps {  
    return EntryComponentProps{value, onChange}
}

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

    return &EntryComponent{entry, potassium.NewComponent(parent), newGtkComponent()}
}
//component callback methods
func (e *EntryComponent) onChange(processor potassium.IComponentProcessor) {
    if props, ok := processor.GetProps().(EntryComponentProps); ok {
        text, err := e.entry.GetText()
        if err == nil && text != props.value {
            props.onChange(text)
        }
    }
}
func (e *EntryComponent) ComponentDidMount(processor potassium.IComponentProcessor) {
    e.Component.ComponentDidMount(processor)
    
    e.entry.Connect("changed", func() { 
        e.onChange(processor)
    })
}
//iGtkComponent
func (e *EntryComponent) getGtkWidget() gtk.IWidget {
    return e.entry
}
func (e *EntryComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if props, ok := processor.GetProps().(EntryComponentProps); ok {
        e.entry.SetText(props.value)
    }

    return &potassium.RenderResult{}
}
