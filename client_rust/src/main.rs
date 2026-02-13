use dioxus::prelude::*;
use reqwest::Client;
use views::{CacheDetailView, CacheListView, HomeView, LeaderBoardView, UserView};

mod components;
mod service;
mod views;

#[derive(Debug, Clone, Routable, PartialEq)]
#[rustfmt::skip]
enum Route {
    #[route("/")]
        HomeView {},
    #[route("/user?:session_id&:user_id")]
        UserView {
            session_id: String,
            user_id: String,
        },

    #[route("/cache?:cache_number")]
        CacheDetailView {
            cache_number: String,
        },

    #[route("/cache_list")]
        CacheListView{},

    #[route("/leaderboard")]
        LeaderBoardView{},
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
