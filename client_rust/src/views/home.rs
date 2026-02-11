use crate::components::{CacheListButton, ConnectionForm};
use dioxus::prelude::*;

/// The Home page component that will be rendered when the current route is `[Route::Home]`
#[component]
pub fn Home() -> Element {
    rsx! {
        ConnectionForm {}
        CacheListButton{}
    }
}
