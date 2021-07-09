!https://www.fortrantutorial.com/arrays-formatted-io/index.php

! Essentialy just use sub modules and don't use functions.  
! Functions can only return 1 value 
program output1
   implicit none
   real,dimension(3) :: a,b,c
   !initialise arrays
   a = 1.7
   b = 2.8
   c = 3.8
   
   write(*,1) 'a',a
  call prompt()
   write(*,1) 'b',b
  call prompt()
   write(*,1) 'c',c
  call prompt()
   write(*,1) 'a*b*c',a * b * c
   1        format(a,3f8.3)
end program output1
!++++++++++++++++++++++++++++++++++++++++++++++     
subroutine prompt()
  !prompts for a keypress
  implicit none
  character answer*1
  print *, 'type y to continue or any other key to finish'
  read *, answer
  if (answer /= 'y') stop
end subroutine prompt