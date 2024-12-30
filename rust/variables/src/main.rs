fn main() {
    let abc: (f32, f32, f32) = (0.1, 0.2, 0.3);
    let xyz: (f64, f64, f64) = (0.1, 0.2, 0.3);

    println!("abc (f32)");
    println!("    0.1 + 0.2: {}", (abc.0 + abc.1).to_bits());
    println!("          0.3: {}", (abc.2).to_bits());
    println!();
    println!("abc (f64)");
    println!("    0.1 + 0.2: {}", (xyz.0 + xyz.1).to_bits());
    println!("          0.3: {}", (xyz.2).to_bits());
    println!();

    assert!(abc.0 + abc.1 == abc.2);
    assert!(xyz.0 + xyz.1 == xyz.2);
}
