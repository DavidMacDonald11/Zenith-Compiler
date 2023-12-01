parser grammar ZenithParser;
options { tokenVocab = ZenithLexer; }

fileStat : endedStat+
         | EOF
         ;

endedStat : NL? stat lineEnd ;

lineEnd : (NL | SEMICOLON)+
        | EOF
        ;

stat : ID type? INIT_ASSIGN NL? expr #defineStat
     | LBRACE NL? endedStat* stat? lineEnd? RBRACE #multiStat
     | expr #exprStat
     ;

type : TYPE #baseType
     | type HASH? Ref=(EXCLAIM | QUESTION) #refType
     ;

expr : NUM #numExpr
     | ID #idExpr
     | Key=(TRUE | FALSE | NULL) #keyExpr
     | LPAREN NL? expr NL? RPAREN #parenExpr
     | Left=expr HASH? Op=(EXCLAIM | QUESTION) #refExpr
     | Left=expr NOT_NULL #notNullExpr
     | ALLOC Ref=(EXCLAIM | QUESTION) LPAREN NL? expr NL? RPAREN #allocExpr
     | DEALLOC LPAREN NL? expr NL? RPAREN #deallocExpr
     | TYPE LPAREN NL? expr NL? RPAREN #castExpr
     | Op=(PLUS | MINUS | EXCLAIM) Right=expr #prefixExpr
     | Left=expr (POW NL? Right=expr)+ #powExpr
     | Left=expr Op=(TIMES | DIVIDE | REM) NL? Right=expr #mulExpr
     | Left=expr Op=(PLUS | MINUS) NL? Right=expr #addExpr
     | Left=expr Op=(LSHIFT | RSHIFT) NL? Right=expr #shiftExpr
     | Left=expr AND NL? Right=expr #bitAndExpr
     | Left=expr DOLLAR NL? Right=expr #bitXorExpr
     | Left=expr PIPE NL? Right=expr #bitOrExpr
     | Left=expr Op=(LT | GT | LTE | GTE | EQ | NEQ) NL? Right=expr #compExpr
     | Left=expr COALESCE Right=expr #coalesceExpr
     | Left=expr IF NL? Condition=expr NL? ELSE NL? Right=expr #ifExpr
     ;
