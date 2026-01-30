use dioxus::prelude::*;

fn handler_submit(event: Event<FormData>) {
    event.prevent_default();
    web_sys::console::log_1(&"Hello from Dioxus!".into());
}

#[component]
pub fn ConnectionForm() -> Element {
    let mut connection_id = use_signal(|| String::new());
    let mut error = use_signal(|| String::new());
    let mut remember_me = use_signal(|| false);
    let mut is_loading = use_signal(|| false);

    let handle_input_change = move |e: FormEvent| {
        connection_id.set(e.value());
        // Clear error when user starts typing
        if !error().is_empty() {
            error.set(String::new());
        }
    };

    let handle_remember_me_change = move |_| {
        remember_me.toggle();
    };

    let handle_submit = move |_| {
        // Handle form submission logic here
        // You'll need to implement your actual submit logic
        is_loading.set(true);
        // Simulate async operation
        spawn(async move {
            // Your async logic here
            // is_loading.set(false);
        });
    };

    rsx! {
        div {
            class: "connection-card",
            div {
                class: "connection-header",
                h1 {
                    class: "connection-title",
                    "Code secret"
                }
                p {
                    class: "connection-subtitle",
                    "Veuillez entrer le code secret"
                }
            }

            div {
                class: "connection-body",
                form {
                    id: "connectionForm",
                    class: "connection-form",
                    onsubmit: handle_submit,
                    prevent_default: "onsubmit",

                    div {
                        class: "form-group",
                        label {
                            r#for: "connectionId",
                            class: "form-label",
                            "Identifiant"
                        }
                        div {
                            class: "input-wrapper",
                            input {
                                r#type: "text",
                                id: "connectionId",
                                class: if error().is_empty() { "form-input" } else { "form-input error" },
                                placeholder: "Entrer votre code personnalisé",
                                value: "{connection_id}",
                                oninput: handle_input_change,
                                required: true,
                                autocomplete: "off",
                                autofocus: true,
                                disabled: is_loading(),
                            }
                            div {
                                class: "input-icon",
                                svg {
                                    xmlns: "http://www.w3.org/2000/svg",
                                    width: "22",
                                    height: "22",
                                    view_box: "0 0 24 24",
                                    fill: "none",
                                    stroke: "currentColor",
                                    stroke_width: "2",
                                    path {
                                        d: "M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"
                                    }
                                    circle {
                                        cx: "12",
                                        cy: "7",
                                        r: "4"
                                    }
                                }
                            }
                        }
                        if !error().is_empty() {
                            div {
                                class: "error-message",
                                "{error}"
                            }
                        }
                    }

                    div {
                        class: "form-options",
                        label {
                            class: "checkbox-label",
                            input {
                                r#type: "checkbox",
                                id: "rememberMe",
                                checked: remember_me(),
                                onchange: handle_remember_me_change,
                                disabled: is_loading(),
                            }
                            span {
                                class: "checkbox-custom"
                            }
                            span {
                                class: "checkbox-text",
                                "Garder votre session enregistrée pour 10 minutes"
                            }
                        }
                    }

                    button {
                        r#type: "submit",
                        class: "connect-button",
                        disabled: is_loading() || connection_id().trim().is_empty(),
                        if is_loading() {
                            span {
                                class: "button-loader",
                                div {
                                    class: "spinner"
                                }
                            }
                        } else {
                            span {
                                class: "button-text",
                                "⏎ Enter the Citadel"
                            }
                        }
                    }
                }

                div {
                    class: "connection-footer",
                    p {
                        class: "footer-text",
                        "Pour toutes questions, veuillez les poser au ..."
                    }
                }
            }
        }
    }
}
