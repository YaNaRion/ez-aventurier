use dioxus::prelude::*;

use crate::service::User;

fn check_if_women(user: &User) -> bool {
    let is_pionne_or_eclaireuses = user.unity == "Pionnières" || user.unity == "Éclaireuses";
    let is_special_case =
        user.player_name != "Alex Labelle" || user.player_name != "Hugo Palardy-Beaud";

    is_pionne_or_eclaireuses && !is_special_case
}

#[component]
pub fn UserHeader(user: User) -> Element {
    rsx! {
            div { class: "connection-header",
                div {
                    class: "success-icon",
                    style: "font-size: 48px; margin-bottom: 20px;",
                    "🏰"
                }

                h1 {
                    class: "connection-title",
                    if check_if_women(&user)  {
                        "Bien le Bonjour Noble Chevalière"
                    } else {
                        "Bien le Bonjour Noble Chevalier"
                    }
                }

                p {
                    class: "connection-subtitle",
                    if check_if_women(&user)  {
                        "Soyez la bienvenue: "
                    } else {
                        "Soyez le bienvenu: "
                    }
                    strong { "{user.player_name}" }
                }
            }
    }
}
