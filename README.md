The standard for the conversion file:
from,to,value;

Example:
rem,px,16;

1 rem is equal to 16 px
"Standard" call:
uconv rempx 5 == 5 * 16 to get 80px

"Reverse" call:
uconv pxrem 80 == 80 / 16 to get 5rem
Will use the same conversion entry but will just reverse it

Uses badger to store the different conversions