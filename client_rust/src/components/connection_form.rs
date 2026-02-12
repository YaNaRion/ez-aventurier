use dioxus::prelude::*;
use reqwest::Client;

use crate::service::{Session, API_BASE_URL};

#[component]
pub fn ConnectionForm() -> Element {
    let mut connection_id = use_signal(|| String::new());
    let mut error = use_signal(|| String::new());
    // let mut remember_me = use_signal(|| false);
    let mut is_loading = use_signal(|| false);

    let handle_input_change = move |e: FormEvent| {
        connection_id.set(e.value());
        if !error().is_empty() {
            error.set(String::new());
        }
    };

    let handle_submit = move |evt: FormEvent| async move {
        evt.prevent_default(); // ← very important in web!

        is_loading.set(true);
        error.set(String::new());

        let client = use_context::<Client>();
        match client
            .get(format!(
                "{}/api/login?user_id={}",
                API_BASE_URL, connection_id
            ))
            .send()
            .await
        {
            Ok(resp) if resp.status().is_success() => {
                // If JSON:
                let data: Session = resp.json().await.unwrap();

                dioxus_router::router().push(format!("/user?session_id={}", data.session_id));
            }

            Ok(resp) => {
                let text = resp.text().await.unwrap_or_default();
                error.set(format!("{}", text));
            }
            Err(e) => {
                error.set(format!("Network error: {}", e));
            }
        }

        is_loading.set(false);
    };

    rsx! {
        div { class: "connection-card",
            div { class: "connection-header",
                h1 { class: "connection-title", "Code secret" }
                p { class: "connection-subtitle", "Veuillez entrer le code secret" }
            }
            div { class: "connection-body",
                form {
                    id: "connectionForm",
                    class: "connection-form",
                    onsubmit: handle_submit,
                    // prevent_default: "onsubmit",

                    div { class: "form-group",
                        label { r#for: "connectionId", class: "form-label", "Identifiant" }
                        div { class: "input-wrapper",
                            input {
                                r#type: "text",
                                id: "connectionId",
                                class: if error().is_empty() { "form-input" } else { "form-input error" },
                                placeholder: "Entrer votre code personnalisé",
                                value: "{connection_id}",
                                oninput: handle_input_change,
                                required: true,
                                autocomplete: "off",
                                autofocus: true,
                                disabled: is_loading(),
                            }
                            // your svg icon...
                        }
                        if !error().is_empty() {
                            div { class: "error-message", "{error}" }
                        }
                    }

                    button {
                        r#type: "submit",
                        class: "connect-button",
                        disabled: is_loading() || connection_id().trim().is_empty(),
                        if is_loading() {
                            span { class: "button-loader", div { class: "spinner" } }
                        } else {
                            span { class: "button-text", "⏎ Enter the Citadel" }
                        }
                    }
                }

                div { class: "connection-footer",
                    p { class: "footer-text", "Pour toutes questions, veuillez les poser au ..." }
                }
            }
        }
    }
}
