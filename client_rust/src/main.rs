use dioxus::prelude::*;
use reqwest::Client;
use views::{CacheList, Home, User};

mod components;
mod service;
mod views;

#[derive(Debug, Clone, Routable, PartialEq)]
#[rustfmt::skip]
enum Route {
    #[route("/")]
        Home {},
    #[route("/user?:session_id")]
        User {
            session_id: String,
        },

    #[route("/cache_list")]
        CacheList{},
}

const FAVICON: Asset = asset!("/assets/favicon.ico");
const MAIN_CSS: Asset = asset!("/assets/styling/main.css");

fn main() {
    dioxus::launch(App);
}

#[component]
fn App() -> Element {
    let client = Client::new();
    use_context_provider(|| client);

    rsx! {
        document::Link { rel: "icon", href: FAVICON }
        document::Link { rel: "stylesheet", href: MAIN_CSS }
        Router::<Route> {}
    }
}
