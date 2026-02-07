use dioxus::prelude::*;
use reqwest::Client;

// const HEADER_SVG: Asset = asset!("/assets/header.svg");

#[component]
pub fn User(user_id: String) -> Element {
    let mut is_loading = use_signal(|| true);
    let mut is_valid_session = use_signal(|| false);
    let mut error = use_signal(|| String::new());

    let value = user_id.clone();
    let client = use_context::<Client>();

    use_future(move || {
        // Move the cloned values into the async block
        let client_for_async = client.clone(); // Assuming Client implements Clone
        let user_id_for_async = value.clone();

        async move {
            let req_string = format!(
                "http://localhost:3000/api/isSessionValid?user_id={}&session_id={}",
                user_id_for_async, user_id_for_async
            );

            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    is_valid_session.set(true);
                }
                Ok(_resp) => {
                    error.set("NOT CONNECTED".to_string());
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
                "USER"
            } else {
                "NOT CONNECTED"
            }
        }
    }
}
