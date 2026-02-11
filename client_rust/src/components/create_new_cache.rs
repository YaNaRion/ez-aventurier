use dioxus::prelude::*;

use crate::{
    components::{InfoCard, MessageCard},
    service::User,
};

#[component]
pub fn CreateNewCache() -> Element {
    rsx! {
        MessageCard {
            input_name: "Entrez le nom de la nouvelle enigme",
        }
    }
}
