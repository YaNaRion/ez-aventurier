use dioxus::prelude::*;
use reqwest::Client;

use crate::{
    components::InfoCard,
    service::{get_base_url, User},
};

#[component]
pub fn LeaderBoardView() -> Element {
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
            div { class: "scrollable-container",
                div { class: "connection-header",
                    div {
                        class: "success-icon",
                        style: "font-size: 48px; margin-bottom: 20px;",
                        "📜"
                    }

                    h1 {
                        class: "connection-title",
                        "Classement"
                    }
                },
                div { class: "user-info",
                    if *is_loading.read() {
                            "Loading..."
                    } else {
                      for user in users.read().iter() {
                        {
                        rsx! {
                            InfoCard {
                                    data: format!("{} des {} avec {} points",user.player_name.to_string(), user.unity, user.score.to_string()),
                                    title: "".to_string(),
                                    icon: "🛡️".to_string(),
                            }
                        }
                    },
                    }
                }
            }
        }
    }
}
