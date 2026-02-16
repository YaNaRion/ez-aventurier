use dioxus::prelude::*;
use reqwest::Client;

use crate::service::{get_base_url, User};

const STYLE: Asset = asset!("./leaderboard.css");

#[component]
pub fn Leaderboard() -> Element {
    let client = use_context::<Client>();
    let mut error = use_signal(|| String::new());
    let mut is_loading = use_signal(|| true);
    let mut users = use_signal(|| Vec::<User>::new());

    use_future(move || {
        // Move the cloned values into the async block
        let client_for_async = client.clone(); // Assuming Client implements Clone

        async move {
            let origin = get_base_url();
            let req_string = format!("{}/leaderboard", origin);

            match client_for_async.get(&req_string).send().await {
                Ok(resp) if resp.status().is_success() => {
                    users.set(resp.json().await.unwrap());
                }

                Ok(_resp) => {
                    error.set("NOT CONNECTED".to_string());
                }

                Err(e) => {
                    error.set(format!("Network error: {}", e));
                }
            }

            is_loading.set(false);
        }
    });

    rsx! {
        document::Link { rel: "stylesheet", href: STYLE }
        div { class: "scrollable-container",
            div { class: "connection-header",
                div {
                    class: "success-icon",
                    style: "font-size: 48px; margin-bottom: 20px;",
                    "🏆"
                }

                h1 {
                    class: "connection-title",
                    "Classement"
                }
            }

            div { class: "leaderboard-table-container",
                if *is_loading.read() {
                    div { class: "loading-message", "Loading..." }
                } else {
                    table { class: "leaderboard-table",
                        thead {
                            tr {
                                // To remove a column, just comment out its <th> line
                                th { class: "rank-col", "Rang" }
                                th { class: "player-col", "Joueur" }
                                th { class: "unity-col", "Unité" }  // Unity column
                                th { class: "ordre-col", "Ordre" }  // Unity column
                                th { class: "score-col", "Points" }

                            }
                        }
                        tbody {
                            for (index, user) in users.read().iter().enumerate() {
                                tr {
                                    // Use a combination of index and player name as key instead of id
                                    key: "{index}_{user.player_name}",

                                    // Rank column (always shows)
                                    td { class: "rank-col", "{index + 1}" }

                                    // Player name column - comment to hide
                                    // td { class: "player-col", "{user.player_name}" }
                                    td { class: "player-col", "{user.player_name}" }

                                    // Unity column - comment to hide
                                    td { class: "unity-col", "{user.unity}" }

                                    // Unity column - comment to hide
                                    td { class: "order-col", "{user.order}" }

                                    // Score column - comment to hide
                                    // td { class: "score-col", "{user.score} points" }
                                    td { class: "score-col", "{user.score} points" }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
