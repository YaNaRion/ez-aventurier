use dioxus::prelude::*;

use crate::components::{InfoCard, MessageCard};

#[component]
pub fn CacheListButton() -> Element {
    rsx! {
        div {
            onclick: move|_| {
                    dioxus_router::router().push("/cache_list");
            },
            div { class: "connection-body",
                div { class: "user-info",
                    InfoCard {
                        title: "Voir la liste de toutes les caches".to_string(),
                        data: "".to_string(),
                        icon: "📜".to_string(),
                    }
                }
            }
        }
    }
}
