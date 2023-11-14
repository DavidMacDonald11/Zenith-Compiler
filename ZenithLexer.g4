lexer grammar ZenithLexer;

PLUS : '+' ;
MINUS : '-' ;
TIMES : '*' ;
DIVIDE : '/' ;

NUM
    : DEC_NUM
    | BIN_NUM
    | OCT_NUM
    | HEX_NUM
    ;

fragment DEC_NUM
    : DEC_SEG
    | [0-9_]* '.' DEC_SEG
    | DEC_SEG '.' [0-9_]*
    ;

fragment DEC_SEG : [0-9] [0-9_]* ;

fragment BIN_NUM : '0b' (
      BIN_SEG
    | [0-1_]* '.' BIN_SEG
    | BIN_SEG '.' [0-1_]*
) ;

fragment BIN_SEG : [0-1_]* [0-1] [0-1_]* ;

fragment OCT_NUM : '0o' (
      OCT_SEG
    | [0-7_]* '.' OCT_SEG
    | OCT_SEG '.' [0-7_]*
) ;

fragment OCT_SEG : [0-7_]* [0-7] [0-7_]* ;

fragment HEX_NUM : '0x' (
      HEX_SEG
    | (HEX_DIGIT | '_')* '.' HEX_SEG
    | HEX_SEG '.' (HEX_DIGIT | '_')*
) ;

fragment HEX_SEG : (HEX_DIGIT | '_')* HEX_DIGIT (HEX_DIGIT | '_')* ;
fragment HEX_DIGIT : [0-9a-fA-F] ;

NL : ([\r]? [\n])+ ;
SPACE : [ \t] -> skip ;
