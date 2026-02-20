use dioxus::prelude::*;
const STYLE: Asset = asset!("./cache_detail.css");

#[component]
pub fn LettreCamp() -> Element {
    let lettre: String = "
Suite au concile du Mont-Royal, vous avez certainement saisi l'ampleur de la situation qui se dessine pour le Royaume. Au fil des derniers siècles, vous avez montré votre courage et votre loyauté en protégeant la paix dans vos régions respectives. Aujourd’hui, malgré vos anciennes divisions, qui ont façonné vos ordres et votre grand sentiment de fierté, c’est un appel à l’unité que nous vous demandons de mettre au service d’un sauvetage extraordinaire.

Nos compatriotes qui vivent dans les pays d’Orient ont besoin d’aide. Un peuple venu de l’est, appelé les Hydres, a envahi leurs territoires. Ils ont pris le contrôle de plusieurs régions, détruits des villes,  des routes commerciales et causé déjà trop de souffrance. Nous ne pouvons pas rester les bras croisés. Tôt ou tard se sera notre ville, Byzance, qui sera sur la route des Hydres. Rassemblons nous pour défendre une cause juste et honorable. Vous aurez l’occasion de servir le Royaume avec fierté, de travailler en équipe et de vous dépasser. Ceux qui s'engagent dans cette mission gagneront respect, honneur et reconnaissance.

C’est pourquoi le 15 mai à 18h30, date de fin de la Paix et de la Trêve de guerre, un convoi militaire vous attendra pour vous mener à Byzance où le périple commencera. C’est le moment de transformer les conflits inutiles en un combat pour une cause noble. C’est le moment de devenir de véritables défenseurs de la Liberté. Nous comptons sur vous.

Le très honorable Urbain II
".to_string();

    rsx! {
        document::Link { rel: "stylesheet", href: STYLE }

        style { {"
            .preserve-line {
                white-space: pre-line;
            }
        "} }

        div { class: "cache-detail-view-header",

            div {
                class: "success-icon",
                style: "font-size: 48px; margin-bottom: 20px;",
                "📜"
            }

            h1 {
                class: "cache-detail-view-title",
                "Lettre de camp"
            }

            p {
                class: "cache-detail-description preserve-line",
                span { class: "value", "{lettre}" }
            }
        }
    }
}
