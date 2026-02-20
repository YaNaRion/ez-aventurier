use crate::components::{button::RoutingButton, ConnectionForm, InfoCard};
use dioxus::prelude::*;

/// The Home page component that will be rendered when the current route is `[Route::Home]`
#[component]
pub fn HomeView() -> Element {
    let navigator = use_navigator();
    rsx! {

        div { class: "connection-header",
            h1 { class: "connection-title", "Camp Aventurier 2026" }
            h2 { class: "connection-subtitle", "Pour toutes questions, veuillez les poser à l'adresse courriel suivante : camp.aventurier.229@gmail.com" }
        }
        ConnectionForm {}

        RoutingButton {
            path: "/letter".to_string(),
            title: "Lettre de camp".to_string(),
            data: "".to_string(),
            icon: "📜".to_string(),
        }
        RoutingButton {
            path: "/cache_list".to_string(),
            title: "Voir les caches".to_string(),
            data: "".to_string(),
            icon: "📜".to_string(),
        }
        RoutingButton {
            path: "/leaderboard".to_string(),
            title: "Voir le classement".to_string(),
            data: "".to_string(),
            icon: "🏆".to_string(),
        }


        div {
            onclick: {
            let navigator = navigator.clone();
                move |_| {
                    navigator.push("https://www.facebook.com/229scoutsNDN");
                }
            },
            InfoCard {
                title: "",
                data: "Ce jeu est réalisé par le 229ème groupe scout Notre-Dame des Neiges",
                icon: "⚜️",
            }
        }
    }
}
