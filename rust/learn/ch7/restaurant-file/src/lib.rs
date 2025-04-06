// include the front_of_house file
mod front_of_house;
mod insite;

pub use crate::front_of_house::hosting;
use crate::insite as insitemod;

pub fn eat_at_restaurant() {
    insitemod::getinto_in_site();
    hosting::add_to_waitlist();
}
