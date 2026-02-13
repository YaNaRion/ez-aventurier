use crate::components::{Button, ConnectionForm};
use dioxus::prelude::*;

/// The Home page component that will be rendered when the current route is `[Route::Home]`
#[component]
pub fn HomeView() -> Element {
    rsx! {
        ConnectionForm {}
        Button {
            path: "/cache_list".to_string(),
            title: "Voir la liste de toutes les caches".to_string(),
            data: "".to_string(),
            icon: "📜".to_string(),
        }
        Button {
            path: "/leaderboard".to_string(),
            title: "Voir le leaderboard".to_string(),
            data: "".to_string(),
            icon: "🏆".to_string(),
        }
    }
}
