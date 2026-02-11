use dioxus::prelude::*;

#[component]
pub fn InfoCard(title: String, data: String, icon: String) -> Element {
    rsx! {
        // Session Info Card
        div { class: "info-card",
            div { class: "info-icon", "{icon}" }
            div { class: "info-content",
                h3 { "{title}" }
                p { "{data}" }
            }
        }
    }
}
