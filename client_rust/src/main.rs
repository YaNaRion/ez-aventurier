use dioxus::prelude::*;
use reqwest::Client;
use views::{Admin, Home, User};

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
}

const FAVICON: Asset = asset!("/assets/favicon.ico");
const MAIN_CSS: Asset = asset!("/assets/styling/main.css");

fn main() {
    // Launch the app
    dioxus::launch(App);
}

/// App is the main component of our app
#[component]
fn App() -> Element {
    // Provide the combined context once
    // use_shared_state();

    let client = Client::new();
    use_context_provider(|| client);

    rsx! {
        document::Link { rel: "icon", href: FAVICON }
        document::Link { rel: "stylesheet", href: MAIN_CSS }
        Router::<Route> {}
    }
}
