lexer grammar ZenithLexer;

PLUS : '+' ;
MINUS : '-' ;
TIMES : '*' ;
DIVIDE : '/' ;

NUM :       DEC_DIGIT+ | DEC_DIGIT* '.' DEC_DIGIT+ | DEC_DIGIT+ '.' DEC_DIGIT*
    | '0b' (BIN_DIGIT+ | BIN_DIGIT* '.' BIN_DIGIT+ | BIN_DIGIT+ '.' BIN_DIGIT*)
    | '0o' (OCT_DIGIT+ | OCT_DIGIT* '.' OCT_DIGIT+ | OCT_DIGIT+ '.' OCT_DIGIT*)
    | '0x' (HEX_DIGIT+ | HEX_DIGIT* '.' HEX_DIGIT+ | HEX_DIGIT+ '.' HEX_DIGIT*)
    ;

fragment BIN_DIGIT : [0-1] ;
fragment OCT_DIGIT : [0-7] ;
fragment DEC_DIGIT : [0-9] ;
fragment HEX_DIGIT : [0-9a-fA-F] ;

NL : ([\r]? [\n])+ ;
SPACE : [ \t] -> skip ;
