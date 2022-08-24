use regex::Regex;

fn main() {

    // Regex
    let re_add = Regex::new(r"(\d+)\s?\+\s?(\d+)").unwrap();
    let re_mult = Regex::new(r"(\d+)\s?\*\s?(\d+)").unwrap();

    // Traer datos del usuario
    println!("Por favor uintroduce tu expresion");
    let mut expresion = String::new();
    std::io::stdin().read_line(&mut expresion).unwrap();

    loop {
        // Aplicar Operaciones
        let caps = re_mult.captures(expresion.as_str());
        if caps.is_none() {
            break
        }
        let caps = caps.unwrap();
        let cap_expression = caps.get(0).unwrap().as_str();
        let left_value: i32= caps.get(1).unwrap().as_str().parse().unwrap();
        let right_value: i32= caps.get(2).unwrap().as_str().parse().unwrap();
        let mult = left_value * right_value;
        expresion = expresion.replace(cap_expression, &mult.to_string())
        
    }

    loop {
        // Aplicar Operaciones
        let caps = re_add.captures(expresion.as_str());
        if caps.is_none() {
            break
        }
        let caps = caps.unwrap();
        let cap_expression = caps.get(0).unwrap().as_str();
        let left_value: i32= caps.get(1).unwrap().as_str().parse().unwrap();
        let right_value: i32= caps.get(2).unwrap().as_str().parse().unwrap();
        let addition = left_value + right_value;
        expresion = expresion.replace(cap_expression, &addition.to_string())
        
    }
    // mostrar el resuktdi


    print!("Resultado {}", expresion)

}
