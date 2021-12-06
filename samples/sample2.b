//Functions

square: function integer ( x: integer ) = {
    return x^2;
}

printarray: function void
 ( a: array [] integer, size: integer ) = {
    i: integer;
    for( i=0 ; i < size ; i++) {
       print a[i], "\n";
    }
}
