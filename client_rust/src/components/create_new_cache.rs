use dioxus::prelude::*;
use dioxus_primitives::alert_dialog::{
    AlertDialogAction, AlertDialogContent, AlertDialogDescription, AlertDialogRoot,
    AlertDialogTitle,
};
use reqwest::Client;
use serde::{Deserialize, Serialize};

use crate::service::get_base_url;

#[derive(Serialize, Deserialize, Debug, Clone, Default, PartialEq)]
pub struct NewCacheRequest {
    #[serde(rename = "name")]
    pub name: String,

    #[serde(rename = "description")]
    pub description: String,
}

#[component]
pub fn CreateNewCache(session_id: String) -> Element {
    let mut open = use_signal(|| false);
    let error = use_signal(|| String::new());
    let is_loading = use_signal(|| false);

    let session_id_value = session_id.clone();
    let name: Signal<String> = use_signal(|| String::new());
    let description: Signal<String> = use_signal(|| String::new());

    let handle_submit = Callback::new(move |_| {
        let session_id_value_copy = session_id_value.clone();
        let mut name_clone = name.clone();
        let mut description_clone = description.clone();
        let mut error_clone = error.clone();
        let mut is_loading_clone = is_loading.clone();
        let mut open_clone = open.clone();

        async move {
            is_loading_clone.set(true);
            error_clone.set(String::new());

            let client = use_context::<Client>();
            let origin = get_base_url();
            let url = format!("{}/cache?session_id={}", origin, session_id_value_copy);

            // Create JSON body with name and description
            let body = NewCacheRequest {
                name: name_clone.read().clone(),
                description: description_clone.read().clone(),
            };

            match client
                .post(&url)
                .header("Content-Type", "application/json")
                .json(&body)
                .send()
                .await
            {
                Ok(resp) if resp.status().is_success() => {
                    open_clone.set(true);
                    error_clone.set(String::new());
                    // Clear form fields
                    name_clone.set(String::new());
                    description_clone.set(String::new());
                }

                Ok(resp) => {
                    let status = resp.status();
                    let error_text = resp
                        .text()
                        .await
                        .unwrap_or_else(|_| "Unknown error".to_string());
                    error_clone.set(format!("Error {}: {}", status, error_text));
                }
                Err(e) => {
                    error_clone.set(format!("Request failed: {}", e));
                }
            }
            is_loading_clone.set(false);
        }
    });

    rsx! {
        FormCard {
            name: name,
            description: description,
            on_submit: handle_submit,
            is_loading: is_loading,
        }

        AlertDialogRoot {
            open: *open.read(),
            on_open_change: move |v| open.set(v),
            AlertDialogContent {
                AlertDialogTitle { "Creation Status" }
                if error.read().is_empty() {
                    AlertDialogDescription { "La cache a été ajoutée avec succès" }
                } else {
                    AlertDialogDescription { "Error: {error}" }
                }
                AlertDialogAction { "OK" }
            }
        }
    }
}

#[component]
pub fn FormCard(
    mut name: Signal<String>,
    mut description: Signal<String>,
    on_submit: Callback<()>,
    is_loading: Signal<bool>,
) -> Element {
    rsx! {
        div { class: "form-card-aligned",
            div { class: "form-card-content",
                div { class: "form-input-wrapper",
                    h3 { "Ajouter une cache" }

                    // Name input
                    div { class: "form-group",
                        label { class: "form-label", "Nom" }
                        input {
                            class: "form-input",
                            placeholder: "Entrer le nom de la cache",
                            value: "{name}",
                            r#type: "text",
                            disabled: is_loading(),
                            oninput: move |evt| name.set(evt.value()),
                        }
                    }

                    // Description input
                    div { class: "form-group",
                        label { class: "form-label", "Description" }
                        textarea {
                            class: "form-textarea",
                            placeholder: "Enter la description de la cache",
                            rows: "3",
                            disabled: is_loading(),
                            oninput: move |evt| description.set(evt.value()),
                            "{description}"
                        }
                    }

                    // Submit button
                    div { class: "form-actions",
                        button {
                            class: "submit-button",
                            r#type: "button",
                            disabled: is_loading() || name.read().is_empty() || description.read().is_empty(),
                            onclick: move |_| {
                                if !is_loading() {
                                    on_submit.call(());
                                }
                            },
                            if is_loading() {
                                "Creating..."
                            } else {
                                "Créer la cache"
                            }
                        }
                    }
                }
            }
        }
    }
}
