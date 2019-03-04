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
type appComponentProps struct {}
func newAppComponentProps() appComponentProps { return appComponentProps{} }

//State
type appComponentState struct {
    clicks int
}
func newAppComponentState() appComponentState { return appComponentState{0} }

//Component construction
type appComponent struct {
    potassium.Component
}
func newAppComponent(parent potassium.IComponentProcessor) potassium.IComponent {  
    return &appComponent{potassium.NewComponent(parent)}
}

func (a *appComponent) SetInitialState(props potassium.IProps) potassium.IState {
    return newAppComponentState()
}

//component callback methods
func (a *appComponent) onClick(processor potassium.IComponentProcessor) {
    state, ok := processor.GetState().(appComponentState)
    if ok {
        state.clicks = state.clicks + 1
        processor.SetState(state)
    }
}

//component render
func (a *appComponent) Render(processor potassium.IComponentProcessor) *potassium.RenderResult {
    state, ok := processor.GetState().(appComponentState)

    if ok {
        /*
        TODO: Setup some jsx like preprocessor to convert this:

        return <potassiumgtk.Window key="window" title={"Window Title " + strconv.Itoa(state.clicks)}>
            <potassiumgtk.Box key="grid">
                <potassiumgtk.Label key={"label"} title={"Total button clicks: " + strconv.Itoa(state.clicks)} />
                <potassiumgtk.Button key="button" title="Button" onClick={func() { a.onClick(processor) }} />
            </potassiumgtk.Box>
        </potassiumgtk.Window>

        To this:
        */
        return &potassium.RenderResult{
            []potassium.IComponentProcessor{
                a.CreateElement(
                    potassium.NewComponentKey("window"),
                    potassiumgtk.NewWindowComponent,
                    potassiumgtk.NewWindowComponentProps("Window Title " + strconv.Itoa(state.clicks)),
                    []potassium.IComponentProcessor{
                        a.CreateElement(
                            potassium.NewComponentKey("row"),
                            potassiumgtk.NewRowComponent,
                            potassiumgtk.NewRowComponentProps(),
                            []potassium.IComponentProcessor{
                                a.CreateElement(
                                    potassium.NewComponentKey("appButtonRow"),
                                    newAComponent,
                                    newAComponentProps(state.clicks, func() { a.onClick(processor) }),
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
