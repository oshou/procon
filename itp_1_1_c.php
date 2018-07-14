<?php
$str = trim(fgets(STDIN));
$x = preg_split("/ /",$str);
echo $x[0]*$x[1]," ",($x[0]+$x[1])*2;
