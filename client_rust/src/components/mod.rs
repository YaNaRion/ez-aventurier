//! The components module contains all shared components for our app. Components are the building blocks of dioxus apps.
//! They can be used to defined common UI elements like buttons, forms, and modals. In this template, we define a Hero
//! component  to be used in our app.

mod connection_form;
pub use connection_form::ConnectionForm;

mod connected_user;
pub use connected_user::ConnectedUser;

mod admin;
pub use admin::Admin;
