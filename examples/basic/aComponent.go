/*
Copyright 2019 Nathan Mentley

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
    "strconv"

    "github.com/nathanmentley/potassium"
    "github.com/nathanmentley/potassiumgtk"
)

//Component construction
type aComponent struct {
    potassium.Component
}
func newAComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    return &aComponent{potassium.NewComponent(parent)}
}
//component callback methods
func (a *aComponent) onAddClick(processor potassium.IComponentProcessor) {
    if onAddClick, ok := processor.GetProps()["onAddClick"].(func()); ok {
        onAddClick()
    }
}
func (a *aComponent) onSubtractClick(processor potassium.IComponentProcessor) {
    if onSubstractClick, ok := processor.GetProps()["onSubstractClick"].(func()); ok {
        onSubstractClick()
    }
}
//component render
func (a *aComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if clicks, ok := processor.GetProps()["clicks"].(int); ok {
        colChildren := []potassium.IComponentProcessor{
            a.CreateElement(
                potassiumgtk.NewButtonComponent,
                map[string]interface{}{
                    "key": "new_button",
                    "title": "Subtract Button",
                    "onClick": func() { a.onSubtractClick(processor) },
                },
                []potassium.IComponentProcessor{
                },
            ),
        }

        if clicks < 3 {
            colChildren = append(
                colChildren,
                a.CreateElement(
                    potassiumgtk.NewLabelComponent,
                    map[string]interface{}{
                        "key": "total_clicks_label",
                        "text": "Total button clicks (only less than three): " + strconv.Itoa(clicks),
                    },
                    []potassium.IComponentProcessor{
                    },
                ),
            )
        }

        return &potassium.RenderResult{
            []potassium.IComponentProcessor{
                a.CreateElement(
                    potassiumgtk.NewRowComponent,
                    map[string]interface{}{},
                    []potassium.IComponentProcessor{
                        a.CreateElement(
                            potassiumgtk.NewColComponent,
                            map[string]interface{}{},
                            colChildren,
                        ),
                        a.CreateElement(
                            potassiumgtk.NewLabelComponent,
                            map[string]interface{}{
                                "text": "Total button clicks: " + strconv.Itoa(clicks),
                            },
                            []potassium.IComponentProcessor{
                            },
                        ),
                        a.CreateElement(
                            potassiumgtk.NewButtonComponent,
                            map[string]interface{}{
                                "title": "Add Button",
                                "onClick": func() { a.onAddClick(processor) },
                            },
                            []potassium.IComponentProcessor{
                            },
                        ),
                    },
                ),
            },
        }
    }

    return &potassium.RenderResult{}
}
