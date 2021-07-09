-- Based on https://riptutorial.com/ada/example/15002/hello-world
-- compile with  :  gnatmake hello_world

with Ada.Text_IO; use Ada.Text_IO;

procedure Hello_World is
begin
    Put_Line ("Hello World");
end Hello_World;