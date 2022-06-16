import operations

fnum = int(input("Enter the first no. : "))
snum = int(input("Enter the second no. : ")) 
operation = int(input(" 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Other\n Enter your choice: "))

if operation==1:
    operations.add(fnum, snum)
elif operation == 2:
    operations.subtract(fnum, snum)
elif operation == 3:
    operations.multiply(fnum, snum)
elif operation == 4:
    operations.divide(fnum, snum)
elif operation == 5:
    eoperation = int(input(" 1. Exponent\n Enter you choice: "))
    if eoperation == 1:
        operations.exponent(fnum, snum)
    else:
        print("Operation not found\n Exiting.....")
else:
    print("Operation not found\n Exiting.....")
