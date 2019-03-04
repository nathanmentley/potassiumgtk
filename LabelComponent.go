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
type LabelComponentProps struct {
    text string
}
func NewLabelComponentProps(text string) LabelComponentProps {  
    return LabelComponentProps{text}
}

//State
type LabelComponentState struct {
}
func NewLabelComponentState() LabelComponentState {  
    return LabelComponentState{}
}

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

    return &LabelComponent{label, potassium.NewComponent(parent), newGtkComponent()}
}

//IComponent
func (l *LabelComponent) SetInitialState(props potassium.IProps) potassium.IState {
    return NewLabelComponentState()
}

func (l *LabelComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    props, ok := processor.GetProps().(LabelComponentProps)

    if ok {
        l.label.SetLabel(props.text)
    }

    return &potassium.RenderResult{nil}
}

//iGtkComponent
func (l *LabelComponent) getGtkWidget() gtk.IWidget {
    return l.label
}
