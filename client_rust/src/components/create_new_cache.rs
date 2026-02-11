use dioxus::prelude::*;
use dioxus_primitives::alert_dialog::{
    AlertDialogAction, AlertDialogContent, AlertDialogDescription, AlertDialogRoot,
    AlertDialogTitle,
};
use reqwest::Client;

use crate::components::MessageCard;

#[component]
pub fn CreateNewCache(session_id: String) -> Element {
    let mut open = use_signal(|| false);

    let session_id_value = session_id.clone();

    let handle_submit = Callback::new(move |text: String| {
        let session_id_value_copy = session_id_value.clone();
        async move {
            open.set(true);

            let client = use_context::<Client>();

            match client
                .post(format!(
                    "http://localhost:3000/api/cache?cache_txt={}&session_id={}",
                    text,
                    session_id_value_copy.clone(),
                ))
                .send()
                .await
            {
                Ok(resp) if resp.status().is_success() => {
                    open.set(true);
                }

                Ok(resp) => {
                    // let text = resp.text().await.unwrap_or_default();
                    // error.set(format!("{}", text));
                }
                Err(e) => {
                    // error.set(format!("Network error: {}", e));
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
                        AlertDialogDescription { "Description" }
                        AlertDialogAction { "Confirm" }
                    }
                }
    }
}
