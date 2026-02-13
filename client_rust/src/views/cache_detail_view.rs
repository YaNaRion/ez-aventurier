use dioxus::prelude::*;
use reqwest::Client;

use crate::service::{get_base_url, Cache};

#[component]
pub fn CacheDetailView(cache_number: String) -> Element {
    let client = use_context::<Client>();
    let mut error = use_signal(|| String::new());
    let mut is_loading = use_signal(|| true);
    let mut cache = use_signal(|| Cache::default());

    use_future(move || {
        // Move the cloned values into the async block
        let client_for_async = client.clone(); // Assuming Client implements Clone

        let cache_number_clone = cache_number.clone();
        async move {
            let origin = get_base_url();
            let req_string = format!(
                "{}/cache?cache_number={}",
                origin,
                cache_number_clone.clone()
            );

            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    cache.set(resp.json().await.unwrap());
                    web_sys::console::log_1(&cache.read().text.clone().into());
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
            } else {
                div { class: "connection-header",
                    div {
                        class: "success-icon",
                        style: "font-size: 48px; margin-bottom: 20px;",
                        "📜"
                    }

                    h1 {
                        class: "connection-title",
                            "Voici les indices pour la cache: {cache.read().cache_number}"
                        }

                    p {
                        class: "connection-subtitle",
                        strong { "Indice: {cache.read().text}" }
                    }
                }
            }
        }
    }
}
