use dioxus::prelude::*;

#[component]
pub fn ClientProvider(client: Client) -> Element {
    use_context_provider(|| client);

    rsx! {
        Outlet::<Route> {}
    }
}
