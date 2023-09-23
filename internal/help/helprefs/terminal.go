package helprefs

// Terminal is a help page for using the Iterm2 terminal (https://iterm2.com/)
const Terminal = `{{.Underline}}SPLIT PANES                                                     {{.Clear}}
{{.Bullet}} New Tab: {{.Yellow}}{{.Command}} + T{{.Clear}}
{{.Bullet}} Split Window (Vertically): {{.Yellow}}{{.Command}} + D{{.Clear}}
{{.Bullet}} Split Window (Horizontally): {{.Yellow}}{{.Command}} + {{.Shift}} + D{{.Clear}}
{{.Bullet}} Close Active Tab: {{.Yellow}}{{.Command}} + W{{.Clear}}
{{.Bullet}} Navigate Panes: {{.Yellow}}{{.Command}} + {{.Option}} + Arrow Key{{.Clear}}
{{.Bullet}} Maximise Pane: {{.Yellow}}{{.Command}} + {{.Shift}} + Enter{{.Clear}}
{{.Bullet}} Move Pane With Mouse: {{.Yellow}}{{.Command}} + Alt + {{.Shift}} + Click And Drag{{.Clear}}
{{.Bullet}} Navigate Tabs: {{.Yellow}}{{.Command}} + Left|Right Arrow{{.Clear}}

{{.Underline}}SHELL KEY COMBOS                                                {{.Clear}}
{{.Bullet}} Clear Line: {{.Yellow}}{{.Control}} + U{{.Clear}}
{{.Bullet}} Clear All Text to Right: {{.Yellow}}{{.Control}} + K{{.Clear}}
{{.Bullet}} Exit Shell: {{.Yellow}}{{.Control}} + D{{.Clear}}

{{.Underline}}SEARCH HISTORY                                                  {{.Clear}}
{{.Bullet}} Search: {{.Yellow}}{{.Control}} + R + Type Search Term{{.Clear}}
    {{.HollowBullet}} Loop Through Search: {{.Yellow}}{{.Control}} + R{{.Clear}}
{{.Bullet}} End Search at Current Entry: {{.Yellow}}{{.Control}} + Y{{.Clear}}
{{.Bullet}} Cancel Search And Restore Original Line: {{.Yellow}}{{.Control}} + G{{.Clear}}

{{.Underline}}MISC                                                            {{.Clear}}
{{.Bullet}} Find Cursor (In Iterm2): {{.Yellow}}{{.Command}} + /{{.Clear}}
`
