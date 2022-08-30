

fn main() {

    let mut nombres : Vec<String> = Vec::new();
    for _i in 0..3 {
        println!("Por favor introduce un nombre");
        let mut nombre  = String::new();
        std::io::stdin().read_line(&mut nombre).unwrap();
        
        nombres.push(nombre.trim().to_string());
    }
    
    for nombre in nombres {
        println!("El nombre es: {}", nombre)
    }

    let  hola = ["H", "O", "L", "A"];

    println!("{}",hola[0]);
    println!("{}",hola[1]);

}
