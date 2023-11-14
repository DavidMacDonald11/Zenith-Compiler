grammar Zenith;

MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUM: [0-9]+;
NL: [\n]+;
WS: [ \r\t]+ -> skip;

fileStat : expr EOF;

expr
    : expr op=(MUL|DIV) expr # MulExpr
    | expr op=(ADD|SUB) expr # AddExpr
    | NUM # Num
    ;
