use std::io;

fn main() {
    loop {
        println!("Please type temperature number!");

        let mut temperature = String::from("");
        io::stdin()
            .read_line(&mut temperature)
            .expect("Please type temperature number!");
        let temperature = temperature.trim().parse::<f64>();

        let temperature = match temperature {
            Ok(num) => num,
            Err(_) => {
                println!("Please type a number!");
                continue;
            }
        };

        let temp_c = (temperature - 32.0) / 1.8;
        let temp_f = (temperature + 32.0) * 1.8;
        println!("Your typing temperature number is {}", temperature);
        println!(
            "if the temperature is Colsius, the Fahrenheit temperature is {}",
            temp_f
        );
        println!(
            "if the temperature is Fahrenheit, the Colsius temperature is {:.2}",
            temp_c
        );
        break;
    }
}
