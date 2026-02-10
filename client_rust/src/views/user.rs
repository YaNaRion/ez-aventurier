use dioxus::prelude::*;
use reqwest::Client;
// use web_sys::console;

use crate::components::ConnectedUser;

// const HEADER_SVG: Asset = asset!("/assets/header.svg");

#[component]
pub fn User(user_id: String, session_id: String) -> Element {
    let mut is_loading = use_signal(|| true);
    let mut is_valid_session = use_signal(|| false);
    let mut error = use_signal(|| String::new());

    let client = use_context::<Client>();

    let user_id_clone = user_id.clone();
    let session_id_clone = session_id.clone();

    use_future(move || {
        // Move the cloned values into the async block
        let client_for_async = client.clone(); // Assuming Client implements Clone

        let user_id_clone_2 = user_id_clone.clone();
        let session_id_clone_2 = session_id_clone.clone();

        async move {
            let req_string = format!(
                "http://localhost:3000/api/isSessionValid?user_id={}&session_id={}",
                user_id_clone_2.clone(),
                session_id_clone_2.clone()
            );

            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    is_valid_session.set(true);
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
    rsx! {
        div {
            if *is_loading.read() {
                "Loading..."
            } else if !error.read().is_empty() {
                "Error: {error}"
            } else if *is_valid_session.read() {
                div { class: "scrollable-container",
                            ConnectedUser {
                                user_id: user_id.to_string(),
                                session_id: session_id.to_string(),
                        }
                }
            } else {
                "NOT CONNECTED"
            }
        }
    }
}
