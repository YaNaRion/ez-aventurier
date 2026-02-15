use dioxus::prelude::*;

use crate::{
    components::{UserBody, UserHeader},
    service::User,
};

#[component]
pub fn UserProfile(user: Signal<User>, session_id: String) -> Element {
    rsx! {
        div { class: "connected-ui",
            UserHeader {
                user: user,
            }

            UserBody {
                user: user,
                session_id: session_id.clone(),
            }
        }
    }
}
