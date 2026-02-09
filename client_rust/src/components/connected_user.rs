use dioxus::prelude::*;
use reqwest::Client;

use crate::service::{ConnectionAPI, User};

#[component]
pub fn ConnectedUser(user_id: String, session_id: String) -> Element {
    let mut is_loading = use_signal(|| false);
    let mut error = use_signal(|| String::new());

    let mut user_data = use_signal(|| User {
        ..Default::default()
    });

    let mut client = use_context::<Client>();
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
                    "Bien le Bonjour Noble Chevalier"
                }

                p {
                    class: "connection-subtitle",
                    "Soyer le bienvenue: "
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
                            h3 { "Unity" }
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

                }

                // Error display
                if !error().is_empty() {
                    div { class: "error-message", "{error}" }
                }
            }
        }
    }
}
