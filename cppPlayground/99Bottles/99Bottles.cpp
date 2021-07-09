#include <string>
#include<iostream>
using namespace std;

// sort of used https://www.w3schools.com/cpp/cpp_classes.asp
class MyBottles {       // The class
  public:             // Access specifier
    int myNum = 99;        // Attribute (int variable)
    string myString = "Bottles Of Beer On The Wall:  ";  // Attribute (string variable)
    string myString2 = "Take one down pass it around "; 
};

int main() {
  MyBottles shelf;  // Create an object of MyBottles

  // Print attribute values
  while ( shelf.myNum > 0) {
      cout << shelf.myString << shelf.myNum << endl;
      shelf.myNum--;
  }

  return 0;
}
