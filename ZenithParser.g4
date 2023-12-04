parser grammar ZenithParser;
options { tokenVocab = ZenithLexer; }

fileStat : endedStat+
         | EOF
         ;

endedStat : NL? stat lineEnd ;

lineEnd : (NL | SEMICOLON)+
        | EOF
        ;

stat : ID Type=fullType? INIT_ASSIGN NL? expr #defineStat
     | LBRACE NL? endedStat* stat? lineEnd? RBRACE #multiStat
     | expr #exprStat
     ;

fullType : partType | ptrType ;

ptrType : Ptr=(HASH | AND) Type=fullType ;

partType : TYPE #baseType
         | LBRACK expr? RBRACK Type=fullType #sliceType
         ;

expr : NUM #numExpr
     | ID #idExpr
     | Key=(TRUE | FALSE | NULL) #keyExpr
     | Type=partType LBRACE initArgs RBRACE #initExpr
     | LPAREN NL? expr NL? RPAREN #parenExpr
     | expr QUESTION #postfixExpr
     | ALLOC LPAREN NL? expr NL? RPAREN #allocExpr
     | DEALLOC LPAREN NL? expr NL? RPAREN #deallocExpr
     | TYPE LPAREN NL? expr NL? RPAREN #castExpr
     | Op=(AND | AT) expr #ptrExpr
     | Op=(PLUS | MINUS | NOT) expr #prefixExpr
     | Left=expr (POW NL? Right=expr)+ #powExpr
     | Left=expr Op=(TIMES | DIVIDE | REM) NL? Right=expr #mulExpr
     | Left=expr Op=(PLUS | MINUS) NL? Right=expr #addExpr
     | Left=expr Op=(LSHIFT | RSHIFT) NL? Right=expr #shiftExpr
     | Left=expr AND NL? Right=expr #bitAndExpr
     | Left=expr DOLLAR NL? Right=expr #bitXorExpr
     | Left=expr PIPE NL? Right=expr #bitOrExpr
     | Left=expr Op=(LT | GT | LTE | GTE | EQ | NEQ) NL? Right=expr #compExpr
     | Left=expr IF NL? Condition=expr NL? ELSE NL? Right=expr #ifExpr
     ;

initArgs : (expr (COMMA expr)* COMMA?)? ;
