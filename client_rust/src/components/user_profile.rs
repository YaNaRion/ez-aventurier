use dioxus::prelude::*;

use crate::{
    components::{UserBody, UserHeader},
    service::User,
};

#[component]
pub fn UserProfile(user: User, session_id: String) -> Element {
    rsx! {
        div { class: "connected-ui",
            UserHeader {
                user: user.clone(),
            }

            UserBody {
                user: user.clone(),
                session_id: session_id.clone(),
            }
        }
    }
}
