/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package potassiumgtk

import (
    "time"

    "github.com/gotk3/gotk3/gtk"
    "github.com/nathanmentley/potassium"
)

//GtkToolkit
type gtkToolkit struct {
    isRunning bool
}

func newGtkToolkit() *gtkToolkit {  
    return &gtkToolkit{false}
}

//IAppToolkit
func (g *gtkToolkit) Setup() {
    gtk.Init(nil)
    g.isRunning = true
}
func (g *gtkToolkit) Step() {
    if gtk.EventsPending() {
        g.isRunning = gtk.MainIterationDo(false)
    } else {
        //sleep for 50 milliseconds to give the cpu some breathing room. not reason to kill it in this loop.
        time.Sleep(50 * time.Millisecond)
    }
}
func (g *gtkToolkit) IsRunning() bool {
    return g.isRunning
}
func (g *gtkToolkit) Shutdown() {}


func (g *gtkToolkit) Mount(parent potassium.IComponentProcessor, child potassium.IComponentProcessor) {
    /*parentGtk, ok := parent.(iGtkComponent)

    if ok {
        childGtk, ok := child.(iGtkComponent)

        if ok {
            //TODO: support more container types.
            parentContainer, ok := parentGtk.getGtkWidget().(*gtk.Window)

            if ok {
                parentContainer.Add(childGtk.getGtkWidget())
            }
            
            parentContainer2, ok := parentGtk.getGtkWidget().(*gtk.Grid)

            if ok {
                parentContainer2.Add(childGtk.getGtkWidget())
            }
        }
    }*/

    parentGtk := g.getNearestGtkParent(parent)

    if parentGtk != nil {
        if childGtk, ok := child.GetComponent().(iGtkComponent); ok {
            //TODO: support more container types.
            parentContainer, ok := parentGtk.getGtkWidget().(*gtk.Window)

            if ok {
                parentContainer.Add(childGtk.getGtkWidget())
            }
            parentContainer2, ok := parentGtk.getGtkWidget().(*gtk.Grid)

            if ok {
                parentContainer2.Add(childGtk.getGtkWidget())
            }
        }
    }
}

func (g *gtkToolkit) Unmount(parent potassium.IComponentProcessor, child potassium.IComponentProcessor) {
    parentGtk := g.getNearestGtkParent(parent)

    if parentGtk != nil {
        if childGtk, ok := child.GetComponent().(iGtkComponent); ok {
            //TODO: support more container types.
            parentContainer, ok := parentGtk.getGtkWidget().(*gtk.Window)

            if ok {
                parentContainer.Remove(childGtk.getGtkWidget())
            }
            parentContainer2, ok := parentGtk.getGtkWidget().(*gtk.Grid)

            if ok {
                parentContainer2.Remove(childGtk.getGtkWidget())
            }
        }
    }
}

func (g *gtkToolkit) getNearestGtkParent(target potassium.IComponentProcessor) iGtkComponent {
    if target != nil && target.GetComponent() != nil {
        if gtk, ok := target.GetComponent().(iGtkComponent); ok {
            return gtk
        } else {
            parent := target.GetParent()
            
            if parent != nil {
                return g.getNearestGtkParent(parent)
            }
        }
    }

    return nil
}
