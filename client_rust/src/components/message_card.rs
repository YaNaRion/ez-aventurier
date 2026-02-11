use dioxus::prelude::*;

// Il faut ajouter la callback pour le submit
#[component]
pub fn MessageCard(input_name: String, callback: Callback<String>) -> Element {
    let mut text = use_signal(String::new);
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
                            oninput: move |evt| text.set(evt.value()),
                        }
                        button {
                            class: "send-button-aligned",
                            r#type: "button",
                            onclick: move |_| {
                                callback.call(text());
                                text.set(String::new()); // Clear input after sending
                            },
                            "Confirmer"
                        }
                    }
                }
            }
        }
    }
}
