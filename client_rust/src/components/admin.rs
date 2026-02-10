use crate::components::ConnectionForm;
use dioxus::prelude::*;

#[component]
pub fn Admin() -> Element {
    rsx! {
        div { class: "connected-ui",
            div { class: "connection-header",
                div {
                    class: "success-icon",
                    style: "font-size: 48px; margin-bottom: 20px;",
                    "🏰"
                }

                h1 {
                    class: "connection-title",
                        "Bien le Bonjour admin"
                }
            }

            // Connection Body
            div { class: "connection-body",
                div { class: "user-info",
                    // Session Info Card
                    div { class: "info-card",
                        div { class: "info-icon", "🕯️" }
                        div { class: "info-content",
                            h3 { "Ajouter une capsule" }
                            p { "Douille" }
                        }
                    }
                }
            }

    }
}
