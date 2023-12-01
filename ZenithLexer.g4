lexer grammar ZenithLexer;

NOT_NULL : '!!' ;
BIT_NOT : '~' ;
POW : '^' ;
TIMES : '*' ;
DIVIDE : '/' ;
REM : '%' ;
PLUS : '+' ;
MINUS : '-' ;
LSHIFT : '<<' ;
RSHIFT : '>>' ;
LT : '<' ;
GT : '>' ;
LTE : '<=' ;
GTE : '>=' ;
EQ : '==' ;
NEQ : '!=' ;
COALESCE : '??';

NULL : 'null' ;
TRUE : 'true' ;
ALLOC : 'alloc' ;
DEALLOC : 'dealloc' ;
FALSE : 'false' ;
IF : 'if' ;
ELSE : 'else' ;

ASSIGN : '=' ;
INIT_ASSIGN : ':=' ;

TYPE
    : 'UInt8' | 'UInt16' |  'UInt32' |  'UInt64' | 'UInt'
    |  'Int8' |  'Int16' |   'Int32' |   'Int64' |  'Int'
                         | 'Float32' | 'Float64'
    | 'Bool'
    ;

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

ID : [_a-zA-Z][_a-zA-Z0-9]* ;

AND : '&' ;
DOLLAR : '$' ;
PIPE : '|' ;
EXCLAIM : '!' ;
QUESTION : '?' ;
HASH : '#' ;
LPAREN : '(' ;
RPAREN : ')' ;
LBRACE : '{' ;
RBRACE : '}' ;
LBRACK : '[' ;
RBRACK : ']' ;
SEMICOLON : ';' ;
NL : ([\r]? [\n])+ ;
IGNORED : ([ \t] | '/*' .*? '*/' | '//' ~[\n]*) -> skip ;
