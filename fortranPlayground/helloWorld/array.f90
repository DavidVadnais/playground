!https://www.fortrantutorial.com/arrays-formatted-io/index.php

program av2
    implicit none
    real ,dimension(5) :: x
    real                :: average,sum
    integer             :: i
    print *, 'enter 5 numbers'
    sum=0.0
    do i=1,5
    	read *, x(i)
    	sum=sum+x(i) 
    end do
    average=sum/10
    print *, 'the average is ',average
    print *, 'the numbers are'
    print *,x
    print *, 'the hyperbolic sin of your first number is :'
    print *, SINH(x(0))
end program av2