use dioxus::prelude::*;

use crate::components::InfoCard;

#[component]
pub fn RoutingButton(path: String, title: String, data: String, icon: String) -> Element {
    rsx! {
        div {
            onclick: move|_| {
                    dioxus_router::router().push(path.clone());
            },
            div { class: "connection-body",
                div { class: "user-info",
                    InfoCard {
                        title: title,
                        data: data,
                        icon: icon,
                    }
                }
            }
        }
    }
}
