use dioxus::prelude::*;

use crate::{
    components::{AdminBody, UserHeader},
    service::User,
};

#[component]
pub fn Admin(user_id: String, session_id: String) -> Element {
    let user = User {
        user_id: "AdminID".to_string(),
        unity: "Comité Aventurier".to_string(),
        order: "Ordre des organisateurs".to_string(),
        player_name: "Adminitrateur".to_string(),
    };

    rsx! {
        div { class: "scrollable-container",
            div { class: "connected-ui",
                UserHeader {
                    user: user.clone(),
                }
                AdminBody {
                    user: user.clone(),
                    session_id: session_id.clone(),
                }
            }
        }
    }
}
