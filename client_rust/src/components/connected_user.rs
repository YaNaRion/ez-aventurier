use dioxus::prelude::*;
use reqwest::Client;

use crate::service::ConnectionAPI;

#[component]
pub fn ConnectedUser(user_id: String, session_id: String) -> Element {
    let mut is_loading = use_signal(|| false);
    let mut error = use_signal(|| String::new());

    let mut client = use_context::<Client>();
    // Refresh session handler

    let user_id_clone = user_id.clone();
    let session_id_clone = session_id.clone();

    async move {
        let client_for_async = client.clone(); // Assuming Client implements Clone
        let req_string = format!(
            "http://localhost:3000/api/user?user_id={}&session_id={}",
            user_id_clone.clone(),
            session_id_clone.clone()
        );

        match client_for_async.get(&req_string).send().await {
            Ok(resp) if resp.status().is_success() => {
                let data: IsSessionValidAPI = resp.json().await.unwrap();
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
    };

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
                    "Bien le Bonjour Noble Chevalier"
                }

                p {
                    class: "connection-subtitle",
                    "Soyer le bienvenue: "
                    strong { "{user_id}" }
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
                            p { "{user_id}" }
                        }
                    }

                    // Session Info Card
                    div { class: "info-card",
                        div { class: "info-icon", "🕯️" }
                        div { class: "info-content",
                            h3 { "Session ID" }
                            p { "{session_id}" }
                        }
                    }

                    // Conversations Card (clickable)
                    div {
                        class: "info-card clickable",
                        onclick: move |_| {
                            dioxus_router::router().push(format!(
                                "/conversations?user_id={}&session_id={}",
                                user_id, session_id
                            ));
                        },
                        div { class: "info-icon", "💬" }
                        div { class: "info-content",
                            h3 { "Voir vos conversations" }
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
