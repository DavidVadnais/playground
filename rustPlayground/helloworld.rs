// Playing with rust
/*based on TutorialsPoint -- rust*/

// crates.io holds dependencies 
// cargo is package manager

fn 
main() {
    println!("Hello, World!");
    
    //let company_string = "TutorialsPoint";  // string type
    let company_string:&str = "TutorialsPoint";  // string literal type
    let rating_float = 4.5;                 // float type
    let is_growing_boolean = true;          // boolean type
    // let icon_char = '♥';                    //unicode character type

    println!("company name is:{}",company_string);
    println!("company rating on 5 is:{}",rating_float);
    println!("company is growing :{}",is_growing_boolean);
    
    
    // how to desig type
    let result = 10;    // i32 by default
    let age:u32 = 20;
    
    // neat syntax sugar
    let int_with_separator = 50_000;
    println!("int value {}",int_with_separator);
    
    //booleans 
    let isfun:bool = true;
    println!("Is Rust Programming Fun ? {}",isfun);
    
    // default immutable
    let mut fees:i32 = 25_000;//the mut here is key
    println!("fees is {} ",fees);
    fees = 35_000;
    println!("fees changed is {}",fees);
    
    // empty string
    let mut z = String::new();
    z.push_str("David");
    println!("{}",z);
    
    //skipping operators because they are normal
    
    // skipping if
    
    //match 
    let state_code = "co";
    let state = match state_code {
      "co" => {println!("Found match for co"); "Colorado"},
      "KL" => "Kerala",
      "KA" => "Karnadaka",
      "GA" => "Goa",
      _ => "Unknown"
    };
    println!("State name is {}",state);
    
    //
    for x in 1..11{ // 11 is not inclusive
      if x==5 {
          continue;
      }
      //println!("x is {}",x);
    }
   
    //skip while
   
    //loop
    let mut x = 0;
    loop {// aka while true
      x+=1;
      
      if x==15 {
        println!("x={}",x);
        break;
      }
   }
   
    //calling a function
    fn_hello();
    
    //tuples
    let tuple:(i32,f64,u8) = (-325,4.9,22);
    println!("{:?}",tuple);
    println!("float is :{:?}",tuple.1);
    
    //arrays
    let arr:[i32;4] = [10,20,30,40];
    println!("array is {:?}",arr);
    println!("array size is :{}",arr.len());
    
    for val in arr.iter(){
      println!("value is :{}",val);
    }
    
    //stack and heap and borrow and struct and enum - skipping
    
    
}

//Defining a function
fn fn_hello(){
   println!("hello from function fn_hello ");
}

//no semicolon means this value is returned
fn get_pi()->f64 {
   22.0/7.0
}
