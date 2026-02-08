#[derive(serde::Deserialize, Default)]
pub struct SessionAPI {
    #[serde(rename = "sessionID")]
    pub session_id: String,
    #[serde(rename = "userID")]
    pub user_id: String,
    #[serde(rename = "createdOn")]
    pub jcreated_on: String, // or use chrono::DateTime
}

#[derive(serde::Deserialize, Default)]
pub struct IsSessionValidAPI {
    #[serde(rename = "session")]
    pub session: SessionAPI,
    #[serde(rename = "isValid")]
    pub is_valid: bool,
}
