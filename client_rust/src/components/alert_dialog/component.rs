use dioxus::prelude::*;
use dioxus_primitives::alert_dialog::{
    self, AlertDialogActionProps, AlertDialogActionsProps, AlertDialogCancelProps,
    AlertDialogContentProps, AlertDialogDescriptionProps, AlertDialogRootProps,
    AlertDialogTitleProps,
};

#[component]
pub fn AlertDialogRoot(props: AlertDialogRootProps) -> Element {
    rsx! {
        // The backdrop should be rendered HERE, not as a class on the root
        if props.open.read().unwrap_or(false) {
            div {
                class: "alert-dialog-backdrop",
                // onclick: move |_| {
                //     // Optional: close on backdrop click
                //     if let Some(on_change) = props.on_open_change {
                //         on_change(false);
                //     }
                // },
                // The actual dialog content
                alert_dialog::AlertDialogRoot {
                    class: "",  // Remove the backdrop class from here
                    id: props.id,
                    default_open: props.default_open,
                    open: props.open,
                    on_open_change: props.on_open_change,
                    attributes: props.attributes,
                    {props.children}
                }
            }
        }
    }
}

#[component]
pub fn AlertDialogContent(props: AlertDialogContentProps) -> Element {
    rsx! {
        alert_dialog::AlertDialogContent {
            id: props.id,
            class: props.class.unwrap_or_default() + " alert-dialog",
            attributes: props.attributes,
            {props.children}
        }
    }
}

// ... rest of your components remain the same

#[component]
pub fn AlertDialogTitle(props: AlertDialogTitleProps) -> Element {
    alert_dialog::AlertDialogTitle(props)
}

#[component]
pub fn AlertDialogDescription(props: AlertDialogDescriptionProps) -> Element {
    alert_dialog::AlertDialogDescription(props)
}

#[component]
pub fn AlertDialogActions(props: AlertDialogActionsProps) -> Element {
    rsx! {
        alert_dialog::AlertDialogActions { class: "alert-dialog-actions", attributes: props.attributes, {props.children} }
    }
}

#[component]
pub fn AlertDialogCancel(props: AlertDialogCancelProps) -> Element {
    rsx! {
        alert_dialog::AlertDialogCancel {
            on_click: props.on_click,
            class: "alert-dialog-cancel",
            attributes: props.attributes,
            {props.children}
        }
    }
}

#[component]
pub fn AlertDialogAction(props: AlertDialogActionProps) -> Element {
    rsx! {
        alert_dialog::AlertDialogAction {
            class: "alert-dialog-action",
            on_click: props.on_click,
            attributes: props.attributes,
            {props.children}
        }
    }
}
