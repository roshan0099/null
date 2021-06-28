var sam = 1;
var numO = 0;
var numT = 0;
var choice = 0;
np("haai");
while(sam > 0 ){
    np("type two numbers : ");
    numO = ni();
    numT = ni();

    choice = ni("Enter your coice 1. Add 2. Subb");
    if(choice == 1){
        np("result be: ", numO+numT);
    }
    if (choice == 2){
        np("result be : ", numO - numT);
    }

    sam = ni("you wanna conntinue ? [0n/1y] :  ");
}

