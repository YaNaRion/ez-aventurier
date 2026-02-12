use dioxus::prelude::*;
use reqwest::Client;

use crate::{
    components::{UserBody, UserHeader},
    service::{User, API_BASE_URL},
};

#[component]
pub fn ConnectedUser(user_id: String, session_id: String) -> Element {
    let mut is_loading = use_signal(|| false);
    let mut error = use_signal(|| String::new());

    let mut user_data = use_signal(|| User {
        ..Default::default()
    });

    let client = use_context::<Client>();
    // Refresh session handler
    let user_id_clone = user_id.clone();
    let session_id_clone = session_id.clone();

    // Get Client

    use_future(move || {
        let client_for_async = client.clone(); // Assuming Client implements Clone
        let req_string = format!(
            "{}api/user?user_id={}&session_id={}",
            API_BASE_URL,
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
            UserHeader {
                user: user_data.read().clone(),
            }

            UserBody {
                user: user_data.read().clone(),
            }
        }
    }
}
