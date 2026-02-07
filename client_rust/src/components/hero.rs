use dioxus::prelude::*;

// const HEADER_SVG: Asset = asset!("/assets/header.svg");

use crate::components::ConnectionForm;

#[component]
pub fn Hero() -> Element {
    rsx! {
        // We can create elements inside the rsx macro with the element name followed by a block of attributes and children.
        div {
            div { class: "connection-container",
                // The RSX macro also supports text nodes surrounded by quotes
                ConnectionForm {}
            }
        }
    }
}
