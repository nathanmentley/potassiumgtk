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

//State
type appComponentState struct {
    clicks int
    textValue string
}
//Component construction
type appComponent struct {
    potassium.Component
}
func newAppComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    return &appComponent{potassium.NewComponent(parent)}
}
func (a *appComponent) SetInitialState(props map[string]interface{}) potassium.IState {
    return appComponentState{0, ""}
}
//component callback methods
func (a *appComponent) onAddClick(processor potassium.IComponentProcessor) {
    if state, ok := processor.GetState().(appComponentState); ok {
        state.clicks = state.clicks + 1
        processor.SetState(state)
    }
}
func (a *appComponent) onSubtractClick(processor potassium.IComponentProcessor) {
    if state, ok := processor.GetState().(appComponentState); ok {
        state.clicks = state.clicks - 1
        processor.SetState(state)
    }
}
func (a *appComponent) onTextChange(processor potassium.IComponentProcessor, text string) {
    if state, ok := processor.GetState().(appComponentState); ok {
        state.textValue = text
        processor.SetState(state)
    }
}
//component render
func (a *appComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    if state, ok := processor.GetState().(appComponentState); ok {
        return &potassium.RenderResult{
            []potassium.IComponentProcessor{
                a.CreateElement(
                    potassiumgtk.NewWindowComponent,
                    map[string]interface{}{
                        "title": "Window Title " + strconv.Itoa(state.clicks),
                    },
                    []potassium.IComponentProcessor{
                        a.CreateElement(
                            potassiumgtk.NewColComponent,
                            map[string]interface{}{},
                            []potassium.IComponentProcessor{
                                a.CreateElement(
                                    potassiumgtk.NewLabelComponent,
                                    map[string]interface{}{
                                        "text": state.textValue,
                                    },
                                    []potassium.IComponentProcessor{
                                    },
                                ),
                                a.CreateElement(
                                    potassiumgtk.NewEntryComponent,
                                    map[string]interface{}{
                                        "value": state.textValue,
                                        "onChange": func(text string) { a.onTextChange(processor, text) },
                                    },
                                    []potassium.IComponentProcessor{
                                    },
                                ),
                                a.CreateElement(
                                    newAComponent,
                                    map[string]interface{}{
                                        "clicks": state.clicks,
                                        "onAddClick": func() { a.onAddClick(processor) },
                                        "onSubstractClick": func() { a.onSubtractClick(processor) },
                                    },
                                    []potassium.IComponentProcessor{
                                    },
                                ),
                            },
                        ),
                    },
                ),
            },
        }
    }

    return &potassium.RenderResult{
        []potassium.IComponentProcessor{},
    }
}
