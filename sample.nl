np("========")
var line = ns("enter the S P A C E  : ");
var point = 0;
var val = [0];
var valPointer = 0;
var arrayExpand = 1;

while(point < len(line)){

    if(line[point] == " "){

        val[valPointer] = val[valPointer]+1;

    } elf(line[point] ==  "^"){

        val[valPointer] = val[valPointer] - 1;

    } elf(line[point] == "_"){
        np(chr(val[valPointer]));

    } elf(line[point] == "}"){
        if(valPointer == len(val) - 1 ){
            valPointer = valPointer+1;
            push(val,0);
            arrayExpand = arrayExpand +1;
        }else{
            valPointer = valPointer +1;
        }
    } elf(line[point] == "\"){
        if(val[valPointer] != 0){
            while(line[point] != "/" ){
                point = point - 1;
            }
        }

    } elf(line[point] == "{"){
        valPointer = valPointer -1;
    }

    point = point+1;

}
