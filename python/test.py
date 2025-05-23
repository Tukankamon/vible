#nix-shell -p python312Packages.prompt-toolkit

from prompt_toolkit.application import Application
from prompt_toolkit.key_binding import KeyBindings
from prompt_toolkit.layout import Layout
from prompt_toolkit.widgets import Box, Frame, TextArea

items = ["Genesis", "Exodus", "Leviticus", "Numbers", "Deuteronomy"]
selected = [0]

def get_text():
    return "\n".join(
        f"{'>' if i == selected[0] else ' '} {item}"
        for i, item in enumerate(items)
    )

text_area = TextArea(text=get_text(), focusable=True, scrollbar=True)

kb = KeyBindings()

@kb.add('j')
def down(event):
    if selected[0] < len(items) - 1:
        selected[0] += 1
        text_area.text = get_text()

@kb.add('k')
def up(event):
    if selected[0] > 0:
        selected[0] -= 1
        text_area.text = get_text()

@kb.add('q')
def exit_(event):
    event.app.exit()

app = Application(
    layout=Layout(Box(Frame(text_area))),
    key_bindings=kb,
    full_screen=True,
)

app.run()