use crate::components::{Admin, ConnectionForm};
use dioxus::prelude::*;

#[component]
pub fn AdminView() -> Element {
    rsx! {
        div { "Admin view" }
        Admin{}
    }
}
