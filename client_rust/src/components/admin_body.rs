use dioxus::prelude::*;

use crate::{
    components::{CreateNewCache, InfoCard},
    service::User,
};

#[component]
pub fn AdminBody(user: User, session_id: String) -> Element {
    rsx! {
            div { class: "connection-body",
                div { class: "user-info",
                    InfoCard {
                        title: "Votre code secrêt".to_string(),
                        data: "Cest pas vrai cest trop dangeureux",
                        icon: "⚔️".to_string(),
                    }

                    CreateNewCache{
                        session_id: session_id.clone(),
                    }
                }
            }
    }
}
