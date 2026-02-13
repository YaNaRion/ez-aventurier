use crate::service::{get_base_url, Cache};
use dioxus::prelude::*;
use reqwest::Client;

use crate::components::InfoCard;

#[component]
pub fn CacheListView() -> Element {
    let client = use_context::<Client>();
    let mut error = use_signal(|| String::new());
    let mut is_loading = use_signal(|| true);
    let mut caches = use_signal(|| Vec::<Cache>::new());

    use_future(move || {
        // Move the cloned values into the async block
        let client_for_async = client.clone(); // Assuming Client implements Clone

        async move {
            let origin = get_base_url();
            let req_string = format!("{}/caches", origin);

            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    caches.set(resp.json().await.unwrap());
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
        div { class: "scrollable-container",
            div { class: "connection-header",
                div {
                    class: "success-icon",
                    style: "font-size: 48px; margin-bottom: 20px;",
                    "📜"
                }

                h1 {
                    class: "connection-title",
                    "LISTE DES CACHES DISPONIBLES",
                }
            },
            div { class: "user-info",
                if *is_loading.read() {
                        "Loading..."
                } else {
                  for cache in caches.read().iter() {
                    {
                    let cache_number = cache.cache_number.clone();
                    rsx! {
                        div {
                                onclick: move |_| {
                                        dioxus_router::router().push(format!("/cache?cache_number={}", cache_number));
                                },
                                InfoCard {
                                    title: format!("Cache numéro: {}", cache.cache_number),
                                    data: cache.text.clone(),
                                    icon: "📜".to_string(),
                                }
                            },
                        }
                    }
                    }
                },
            }
        }
    }
}
