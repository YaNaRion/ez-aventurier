use dioxus::prelude::*;
use reqwest::Client;
use views::{Home, Navbar, User};

mod components;
mod views;

struct GlobalSharedState {
    session_id: Option<String>,
    user_id: Option<String>,
}

#[derive(Debug, Clone, Routable, PartialEq)]
#[rustfmt::skip]
enum Route {
    #[layout(Navbar)]
        #[route("/")]
            Home {},
        #[route("/user?:user_id")]
            User {
                user_id: String,
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
    // Create a single Client instance for the entire app
    let client = Client::new();
    use_context_provider(|| client);

    let isConnedted = false;
    use_context_provider(|| isConnedted);

    let session_id = "";
    use_context_provider(|| session_id);

    rsx! {
        document::Link { rel: "icon", href: FAVICON }
        document::Link { rel: "stylesheet", href: MAIN_CSS }
        Router::<Route> {}
    }
}
