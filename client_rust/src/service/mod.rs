use serde::{Deserialize, Serialize};

#[derive(serde::Deserialize, serde::Serialize, Default)]
pub struct ConnectionAPI {
    #[serde(rename = "sessionID")]
    pub session_id: String,
    #[serde(rename = "userID")]
    pub user_id: String,
    #[serde(rename = "createdOn")]
    pub jcreated_on: String, // or use chrono::DateTime
}

#[derive(serde::Deserialize, serde::Serialize, Default)]
pub struct IsSessionValidAPI {
    #[serde(rename = "session")]
    pub session: ConnectionAPI,
    #[serde(rename = "isValid")]
    pub is_valid: bool,
}

#[derive(Serialize, Deserialize, Debug, Clone, Copy, PartialEq, Eq)]
#[serde(rename_all = "UPPERCASE")]
pub enum Unity {
    #[serde(rename = "ECLAIREUR")]
    Eclaireur,

    #[serde(rename = "ECLAIREUSE")]
    Eclaireuse,

    #[serde(rename = "PION")]
    Pion,

    #[serde(rename = "PIONNE")]
    Pionne,

    #[serde(rename = "ANIM")]
    Anim,
}

// Order enum
#[derive(Serialize, Deserialize, Debug, Clone, Copy, PartialEq, Eq)]
#[serde(rename_all = "UPPERCASE")]
pub enum Order {
    #[serde(rename = "TEMPLAR")]
    Templar,

    #[serde(rename = "TEUTONIC")]
    Teutonic,

    #[serde(rename = "HOSPITAL")]
    Hospital,

    #[serde(rename = "LAZARITE")]
    Lazarite,

    #[serde(rename = "SECULAR")]
    Secular,
}

// CacheID as a type alias for String
pub type CacheID = String;

// User struct
#[derive(Serialize, Deserialize, Debug, Clone, Default)]
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
