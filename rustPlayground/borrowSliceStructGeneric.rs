/******************************************************************************

https://www.tutorialspoint.com/rust/rust_slices.htm

*******************************************************************************/
use std::fmt::Display;

fn main(){
    //////////Borrowing
   // a list of nos
   let v = vec![10,20,30];
   print_vector(&v); // passing reference
   println!("Printing the value from main() v[0]={}",v[0]);
   
    //////////slices
    let data = [10,20,30,40,50];
    use_slice(&data[1..4]);

    ///////// struct
    let fido = Dog {
        name:String::from("fido"),
        age:5,
        teeth:12,
        gender:GenderCategory::Male
    };
    println!("Name is : {} age is {} teeth are {} gender is {:?}",fido.name,fido.age,fido.teeth, fido.gender);
    
    ///////// generics
    print_pro(42 as u16);
    print_pro("Hello David");
}
fn print_pro<T:Display>(t:T){
   println!("Inside print_pro generic function:");
   println!("{}",t);
}

fn print_vector(x:&Vec<i32>){
   println!("Inside print_vector function {:?}",x);
}

fn use_slice(slice:&[i32]) { 
   // is taking a slice or borrowing a part of an array of i32s
   println!("length of slice is {:?}",slice.len());
   println!("{:?}",slice);
}
struct Dog{
    name:String,
    age:u32,
    teeth:u32,
    gender:GenderCategory
}

// The `derive` attribute automatically creates the implementation
// required to make this `enum` printable with `fmt::Debug`.
#[derive(Debug)]
enum GenderCategory {
   Male,Female
}
