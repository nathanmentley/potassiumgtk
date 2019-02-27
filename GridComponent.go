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
type GridComponentProps struct {}
func NewGridComponentProps() GridComponentProps {  
    return GridComponentProps{}
}

//State
type GridComponentState struct {
}
func NewGridComponentState() GridComponentState {  
    return GridComponentState{}
}

//Component
type GridComponent struct {
    grid *gtk.Grid

    potassium.Component
}
func NewGridComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    grid, err := gtk.GridNew()
    if err != nil {
        log.Fatal("Unable to create grid:", err)
    }

    return &GridComponent{grid, potassium.NewComponent(parent)}
}

//IComponent
func (g *GridComponent) SetInitialState(props potassium.IProps) potassium.IState {
    return NewGridComponentState()
}

func (g *GridComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    return &potassium.RenderResult{processor.GetChildren()}
}

//iGtkComponent
func (g *GridComponent) getGtkWidget() gtk.IWidget {
    return g.grid
}
