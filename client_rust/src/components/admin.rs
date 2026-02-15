use dioxus::prelude::*;

use crate::{
    components::{AdminBody, UserHeader},
    service::User,
};

#[component]
pub fn Admin(user: Signal<User>, session_id: String) -> Element {
    rsx! {
        div { class: "connected-ui",
            UserHeader {
                user: user,
            }
            AdminBody {
                user: user,
                session_id: session_id.clone(),
            }
        }
    }
}
