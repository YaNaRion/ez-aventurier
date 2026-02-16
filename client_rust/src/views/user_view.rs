use dioxus::prelude::*;
use reqwest::Client;
// use web_sys::console;

use crate::{
    components::{
        admin::Admin,
        user::{UserHeader, UserProfile},
    },
    service::{get_base_url, User},
};

// const HEADER_SVG: Asset = asset!("/assets/header.svg");

#[component]
pub fn UserView(session_id: String, user_id: String) -> Element {
    let mut is_loading = use_signal(|| true);
    let mut is_admin = use_signal(|| true);
    let mut error = use_signal(|| String::new());
    let mut user = use_signal(|| User::default());

    let client = use_context::<Client>();

    let session_id_clone = session_id.clone();
    let user_id_clone = user_id.clone();

    use_future(move || {
        // Move the cloned values into the async block
        let client_for_async = client.clone(); // Assuming Client implements Clone

        let session_id_clone_2 = session_id_clone.clone();
        let user_id_clone_2 = user_id_clone.clone();

        async move {
            let origin = get_base_url();
            let req_string = format!(
                "{}/user?user_id={}&session_id={}",
                origin,
                user_id_clone_2.clone(),
                session_id_clone_2.clone(),
            );

            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    user.set(resp.json().await.unwrap());
                    is_admin.set(user.read().user_id == "admin");
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
            } else {
                div { class: "scrollable-container",
                    if *is_admin.read() {
                        Admin {
                            user: user,
                            session_id: session_id.clone(),
                        }
                    } else {
                        UserProfile {
                            user: user,
                            session_id: session_id.clone(),
                        }
                    }
                }
            }
        }
    }
}
