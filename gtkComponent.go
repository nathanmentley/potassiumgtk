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

type gtkComponent struct {
    index int
    widget *gtk.Widget
}
func newGtkComponent(widget *gtk.Widget) gtkComponent {
    return gtkComponent{-1, widget}
}

func (g *gtkComponent) getIndex() int {
    return g.index
}
func (g *gtkComponent) setIndex(index int) {
    g.index = index
}
func (g *gtkComponent) componentWillUpdate(processor potassium.IComponentProcessor) {}
func (g *gtkComponent) componentDidMount(processor potassium.IComponentProcessor) {}
func (g *gtkComponent) componentDidUpdate(processor potassium.IComponentProcessor) {
    if g.widget != nil {
        if hide, ok := processor.GetProps()["hide"].(bool); ok {
            if hide {
                g.widget.Hide()
            } else {
                g.widget.Show()
            }
        } else {
            g.widget.Show()
        }
    }
}
func (g *gtkComponent) componentWillUnmount(processor potassium.IComponentProcessor) {}
