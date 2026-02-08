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

#[derive(serde::Deserialize, serde::Serialize, Default)]
pub struct UserAPI {}

// Unity enum
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
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct User {
    #[serde(rename = "id")]
    pub id: String,

    #[serde(rename = "player_name")]
    pub player_name: String,

    #[serde(rename = "unity")]
    pub unity: Unity,

    #[serde(rename = "order")]
    pub order: Order,

    #[serde(rename = "caches")]
    pub caches: Vec<CacheID>,
}

impl Unity {
    // Get display name
    pub fn display_name(&self) -> &'static str {
        match self {
            Unity::Eclaireur => "Éclaireur",
            Unity::Eclaireuse => "Éclaireuse",
            Unity::Pion => "Pion",
            Unity::Pionne => "Pionne",
            Unity::Anim => "Animateur",
        }
    }

    // Convert from integer (like Go's iota)
    pub fn from_int(value: i32) -> Option<Self> {
        match value {
            0 => Some(Unity::Eclaireur),
            1 => Some(Unity::Eclaireuse),
            2 => Some(Unity::Pion),
            3 => Some(Unity::Pionne),
            4 => Some(Unity::Anim),
            _ => None,
        }
    }

    // Convert to integer
    pub fn to_int(&self) -> i32 {
        match self {
            Unity::Eclaireur => 0,
            Unity::Eclaireuse => 1,
            Unity::Pion => 2,
            Unity::Pionne => 3,
            Unity::Anim => 4,
        }
    }
}

impl Order {
    // Get display name
    pub fn display_name(&self) -> &'static str {
        match self {
            Order::Templar => "Ordre du Temple",
            Order::Teutonic => "Ordre Teutonique",
            Order::Hospital => "Ordre de l'Hôpital",
            Order::Lazarite => "Ordre de Saint-Lazare",
            Order::Secular => "Séculier",
        }
    }

    // Convert from integer
    pub fn from_int(value: i32) -> Option<Self> {
        match value {
            0 => Some(Order::Templar),
            1 => Some(Order::Teutonic),
            2 => Some(Order::Hospital),
            3 => Some(Order::Lazarite),
            4 => Some(Order::Secular),
            _ => None,
        }
    }

    // Convert to integer
    pub fn to_int(&self) -> i32 {
        match self {
            Order::Templar => 0,
            Order::Teutonic => 1,
            Order::Hospital => 2,
            Order::Lazarite => 3,
            Order::Secular => 4,
        }
    }
}

impl User {
    // Constructor
    pub fn new(id: String, player_name: String, unity: Unity, order: Order) -> Self {
        User {
            id,
            player_name,
            unity,
            order,
            caches: Vec::new(),
        }
    }

    // Check if user has specific cache
    pub fn has_cache(&self, cache_id: &str) -> bool {
        self.caches.iter().any(|cache| cache == cache_id)
    }

    // Add a cache if not already present
    pub fn add_cache(&mut self, cache_id: CacheID) -> bool {
        if !self.has_cache(&cache_id) {
            self.caches.push(cache_id);
            true
        } else {
            false
        }
    }

    // Remove a cache
    pub fn remove_cache(&mut self, cache_id: &str) -> bool {
        if let Some(pos) = self.caches.iter().position(|c| c == cache_id) {
            self.caches.remove(pos);
            true
        } else {
            false
        }
    }
}
