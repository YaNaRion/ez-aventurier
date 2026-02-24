use serde::{Deserialize, Serialize};

// OR if you want direct backend access (bypass Caddy):
pub const API_BASE_URL: &str = if cfg!(debug_assertions) {
    ":3000" // Development: direct to backend
} else {
    "/api" // Production: through Caddy
};

pub fn get_base_url() -> String {
    format!(
        "{}{}",
        web_sys::window()
            .and_then(|w| w.location().origin().ok())
            .map(|origin| {
                // Remove port number
                origin
                    .split(':')
                    .take(2) // Take protocol and hostname
                    .collect::<Vec<&str>>()
                    .join(":")
            })
            .unwrap_or_else(|| "https://yourdomain.com".to_string()),
        API_BASE_URL
    )
}

#[derive(serde::Deserialize, serde::Serialize, Default)]
pub struct Session {
    #[serde(rename = "sessionID")]
    pub session_id: String,
    #[serde(rename = "userID")]
    pub user_id: String,
    #[serde(rename = "createdOn")]
    pub jcreated_on: String, // or use chrono::DateTime
}

#[derive(Serialize, Deserialize, Debug, Clone, Default, PartialEq)]
pub struct User {
    #[serde(rename = "userId")]
    pub user_id: String,

    #[serde(rename = "name")]
    pub player_name: String,

    #[serde(rename = "unity")]
    pub unity: String,

    #[serde(rename = "order")]
    pub order: String,

    #[serde(rename = "score")]
    pub score: i32,
}

#[derive(Serialize, Deserialize, Debug, Clone, Default, PartialEq)]
pub struct Cache {
    #[serde(rename = "cacheNumber")]
    pub cache_number: i32,

    #[serde(rename = "name")]
    pub name: String,

    #[serde(rename = "description")]
    pub description: String,

    #[serde(rename = "createdAt")]
    pub jcreated_on: String, // or use chrono::DateTime
}
