use dioxus::prelude::*;

// Il faut ajouter la callback pour le submit
#[component]
pub fn MessageCard(input_name: String) -> Element {
    rsx! {
        div { class: "message-card-aligned",
            div { class: "message-card-content",
                div { class: "message-icon", "📜" }
                div { class: "message-input-wrapper",
                    h3 { "{input_name}" }
                    div { class: "input-button-group",
                        input {
                            class: "message-input-aligned",
                            placeholder: "Écrivez...",
                            r#type: "text",
                            // Add your bind:value logic here when implementing
                        }
                        button {
                            class: "send-button-aligned",
                            r#type: "button",
                            "Confirmer"
                        }
                    }
                }
            }
        }
    }
}
