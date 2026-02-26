use dioxus::prelude::*;
use reqwest::Client;

use dioxus_primitives::alert_dialog::{
    AlertDialogAction, AlertDialogContent, AlertDialogDescription, AlertDialogRoot,
};

use crate::{
    components::{input::Input, InfoCard},
    service::{get_base_url, User},
};

#[component]
pub fn UserBody(user: Signal<User>, session_id: String) -> Element {
    let mut error = use_signal(|| String::new());
    let mut open = use_signal(|| false);

    let session_id_clone = session_id.clone();
    let user_id_clone = user.read().user_id.clone();
    let handle_submit = Callback::new(move |text: String| {
        let session_id_clone_2 = session_id_clone.clone();
        let user_id_clone_2 = user_id_clone.clone();
        async move {
            let client = use_context::<Client>();

            let origin = get_base_url();
            let req_string = format!(
                "{}/claimCache?session_id={}&answer_id={}&user_id={}",
                origin, session_id_clone_2, text, user_id_clone_2,
            );

            match client.put(req_string).body(text).send().await {
                Ok(resp) if resp.status().is_success() => {
                    user.set(resp.json().await.unwrap());
                    open.set(true);
                }

                Ok(resp) => {
                    let text = resp.text().await.unwrap_or_default();
                    error.set(format!("{}", text));
                    open.set(true)
                }

                Err(e) => {
                    error.set(format!("{}", e));
                    open.set(true)
                }
            }
        }
    });

    // Function to handle dialog close
    let handle_close = move |v: bool| {
        open.set(v);
        if !v {
            // Reset error when closing
            error.set(String::new());
        }
    };

    rsx! {
            AlertDialogRoot {
                open: *open.read(),
                on_open_change: handle_close,
                AlertDialogContent {
                    if error.read().is_empty() {
                        AlertDialogDescription { "La cache a été ajoutée avec succès" }
                    } else {
                        AlertDialogDescription { "{error}" }
                    }
                    AlertDialogAction {
                        class: "AlertDialogAction",
                        on_click: move |_| {
                            open.set(false);
                        },
                        "Confirm"
                    }
                }
            }
            div { class: "connection-body",
                div { class: "user-info",

                    InfoCard {
                        title: "Votre code secret".to_string(),
                        data: user.read().user_id.clone(),
                        icon: "⚔️".to_string(),
                    }

                    InfoCard {
                        title: "Unité Scout".to_string(),
                        data: user.read().unity.clone(),
                        icon: "🕯️".to_string(),
                    }

                    // InfoCard {
                    //     title: "Votre ordre".to_string(),
                    //     data: user.read().order.clone(),
                    //     icon: "🕯️".to_string(),
                    // }

                    InfoCard {
                        title: "Votre score".to_string(),
                        data: user.read().score.clone(),
                        icon: "🏆".to_string(),
                    }

                    Input {
                        input_name: "Entrer le code de la cache pour la confirmer".to_string(),
                        callback: handle_submit,
                    }

                }

            }
    }
}
