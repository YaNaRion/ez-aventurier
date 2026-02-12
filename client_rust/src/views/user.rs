use dioxus::prelude::*;
use reqwest::Client;
// use web_sys::console;

use crate::{
    components::ConnectedUser,
    service::{CheckSessionValid, API_BASE_URL},
    views::Admin,
};

// const HEADER_SVG: Asset = asset!("/assets/header.svg");

#[component]
pub fn User(session_id: String) -> Element {
    let mut is_loading = use_signal(|| true);
    let mut is_admin = use_signal(|| true);
    let mut error = use_signal(|| String::new());
    let mut session_data = use_signal(|| CheckSessionValid::default());

    let client = use_context::<Client>();

    let session_id_clone = session_id.clone();

    use_future(move || {
        // Move the cloned values into the async block
        let client_for_async = client.clone(); // Assuming Client implements Clone

        let session_id_clone_2 = session_id_clone.clone();

        async move {
            let req_string = format!(
                "{}api/isSessionValid?session_id={}",
                API_BASE_URL,
                session_id_clone_2.clone()
            );

            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    session_data.set(resp.json().await.unwrap());
                    is_admin.set(session_data.read().session.user_id == "admin");
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
            } else if session_data.read().is_valid {
                div { class: "scrollable-container",
                if *is_admin.read() {
                    Admin {
                        user_id: session_data.read().session.user_id.clone(),
                        session_id: session_data.read().session.session_id.clone(),
                    }
                } else {
                        ConnectedUser {
                            user_id: session_data.read().session.user_id.clone(),
                            session_id: session_data.read().session.session_id.clone(),
                        }
                    }
                }
            } else {
                "NOT CONNECTED"
            }
        }
    }
}
