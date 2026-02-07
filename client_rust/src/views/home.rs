use crate::components::{ConnectionForm, Hero};
use dioxus::prelude::*;
use reqwest::Client;

/// The Home page component that will be rendered when the current route is `[Route::Home]`
#[component]
pub fn Home() -> Element {
    rsx! {
        ConnectionForm {}
    }
}
