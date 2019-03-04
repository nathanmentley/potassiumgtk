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

//AppComponent

//Props
type aComponentProps struct {
    clicks int
    onClick func()
}
func newAComponentProps(clicks int, onClick func()) aComponentProps { return aComponentProps{clicks, onClick} }

//State
type aComponentState struct {}
func newAComponentState() aComponentState { return aComponentState{}}

//Component construction
type aComponent struct {
    potassium.Component
}
func newAComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    return &aComponent{potassium.NewComponent(parent)}
}

func (a *aComponent) SetInitialState(props potassium.IProps) potassium.IState {
    return newAComponentState()
}

//component callback methods
func (a *aComponent) onClick(processor potassium.IComponentProcessor) {
    props, ok := processor.GetProps().(aComponentProps)
    if ok {
        props.onClick()
    }
}

//component render
func (a *aComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if props, ok := processor.GetProps().(aComponentProps); ok {
        children := []potassium.IComponentProcessor{
            a.CreateElement(
                potassium.NewComponentKey("button"),
                potassiumgtk.NewButtonComponent,
                potassiumgtk.NewButtonComponentProps("Button", func() { a.onClick(processor) }),
                []potassium.IComponentProcessor{
                },
            ),
        }

        if props.clicks < 10 {
            children = append(
                children, 
                a.CreateElement(
                    potassium.NewComponentKey("label2_" + strconv.Itoa(props.clicks)),
                    potassiumgtk.NewLabelComponent,
                    potassiumgtk.NewLabelComponentProps("Total button clicks (only less than ten): " + strconv.Itoa(props.clicks)),
                    []potassium.IComponentProcessor{
                    },
                ),
            )
        }

        return &potassium.RenderResult{
            []potassium.IComponentProcessor{
                a.CreateElement(
                    potassium.NewComponentKey("col"),
                    potassiumgtk.NewColComponent,
                    potassiumgtk.NewColComponentProps(),
                    children,
                ),
                a.CreateElement(
                    potassium.NewComponentKey("label_" + strconv.Itoa(props.clicks)),
                    potassiumgtk.NewLabelComponent,
                    potassiumgtk.NewLabelComponentProps("Total button clicks: " + strconv.Itoa(props.clicks)),
                    []potassium.IComponentProcessor{
                    },
                ),
                a.CreateElement(
                    potassium.NewComponentKey("button2"),
                    potassiumgtk.NewButtonComponent,
                    potassiumgtk.NewButtonComponentProps("Button", func() { a.onClick(processor) }),
                    []potassium.IComponentProcessor{
                    },
                ),
            },
        }
    }

    return &potassium.RenderResult{
        []potassium.IComponentProcessor{},
    }
}
