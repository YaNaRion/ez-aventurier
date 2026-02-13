use dioxus::prelude::*;

#[component]
pub fn LeaderBoardView() -> Element {
    rsx! {
        div { class: "scrollable-container",
            "leader board",
        }
    }
}
