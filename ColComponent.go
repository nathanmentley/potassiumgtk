/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassiumgtk

import (
    "github.com/gotk3/gotk3/gtk"

    "github.com/nathanmentley/potassium"
)

//Component construction
type ColComponent struct {
    potassium.Component
}
func NewColComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    return &ColComponent{potassium.NewComponent(parent)}
}
//component render
func (r *ColComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    return &potassium.RenderResult{
        []potassium.IComponentProcessor{
            r.CreateElement(
                potassium.NewComponentKey("Box"),
                newBoxComponent,
                map[string]interface{}{
                    "orientation": gtk.ORIENTATION_VERTICAL,
                },
                processor.GetChildren(),
            ),
        },
    }
}
