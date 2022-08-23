fn main() {
   let mut nombre: String = String::new();
   println!("Hola Soy {}", nombre);

   loop {
       std::io::stdin().read_line(&mut nombre).unwrap();
       println!("Hola Soy {}", nombre);
    }

}
