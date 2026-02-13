//! The components module contains all shared components for our app. Components are the building blocks of dioxus apps.
//! They can be used to defined common UI elements like buttons, forms, and modals. In this template, we define a Hero
//! component  to be used in our app.

mod connection_form;
pub use connection_form::ConnectionForm;

mod user_profile;
pub use user_profile::UserProfile;

mod admin_body;
pub use admin_body::AdminBody;

mod user_header;
pub use user_header::UserHeader;

mod user_body;
pub use user_body::UserBody;

mod info_card;
pub use info_card::InfoCard;

mod message_card;
pub use message_card::MessageCard;

mod create_new_cache;
pub use create_new_cache::CreateNewCache;

mod button;
pub use button::Button;

pub mod alert_dialog;
