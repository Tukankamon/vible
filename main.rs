use std::fs;

fn main(){
    
    let content = fs::read_to_string("bible/kjv.txt")
        .expect("There has been an error reading the file"); 

    // Print the contents
    println!("File Contents:\n{}", content);
}