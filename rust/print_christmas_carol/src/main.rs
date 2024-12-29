fn main() {
    let days_action = [
        "A Partridge in a Pear Tree",
        "Two Turtle Doves",
        "Three French Hens",
        "Four Calling Birds",
        "Five Gold Rings",
        "Six Geese a-Laying",
        "Seven Swans a-Swimming",
        "Eight Maids a-Milking",
        "Nine Ladies Dancing",
        "Ten Lords a-Leaping",
        "Eleven Pipers Piping",
        "Twelve Drummers Drumming",
    ];
    let days = [
        "first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth",
        "tenth", "eleventh", "twelfth",
    ];

    for i in 0..12 {
        print!(
            "On the {} day of Christmas, my true love sent to me: ",
            days[i]
        );
        let mut n = i;
        while n > 0 {
            print!("{}, ", days_action[n]);
            n -= 1;
        }
        if i == 0 {
            println!("{}.", days_action[n]);
        } else {
            println!("And {}.", days_action[n].replace("A", "a"));
        }
    }
}
