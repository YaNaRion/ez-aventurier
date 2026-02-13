use dioxus::prelude::*;

use crate::{
    components::{AdminBody, UserHeader},
    service::User,
};

#[component]
pub fn Admin(user: User, session_id: String) -> Element {
    rsx! {
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
