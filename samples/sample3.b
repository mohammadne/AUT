intFunction: function integer(intVar:integer, boolVar: boolean)
{
	if(boolVar)
	{
		intVar = intVar * 2;
	}
	return intVar;
}
// finds the largest value in the array
getLargest: function integer(intArrVar:array [] integer, int len) =
{
	largest:integer = 0;
	x:integer;
	for(x = 0; x < len; x++)
	{
		if(intArrVar[x] > largest)
		{
			largest = intArrVar[x];
		}
	}
	return largest;

}

main: function integer ( argc: integer, argv: array [] string ) =
{
	b:array [10] integer = {0,1,2,3,4,5};
	b[4] = 6;
	i:integer;
	z:integer = intFunction(i, true);
	print intArrFunction(v, z);
}
