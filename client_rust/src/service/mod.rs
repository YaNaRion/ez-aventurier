use serde::{Deserialize, Serialize};

#[derive(serde::Deserialize, serde::Serialize, Default)]
pub struct Session {
    #[serde(rename = "sessionID")]
    pub session_id: String,
    #[serde(rename = "userID")]
    pub user_id: String,
    #[serde(rename = "createdOn")]
    pub jcreated_on: String, // or use chrono::DateTime
}

#[derive(serde::Deserialize, serde::Serialize, Default)]
pub struct CheckSessionValid {
    #[serde(rename = "session")]
    pub session: Session,
    #[serde(rename = "isValid")]
    pub is_valid: bool,
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
}
