!https://www.fortrantutorial.com/loops/index.php

Program Hello
  implicit none                                        ! checks if you did good
  !an example of program  structure                    !b: a  comment  
  real :: answer,x,y                                   !c: declarations
  integer  ::  i

  print  *, 'Enter two numbers'                        !d: output
  read  *, x                                           !e: input
  read  *, y                                           !e: input
  answer=x+y                                           !f: arithmetic
  print  *, 'The total is ', answer                    !g: output
  print *, "Hello World"
  
  !lOOPS
  
  do  i=0,20
  	print  *,i
  end  do
End Program Hello