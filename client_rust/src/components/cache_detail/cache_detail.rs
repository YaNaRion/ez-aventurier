use dioxus::prelude::*;
const STYLE: Asset = asset!("./cache_detail.css");

use crate::service::Cache;

#[component]
pub fn CacheDetail(cache: Signal<Cache>) -> Element {
    rsx! {

        document::Link { rel: "stylesheet", href: STYLE }
    div { class: "cache-detail-view-header",
                        div {
                            class: "success-icon",
                            style: "font-size: 48px; margin-bottom: 20px;",
                            "📜"
                        }

                        h1 {
                            class: "cache-detail-view-title",
                                "Voici les indices pour la cache: {cache.read().cache_number}"
                            }

                        h2 {
                            class: "cache-detail-name",
                            strong { "Nom: {cache.read().name}" }
                        }

                        p {
                            class: "cache-detail-description",
                            strong { "Description: {cache.read().description}" }
                        }
                    }
        }
}
