use dioxus::prelude::*;

use crate::{
    components::{InfoCard, MessageCard},
    service::User,
};

#[component]
pub fn UserBody(user: User) -> Element {
    rsx! {
            div { class: "connection-body",
                div { class: "user-info",
                    InfoCard {
                        title: "Votre code secrêt".to_string(),
                        data: user.user_id.clone(),
                        icon: "⚔️".to_string(),
                    }

                    InfoCard {
                        title: "Unité Scout".to_string(),
                        data: user.unity.clone(),
                        icon: "🕯️".to_string(),
                    }

                    InfoCard {
                        title: "Votre ordre".to_string(),
                        data: user.order.clone(),
                        icon: "🕯️".to_string(),
                    }

                    MessageCard {
                        input_name: "Entrer le code secret pour confirmer votre quête".to_string(),
                    }

                }
            }
    }
}
