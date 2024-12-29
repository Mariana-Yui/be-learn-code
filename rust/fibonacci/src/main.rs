use std::io;

fn main() {
    println!("Please type fibonacci's n!");

    let mut n = String::from("");
    io::stdin().read_line(&mut n).expect("Please type a number");

    let n = n.trim().parse().expect("Please type a number");

    let result = fibonacci(n);
    println!("the result of fibonacci({}) is {}", n, result);
}

fn fibonacci(n: i32) -> i32 {
    if n == 0 {
        0
    } else if n == 1 {
        1
    } else {
        fibonacci(n - 1) + fibonacci(n - 2)
    }
}
