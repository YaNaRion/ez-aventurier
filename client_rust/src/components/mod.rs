//! The components module contains all shared components for our app. Components are the building blocks of dioxus apps.
//! They can be used to defined common UI elements like buttons, forms, and modals. In this template, we define a Hero
//! component  to be used in our app.

mod connection_form;
pub use connection_form::ConnectionForm;

mod info_card;
pub use info_card::InfoCard;

pub mod admin;
pub use admin::*;

pub mod alert_dialog;
pub mod button;
pub mod cache_detail;
pub mod input;
pub mod user;
