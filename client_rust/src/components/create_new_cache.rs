use dioxus::prelude::*;
use dioxus_primitives::alert_dialog::{
    AlertDialogAction, AlertDialogContent, AlertDialogDescription, AlertDialogRoot,
    AlertDialogTitle,
};
use reqwest::Client;

use crate::{
    components::MessageCard,
    service::{get_base_url, API_BASE_URL},
};

#[component]
pub fn CreateNewCache(session_id: String) -> Element {
    let mut open = use_signal(|| false);

    let mut error = use_signal(|| String::new());
    let session_id_value = session_id.clone();

    let handle_submit = Callback::new(move |text: String| {
        let session_id_value_copy = session_id_value.clone();
        async move {
            open.set(true);

            let client = use_context::<Client>();

            let origin = get_base_url();
            let req_string = format!("{}/cache?&session_id={}", origin, session_id_value_copy);

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
            MessageCard {
                input_name: "Entrez le nom de la nouvelle enigme",
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
