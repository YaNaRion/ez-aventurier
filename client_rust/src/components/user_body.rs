use dioxus::prelude::*;
use reqwest::Client;

use dioxus_primitives::alert_dialog::{
    AlertDialogAction, AlertDialogContent, AlertDialogDescription, AlertDialogRoot,
    AlertDialogTitle,
};

use crate::{
    components::{InfoCard, MessageCard},
    service::{get_base_url, User},
};

#[component]
pub fn UserBody(user: User, session_id: String) -> Element {
    let mut error = use_signal(|| String::new());
    let mut open = use_signal(|| false);

    let session_id_clone = session_id.clone();
    let user_id_clone = user.user_id.clone();
    let handle_submit = Callback::new(move |text: String| {
        let session_id_clone_2 = session_id_clone.clone();
        let user_id_clone_2 = user_id_clone.clone();
        async move {
            open.set(true);

            let client = use_context::<Client>();

            let origin = get_base_url();
            let req_string = format!(
                "{}/claimCache?session_id={}&answer_id={}&user_id={}",
                origin, session_id_clone_2, text, user_id_clone_2,
            );

            match client.put(req_string).body(text).send().await {
                Ok(resp) if resp.status().is_success() => {
                    open.set(true);
                }

                Ok(resp) => {
                    let text = resp.text().await.unwrap_or_default();
                    error.set(format!("{}", text));
                }

                Err(e) => {
                    error.set(format!("{}", e));
                }
            }
        }
    });

    rsx! {
            div { class: "connection-body",
                div { class: "user-info",
                    InfoCard {
                        title: "Votre code secrêt".to_string(),
                        data: user.user_id.clone(),
                        icon: "⚔️".to_string(),
                    }

                    InfoCard {
                        title: "Unité Scout".to_string(),
                        data: user.unity.clone(),
                        icon: "🕯️".to_string(),
                    }

                    InfoCard {
                        title: "Votre ordre".to_string(),
                        data: user.order.clone(),
                        icon: "🕯️".to_string(),
                    }

                    InfoCard {
                        title: "Votre score".to_string(),
                        data: user.score.clone(),
                        icon: "🏆".to_string(),
                    }

                    MessageCard {
                        input_name: "Entrer le code secret pour confirmer votre quête".to_string(),
                        callback: handle_submit,
                    }

                    AlertDialogRoot { open: *open.read(), on_open_change: move |v| open.set(v),
                            AlertDialogContent {
                                // You may pass class/style for custom appearance
                                AlertDialogTitle { "Title" }
                                if error.read().is_empty() {
                                    AlertDialogDescription { "La cache a été ajoutée avec succès" }
                                } else {
                                    AlertDialogDescription { "Error: {error}" }
                                }
                                AlertDialogAction { "Confirm" }
                            }
                        }

                }
            }
    }
}
