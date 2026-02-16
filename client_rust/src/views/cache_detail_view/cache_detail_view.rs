use dioxus::prelude::*;
use reqwest::Client;

use crate::{
    components::cache_detail::CacheDetail,
    service::{get_base_url, Cache},
};

const STYLE: Asset = asset!("./cache_detail_view.css");

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
                }

                Ok(_resp) => {
                    dioxus_router::router().push("/cache_list");
                }

                Err(e) => {
                    error.set(format!("Network error: {}", e));
                }
            }

            is_loading.set(false);
        }
    });

    rsx! {

        document::Link { rel: "stylesheet", href: STYLE }

        div {
            if *is_loading.read() {
                "Loading..."
            } else if !error.read().is_empty() {
                "Error: {error}"
            } else {
                CacheDetail {
                    cache: cache,
                }
            }
        }
    }
}
