use crate::components::{Button, ConnectionForm, InfoCard};
use dioxus::prelude::*;

/// The Home page component that will be rendered when the current route is `[Route::Home]`
#[component]
pub fn HomeView() -> Element {
    rsx! {
        div { class: "connection-header",
            h1 { class: "connection-title", "Camp Aventurier 2026" }
            h2 { class: "connection-subtitle", "Pour toutes questions, veuillez les poser à l'adresse courriel suivante : camp.aventurier229@gmail.com" }
        }

        ConnectionForm {}
        Button {
            path: "/cache_list".to_string(),
            title: "Voir les caches".to_string(),
            data: "".to_string(),
            icon: "📜".to_string(),
        }
        Button {
            path: "/leaderboard".to_string(),
            title: "Voir le classement".to_string(),
            data: "".to_string(),
            icon: "🏆".to_string(),
        }

        InfoCard {
            title: "",
            data: "Ce jeu est réalisé par le 229ème groupe scout Notre-Dame des Neiges",
            icon: "⚜️",
        }
    }
}
