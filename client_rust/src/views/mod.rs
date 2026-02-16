//! The views module contains the components for all Layouts and Routes for our app. Each layout and route in our [`Route`]
//! enum will render one of these components.
//!
//!
//! The [`Home`] and [`Blog`] components will be rendered when the current route is [`Route::Home`] or [`Route::Blog`] respectively.
//!
//!
//! The [`Navbar`] component will be rendered on all pages of our app since every page is under the layout. The layout defines
//! a common wrapper around all child routes.

mod home_view;
pub use home_view::HomeView;

mod user_view;
pub use user_view::UserView;

mod cache_detail_view;
pub use cache_detail_view::CacheDetailView;

mod cache_list_view;
pub use cache_list_view::CacheListView;

mod leaderboard;
pub use leaderboard::Leaderboard;
