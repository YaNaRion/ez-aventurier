use dioxus::prelude::*;
use reqwest::Client;

use crate::service::User;

fn check_if_women(user: &User) -> bool {
    let is_pionne_or_eclaireuses = user.unity == "Pionnières" || user.unity == "Éclaireuses";
    let is_special_case =
        user.player_name != "Alex Labelle" || user.player_name != "Hugo Palardy-Beaud";

    is_pionne_or_eclaireuses && !is_special_case
}

#[component]
pub fn ConnectedUser(user_id: String, session_id: String) -> Element {
    let mut is_loading = use_signal(|| false);
    let mut error = use_signal(|| String::new());

    let mut user_data = use_signal(|| User {
        ..Default::default()
    });

    let client = use_context::<Client>();
    // Refresh session handler
    let user_id_clone = user_id.clone();
    let session_id_clone = session_id.clone();

    // Get Client

    use_future(move || {
        let client_for_async = client.clone(); // Assuming Client implements Clone
        let req_string = format!(
            "http://localhost:3000/api/user?user_id={}&session_id={}",
            user_id_clone.clone(),
            session_id_clone.clone()
        );

        async move {
            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    let data: User = resp.json().await.unwrap();
                    user_data.set(data);
                }

                Ok(_resp) => {
                    error.set("NOT CONNECTED".to_string());
                    dioxus_router::router().push("/");
                }

                Err(e) => {
                    error.set(format!("Network error: {}", e));
                }
            }

            is_loading.set(false);
        }
    });

    // let user_data_clone = user_data.read().unity;

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
                    if check_if_women(&user_data.read())  {
                        "Bien le Bonjour Noble Chevalière"
                    } else {
                        "Bien le Bonjour Noble Chevalier"
                    }
                }

                p {
                    class: "connection-subtitle",
                    if check_if_women(&user_data.read())  {
                        "Soyez la bienvenue: "
                    } else {
                        "Soyez le bienvenu: "
                    }
                    strong { "{user_data.read().player_name}" }
                }
            }

            // Connection Body
            div { class: "connection-body",
                div { class: "user-info",
                    // Secret Code Card
                    div { class: "info-card",
                        div { class: "info-icon", "⚔️" }
                        div { class: "info-content",
                            h3 { "Votre code secrêt" }
                            p { "{user_data.read().user_id}" }
                        }
                    }

                    // Session Info Card
                    div { class: "info-card",
                        div { class: "info-icon", "🕯️" }
                        div { class: "info-content",
                            h3 { "Unité scout" }
                            p { "{user_data.read().unity}" }
                        }
                    }

                    div { class: "info-card",
                        div { class: "info-icon", "🕯️" }
                        div { class: "info-content",
                            h3 { "Ordre" }
                            p { "{user_data.read().order}" }
                        }
                    }
                    // Message Card with Send Button on the right
                    // Message Card with aligned Send Button
                    div { class: "message-card-aligned",
                        div { class: "message-card-content",
                            div { class: "message-icon", "📜" }
                            div { class: "message-input-wrapper",
                                h3 { "Entrer le code secret pour confirmer votre quête" }
                                div { class: "input-button-group",
                                    input {
                                        class: "message-input-aligned",
                                        placeholder: "Écrivez votre code...",
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

                // Error display
                if !error().is_empty() {
                    div { class: "error-message", "{error}" }
                }
            }
        }
    }
}
